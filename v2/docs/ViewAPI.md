# \ViewAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddViewUserQuotaOverrides**](ViewAPI.md#AddViewUserQuotaOverrides) | **Post** /file-services/views/{viewId}/user-quotas | Add User Quota overrides.
[**ClearNlmLocks**](ViewAPI.md#ClearNlmLocks) | **Delete** /file-services/nlm-locks | Clear NLM locks.
[**CloneView**](ViewAPI.md#CloneView) | **Post** /file-services/views/{id}/clone | Clone View.
[**CloneViewDirectory**](ViewAPI.md#CloneViewDirectory) | **Post** /file-services/views/clone-directory | Clone View Directory.
[**CloseSmbFileOpen**](ViewAPI.md#CloseSmbFileOpen) | **Delete** /file-services/smb-file-opens | Close SMB File open.
[**CreateShare**](ViewAPI.md#CreateShare) | **Post** /file-services/shares | Create a Share.
[**CreateView**](ViewAPI.md#CreateView) | **Post** /file-services/views | Create a View
[**CreateViewTemplate**](ViewAPI.md#CreateViewTemplate) | **Post** /file-services/view-template | Create a View Template
[**DeleteShare**](ViewAPI.md#DeleteShare) | **Delete** /file-services/shares/{name} | Delete a Share.
[**DeleteView**](ViewAPI.md#DeleteView) | **Delete** /file-services/views/{id} | Delete a View
[**DeleteViewDirectoryQuota**](ViewAPI.md#DeleteViewDirectoryQuota) | **Delete** /file-services/views/{id}/directory-quotas | Delete directory quota for the View.
[**DeleteViewTemplate**](ViewAPI.md#DeleteViewTemplate) | **Delete** /file-services/view-template/{id} | Delete a View Template
[**DeleteViewUserQuotaOverrides**](ViewAPI.md#DeleteViewUserQuotaOverrides) | **Delete** /file-services/views/{viewId}/user-quotas | Delete user quota overrides.
[**GetFileLockStatus**](ViewAPI.md#GetFileLockStatus) | **Get** /file-services/views/{id}/file-lock | Get file lock status
[**GetNlmLocks**](ViewAPI.md#GetNlmLocks) | **Get** /file-services/nlm-locks | Get NLM locks.
[**GetQosPolicies**](ViewAPI.md#GetQosPolicies) | **Get** /file-services/qos-policies | Get QoS Policies.
[**GetShares**](ViewAPI.md#GetShares) | **Get** /file-services/shares | Get Shares.
[**GetViewById**](ViewAPI.md#GetViewById) | **Get** /file-services/views/{id} | Get a View by Id
[**GetViewClients**](ViewAPI.md#GetViewClients) | **Get** /file-services/view-clients | Get View Clients.
[**GetViewClientsSummary**](ViewAPI.md#GetViewClientsSummary) | **Get** /file-services/view-clients/summary | Get View Clients Summary.
[**GetViewDirectoryQuotas**](ViewAPI.md#GetViewDirectoryQuotas) | **Get** /file-services/views/{id}/directory-quotas | Get directory quotas for the View.
[**GetViewUserQuotas**](ViewAPI.md#GetViewUserQuotas) | **Get** /file-services/views/{viewId}/user-quotas | Get View user quotas.
[**GetViews**](ViewAPI.md#GetViews) | **Get** /file-services/views | List Views
[**GetViewsSummary**](ViewAPI.md#GetViewsSummary) | **Get** /file-services/views-summary | Get Views summary.
[**ListSmbFileOpens**](ViewAPI.md#ListSmbFileOpens) | **Get** /file-services/smb-file-opens | Get SMB File opens.
[**LockFile**](ViewAPI.md#LockFile) | **Post** /file-services/views/{id}/file-lock | Create a file-lock
[**MigrateS3Views**](ViewAPI.md#MigrateS3Views) | **Post** /file-services/migrate-s3-views | Migrate S3 Views.
[**OverwriteView**](ViewAPI.md#OverwriteView) | **Post** /file-services/views/{id}/overwrite | Overwrite View.
[**ReadViewTemplateById**](ViewAPI.md#ReadViewTemplateById) | **Get** /file-services/view-template/{id} | Read a View Template by Id
[**ReadViewTemplates**](ViewAPI.md#ReadViewTemplates) | **Get** /file-services/view-template | List View Templates
[**UpdateShare**](ViewAPI.md#UpdateShare) | **Put** /file-services/shares/{name} | Update a Share.
[**UpdateView**](ViewAPI.md#UpdateView) | **Put** /file-services/views/{id} | Update a View
[**UpdateViewDirectoryQuota**](ViewAPI.md#UpdateViewDirectoryQuota) | **Put** /file-services/views/{id}/directory-quotas | Update directory quota for the View.
[**UpdateViewTemplate**](ViewAPI.md#UpdateViewTemplate) | **Put** /file-services/view-template/{id} | Update a View Template
[**UpdateViewUserQuotaOverride**](ViewAPI.md#UpdateViewUserQuotaOverride) | **Put** /file-services/views/{viewId}/user-quotas/{userId} | Update user quota override.
[**UpdateViewUserQuotaSettings**](ViewAPI.md#UpdateViewUserQuotaSettings) | **Put** /file-services/views/{viewId}/user-quotas | Update View user quota settings.



## AddViewUserQuotaOverrides

> UserQuotaOverrides AddViewUserQuotaOverrides(ctx, viewId).Body(body).Execute()

Add User Quota overrides.



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
	viewId := int64(789) // int64 | Specifies the id of a view.
	body := *openapiclient.NewUserQuotaOverrides([]openapiclient.UserQuota{*openapiclient.NewUserQuota()}) // UserQuotaOverrides | Specifies the parameters to override the default user quota on the view.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.AddViewUserQuotaOverrides(context.Background(), viewId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.AddViewUserQuotaOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddViewUserQuotaOverrides`: UserQuotaOverrides
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.AddViewUserQuotaOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **int64** | Specifies the id of a view. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddViewUserQuotaOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UserQuotaOverrides**](UserQuotaOverrides.md) | Specifies the parameters to override the default user quota on the view. | 

### Return type

[**UserQuotaOverrides**](UserQuotaOverrides.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearNlmLocks

> ClearNlmLocks(ctx).Body(body).Execute()

Clear NLM locks.



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
	body := *openapiclient.NewClearNlmLockRequest() // ClearNlmLockRequest | Request to clear NLM lock.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ViewAPI.ClearNlmLocks(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.ClearNlmLocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClearNlmLocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClearNlmLockRequest**](ClearNlmLockRequest.md) | Request to clear NLM lock. | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneView

> View CloneView(ctx, id).Body(body).Execute()

Clone View.



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
	id := int64(789) // int64 | Specifies the View id to clone.
	body := *openapiclient.NewCloneViewParams("Name_example") // CloneViewParams | Specifies the request to clone the View.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.CloneView(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.CloneView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloneView`: View
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.CloneView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the View id to clone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CloneViewParams**](CloneViewParams.md) | Specifies the request to clone the View. | 

### Return type

[**View**](View.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneViewDirectory

> CloneViewDirectory(ctx).Body(body).Execute()

Clone View Directory.



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
	body := *openapiclient.NewCloneViewDirectoryParams("SourceDirectoryPath_example", "TargetDirectoryName_example", "TargetParentDirectoryPath_example") // CloneViewDirectoryParams | Specifies the request to clone View directory.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ViewAPI.CloneViewDirectory(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.CloneViewDirectory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloneViewDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CloneViewDirectoryParams**](CloneViewDirectoryParams.md) | Specifies the request to clone View directory. | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloseSmbFileOpen

> CloseSmbFileOpen(ctx).Body(body).Execute()

Close SMB File open.



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
	body := *openapiclient.NewCloseSmbFileOpenParams("FilePath_example", NullableInt64(123), "ViewName_example") // CloseSmbFileOpenParams | Specifies parameters to close active  SMB file open. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ViewAPI.CloseSmbFileOpen(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.CloseSmbFileOpen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloseSmbFileOpenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CloseSmbFileOpenParams**](CloseSmbFileOpenParams.md) | Specifies parameters to close active  SMB file open. | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateShare

> Share CreateShare(ctx).Body(body).Execute()

Create a Share.



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
	body := *openapiclient.NewShare("Name_example", "ViewName_example", "ViewPath_example") // Share | Specifies the request to create a Share.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.CreateShare(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.CreateShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateShare`: Share
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.CreateShare`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Share**](Share.md) | Specifies the request to create a Share. | 

### Return type

[**Share**](Share.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateView

> View CreateView(ctx).Body(body).Execute()

Create a View



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
	body := *openapiclient.NewCreateViewRequest(NullableInt64(123), "Category_example", []openapiclient.ViewProtocol{*openapiclient.NewViewProtocol()}, map[string]interface{}(123)) // CreateViewRequest | Request to create a View.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.CreateView(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.CreateView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateView`: View
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.CreateView`: %v\n", resp)
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

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateViewTemplate

> Template CreateViewTemplate(ctx).Body(body).Execute()

Create a View Template



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
	body := *openapiclient.NewTemplate() // Template | Request to create a view template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.CreateViewTemplate(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.CreateViewTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateViewTemplate`: Template
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.CreateViewTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateViewTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Template**](Template.md) | Request to create a view template. | 

### Return type

[**Template**](Template.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteShare

> DeleteShare(ctx, name).Execute()

Delete a Share.



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
	name := "name_example" // string | Specifies the Share name to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ViewAPI.DeleteShare(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.DeleteShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the Share name to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteShareRequest struct via the builder pattern


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


## DeleteView

> DeleteView(ctx, id).Execute()

Delete a View



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
	id := int64(789) // int64 | Specifies a unique id of the View to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ViewAPI.DeleteView(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.DeleteView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the View to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteViewRequest struct via the builder pattern


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


## DeleteViewDirectoryQuota

> DeleteViewDirectoryQuota(ctx, id).DirectoryPath(directoryPath).DeleteAllDirectoryQuotas(deleteAllDirectoryQuotas).Execute()

Delete directory quota for the View.



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
	id := int64(789) // int64 | Specifies the View id.
	directoryPath := "directoryPath_example" // string | Specifies the directory path to delete. Exactly one of 'directoryPath' and 'deleteAllDirectoryQuotas' should be provided. (optional)
	deleteAllDirectoryQuotas := true // bool | Specifies whether to delete all directory quotas for this view. Exactly one of 'directoryPath' and 'deleteAllDirectoryQuotas' should be provided. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ViewAPI.DeleteViewDirectoryQuota(context.Background(), id).DirectoryPath(directoryPath).DeleteAllDirectoryQuotas(deleteAllDirectoryQuotas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.DeleteViewDirectoryQuota``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteViewDirectoryQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **directoryPath** | **string** | Specifies the directory path to delete. Exactly one of &#39;directoryPath&#39; and &#39;deleteAllDirectoryQuotas&#39; should be provided. | 
 **deleteAllDirectoryQuotas** | **bool** | Specifies whether to delete all directory quotas for this view. Exactly one of &#39;directoryPath&#39; and &#39;deleteAllDirectoryQuotas&#39; should be provided. | 

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


## DeleteViewTemplate

> DeleteViewTemplate(ctx, id).Execute()

Delete a View Template



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
	id := int64(789) // int64 | Specifies a unique id of the view template to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ViewAPI.DeleteViewTemplate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.DeleteViewTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the view template to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteViewTemplateRequest struct via the builder pattern


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


## DeleteViewUserQuotaOverrides

> DeleteViewUserQuotaOverrides(ctx, viewId).Body(body).Execute()

Delete user quota overrides.



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
	viewId := int64(789) // int64 | Specifies the id of a view.
	body := *openapiclient.NewUserQuotaDeleteParams() // UserQuotaDeleteParams | Specifies parameters to delete user quotas.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ViewAPI.DeleteViewUserQuotaOverrides(context.Background(), viewId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.DeleteViewUserQuotaOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **int64** | Specifies the id of a view. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteViewUserQuotaOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UserQuotaDeleteParams**](UserQuotaDeleteParams.md) | Specifies parameters to delete user quotas. | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileLockStatus

> FileLockStatus GetFileLockStatus(ctx, id).Path(path).Execute()

Get file lock status



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
	id := int64(789) // int64 | Specifies the id of a view.
	path := "path_example" // string | Specifies the file path relative to root of the view.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.GetFileLockStatus(context.Background(), id).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.GetFileLockStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileLockStatus`: FileLockStatus
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.GetFileLockStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of a view. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileLockStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | Specifies the file path relative to root of the view. | 

### Return type

[**FileLockStatus**](FileLockStatus.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNlmLocks

> GetNlmLocksResult GetNlmLocks(ctx).FilePath(filePath).ViewName(viewName).MaxCount(maxCount).Cookie(cookie).Execute()

Get NLM locks.



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
	filePath := "filePath_example" // string | Specifies the filepath in the view relative to the root filesystem. If this field is specified, viewName field must also be specified. (optional)
	viewName := "viewName_example" // string | Specifies the name of the View in which to search. If a view name is not specified, all the views in the Cluster is searched. This field is mandatory if filePath field is specified. (optional)
	maxCount := int32(56) // int32 | Specifies the maximum number of NLM locks to return in the response. By default, maxCount is set to 1000. At any given instance, maxCount value cannot be set to more than 1000. (optional)
	cookie := "cookie_example" // string | Specifies the pagination cookie. If this is set, next set of locks just after the previous response are returned. If this is not set, first set of NLM locks are returned.\" (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.GetNlmLocks(context.Background()).FilePath(filePath).ViewName(viewName).MaxCount(maxCount).Cookie(cookie).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.GetNlmLocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNlmLocks`: GetNlmLocksResult
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.GetNlmLocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNlmLocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filePath** | **string** | Specifies the filepath in the view relative to the root filesystem. If this field is specified, viewName field must also be specified. | 
 **viewName** | **string** | Specifies the name of the View in which to search. If a view name is not specified, all the views in the Cluster is searched. This field is mandatory if filePath field is specified. | 
 **maxCount** | **int32** | Specifies the maximum number of NLM locks to return in the response. By default, maxCount is set to 1000. At any given instance, maxCount value cannot be set to more than 1000. | 
 **cookie** | **string** | Specifies the pagination cookie. If this is set, next set of locks just after the previous response are returned. If this is not set, first set of NLM locks are returned.\&quot; | 

### Return type

[**GetNlmLocksResult**](GetNlmLocksResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQosPolicies

> QosPoliciesResult GetQosPolicies(ctx).Execute()

Get QoS Policies.



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
	resp, r, err := apiClient.ViewAPI.GetQosPolicies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.GetQosPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQosPolicies`: QosPoliciesResult
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.GetQosPolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQosPoliciesRequest struct via the builder pattern


### Return type

[**QosPoliciesResult**](QosPoliciesResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShares

> Shares GetShares(ctx).Name(name).MatchPartialName(matchPartialName).MaxCount(maxCount).Cookie(cookie).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()

Get Shares.



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
	name := "name_example" // string | Specifies the Share name. (optional)
	matchPartialName := true // bool | If true, the share name is matched by any partial rather than exactly matched. (optional)
	maxCount := int32(56) // int32 | Specifies a limit on the number of Shares returned. If maxCount is not specified, the first 2000 Shares. (optional)
	cookie := "cookie_example" // string | Specifies the pagination cookie. Expected to be empty in the first call to the API. To get the next set of results, set this value to the pagination cookie value returned in the response of the previous call. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
	includeTenants := true // bool | IncludeTenants specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.GetShares(context.Background()).Name(name).MatchPartialName(matchPartialName).MaxCount(maxCount).Cookie(cookie).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.GetShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShares`: Shares
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.GetShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Specifies the Share name. | 
 **matchPartialName** | **bool** | If true, the share name is matched by any partial rather than exactly matched. | 
 **maxCount** | **int32** | Specifies a limit on the number of Shares returned. If maxCount is not specified, the first 2000 Shares. | 
 **cookie** | **string** | Specifies the pagination cookie. Expected to be empty in the first call to the API. To get the next set of results, set this value to the pagination cookie value returned in the response of the previous call. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | IncludeTenants specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 

### Return type

[**Shares**](Shares.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewById

> View GetViewById(ctx, id).Execute()

Get a View by Id



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
	id := int64(789) // int64 | Specifies a unique id of the View to fetch.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.GetViewById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.GetViewById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetViewById`: View
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.GetViewById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the View to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetViewByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**View**](View.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewClients

> ViewClients GetViewClients(ctx).Protocols(protocols).ViewIds(viewIds).NodeIp(nodeIp).MaxCount(maxCount).Execute()

Get View Clients.



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
	protocols := []string{"Protocols_example"} // []string | Specifies a list of protocols to filter the clients. (optional)
	viewIds := []int64{int64(123)} // []int64 | Specifies a list of View ids. Only clients connected to these Views will be returned. (optional)
	nodeIp := "nodeIp_example" // string | Specifies a node ip. Only clients connected to this node will be returned. (optional)
	maxCount := int32(56) // int32 | Specifies the maximum number of connections to return for SMB and NFS protocols respectively. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.GetViewClients(context.Background()).Protocols(protocols).ViewIds(viewIds).NodeIp(nodeIp).MaxCount(maxCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.GetViewClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetViewClients`: ViewClients
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.GetViewClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetViewClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **protocols** | **[]string** | Specifies a list of protocols to filter the clients. | 
 **viewIds** | **[]int64** | Specifies a list of View ids. Only clients connected to these Views will be returned. | 
 **nodeIp** | **string** | Specifies a node ip. Only clients connected to this node will be returned. | 
 **maxCount** | **int32** | Specifies the maximum number of connections to return for SMB and NFS protocols respectively. | 

### Return type

[**ViewClients**](ViewClients.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewClientsSummary

> ViewClientsSummary GetViewClientsSummary(ctx).ViewIds(viewIds).Execute()

Get View Clients Summary.



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
	viewIds := []int64{int64(123)} // []int64 | Specifies a list of View ids. Only clients connected to these Views will be included in the summary. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.GetViewClientsSummary(context.Background()).ViewIds(viewIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.GetViewClientsSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetViewClientsSummary`: ViewClientsSummary
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.GetViewClientsSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetViewClientsSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewIds** | **[]int64** | Specifies a list of View ids. Only clients connected to these Views will be included in the summary. | 

### Return type

[**ViewClientsSummary**](ViewClientsSummary.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewDirectoryQuotas

> ViewDirectoryQuotas GetViewDirectoryQuotas(ctx, id).MaxCount(maxCount).Cookie(cookie).Execute()

Get directory quotas for the View.



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
	id := int64(789) // int64 | Specifies the View id.
	maxCount := int64(789) // int64 | Specifies a limit on the number of quotas returned. (optional)
	cookie := int64(789) // int64 | Specifies the cookie. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.GetViewDirectoryQuotas(context.Background(), id).MaxCount(maxCount).Cookie(cookie).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.GetViewDirectoryQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetViewDirectoryQuotas`: ViewDirectoryQuotas
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.GetViewDirectoryQuotas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the View id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetViewDirectoryQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxCount** | **int64** | Specifies a limit on the number of quotas returned. | 
 **cookie** | **int64** | Specifies the cookie. | 

### Return type

[**ViewDirectoryQuotas**](ViewDirectoryQuotas.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewUserQuotas

> ViewUserQuotas GetViewUserQuotas(ctx, viewId).IncludeUsage(includeUsage).IncludeUserWithQuotaOverridesOnly(includeUserWithQuotaOverridesOnly).ExcludeUsersWithinAlertThreshold(excludeUsersWithinAlertThreshold).SummaryOnly(summaryOnly).OutputFormat(outputFormat).TopQuotas(topQuotas).MaxCount(maxCount).Cookie(cookie).UnixUid(unixUid).Sid(sid).Execute()

Get View user quotas.



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
	viewId := int64(789) // int64 | Specifies the View id.
	includeUsage := true // bool | If set to true, the logical usage info is included only for users with quota overrides. By default, it is set to false (optional)
	includeUserWithQuotaOverridesOnly := true // bool | If set to true, the result will only contain user with user quota override enabled. By default, this field is set to false, and it's only in effect when 'SummaryOnly' is set to false. (optional)
	excludeUsersWithinAlertThreshold := true // bool | This field can be set only when includeUsage is set to true. By default, all the users with logical usage > 0 will be returned in the result. If this field is set to true, only the list of users who has exceeded the alert threshold will be returned. (optional)
	summaryOnly := true // bool | Specifies a flag to just return a summary. If set to true, it returns the summary of users for a view. By default, it is set to false. (optional)
	outputFormat := "outputFormat_example" // string | OutputFormat is the Output format for the output. If it is not specified, default is json. (optional)
	topQuotas := int64(789) // int64 | TopQuotas is the quotas sorted by quota usage in descending order. This parameter defines number of results to be returned. No pagination cookie is returned if this parameter is set. (optional)
	maxCount := int64(789) // int64 | Specifies a limit on the number of quotas returned. If maxCount is not set, response will have a maximum of 100 results. This parameter will be ignored if 'topQuotas' is set. (optional)
	cookie := "cookie_example" // string | Specifies the cookie. If there are more results than maxCount, response will include a cookie with has to be set as part of the next GET request. (optional)
	unixUid := int32(56) // int32 | Specifies the user identifier of an Unix user. If a valid unix-id to SID mappings are available (i.e., when mixed mode is enabled) the server will perform the necessary id mapping and return the correct usage irrespective of whether the unix id / SID is provided. (optional)
	sid := "sid_example" // string | Specifies the user identifier of a SMB user. If a valid unix-id to SID mappings are available (i.e., when mixed mode is enabled) the server will perform the necessary id mapping and return the correct usage irrespective of whether the unix id / SID is provided. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.GetViewUserQuotas(context.Background(), viewId).IncludeUsage(includeUsage).IncludeUserWithQuotaOverridesOnly(includeUserWithQuotaOverridesOnly).ExcludeUsersWithinAlertThreshold(excludeUsersWithinAlertThreshold).SummaryOnly(summaryOnly).OutputFormat(outputFormat).TopQuotas(topQuotas).MaxCount(maxCount).Cookie(cookie).UnixUid(unixUid).Sid(sid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.GetViewUserQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetViewUserQuotas`: ViewUserQuotas
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.GetViewUserQuotas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **int64** | Specifies the View id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetViewUserQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeUsage** | **bool** | If set to true, the logical usage info is included only for users with quota overrides. By default, it is set to false | 
 **includeUserWithQuotaOverridesOnly** | **bool** | If set to true, the result will only contain user with user quota override enabled. By default, this field is set to false, and it&#39;s only in effect when &#39;SummaryOnly&#39; is set to false. | 
 **excludeUsersWithinAlertThreshold** | **bool** | This field can be set only when includeUsage is set to true. By default, all the users with logical usage &gt; 0 will be returned in the result. If this field is set to true, only the list of users who has exceeded the alert threshold will be returned. | 
 **summaryOnly** | **bool** | Specifies a flag to just return a summary. If set to true, it returns the summary of users for a view. By default, it is set to false. | 
 **outputFormat** | **string** | OutputFormat is the Output format for the output. If it is not specified, default is json. | 
 **topQuotas** | **int64** | TopQuotas is the quotas sorted by quota usage in descending order. This parameter defines number of results to be returned. No pagination cookie is returned if this parameter is set. | 
 **maxCount** | **int64** | Specifies a limit on the number of quotas returned. If maxCount is not set, response will have a maximum of 100 results. This parameter will be ignored if &#39;topQuotas&#39; is set. | 
 **cookie** | **string** | Specifies the cookie. If there are more results than maxCount, response will include a cookie with has to be set as part of the next GET request. | 
 **unixUid** | **int32** | Specifies the user identifier of an Unix user. If a valid unix-id to SID mappings are available (i.e., when mixed mode is enabled) the server will perform the necessary id mapping and return the correct usage irrespective of whether the unix id / SID is provided. | 
 **sid** | **string** | Specifies the user identifier of a SMB user. If a valid unix-id to SID mappings are available (i.e., when mixed mode is enabled) the server will perform the necessary id mapping and return the correct usage irrespective of whether the unix id / SID is provided. | 

### Return type

[**ViewUserQuotas**](ViewUserQuotas.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViews

> GetViewsResult GetViews(ctx).ViewNames(viewNames).ViewIds(viewIds).StorageDomainIds(storageDomainIds).StorageDomainNames(storageDomainNames).ProtocolAccesses(protocolAccesses).MatchPartialNames(matchPartialNames).MaxCount(maxCount).IncludeInternalViews(includeInternalViews).IncludeProtectionGroups(includeProtectionGroups).MaxViewId(maxViewId).IncludeInactive(includeInactive).ProtectionGroupIds(protectionGroupIds).ViewProtectionGroupIds(viewProtectionGroupIds).ViewCountOnly(viewCountOnly).SummaryOnly(summaryOnly).SortByLogicalUsage(sortByLogicalUsage).InternalAccessSids(internalAccessSids).MatchAliasNames(matchAliasNames).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeStats(includeStats).IncludeFileCountBySize(includeFileCountBySize).IncludeViewsWithAntivirusEnabledOnly(includeViewsWithAntivirusEnabledOnly).IncludeViewsWithDataLockEnabledOnly(includeViewsWithDataLockEnabledOnly).FilerAuditLogEnabled(filerAuditLogEnabled).Categories(categories).ViewProtectionTypes(viewProtectionTypes).LastRunAnyStatuses(lastRunAnyStatuses).LastRunLocalBackupStatuses(lastRunLocalBackupStatuses).LastRunReplicationStatuses(lastRunReplicationStatuses).LastRunArchivalStatuses(lastRunArchivalStatuses).IsProtected(isProtected).QosPrincipalIds(qosPrincipalIds).QosPolicies(qosPolicies).UseCachedData(useCachedData).IncludeDeletedProtectionGroups(includeDeletedProtectionGroups).ReturnAllViews(returnAllViews).IncludeS3MigrationOnly(includeS3MigrationOnly).S3MigrationState(s3MigrationState).Execute()

List Views



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
	viewNames := []string{"Inner_example"} // []string | Filter by a list of View names. (optional)
	viewIds := []int64{int64(123)} // []int64 | Filter by a list of View ids. (optional)
	storageDomainIds := []int64{int64(123)} // []int64 | Filter by a list of Storage Domains (View Boxes) specified by id. (optional)
	storageDomainNames := []string{"Inner_example"} // []string | Filter by a list of View Box names. (optional)
	protocolAccesses := []string{"ProtocolAccesses_example"} // []string | Filter by a list of protocol accesses. Only views with protocol accesses in these specified accesses list will be returned. (optional)
	matchPartialNames := true // bool | If true, the names in viewNames are matched by any partial rather than exactly matched. (optional)
	maxCount := int32(56) // int32 | Specifies a limit on the number of Views returned. (optional)
	includeInternalViews := true // bool | Specifies if internal Views created by the Cohesity Cluster are also returned. In addition, regular Views are returned. (optional)
	includeProtectionGroups := true // bool | Specifies if Protection Groups information needs to be returned along with view metadata. By default, if not set or set to true, Group information is returned. (optional)
	maxViewId := int64(789) // int64 | If the number of Views to return exceeds the maxCount specified in the original request, specify the id of the last View from the viewList in the previous response to get the next set of Views. (optional)
	includeInactive := true // bool | Specifies if inactive Views on this Remote Cluster (which have Snapshots copied by replication) should also be returned. Inactive Views are not counted towards the maxCount. By default, this field is set to false. (optional)
	protectionGroupIds := []int64{int64(123)} // []int64 | This field will be deprecated. Filter by Protection Group ids. Return Views that are being protected by listed Groups, which are specified by ids. If both protectionGroupIds and viewProtectionGroupIds are specified, only viewProtectionGroupIds will be used. (optional)
	viewProtectionGroupIds := []string{"Inner_example"} // []string | Filter by Protection Group ids. Return Views that are being protected by listed Groups, which are specified by ids. (optional)
	viewCountOnly := true // bool | Whether to get just the total number of views with the given input filters. If the flag is true, we ignore the parameter 'maxViews' for the count. Also, if flag is true, list of views will not be returned. (optional)
	summaryOnly := true // bool | Whether to get only view summary including 'name', 'viewId', 'storageDomainName', 'storageDomainId' and 'tenantId'. (optional)
	sortByLogicalUsage := true // bool | If set to true, the list is sorted descending by logical usage. (optional)
	internalAccessSids := []string{"Inner_example"} // []string | Sids of restricted principals who can access the view. This is an internal field and therefore does not have json tag. (optional)
	matchAliasNames := true // bool | If true, view aliases are also matched with the names in viewNames. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
	includeTenants := true // bool | IncludeTenants specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)
	includeStats := true // bool | If set to true, stats of views will be returned. By default this parameter is set to false. (optional)
	includeFileCountBySize := true // bool | Whether to include View file count by size. (optional)
	includeViewsWithAntivirusEnabledOnly := true // bool | If set to true, the list will contain only the views for which antivirus scan is enabled. (optional)
	includeViewsWithDataLockEnabledOnly := true // bool | If set to true, the list will contain only the views for which either file level data lock is enabled or view level data lock is enabled. (optional)
	filerAuditLogEnabled := true // bool | If set to true, only views with filer audit log enabled will be returned. If set to false, only views with filer audit log disabled will be returned. (optional)
	categories := []string{"Categories_example"} // []string | Filter by a list of View categories. (optional)
	viewProtectionTypes := []string{"ViewProtectionTypes_example"} // []string | Filter by a list of View protection types. Supported types: [Local Archival ReplicationOut ReplicationIn UnProtected]. UnProtected is mutually exclusive from remaining types. (optional)
	lastRunAnyStatuses := []string{"LastRunAnyStatuses_example"} // []string | Filter by last any run status of the view.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Skipped' indicates that the run was skipped. (optional)
	lastRunLocalBackupStatuses := []string{"LastRunLocalBackupStatuses_example"} // []string | Filter by last local backup run status of the view.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Skipped' indicates that the run was skipped. (optional)
	lastRunReplicationStatuses := []string{"LastRunReplicationStatuses_example"} // []string | Filter by last remote replication run status of the view.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Skipped' indicates that the run was skipped. (optional)
	lastRunArchivalStatuses := []string{"LastRunArchivalStatuses_example"} // []string | Filter by last cloud archival run status of the view.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Skipped' indicates that the run was skipped. (optional)
	isProtected := true // bool | Specifies the protection status of Views. If set to true, only protected Views will be returned. If set to false, only unprotected Views will be returned. (optional)
	qosPrincipalIds := []int64{int64(123)} // []int64 | qosPrincipalIds contains ids of the QoS principal for which views are to be returned. This field is deprecated. (optional)
	qosPolicies := []string{"QosPolicies_example"} // []string | Specifies a filter for Views based on the qosPolicies. This param will be prioritized if qosPrincipalIds is also specified. (optional)
	useCachedData := true // bool | Specifies whether we can serve the GET request to the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. (optional)
	includeDeletedProtectionGroups := true // bool | Specifies if deleted Protection Groups information needs to be returned along with view metadata. By default, deleted Protection Groups are not returned. This is only applied if used along with any view protection related parameter. (optional)
	returnAllViews := true // bool | Specifies if all the Views should be returned as part of the response. (optional)
	includeS3MigrationOnly := true // bool | Specifies whether to return only views which have a s3 migration state. (optional)
	s3MigrationState := "s3MigrationState_example" // string | Filter the list of Views by S3 Migration Statuses. Supported filter values are [Enabled, UnderMigration, Paused, Completed, Eligible].\" If `s3MigrationState` is specified then `includeS3MigrationOnly` param should also be set to true. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.GetViews(context.Background()).ViewNames(viewNames).ViewIds(viewIds).StorageDomainIds(storageDomainIds).StorageDomainNames(storageDomainNames).ProtocolAccesses(protocolAccesses).MatchPartialNames(matchPartialNames).MaxCount(maxCount).IncludeInternalViews(includeInternalViews).IncludeProtectionGroups(includeProtectionGroups).MaxViewId(maxViewId).IncludeInactive(includeInactive).ProtectionGroupIds(protectionGroupIds).ViewProtectionGroupIds(viewProtectionGroupIds).ViewCountOnly(viewCountOnly).SummaryOnly(summaryOnly).SortByLogicalUsage(sortByLogicalUsage).InternalAccessSids(internalAccessSids).MatchAliasNames(matchAliasNames).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeStats(includeStats).IncludeFileCountBySize(includeFileCountBySize).IncludeViewsWithAntivirusEnabledOnly(includeViewsWithAntivirusEnabledOnly).IncludeViewsWithDataLockEnabledOnly(includeViewsWithDataLockEnabledOnly).FilerAuditLogEnabled(filerAuditLogEnabled).Categories(categories).ViewProtectionTypes(viewProtectionTypes).LastRunAnyStatuses(lastRunAnyStatuses).LastRunLocalBackupStatuses(lastRunLocalBackupStatuses).LastRunReplicationStatuses(lastRunReplicationStatuses).LastRunArchivalStatuses(lastRunArchivalStatuses).IsProtected(isProtected).QosPrincipalIds(qosPrincipalIds).QosPolicies(qosPolicies).UseCachedData(useCachedData).IncludeDeletedProtectionGroups(includeDeletedProtectionGroups).ReturnAllViews(returnAllViews).IncludeS3MigrationOnly(includeS3MigrationOnly).S3MigrationState(s3MigrationState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.GetViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetViews`: GetViewsResult
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.GetViews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewNames** | **[]string** | Filter by a list of View names. | 
 **viewIds** | **[]int64** | Filter by a list of View ids. | 
 **storageDomainIds** | **[]int64** | Filter by a list of Storage Domains (View Boxes) specified by id. | 
 **storageDomainNames** | **[]string** | Filter by a list of View Box names. | 
 **protocolAccesses** | **[]string** | Filter by a list of protocol accesses. Only views with protocol accesses in these specified accesses list will be returned. | 
 **matchPartialNames** | **bool** | If true, the names in viewNames are matched by any partial rather than exactly matched. | 
 **maxCount** | **int32** | Specifies a limit on the number of Views returned. | 
 **includeInternalViews** | **bool** | Specifies if internal Views created by the Cohesity Cluster are also returned. In addition, regular Views are returned. | 
 **includeProtectionGroups** | **bool** | Specifies if Protection Groups information needs to be returned along with view metadata. By default, if not set or set to true, Group information is returned. | 
 **maxViewId** | **int64** | If the number of Views to return exceeds the maxCount specified in the original request, specify the id of the last View from the viewList in the previous response to get the next set of Views. | 
 **includeInactive** | **bool** | Specifies if inactive Views on this Remote Cluster (which have Snapshots copied by replication) should also be returned. Inactive Views are not counted towards the maxCount. By default, this field is set to false. | 
 **protectionGroupIds** | **[]int64** | This field will be deprecated. Filter by Protection Group ids. Return Views that are being protected by listed Groups, which are specified by ids. If both protectionGroupIds and viewProtectionGroupIds are specified, only viewProtectionGroupIds will be used. | 
 **viewProtectionGroupIds** | **[]string** | Filter by Protection Group ids. Return Views that are being protected by listed Groups, which are specified by ids. | 
 **viewCountOnly** | **bool** | Whether to get just the total number of views with the given input filters. If the flag is true, we ignore the parameter &#39;maxViews&#39; for the count. Also, if flag is true, list of views will not be returned. | 
 **summaryOnly** | **bool** | Whether to get only view summary including &#39;name&#39;, &#39;viewId&#39;, &#39;storageDomainName&#39;, &#39;storageDomainId&#39; and &#39;tenantId&#39;. | 
 **sortByLogicalUsage** | **bool** | If set to true, the list is sorted descending by logical usage. | 
 **internalAccessSids** | **[]string** | Sids of restricted principals who can access the view. This is an internal field and therefore does not have json tag. | 
 **matchAliasNames** | **bool** | If true, view aliases are also matched with the names in viewNames. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | IncludeTenants specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 
 **includeStats** | **bool** | If set to true, stats of views will be returned. By default this parameter is set to false. | 
 **includeFileCountBySize** | **bool** | Whether to include View file count by size. | 
 **includeViewsWithAntivirusEnabledOnly** | **bool** | If set to true, the list will contain only the views for which antivirus scan is enabled. | 
 **includeViewsWithDataLockEnabledOnly** | **bool** | If set to true, the list will contain only the views for which either file level data lock is enabled or view level data lock is enabled. | 
 **filerAuditLogEnabled** | **bool** | If set to true, only views with filer audit log enabled will be returned. If set to false, only views with filer audit log disabled will be returned. | 
 **categories** | **[]string** | Filter by a list of View categories. | 
 **viewProtectionTypes** | **[]string** | Filter by a list of View protection types. Supported types: [Local Archival ReplicationOut ReplicationIn UnProtected]. UnProtected is mutually exclusive from remaining types. | 
 **lastRunAnyStatuses** | **[]string** | Filter by last any run status of the view.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages.&lt;br&gt; &#39;Skipped&#39; indicates that the run was skipped. | 
 **lastRunLocalBackupStatuses** | **[]string** | Filter by last local backup run status of the view.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages.&lt;br&gt; &#39;Skipped&#39; indicates that the run was skipped. | 
 **lastRunReplicationStatuses** | **[]string** | Filter by last remote replication run status of the view.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages.&lt;br&gt; &#39;Skipped&#39; indicates that the run was skipped. | 
 **lastRunArchivalStatuses** | **[]string** | Filter by last cloud archival run status of the view.&lt;br&gt; &#39;Running&#39; indicates that the run is still running.&lt;br&gt; &#39;Canceled&#39; indicates that the run has been canceled.&lt;br&gt; &#39;Canceling&#39; indicates that the run is in the process of being canceled.&lt;br&gt; &#39;Failed&#39; indicates that the run has failed.&lt;br&gt; &#39;Missed&#39; indicates that the run was unable to take place at the scheduled time because the previous run was still happening.&lt;br&gt; &#39;Succeeded&#39; indicates that the run has finished successfully.&lt;br&gt; &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages.&lt;br&gt; &#39;Skipped&#39; indicates that the run was skipped. | 
 **isProtected** | **bool** | Specifies the protection status of Views. If set to true, only protected Views will be returned. If set to false, only unprotected Views will be returned. | 
 **qosPrincipalIds** | **[]int64** | qosPrincipalIds contains ids of the QoS principal for which views are to be returned. This field is deprecated. | 
 **qosPolicies** | **[]string** | Specifies a filter for Views based on the qosPolicies. This param will be prioritized if qosPrincipalIds is also specified. | 
 **useCachedData** | **bool** | Specifies whether we can serve the GET request to the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. | 
 **includeDeletedProtectionGroups** | **bool** | Specifies if deleted Protection Groups information needs to be returned along with view metadata. By default, deleted Protection Groups are not returned. This is only applied if used along with any view protection related parameter. | 
 **returnAllViews** | **bool** | Specifies if all the Views should be returned as part of the response. | 
 **includeS3MigrationOnly** | **bool** | Specifies whether to return only views which have a s3 migration state. | 
 **s3MigrationState** | **string** | Filter the list of Views by S3 Migration Statuses. Supported filter values are [Enabled, UnderMigration, Paused, Completed, Eligible].\&quot; If &#x60;s3MigrationState&#x60; is specified then &#x60;includeS3MigrationOnly&#x60; param should also be set to true. | 

### Return type

[**GetViewsResult**](GetViewsResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewsSummary

> ViewsSummary GetViewsSummary(ctx).MsecsBeforeCurrentTimeToCompare(msecsBeforeCurrentTimeToCompare).UseCachedData(useCachedData).IncludeInternalViews(includeInternalViews).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeDeletedProtectionGroups(includeDeletedProtectionGroups).Execute()

Get Views summary.



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
	msecsBeforeCurrentTimeToCompare := int64(789) // int64 | Specifies the time in msecs before current time to compare with. (optional)
	useCachedData := true // bool | Specifies whether we can serve the GET request to the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. (optional)
	includeInternalViews := true // bool | Specifies if internal Views created by the Cohesity Cluster are also returned. In addition, regular Views are returned. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
	includeTenants := true // bool | IncludeTenants specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)
	includeDeletedProtectionGroups := true // bool | Specifies if deleted Protection Groups information needs to be returned along with view metadata. By default, deleted Protection Groups are not returned. This is only applied if used along with any view protection related parameter. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.GetViewsSummary(context.Background()).MsecsBeforeCurrentTimeToCompare(msecsBeforeCurrentTimeToCompare).UseCachedData(useCachedData).IncludeInternalViews(includeInternalViews).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeDeletedProtectionGroups(includeDeletedProtectionGroups).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.GetViewsSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetViewsSummary`: ViewsSummary
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.GetViewsSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetViewsSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msecsBeforeCurrentTimeToCompare** | **int64** | Specifies the time in msecs before current time to compare with. | 
 **useCachedData** | **bool** | Specifies whether we can serve the GET request to the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. | 
 **includeInternalViews** | **bool** | Specifies if internal Views created by the Cohesity Cluster are also returned. In addition, regular Views are returned. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | IncludeTenants specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 
 **includeDeletedProtectionGroups** | **bool** | Specifies if deleted Protection Groups information needs to be returned along with view metadata. By default, deleted Protection Groups are not returned. This is only applied if used along with any view protection related parameter. | 

### Return type

[**ViewsSummary**](ViewsSummary.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSmbFileOpens

> SmbFileOpens ListSmbFileOpens(ctx).FilePath(filePath).ViewName(viewName).MaxCount(maxCount).Cookie(cookie).Execute()

Get SMB File opens.



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
	filePath := "filePath_example" // string | Specifies the filepath in the Cohesity View relative to the root filesystem. If this field is specified, viewName field must also be specified. (optional)
	viewName := "viewName_example" // string | Specifies the name of the Cohesity View in which to search. If a view name is not specified, all the views in the Cluster are searched. This field is mandatory if filePath field is specified. (optional)
	maxCount := int32(56) // int32 | Specifies the maximum number of active file opens to return in the response. This field cannot be set above 1000. If this is not set, maximum of 1000 entries are returned. (optional)
	cookie := "cookie_example" // string | Specifies the Pagination Cookie returned in the previous response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.ListSmbFileOpens(context.Background()).FilePath(filePath).ViewName(viewName).MaxCount(maxCount).Cookie(cookie).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.ListSmbFileOpens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSmbFileOpens`: SmbFileOpens
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.ListSmbFileOpens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSmbFileOpensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filePath** | **string** | Specifies the filepath in the Cohesity View relative to the root filesystem. If this field is specified, viewName field must also be specified. | 
 **viewName** | **string** | Specifies the name of the Cohesity View in which to search. If a view name is not specified, all the views in the Cluster are searched. This field is mandatory if filePath field is specified. | 
 **maxCount** | **int32** | Specifies the maximum number of active file opens to return in the response. This field cannot be set above 1000. If this is not set, maximum of 1000 entries are returned. | 
 **cookie** | **string** | Specifies the Pagination Cookie returned in the previous response. | 

### Return type

[**SmbFileOpens**](SmbFileOpens.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockFile

> FileLockStatus LockFile(ctx, id).Body(body).Execute()

Create a file-lock



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
	id := int64(789) // int64 | Specifies the id of a view.
	body := *openapiclient.NewLockFileParams(NullableInt32(123), "FilePath_example") // LockFileParams | Specifies the request params to lock a file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.LockFile(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.LockFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LockFile`: FileLockStatus
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.LockFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of a view. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**LockFileParams**](LockFileParams.md) | Specifies the request params to lock a file | 

### Return type

[**FileLockStatus**](FileLockStatus.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrateS3Views

> MultipleViewsUpdateSuccessFailureIds MigrateS3Views(ctx).Body(body).Execute()

Migrate S3 Views.



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
	body := *openapiclient.NewMigrateS3Views("S3MigrationAction_example", []int32{int32(123)}) // MigrateS3Views | Specifies the request body to Migrate S3 Views.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.MigrateS3Views(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.MigrateS3Views``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MigrateS3Views`: MultipleViewsUpdateSuccessFailureIds
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.MigrateS3Views`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMigrateS3ViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MigrateS3Views**](MigrateS3Views.md) | Specifies the request body to Migrate S3 Views. | 

### Return type

[**MultipleViewsUpdateSuccessFailureIds**](MultipleViewsUpdateSuccessFailureIds.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverwriteView

> OverwriteView(ctx, id).Body(body).Execute()

Overwrite View.



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
	id := int64(789) // int64 | Specifies the View id to be overwritten.
	body := *openapiclient.NewOverwriteViewParams(NullableInt64(123)) // OverwriteViewParams | Specifies the request to overwrite the View.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ViewAPI.OverwriteView(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.OverwriteView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the View id to be overwritten. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverwriteViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OverwriteViewParams**](OverwriteViewParams.md) | Specifies the request to overwrite the View. | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadViewTemplateById

> Template ReadViewTemplateById(ctx, id).Execute()

Read a View Template by Id



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
	id := int64(789) // int64 | Specifies a unique id of the view template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.ReadViewTemplateById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.ReadViewTemplateById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadViewTemplateById`: Template
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.ReadViewTemplateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the view template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadViewTemplateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Template**](Template.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadViewTemplates

> GetViewTemplatesResult ReadViewTemplates(ctx).Execute()

List View Templates



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
	resp, r, err := apiClient.ViewAPI.ReadViewTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.ReadViewTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadViewTemplates`: GetViewTemplatesResult
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.ReadViewTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReadViewTemplatesRequest struct via the builder pattern


### Return type

[**GetViewTemplatesResult**](GetViewTemplatesResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShare

> Share UpdateShare(ctx, name).Body(body).Execute()

Update a Share.



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
	name := "name_example" // string | Specifies the Share name to update.
	body := *openapiclient.NewUpdateShareParam() // UpdateShareParam | Specifies the request to update a Share.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.UpdateShare(context.Background(), name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.UpdateShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateShare`: Share
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.UpdateShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the Share name to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateShareParam**](UpdateShareParam.md) | Specifies the request to update a Share. | 

### Return type

[**Share**](Share.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateView

> View UpdateView(ctx, id).Body(body).Execute()

Update a View



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
	id := int64(789) // int64 | Specifies a unique id of the View to update.
	body := *openapiclient.NewView("Category_example", []openapiclient.ViewProtocol{*openapiclient.NewViewProtocol()}, map[string]interface{}(123)) // View | Request to update a view.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.UpdateView(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.UpdateView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateView`: View
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.UpdateView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the View to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**View**](View.md) | Request to update a view. | 

### Return type

[**View**](View.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateViewDirectoryQuota

> ViewDirectoryQuota UpdateViewDirectoryQuota(ctx, id).Body(body).Execute()

Update directory quota for the View.



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
	id := int64(789) // int64 | Specifies the View id.
	body := *openapiclient.NewViewDirectoryQuota("DirectoryPath_example", *openapiclient.NewViewDirectoryQuotaQuotaPolicy()) // ViewDirectoryQuota | Specifies the request to update directory quota.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.UpdateViewDirectoryQuota(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.UpdateViewDirectoryQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateViewDirectoryQuota`: ViewDirectoryQuota
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.UpdateViewDirectoryQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the View id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateViewDirectoryQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ViewDirectoryQuota**](ViewDirectoryQuota.md) | Specifies the request to update directory quota. | 

### Return type

[**ViewDirectoryQuota**](ViewDirectoryQuota.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateViewTemplate

> Template UpdateViewTemplate(ctx, id).Body(body).Execute()

Update a View Template



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
	id := int64(789) // int64 | Specifies a unique id of the view template.
	body := *openapiclient.NewTemplate() // Template | Request to update a view template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.UpdateViewTemplate(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.UpdateViewTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateViewTemplate`: Template
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.UpdateViewTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the view template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateViewTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Template**](Template.md) | Request to update a view template. | 

### Return type

[**Template**](Template.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateViewUserQuotaOverride

> UserQuota UpdateViewUserQuotaOverride(ctx, viewId, userId).Body(body).Execute()

Update user quota override.



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
	viewId := int64(789) // int64 | Specifies the View id.
	userId := "userId_example" // string | Specifies the unixUid or sid or an user.
	body := *openapiclient.NewQuotaPolicy() // QuotaPolicy | Specifies the user quota policy of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.UpdateViewUserQuotaOverride(context.Background(), viewId, userId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.UpdateViewUserQuotaOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateViewUserQuotaOverride`: UserQuota
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.UpdateViewUserQuotaOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **int64** | Specifies the View id. | 
**userId** | **string** | Specifies the unixUid or sid or an user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateViewUserQuotaOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**QuotaPolicy**](QuotaPolicy.md) | Specifies the user quota policy of the user. | 

### Return type

[**UserQuota**](UserQuota.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateViewUserQuotaSettings

> ViewUserQuotas UpdateViewUserQuotaSettings(ctx, viewId).Body(body).Execute()

Update View user quota settings.



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
	viewId := int64(789) // int64 | Specifies the View id.
	body := *openapiclient.NewViewUserQuotaSettings(false) // ViewUserQuotaSettings | Specifies the parameters to enable/disable or update the default quota config on the view.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewAPI.UpdateViewUserQuotaSettings(context.Background(), viewId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewAPI.UpdateViewUserQuotaSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateViewUserQuotaSettings`: ViewUserQuotas
	fmt.Fprintf(os.Stdout, "Response from `ViewAPI.UpdateViewUserQuotaSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **int64** | Specifies the View id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateViewUserQuotaSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ViewUserQuotaSettings**](ViewUserQuotaSettings.md) | Specifies the parameters to enable/disable or update the default quota config on the view. | 

### Return type

[**ViewUserQuotas**](ViewUserQuotas.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

