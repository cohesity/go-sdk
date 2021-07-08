# \ViewsApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateViewAliases**](ViewsApi.md#ActivateViewAliases) | **Post** /public/viewAliases/{name}/activate | Activates the view aliases of a view from it&#39;s source view.
[**ClearNlmLocks**](ViewsApi.md#ClearNlmLocks) | **Delete** /public/nlmLocks | Clear NLM locks that match the filter criteria specified using parameters.
[**CloneDirectory**](ViewsApi.md#CloneDirectory) | **Post** /public/views/cloneDirectory | Clone a directory of a view.
[**CloneView**](ViewsApi.md#CloneView) | **Post** /public/views/clone | Clone a View.
[**CreateView**](ViewsApi.md#CreateView) | **Post** /public/views | Create a View.
[**CreateViewAlias**](ViewsApi.md#CreateViewAlias) | **Post** /public/viewAliases | Create a View Alias. A View Alias allows a directory inside the view to be mounted without specifying the entire path.
[**CreateViewUserQuota**](ViewsApi.md#CreateViewUserQuota) | **Post** /public/viewUserQuotas | Create a new quota policy for a user in a view.
[**DeleteView**](ViewsApi.md#DeleteView) | **Delete** /public/views/{name} | Delete a View.
[**DeleteViewAlias**](ViewsApi.md#DeleteViewAlias) | **Delete** /public/viewAliases/{name} | Delete a View Alias.
[**DeleteViewById**](ViewsApi.md#DeleteViewById) | **Delete** /public/views/id/{id} | Delete a View.
[**DeleteViewUsersQuota**](ViewsApi.md#DeleteViewUsersQuota) | **Delete** /public/viewUserQuotas | Delete the quota policy overrides for users in a view.
[**GetFileLockStatus**](ViewsApi.md#GetFileLockStatus) | **Get** /public/views/{name}/fileLocks | Fetches the lock status of a file in a view.
[**GetFileLockStatusById**](ViewsApi.md#GetFileLockStatusById) | **Get** /public/views/id/{id}/fileLocks | Fetches the lock status of a file in a view.
[**GetQoSPolicies**](ViewsApi.md#GetQoSPolicies) | **Get** /public/qosPolicies | Fetches QoS Policies of a view.
[**GetSmbConnections**](ViewsApi.md#GetSmbConnections) | **Get** /public/smbConnections | Fetches Smb Connection Status of a view.
[**GetViewById**](ViewsApi.md#GetViewById) | **Get** /public/views/id/{id} | List details about a single View.
[**GetViewByName**](ViewsApi.md#GetViewByName) | **Get** /public/views/{name} | List details about a single View.
[**GetViewDirQuotaInfo**](ViewsApi.md#GetViewDirQuotaInfo) | **Get** /public/viewDirectoryQuotas | Gets directory quota info for a view. Returns error if op fails.
[**GetViewUserQuotas**](ViewsApi.md#GetViewUserQuotas) | **Get** /public/viewUserQuotas | Get the quota policies, usage and summary for a view for all its users. It can also fetch the quota policies, usage and summary for a user in all his views.
[**GetViews**](ViewsApi.md#GetViews) | **Get** /public/views | List Views filtered by some parameters.
[**GetViewsByShareName**](ViewsApi.md#GetViewsByShareName) | **Get** /public/shares | List shares filtered by name.
[**ListNlmLocks**](ViewsApi.md#ListNlmLocks) | **Get** /public/nlmLocks | List the NLM locks that match the filter criteria specified using parameters.
[**LockFile**](ViewsApi.md#LockFile) | **Post** /public/views/{name}/fileLocks | Lock a file in a view and optionally update its expiry timestamp.
[**LockFileById**](ViewsApi.md#LockFileById) | **Post** /public/views/id/{id}/fileLocks | Lock a file in a view and optionally update its expiry timestamp.
[**OverwriteView**](ViewsApi.md#OverwriteView) | **Post** /public/views/overwrite | Overwrites a Target view with contents of a Source view.
[**RenameView**](ViewsApi.md#RenameView) | **Post** /public/views/rename/{name} | Rename a View.
[**RenameViewById**](ViewsApi.md#RenameViewById) | **Post** /public/views/rename/id/{id} | Rename a View.
[**UpdateUserQuotaSettings**](ViewsApi.md#UpdateUserQuotaSettings) | **Put** /public/viewUserQuotasSettings | Update the user quota settings in a view.
[**UpdateView**](ViewsApi.md#UpdateView) | **Put** /public/views | Update a View.
[**UpdateViewAlias**](ViewsApi.md#UpdateViewAlias) | **Put** /public/viewAliases | Update a View Alias. A View Alias allows a directory inside the view to be mounted without specifying the entire path.
[**UpdateViewByName**](ViewsApi.md#UpdateViewByName) | **Put** /public/views/{name} | Update a View.
[**UpdateViewDirQuota**](ViewsApi.md#UpdateViewDirQuota) | **Put** /public/viewDirectoryQuotas | Updates a directory quota policy for a view. Returns error if op fails.
[**UpdateViewUserQuota**](ViewsApi.md#UpdateViewUserQuota) | **Put** /public/viewUserQuotas | Update a new quota policy for a user in a view.



## ActivateViewAliases

> ActivateViewAliasesResult ActivateViewAliases(ctx, name).Execute()

Activates the view aliases of a view from it's source view.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    name := "name_example" // string | Specifies the View name.

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

    request := cohesitysdk.ApiActivateViewAliasesRequest{
        Name: &name
    }

    resp, r, err := api_client.ViewsApi.ActivateViewAliases(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.ActivateViewAliases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateViewAliases`: ActivateViewAliasesResult
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.ActivateViewAliases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the View name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateViewAliasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActivateViewAliasesResult**](ActivateViewAliasesResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearNlmLocks

> ClearNlmLocks(ctx).Body(body).Execute()

Clear NLM locks that match the filter criteria specified using parameters.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewClearNlmLocksParameters() // ClearNlmLocksParameters | Request to clear NLM locks.

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

    request := cohesitysdk.ApiClearNlmLocksRequest{
        Body: &body
    }

    resp, r, err := api_client.ViewsApi.ClearNlmLocks(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.ClearNlmLocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClearNlmLocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClearNlmLocksParameters**](ClearNlmLocksParameters.md) | Request to clear NLM locks. | 

### Return type

 (empty response body)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneDirectory

> CloneDirectory(ctx).Body(body).Execute()

Clone a directory of a view.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewCloneDirectoryParams() // CloneDirectoryParams | Request to clone a directory.

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

    request := cohesitysdk.ApiCloneDirectoryRequest{
        Body: &body
    }

    resp, r, err := api_client.ViewsApi.CloneDirectory(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.CloneDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloneDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CloneDirectoryParams**](CloneDirectoryParams.md) | Request to clone a directory. | 

### Return type

 (empty response body)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneView

> View CloneView(ctx).Body(body).Execute()

Clone a View.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewCloneViewRequest() // CloneViewRequest | Request to clone a View.

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

    request := cohesitysdk.ApiCloneViewRequest{
        Body: &body
    }

    resp, r, err := api_client.ViewsApi.CloneView(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.CloneView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneView`: View
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.CloneView`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloneViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CloneViewRequest**](CloneViewRequest.md) | Request to clone a View. | 

### Return type

[**View**](View.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateView

> View CreateView(ctx).Body(body).Execute()

Create a View.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewCreateViewRequest("Name_example", NullableInt64(123)) // CreateViewRequest | Request to create a View.

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

    request := cohesitysdk.ApiCreateViewRequest{
        Body: &body
    }

    resp, r, err := api_client.ViewsApi.CreateView(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.CreateView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateView`: View
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.CreateView`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateViewRequest**](CreateViewRequest.md) | Request to create a View. | 

### Return type

[**View**](View.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateViewAlias

> ViewAlias CreateViewAlias(ctx).Body(body).Execute()

Create a View Alias. A View Alias allows a directory inside the view to be mounted without specifying the entire path.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewViewAlias() // ViewAlias | Request to create a View Alias.

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

    request := cohesitysdk.ApiCreateViewAliasRequest{
        Body: &body
    }

    resp, r, err := api_client.ViewsApi.CreateViewAlias(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.CreateViewAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateViewAlias`: ViewAlias
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.CreateViewAlias`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateViewAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ViewAlias**](ViewAlias.md) | Request to create a View Alias. | 

### Return type

[**ViewAlias**](ViewAlias.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateViewUserQuota

> UserQuotaAndUsage CreateViewUserQuota(ctx).Body(body).Execute()

Create a new quota policy for a user in a view.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewViewUserQuotaParameters() // ViewUserQuotaParameters | update user quota params. (optional)

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

    request := cohesitysdk.ApiCreateViewUserQuotaRequest{
        Body: &body
    }

    resp, r, err := api_client.ViewsApi.CreateViewUserQuota(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.CreateViewUserQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateViewUserQuota`: UserQuotaAndUsage
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.CreateViewUserQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateViewUserQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ViewUserQuotaParameters**](ViewUserQuotaParameters.md) | update user quota params. | 

### Return type

[**UserQuotaAndUsage**](UserQuotaAndUsage.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteView

> DeleteView(ctx, name).Execute()

Delete a View.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    name := "name_example" // string | Specifies the View name.

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

    request := cohesitysdk.ApiDeleteViewRequest{
        Name: &name
    }

    resp, r, err := api_client.ViewsApi.DeleteView(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.DeleteView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the View name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteViewAlias

> DeleteViewAlias(ctx, name).Execute()

Delete a View Alias.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    name := "name_example" // string | Specifies the View Alias name.

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

    request := cohesitysdk.ApiDeleteViewAliasRequest{
        Name: &name
    }

    resp, r, err := api_client.ViewsApi.DeleteViewAlias(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.DeleteViewAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the View Alias name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteViewAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteViewById

> DeleteViewById(ctx, id).Execute()

Delete a View.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    id := int64(789) // int64 | Specifies the View id.

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

    request := cohesitysdk.ApiDeleteViewByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.ViewsApi.DeleteViewById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.DeleteViewById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the View id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteViewByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteViewUsersQuota

> DeleteViewUsersQuota(ctx).Body(body).Execute()

Delete the quota policy overrides for users in a view.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewDeleteViewUsersQuotaParameters() // DeleteViewUsersQuotaParameters | update user quota params. (optional)

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

    request := cohesitysdk.ApiDeleteViewUsersQuotaRequest{
        Body: &body
    }

    resp, r, err := api_client.ViewsApi.DeleteViewUsersQuota(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.DeleteViewUsersQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteViewUsersQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteViewUsersQuotaParameters**](DeleteViewUsersQuotaParameters.md) | update user quota params. | 

### Return type

 (empty response body)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileLockStatus

> FileLockStatus GetFileLockStatus(ctx, name).Path(path).Execute()

Fetches the lock status of a file in a view.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    name := "name_example" // string | Specifies the View name.
    path := "path_example" // string | Specifies the file path that needs to be locked. (optional)

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

    request := cohesitysdk.ApiGetFileLockStatusRequest{
        Name: &name
        Path: &path
    }

    resp, r, err := api_client.ViewsApi.GetFileLockStatus(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.GetFileLockStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFileLockStatus`: FileLockStatus
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.GetFileLockStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the View name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileLockStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | Specifies the file path that needs to be locked. | 

### Return type

[**FileLockStatus**](FileLockStatus.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileLockStatusById

> FileLockStatus GetFileLockStatusById(ctx, id).Path(path).Execute()

Fetches the lock status of a file in a view.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    id := int64(789) // int64 | Specifies the View id.
    path := "path_example" // string | Specifies the file path that needs to be locked. (optional)

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

    request := cohesitysdk.ApiGetFileLockStatusByIdRequest{
        Id: &id
        Path: &path
    }

    resp, r, err := api_client.ViewsApi.GetFileLockStatusById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.GetFileLockStatusById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFileLockStatusById`: FileLockStatus
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.GetFileLockStatusById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the View id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileLockStatusByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | Specifies the file path that needs to be locked. | 

### Return type

[**FileLockStatus**](FileLockStatus.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQoSPolicies

> []QoSPolicy GetQoSPolicies(ctx).Ids(ids).Names(names).Execute()

Fetches QoS Policies of a view.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    ids := []int64{int64(123)} // []int64 | Specifies the Ids of QoS Policies to filter by. (optional)
    names := []string{"Inner_example"} // []string | Specifies the names of QoS Policies to filter by. (optional)

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

    request := cohesitysdk.ApiGetQoSPoliciesRequest{
        Ids: &ids
        Names: &names
    }

    resp, r, err := api_client.ViewsApi.GetQoSPolicies(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.GetQoSPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQoSPolicies`: []QoSPolicy
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.GetQoSPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQoSPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | Specifies the Ids of QoS Policies to filter by. | 
 **names** | **[]string** | Specifies the names of QoS Policies to filter by. | 

### Return type

[**[]QoSPolicy**](QoSPolicy.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmbConnections

> []SmbConnection GetSmbConnections(ctx).ViewNames(viewNames).ViewIds(viewIds).MaxCount(maxCount).IncludeSid(includeSid).Execute()

Fetches Smb Connection Status of a view.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    viewNames := []string{"Inner_example"} // []string | Parameters to get Smb Connections. List of names of views whose connections are to be fetched. (optional)
    viewIds := []int64{int64(123)} // []int64 | List of ids of views whose connections are to be fetched. (optional)
    maxCount := int32(56) // int32 | Maximum number of results to return. (optional)
    includeSid := true // bool | Specifies whether to include list of sids in the result. (optional)

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

    request := cohesitysdk.ApiGetSmbConnectionsRequest{
        ViewNames: &viewNames
        ViewIds: &viewIds
        MaxCount: &maxCount
        IncludeSid: &includeSid
    }

    resp, r, err := api_client.ViewsApi.GetSmbConnections(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.GetSmbConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmbConnections`: []SmbConnection
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.GetSmbConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSmbConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewNames** | **[]string** | Parameters to get Smb Connections. List of names of views whose connections are to be fetched. | 
 **viewIds** | **[]int64** | List of ids of views whose connections are to be fetched. | 
 **maxCount** | **int32** | Maximum number of results to return. | 
 **includeSid** | **bool** | Specifies whether to include list of sids in the result. | 

### Return type

[**[]SmbConnection**](SmbConnection.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewById

> View GetViewById(ctx, id).Execute()

List details about a single View.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    id := int64(789) // int64 | Specifies the View id.

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

    request := cohesitysdk.ApiGetViewByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.ViewsApi.GetViewById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.GetViewById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetViewById`: View
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.GetViewById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the View id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetViewByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**View**](View.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewByName

> View GetViewByName(ctx, name).Execute()

List details about a single View.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    name := "name_example" // string | Specifies the View name.

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

    request := cohesitysdk.ApiGetViewByNameRequest{
        Name: &name
    }

    resp, r, err := api_client.ViewsApi.GetViewByName(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.GetViewByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetViewByName`: View
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.GetViewByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the View name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetViewByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**View**](View.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewDirQuotaInfo

> DirQuotaInfo GetViewDirQuotaInfo(ctx).ViewName(viewName).Execute()

Gets directory quota info for a view. Returns error if op fails.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    viewName := "viewName_example" // string | The name of the view.

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

    request := cohesitysdk.ApiGetViewDirQuotaInfoRequest{
        ViewName: &viewName
    }

    resp, r, err := api_client.ViewsApi.GetViewDirQuotaInfo(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.GetViewDirQuotaInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetViewDirQuotaInfo`: DirQuotaInfo
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.GetViewDirQuotaInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetViewDirQuotaInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewName** | **string** | The name of the view. | 

### Return type

[**DirQuotaInfo**](DirQuotaInfo.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewUserQuotas

> ViewUserQuotas GetViewUserQuotas(ctx).ViewName(viewName).IncludeUsage(includeUsage).IncludeUserWithQuotaOnly(includeUserWithQuotaOnly).ExcludeUsersWithinAlertThreshold(excludeUsersWithinAlertThreshold).UnixUid(unixUid).Sid(sid).UserUnixIdsForView(userUnixIdsForView).UserSidsForView(userSidsForView).ViewNamesForUser(viewNamesForUser).SummaryOnly(summaryOnly).PageCount(pageCount).MaxViewId(maxViewId).Cookie(cookie).OutputFormat(outputFormat).Execute()

Get the quota policies, usage and summary for a view for all its users. It can also fetch the quota policies, usage and summary for a user in all his views.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    viewName := "viewName_example" // string | Specifies the name of the input view. If given, there could be three scenarios with the viewName input parameter: It gives the user quota overrides for this view, and the user quota settings. Returns 'usersQuotaAndUsage'. If given along with the user id, it returns the quota policy for this user on this view. Returns 'usersQuotaAndUsage'. If given along with SummaryOnly as true, a user quota summary for this view would be returned. Returns 'summaryForView'. If not given, then the user id is checked. (optional)
    includeUsage := true // bool | If set to true, the logical usage info is included only for users with quota overrides. By default, it is set to false. (optional)
    includeUserWithQuotaOnly := true // bool | If set to true, the result will only contain user with user quota enabled. By default, this field is set to false, and it's only in effect when 'SummaryOnly' is set to false and 'ViewName' is specified. (optional)
    excludeUsersWithinAlertThreshold := true // bool | This field can be set only when includeUsage is set to true. By default, all the users with logical usage > 0 will be returned in the result. If this field is set to true, only the list of users who has exceeded the alert threshold will be returned. (optional)
    unixUid := int32(56) // int32 | If interested in a user via unix-identifier, include UnixUid. Otherwise, If a valid unix-id to SID mappings are available (i.e., when mixed mode is enabled) the server will perform the necessary id mapping and return the correct usage irrespective of whether the unix id / SID is provided. (optional)
    sid := "sid_example" // string | If interested in a user via smb_client, include SID. Otherwise, If a valid unix-id to SID mappings are available (i.e., when mixed mode is enabled) the server will perform the necessary id mapping and return the correct usage irrespective of whether the unix id / SID is provided. The string is of following format - S-1-IdentifierAuthority-SubAuthority1-SubAuthority2-...-SubAuthorityn. (optional)
    userUnixIdsForView := []int32{int32(123)} // []int32 | While making a query for a view, this specifies a list of specific users with their unix uid for the result. (optional)
    userSidsForView := []string{"Inner_example"} // []string | While making a query for a view, this specifies a list of specific users with their Sid for the result. (optional)
    viewNamesForUser := []string{"Inner_example"} // []string | While making a query for a user, this specifies a list of specific views for the result. (optional)
    summaryOnly := true // bool | Specifies a flag to just return a summary. If set to true, and if ViewName is not nil, it returns the summary of users for a view. Otherwise if UserId not nil, and ViewName is nil then it fetches the summary for a user in his views.  By default, it is set to false. (optional)
    pageCount := int64(789) // int64 | Specifies the max entries that should be returned in the result. (optional)
    maxViewId := int64(789) // int64 | Related to fetching a particular user's quota and usage in all his views. It only pertains to the scenario where either UnixUid or Sid is specified, and ViewName is nil. Specify the maxViewId for All the views returned would have view_id's less than or equal to the given MaxViewId if it is >= 0. (optional)
    cookie := "cookie_example" // string | Cookie should be used from previous call to list user quota overrides. It resumes (or gives the next set of values) from the result of the previous call. (optional)
    outputFormat := "outputFormat_example" // string | OutputFormat is the Output format for the output. If it is not specified, default is json. (optional)

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

    request := cohesitysdk.ApiGetViewUserQuotasRequest{
        ViewName: &viewName
        IncludeUsage: &includeUsage
        IncludeUserWithQuotaOnly: &includeUserWithQuotaOnly
        ExcludeUsersWithinAlertThreshold: &excludeUsersWithinAlertThreshold
        UnixUid: &unixUid
        Sid: &sid
        UserUnixIdsForView: &userUnixIdsForView
        UserSidsForView: &userSidsForView
        ViewNamesForUser: &viewNamesForUser
        SummaryOnly: &summaryOnly
        PageCount: &pageCount
        MaxViewId: &maxViewId
        Cookie: &cookie
        OutputFormat: &outputFormat
    }

    resp, r, err := api_client.ViewsApi.GetViewUserQuotas(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.GetViewUserQuotas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetViewUserQuotas`: ViewUserQuotas
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.GetViewUserQuotas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetViewUserQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewName** | **string** | Specifies the name of the input view. If given, there could be three scenarios with the viewName input parameter: It gives the user quota overrides for this view, and the user quota settings. Returns &#39;usersQuotaAndUsage&#39;. If given along with the user id, it returns the quota policy for this user on this view. Returns &#39;usersQuotaAndUsage&#39;. If given along with SummaryOnly as true, a user quota summary for this view would be returned. Returns &#39;summaryForView&#39;. If not given, then the user id is checked. | 
 **includeUsage** | **bool** | If set to true, the logical usage info is included only for users with quota overrides. By default, it is set to false. | 
 **includeUserWithQuotaOnly** | **bool** | If set to true, the result will only contain user with user quota enabled. By default, this field is set to false, and it&#39;s only in effect when &#39;SummaryOnly&#39; is set to false and &#39;ViewName&#39; is specified. | 
 **excludeUsersWithinAlertThreshold** | **bool** | This field can be set only when includeUsage is set to true. By default, all the users with logical usage &gt; 0 will be returned in the result. If this field is set to true, only the list of users who has exceeded the alert threshold will be returned. | 
 **unixUid** | **int32** | If interested in a user via unix-identifier, include UnixUid. Otherwise, If a valid unix-id to SID mappings are available (i.e., when mixed mode is enabled) the server will perform the necessary id mapping and return the correct usage irrespective of whether the unix id / SID is provided. | 
 **sid** | **string** | If interested in a user via smb_client, include SID. Otherwise, If a valid unix-id to SID mappings are available (i.e., when mixed mode is enabled) the server will perform the necessary id mapping and return the correct usage irrespective of whether the unix id / SID is provided. The string is of following format - S-1-IdentifierAuthority-SubAuthority1-SubAuthority2-...-SubAuthorityn. | 
 **userUnixIdsForView** | **[]int32** | While making a query for a view, this specifies a list of specific users with their unix uid for the result. | 
 **userSidsForView** | **[]string** | While making a query for a view, this specifies a list of specific users with their Sid for the result. | 
 **viewNamesForUser** | **[]string** | While making a query for a user, this specifies a list of specific views for the result. | 
 **summaryOnly** | **bool** | Specifies a flag to just return a summary. If set to true, and if ViewName is not nil, it returns the summary of users for a view. Otherwise if UserId not nil, and ViewName is nil then it fetches the summary for a user in his views.  By default, it is set to false. | 
 **pageCount** | **int64** | Specifies the max entries that should be returned in the result. | 
 **maxViewId** | **int64** | Related to fetching a particular user&#39;s quota and usage in all his views. It only pertains to the scenario where either UnixUid or Sid is specified, and ViewName is nil. Specify the maxViewId for All the views returned would have view_id&#39;s less than or equal to the given MaxViewId if it is &gt;&#x3D; 0. | 
 **cookie** | **string** | Cookie should be used from previous call to list user quota overrides. It resumes (or gives the next set of values) from the result of the previous call. | 
 **outputFormat** | **string** | OutputFormat is the Output format for the output. If it is not specified, default is json. | 

### Return type

[**ViewUserQuotas**](ViewUserQuotas.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViews

> GetViewsResult GetViews(ctx).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).ViewNames(viewNames).ViewIds(viewIds).ViewBoxIds(viewBoxIds).ViewBoxNames(viewBoxNames).MatchPartialNames(matchPartialNames).MaxCount(maxCount).IncludeProtectionJobs(includeProtectionJobs).MaxViewId(maxViewId).IncludeInactive(includeInactive).JobIds(jobIds).SortByLogicalUsage(sortByLogicalUsage).MatchAliasNames(matchAliasNames).IncludeViewsWithAntivirusEnabledOnly(includeViewsWithAntivirusEnabledOnly).IncludeStats(includeStats).IncludeViewsWithDataLockEnabledOnly(includeViewsWithDataLockEnabledOnly).Execute()

List Views filtered by some parameters.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    allUnderHierarchy := true // bool | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)
    viewNames := []string{"Inner_example"} // []string | Filter by a list of View names. (optional)
    viewIds := []int64{int64(123)} // []int64 | Filter by a list of View ids. (optional)
    viewBoxIds := []int64{int64(123)} // []int64 | Filter by a list of Storage Domains (View Boxes) specified by id. (optional)
    viewBoxNames := []string{"Inner_example"} // []string | Filter by a list of View Box names. (optional)
    matchPartialNames := true // bool | If true, the names in viewNames are matched by prefix rather than exactly matched. (optional)
    maxCount := int32(56) // int32 | Specifies a limit on the number of Views returned. (optional)
    includeProtectionJobs := true // bool | Specifies if Protection Jobs information needs to be returned along with view metadata. By default, if not set or set to true, Job information is returned. (optional)
    maxViewId := int64(789) // int64 | If the number of Views to return exceeds the maxCount specified in the original request, specify the id of the last View from the viewList in the previous response to get the next set of Views. (optional)
    includeInactive := true // bool | Specifies if inactive Views on this Remote Cluster (which have Snapshots copied by replication) should also be returned. Inactive Views are not counted towards the maxCount. By default, this field is set to false. (optional)
    jobIds := []int64{int64(123)} // []int64 | Filter by Protection Job ids. Return Views that are being protected by listed Jobs, which are specified by ids. (optional)
    sortByLogicalUsage := true // bool | If set to true, the list is sorted descending by logical usage. (optional)
    matchAliasNames := true // bool | If true, view aliases are also matched with the names in viewNames. (optional)
    includeViewsWithAntivirusEnabledOnly := true // bool | If set to true, the list will contain only the views for which antivirus scan is enabled. (optional)
    includeStats := true // bool | If set to true, stats of views will be returned. By default this parameter is set to false. (optional)
    includeViewsWithDataLockEnabledOnly := true // bool | If set to true, the list will contain only the views for which either file level data lock is enabled or view level data lock is enabled. (optional)

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

    request := cohesitysdk.ApiGetViewsRequest{
        TenantIds: &tenantIds
        AllUnderHierarchy: &allUnderHierarchy
        ViewNames: &viewNames
        ViewIds: &viewIds
        ViewBoxIds: &viewBoxIds
        ViewBoxNames: &viewBoxNames
        MatchPartialNames: &matchPartialNames
        MaxCount: &maxCount
        IncludeProtectionJobs: &includeProtectionJobs
        MaxViewId: &maxViewId
        IncludeInactive: &includeInactive
        JobIds: &jobIds
        SortByLogicalUsage: &sortByLogicalUsage
        MatchAliasNames: &matchAliasNames
        IncludeViewsWithAntivirusEnabledOnly: &includeViewsWithAntivirusEnabledOnly
        IncludeStats: &includeStats
        IncludeViewsWithDataLockEnabledOnly: &includeViewsWithDataLockEnabledOnly
    }

    resp, r, err := api_client.ViewsApi.GetViews(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.GetViews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetViews`: GetViewsResult
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.GetViews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 
 **viewNames** | **[]string** | Filter by a list of View names. | 
 **viewIds** | **[]int64** | Filter by a list of View ids. | 
 **viewBoxIds** | **[]int64** | Filter by a list of Storage Domains (View Boxes) specified by id. | 
 **viewBoxNames** | **[]string** | Filter by a list of View Box names. | 
 **matchPartialNames** | **bool** | If true, the names in viewNames are matched by prefix rather than exactly matched. | 
 **maxCount** | **int32** | Specifies a limit on the number of Views returned. | 
 **includeProtectionJobs** | **bool** | Specifies if Protection Jobs information needs to be returned along with view metadata. By default, if not set or set to true, Job information is returned. | 
 **maxViewId** | **int64** | If the number of Views to return exceeds the maxCount specified in the original request, specify the id of the last View from the viewList in the previous response to get the next set of Views. | 
 **includeInactive** | **bool** | Specifies if inactive Views on this Remote Cluster (which have Snapshots copied by replication) should also be returned. Inactive Views are not counted towards the maxCount. By default, this field is set to false. | 
 **jobIds** | **[]int64** | Filter by Protection Job ids. Return Views that are being protected by listed Jobs, which are specified by ids. | 
 **sortByLogicalUsage** | **bool** | If set to true, the list is sorted descending by logical usage. | 
 **matchAliasNames** | **bool** | If true, view aliases are also matched with the names in viewNames. | 
 **includeViewsWithAntivirusEnabledOnly** | **bool** | If set to true, the list will contain only the views for which antivirus scan is enabled. | 
 **includeStats** | **bool** | If set to true, stats of views will be returned. By default this parameter is set to false. | 
 **includeViewsWithDataLockEnabledOnly** | **bool** | If set to true, the list will contain only the views for which either file level data lock is enabled or view level data lock is enabled. | 

### Return type

[**GetViewsResult**](GetViewsResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewsByShareName

> GetViewsByShareNameResult GetViewsByShareName(ctx).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).ShareName(shareName).MaxCount(maxCount).PaginationCookie(paginationCookie).MatchPartialNames(matchPartialNames).Execute()

List shares filtered by name.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    allUnderHierarchy := true // bool | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)
    shareName := "shareName_example" // string | The share name(substring) that needs to be searched against existing views and aliases. (optional)
    maxCount := int32(56) // int32 | Specifies a limit on the number of Views returned. (optional)
    paginationCookie := "paginationCookie_example" // string | Expected to be empty in the first call to GetViewsByShareName. To get the next set of results, set this value to the pagination cookie value returned  in the response of the previous call. (optional)
    matchPartialNames := true // bool | If true, the names in viewNames are matched by prefix rather than exactly matched. (optional)

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

    request := cohesitysdk.ApiGetViewsByShareNameRequest{
        TenantIds: &tenantIds
        AllUnderHierarchy: &allUnderHierarchy
        ShareName: &shareName
        MaxCount: &maxCount
        PaginationCookie: &paginationCookie
        MatchPartialNames: &matchPartialNames
    }

    resp, r, err := api_client.ViewsApi.GetViewsByShareName(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.GetViewsByShareName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetViewsByShareName`: GetViewsByShareNameResult
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.GetViewsByShareName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetViewsByShareNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 
 **shareName** | **string** | The share name(substring) that needs to be searched against existing views and aliases. | 
 **maxCount** | **int32** | Specifies a limit on the number of Views returned. | 
 **paginationCookie** | **string** | Expected to be empty in the first call to GetViewsByShareName. To get the next set of results, set this value to the pagination cookie value returned  in the response of the previous call. | 
 **matchPartialNames** | **bool** | If true, the names in viewNames are matched by prefix rather than exactly matched. | 

### Return type

[**GetViewsByShareNameResult**](GetViewsByShareNameResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNlmLocks

> ListNlmLocksResponse ListNlmLocks(ctx).FilePath(filePath).ViewName(viewName).PageCount(pageCount).Cookie(cookie).Execute()

List the NLM locks that match the filter criteria specified using parameters.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    filePath := "filePath_example" // string | Specifies the filepath in the view relative to the root filesystem. If this field is specified, viewName field must also be specified. (optional)
    viewName := "viewName_example" // string | Specifies the name of the View in which to search. If a view name is not specified, all the views in the Cluster is searched. This field is mandatory if filePath field is specified. (optional)
    pageCount := int32(56) // int32 | Specifies the maximum number of NLM locks to return in the response. This field cannot be set above 1000. If this is not set, maximum of 1000 entries are returned. (optional)
    cookie := "cookie_example" // string | Specifies the opaque string returned in the previous response. If this is set, next set of active opens just after the previous response are returned. If this is not set, first set of NLM locks are returned. (optional)

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

    request := cohesitysdk.ApiListNlmLocksRequest{
        FilePath: &filePath
        ViewName: &viewName
        PageCount: &pageCount
        Cookie: &cookie
    }

    resp, r, err := api_client.ViewsApi.ListNlmLocks(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.ListNlmLocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNlmLocks`: ListNlmLocksResponse
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.ListNlmLocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNlmLocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filePath** | **string** | Specifies the filepath in the view relative to the root filesystem. If this field is specified, viewName field must also be specified. | 
 **viewName** | **string** | Specifies the name of the View in which to search. If a view name is not specified, all the views in the Cluster is searched. This field is mandatory if filePath field is specified. | 
 **pageCount** | **int32** | Specifies the maximum number of NLM locks to return in the response. This field cannot be set above 1000. If this is not set, maximum of 1000 entries are returned. | 
 **cookie** | **string** | Specifies the opaque string returned in the previous response. If this is set, next set of active opens just after the previous response are returned. If this is not set, first set of NLM locks are returned. | 

### Return type

[**ListNlmLocksResponse**](ListNlmLocksResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockFile

> FileLockStatus LockFile(ctx, name).Body(body).Execute()

Lock a file in a view and optionally update its expiry timestamp.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    name := "name_example" // string | Specifies the View name.
    body := *openapiclient.NewLockFileParams() // LockFileParams | Request to lock a file. (optional)

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

    request := cohesitysdk.ApiLockFileRequest{
        Name: &name
        Body: &body
    }

    resp, r, err := api_client.ViewsApi.LockFile(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.LockFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LockFile`: FileLockStatus
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.LockFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the View name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**LockFileParams**](LockFileParams.md) | Request to lock a file. | 

### Return type

[**FileLockStatus**](FileLockStatus.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockFileById

> FileLockStatus LockFileById(ctx, id).Body(body).Execute()

Lock a file in a view and optionally update its expiry timestamp.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    id := int64(789) // int64 | Specifies the View id.
    body := *openapiclient.NewLockFileParams() // LockFileParams | Request to lock a file. (optional)

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

    request := cohesitysdk.ApiLockFileByIdRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.ViewsApi.LockFileById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.LockFileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LockFileById`: FileLockStatus
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.LockFileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the View id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockFileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**LockFileParams**](LockFileParams.md) | Request to lock a file. | 

### Return type

[**FileLockStatus**](FileLockStatus.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverwriteView

> View OverwriteView(ctx).Body(body).Execute()

Overwrites a Target view with contents of a Source view.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewOverwriteViewParam("SourceViewName_example", "TargetViewName_example") // OverwriteViewParam | Request to overwrite a Target view with contents of a Source view.

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

    request := cohesitysdk.ApiOverwriteViewRequest{
        Body: &body
    }

    resp, r, err := api_client.ViewsApi.OverwriteView(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.OverwriteView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OverwriteView`: View
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.OverwriteView`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOverwriteViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OverwriteViewParam**](OverwriteViewParam.md) | Request to overwrite a Target view with contents of a Source view. | 

### Return type

[**View**](View.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameView

> View RenameView(ctx, name).Body(body).Execute()

Rename a View.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    name := "name_example" // string | Specifies the View name.
    body := *openapiclient.NewRenameViewParam("NewViewName_example") // RenameViewParam | Request to rename a View.

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

    request := cohesitysdk.ApiRenameViewRequest{
        Name: &name
        Body: &body
    }

    resp, r, err := api_client.ViewsApi.RenameView(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.RenameView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenameView`: View
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.RenameView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the View name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RenameViewParam**](RenameViewParam.md) | Request to rename a View. | 

### Return type

[**View**](View.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameViewById

> View RenameViewById(ctx, id).Body(body).Execute()

Rename a View.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    id := int64(789) // int64 | Specifies the View id.
    body := *openapiclient.NewRenameViewParam("NewViewName_example") // RenameViewParam | Request to rename a View.

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

    request := cohesitysdk.ApiRenameViewByIdRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.ViewsApi.RenameViewById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.RenameViewById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenameViewById`: View
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.RenameViewById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the View id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameViewByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RenameViewParam**](RenameViewParam.md) | Request to rename a View. | 

### Return type

[**View**](View.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserQuotaSettings

> UserQuotaSettings UpdateUserQuotaSettings(ctx).Body(body).Execute()

Update the user quota settings in a view.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewUpdateUserQuotaSettingsForView() // UpdateUserQuotaSettingsForView | update user quota metadata params. (optional)

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

    request := cohesitysdk.ApiUpdateUserQuotaSettingsRequest{
        Body: &body
    }

    resp, r, err := api_client.ViewsApi.UpdateUserQuotaSettings(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.UpdateUserQuotaSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserQuotaSettings`: UserQuotaSettings
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.UpdateUserQuotaSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserQuotaSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateUserQuotaSettingsForView**](UpdateUserQuotaSettingsForView.md) | update user quota metadata params. | 

### Return type

[**UserQuotaSettings**](UserQuotaSettings.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateView

> View UpdateView(ctx).Body(body).Execute()

Update a View.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewUpdateViewParam() // UpdateViewParam | Request to update a view.

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

    request := cohesitysdk.ApiUpdateViewRequest{
        Body: &body
    }

    resp, r, err := api_client.ViewsApi.UpdateView(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.UpdateView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateView`: View
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.UpdateView`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateViewParam**](UpdateViewParam.md) | Request to update a view. | 

### Return type

[**View**](View.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateViewAlias

> ViewAlias UpdateViewAlias(ctx).Body(body).Execute()

Update a View Alias. A View Alias allows a directory inside the view to be mounted without specifying the entire path.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewUpdateViewAliasParam() // UpdateViewAliasParam | Request to update a View.

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

    request := cohesitysdk.ApiUpdateViewAliasRequest{
        Body: &body
    }

    resp, r, err := api_client.ViewsApi.UpdateViewAlias(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.UpdateViewAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateViewAlias`: ViewAlias
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.UpdateViewAlias`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateViewAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateViewAliasParam**](UpdateViewAliasParam.md) | Request to update a View. | 

### Return type

[**ViewAlias**](ViewAlias.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateViewByName

> View UpdateViewByName(ctx, name).Body(body).Execute()

Update a View.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    name := "name_example" // string | Specifies the View name.
    body := *openapiclient.NewUpdateViewParam() // UpdateViewParam | Request to update a view.

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

    request := cohesitysdk.ApiUpdateViewByNameRequest{
        Name: &name
        Body: &body
    }

    resp, r, err := api_client.ViewsApi.UpdateViewByName(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.UpdateViewByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateViewByName`: View
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.UpdateViewByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the View name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateViewByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateViewParam**](UpdateViewParam.md) | Request to update a view. | 

### Return type

[**View**](View.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateViewDirQuota

> DirQuotaInfo UpdateViewDirQuota(ctx).Body(body).Execute()

Updates a directory quota policy for a view. Returns error if op fails.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewUpdateDirQuotaArgs() // UpdateDirQuotaArgs | Request to update view directory quota policy.

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

    request := cohesitysdk.ApiUpdateViewDirQuotaRequest{
        Body: &body
    }

    resp, r, err := api_client.ViewsApi.UpdateViewDirQuota(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.UpdateViewDirQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateViewDirQuota`: DirQuotaInfo
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.UpdateViewDirQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateViewDirQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateDirQuotaArgs**](UpdateDirQuotaArgs.md) | Request to update view directory quota policy. | 

### Return type

[**DirQuotaInfo**](DirQuotaInfo.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateViewUserQuota

> UserQuotaAndUsage UpdateViewUserQuota(ctx).Body(body).Execute()

Update a new quota policy for a user in a view.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cohesitysdk "cohesitysdk"
)

func main() {
    body := *openapiclient.NewViewUserQuotaParameters() // ViewUserQuotaParameters | update user quota params. (optional)

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

    request := cohesitysdk.ApiUpdateViewUserQuotaRequest{
        Body: &body
    }

    resp, r, err := api_client.ViewsApi.UpdateViewUserQuota(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.UpdateViewUserQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateViewUserQuota`: UserQuotaAndUsage
    fmt.Fprintf(os.Stdout, "Response from `ViewsApi.UpdateViewUserQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateViewUserQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ViewUserQuotaParameters**](ViewUserQuotaParameters.md) | update user quota params. | 

### Return type

[**UserQuotaAndUsage**](UserQuotaAndUsage.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

