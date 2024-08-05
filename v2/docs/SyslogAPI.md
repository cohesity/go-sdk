# \SyslogAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSyslogServer**](SyslogAPI.md#AddSyslogServer) | **Post** /syslog | Add Syslog Server
[**GetSupportedSyslogProgramNames**](SyslogAPI.md#GetSupportedSyslogProgramNames) | **Get** /syslog/program-names | Get supported program names.
[**GetSyslogAuditTags**](SyslogAPI.md#GetSyslogAuditTags) | **Get** /syslog/audit-tags | Get cluster audit tags.
[**GetSyslogServerById**](SyslogAPI.md#GetSyslogServerById) | **Get** /syslog/{id} | Get a syslog server by id.
[**GetSyslogServerStatusById**](SyslogAPI.md#GetSyslogServerStatusById) | **Get** /syslog/{id}/status | Get a syslog server reachability status.
[**GetSyslogServers**](SyslogAPI.md#GetSyslogServers) | **Get** /syslog | Get list of syslog servers.
[**PatchSyslogServerById**](SyslogAPI.md#PatchSyslogServerById) | **Patch** /syslog/{id} | Patch a syslog server by id.
[**RemoveSyslogServer**](SyslogAPI.md#RemoveSyslogServer) | **Delete** /syslog/{id} | Remove syslog server by id
[**RemoveSyslogServers**](SyslogAPI.md#RemoveSyslogServers) | **Delete** /syslog | Remove syslog servers
[**UpdateSyslogAuditTags**](SyslogAPI.md#UpdateSyslogAuditTags) | **Post** /syslog/audit-tags | Update cluster audit tags.
[**UpdateSyslogServerById**](SyslogAPI.md#UpdateSyslogServerById) | **Put** /syslog/{id} | Update a syslog server by id.



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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewSyslogServer() // SyslogServer | Specifies parameters to add syslog server.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyslogAPI.AddSyslogServer(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogAPI.AddSyslogServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSyslogServer`: SyslogServer
	fmt.Fprintf(os.Stdout, "Response from `SyslogAPI.AddSyslogServer`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyslogAPI.GetSupportedSyslogProgramNames(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogAPI.GetSupportedSyslogProgramNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportedSyslogProgramNames`: []string
	fmt.Fprintf(os.Stdout, "Response from `SyslogAPI.GetSupportedSyslogProgramNames`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyslogAPI.GetSyslogAuditTags(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogAPI.GetSyslogAuditTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyslogAuditTags`: SyslogAuditTag
	fmt.Fprintf(os.Stdout, "Response from `SyslogAPI.GetSyslogAuditTags`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := int32(56) // int32 | Specifies the id of syslog server.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyslogAPI.GetSyslogServerById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogAPI.GetSyslogServerById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyslogServerById`: SyslogServer
	fmt.Fprintf(os.Stdout, "Response from `SyslogAPI.GetSyslogServerById`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := int32(56) // int32 | Specifies the id of syslog server.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyslogAPI.GetSyslogServerStatusById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogAPI.GetSyslogServerStatusById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyslogServerStatusById`: SyslogServerStatus
	fmt.Fprintf(os.Stdout, "Response from `SyslogAPI.GetSyslogServerStatusById`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyslogAPI.GetSyslogServers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogAPI.GetSyslogServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyslogServers`: SyslogServers
	fmt.Fprintf(os.Stdout, "Response from `SyslogAPI.GetSyslogServers`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := int32(56) // int32 | Specifies the id of syslog server.
	body := *openapiclient.NewSyslogServer() // SyslogServer | Specifies the body of syslog server fields to patch. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyslogAPI.PatchSyslogServerById(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogAPI.PatchSyslogServerById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSyslogServerById`: SyslogServer
	fmt.Fprintf(os.Stdout, "Response from `SyslogAPI.PatchSyslogServerById`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := int32(56) // int32 | Specifies a unique id of the syslog server.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyslogAPI.RemoveSyslogServer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogAPI.RemoveSyslogServer``: %v\n", err)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyslogAPI.RemoveSyslogServers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogAPI.RemoveSyslogServers``: %v\n", err)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewSyslogAuditTag() // SyslogAuditTag | Specifies syslog audit tag to update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyslogAPI.UpdateSyslogAuditTags(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogAPI.UpdateSyslogAuditTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSyslogAuditTags`: SyslogAuditTag
	fmt.Fprintf(os.Stdout, "Response from `SyslogAPI.UpdateSyslogAuditTags`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := int32(56) // int32 | Specifies the id of syslog server.
	body := *openapiclient.NewSyslogServer() // SyslogServer | Specifies the body of syslog server body to update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyslogAPI.UpdateSyslogServerById(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogAPI.UpdateSyslogServerById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSyslogServerById`: SyslogServer
	fmt.Fprintf(os.Stdout, "Response from `SyslogAPI.UpdateSyslogServerById`: %v\n", resp)
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

