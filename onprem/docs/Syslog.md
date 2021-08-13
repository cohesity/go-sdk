# \Syslog

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSyslogServer**](Syslog.md#AddSyslogServer) | **Post** /syslog | Add Syslog Server
[**GetSupportedSyslogProgramNames**](Syslog.md#GetSupportedSyslogProgramNames) | **Get** /syslog/program-names | Get supported program names.
[**GetSyslogAuditTags**](Syslog.md#GetSyslogAuditTags) | **Get** /syslog/audit-tags | Get cluster audit tags.
[**GetSyslogServerById**](Syslog.md#GetSyslogServerById) | **Get** /syslog/{id} | Get a syslog server by id.
[**GetSyslogServerStatusById**](Syslog.md#GetSyslogServerStatusById) | **Get** /syslog/{id}/status | Get a syslog server reachability status.
[**GetSyslogServers**](Syslog.md#GetSyslogServers) | **Get** /syslog | Get list of syslog servers.
[**PatchSyslogServerById**](Syslog.md#PatchSyslogServerById) | **Patch** /syslog/{id} | Patch a syslog server by id.
[**RemoveSyslogServer**](Syslog.md#RemoveSyslogServer) | **Delete** /syslog/{id} | Remove syslog server by id
[**RemoveSyslogServers**](Syslog.md#RemoveSyslogServers) | **Delete** /syslog | Remove syslog servers
[**UpdateSyslogAuditTags**](Syslog.md#UpdateSyslogAuditTags) | **Post** /syslog/audit-tags | Update cluster audit tags.
[**UpdateSyslogServerById**](Syslog.md#UpdateSyslogServerById) | **Put** /syslog/{id} | Update a syslog server by id.



## AddSyslogServer

> SyslogServer AddSyslogServer(ctx).Body(body).Execute()

Add Syslog Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewSyslogServer() // SyslogServer | Specifies parameters to add syslog server.

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiAddSyslogServerRequest{
        Body: &body
    }

    resp, r, err := api_client.Syslog.AddSyslogServer(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Syslog.AddSyslogServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSyslogServer`: SyslogServer
    fmt.Fprintf(os.Stdout, "Response from `Syslog.AddSyslogServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSyslogServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SyslogServer**](SyslogServer.md) | Specifies parameters to add syslog server. | 

### Return type

[**SyslogServer**](SyslogServer.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportedSyslogProgramNames

> []string GetSupportedSyslogProgramNames(ctx).Execute()

Get supported program names.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiGetSupportedSyslogProgramNamesRequest{
    }

    resp, r, err := api_client.Syslog.GetSupportedSyslogProgramNames(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Syslog.GetSupportedSyslogProgramNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupportedSyslogProgramNames`: []string
    fmt.Fprintf(os.Stdout, "Response from `Syslog.GetSupportedSyslogProgramNames`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportedSyslogProgramNamesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyslogAuditTags

> SyslogAuditTag GetSyslogAuditTags(ctx).Execute()

Get cluster audit tags.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiGetSyslogAuditTagsRequest{
    }

    resp, r, err := api_client.Syslog.GetSyslogAuditTags(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Syslog.GetSyslogAuditTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSyslogAuditTags`: SyslogAuditTag
    fmt.Fprintf(os.Stdout, "Response from `Syslog.GetSyslogAuditTags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyslogAuditTagsRequest struct via the builder pattern


### Return type

[**SyslogAuditTag**](SyslogAuditTag.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyslogServerById

> SyslogServer GetSyslogServerById(ctx, id).Execute()

Get a syslog server by id.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    id := int32(56) // int32 | Specifies the id of syslog server.

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiGetSyslogServerByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.Syslog.GetSyslogServerById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Syslog.GetSyslogServerById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSyslogServerById`: SyslogServer
    fmt.Fprintf(os.Stdout, "Response from `Syslog.GetSyslogServerById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Specifies the id of syslog server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyslogServerByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyslogServer**](SyslogServer.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyslogServerStatusById

> SyslogServerStatus GetSyslogServerStatusById(ctx, id).Execute()

Get a syslog server reachability status.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    id := int32(56) // int32 | Specifies the id of syslog server.

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiGetSyslogServerStatusByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.Syslog.GetSyslogServerStatusById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Syslog.GetSyslogServerStatusById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSyslogServerStatusById`: SyslogServerStatus
    fmt.Fprintf(os.Stdout, "Response from `Syslog.GetSyslogServerStatusById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Specifies the id of syslog server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyslogServerStatusByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyslogServerStatus**](SyslogServerStatus.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyslogServers

> SyslogServers GetSyslogServers(ctx).Execute()

Get list of syslog servers.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiGetSyslogServersRequest{
    }

    resp, r, err := api_client.Syslog.GetSyslogServers(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Syslog.GetSyslogServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSyslogServers`: SyslogServers
    fmt.Fprintf(os.Stdout, "Response from `Syslog.GetSyslogServers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyslogServersRequest struct via the builder pattern


### Return type

[**SyslogServers**](SyslogServers.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSyslogServerById

> SyslogServer PatchSyslogServerById(ctx, id).Body(body).Execute()

Patch a syslog server by id.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    id := int32(56) // int32 | Specifies the id of syslog server.
    body := *openapiclient.NewSyslogServer() // SyslogServer | Specifies the body of syslog server fields to patch. (optional)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiPatchSyslogServerByIdRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.Syslog.PatchSyslogServerById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Syslog.PatchSyslogServerById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSyslogServerById`: SyslogServer
    fmt.Fprintf(os.Stdout, "Response from `Syslog.PatchSyslogServerById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Specifies the id of syslog server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSyslogServerByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SyslogServer**](SyslogServer.md) | Specifies the body of syslog server fields to patch. | 

### Return type

[**SyslogServer**](SyslogServer.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSyslogServer

> RemoveSyslogServer(ctx, id).Execute()

Remove syslog server by id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    id := int32(56) // int32 | Specifies a unique id of the syslog server.

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiRemoveSyslogServerRequest{
        Id: &id
    }

    resp, r, err := api_client.Syslog.RemoveSyslogServer(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Syslog.RemoveSyslogServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Specifies a unique id of the syslog server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSyslogServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSyslogServers

> RemoveSyslogServers(ctx).Execute()

Remove syslog servers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiRemoveSyslogServersRequest{
    }

    resp, r, err := api_client.Syslog.RemoveSyslogServers(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Syslog.RemoveSyslogServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSyslogServersRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyslogAuditTags

> SyslogAuditTag UpdateSyslogAuditTags(ctx).Body(body).Execute()

Update cluster audit tags.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewSyslogAuditTag() // SyslogAuditTag | Specifies syslog audit tag to update. (optional)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiUpdateSyslogAuditTagsRequest{
        Body: &body
    }

    resp, r, err := api_client.Syslog.UpdateSyslogAuditTags(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Syslog.UpdateSyslogAuditTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSyslogAuditTags`: SyslogAuditTag
    fmt.Fprintf(os.Stdout, "Response from `Syslog.UpdateSyslogAuditTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSyslogAuditTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SyslogAuditTag**](SyslogAuditTag.md) | Specifies syslog audit tag to update. | 

### Return type

[**SyslogAuditTag**](SyslogAuditTag.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyslogServerById

> SyslogServer UpdateSyslogServerById(ctx, id).Body(body).Execute()

Update a syslog server by id.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    id := int32(56) // int32 | Specifies the id of syslog server.
    body := *openapiclient.NewSyslogServer() // SyslogServer | Specifies the body of syslog server body to update. (optional)

    // create a CohesityClient
    clientConfig := cohesitysdk.CohesityClientConfig{
        Username: "username",
        Password: "password",
        Domain: "LOCAL",
        ClusterVIP: "0.0.0.0",
        // Timeout: 15 * time.Second, // Optional 
    }

    client, err := cohesitysdk.NewCohesityClient(clientConfig)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        return
    }

    request := onprem.ApiUpdateSyslogServerByIdRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.Syslog.UpdateSyslogServerById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Syslog.UpdateSyslogServerById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSyslogServerById`: SyslogServer
    fmt.Fprintf(os.Stdout, "Response from `Syslog.UpdateSyslogServerById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Specifies the id of syslog server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSyslogServerByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SyslogServer**](SyslogServer.md) | Specifies the body of syslog server body to update. | 

### Return type

[**SyslogServer**](SyslogServer.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

