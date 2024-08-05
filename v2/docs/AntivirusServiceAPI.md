# \AntivirusServiceAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAntivirusGroup**](AntivirusServiceAPI.md#CreateAntivirusGroup) | **Post** /antivirus-service/groups | Create an Antivirus Service group.
[**DeleteAntivirusGroup**](AntivirusServiceAPI.md#DeleteAntivirusGroup) | **Delete** /antivirus-service/groups/{id} | Delete an Antivirus Service group
[**DeleteInfectedFiles**](AntivirusServiceAPI.md#DeleteInfectedFiles) | **Delete** /antivirus-service/infected-files | Delete infected files.
[**DeleteInfectedObjects**](AntivirusServiceAPI.md#DeleteInfectedObjects) | **Delete** /antivirus-service/infected-objects | Delete infected objects permanently.
[**GetAntivirusServiceGroups**](AntivirusServiceAPI.md#GetAntivirusServiceGroups) | **Get** /antivirus-service/groups | Get Antivirus Service groups.
[**GetIcapUriConnectionStatus**](AntivirusServiceAPI.md#GetIcapUriConnectionStatus) | **Get** /antivirus-service/icap-uri-connection-status | Get ICAP Uri connection status.
[**GetInfectedFiles**](AntivirusServiceAPI.md#GetInfectedFiles) | **Get** /antivirus-service/infected-files | Get infected entities.
[**UpdateAntivirusGroup**](AntivirusServiceAPI.md#UpdateAntivirusGroup) | **Put** /antivirus-service/groups/{id} | Update an Antivirus Service group with given parameters or if state is specified, enable or disable given group.
[**UpdateInfectedFiles**](AntivirusServiceAPI.md#UpdateInfectedFiles) | **Put** /antivirus-service/infected-files | Update infected entities state.



## CreateAntivirusGroup

> AntivirusServiceGroup CreateAntivirusGroup(ctx).Body(body).Execute()

Create an Antivirus Service group.



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
	body := *openapiclient.NewCreateAntivirusServiceGroupParams([]openapiclient.AntivirusService{*openapiclient.NewAntivirusService("IcapUri_example")}, "Name_example") // CreateAntivirusServiceGroupParams | Specifies the parameters to create antivirus service group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AntivirusServiceAPI.CreateAntivirusGroup(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AntivirusServiceAPI.CreateAntivirusGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAntivirusGroup`: AntivirusServiceGroup
	fmt.Fprintf(os.Stdout, "Response from `AntivirusServiceAPI.CreateAntivirusGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAntivirusGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateAntivirusServiceGroupParams**](CreateAntivirusServiceGroupParams.md) | Specifies the parameters to create antivirus service group. | 

### Return type

[**AntivirusServiceGroup**](AntivirusServiceGroup.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAntivirusGroup

> DeleteAntivirusGroup(ctx, id).Execute()

Delete an Antivirus Service group



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
	id := int64(789) // int64 | Specifies a unique id of the Antivirus Group to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AntivirusServiceAPI.DeleteAntivirusGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AntivirusServiceAPI.DeleteAntivirusGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the Antivirus Group to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAntivirusGroupRequest struct via the builder pattern


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


## DeleteInfectedFiles

> DeleteInfectedFiles DeleteInfectedFiles(ctx).Body(body).Execute()

Delete infected files.



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
	body := *openapiclient.NewDeleteInfectedFilesParameters([]openapiclient.InfectedFile{*openapiclient.NewInfectedFile(NullableInt64(123), NullableInt64(123), NullableInt64(123))}) // DeleteInfectedFilesParameters | Specifies the parameters of infected files to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AntivirusServiceAPI.DeleteInfectedFiles(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AntivirusServiceAPI.DeleteInfectedFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInfectedFiles`: DeleteInfectedFiles
	fmt.Fprintf(os.Stdout, "Response from `AntivirusServiceAPI.DeleteInfectedFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInfectedFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteInfectedFilesParameters**](DeleteInfectedFilesParameters.md) | Specifies the parameters of infected files to be deleted. | 

### Return type

[**DeleteInfectedFiles**](DeleteInfectedFiles.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInfectedObjects

> DeleteInfectedObjects DeleteInfectedObjects(ctx).Body(body).Execute()

Delete infected objects permanently.



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
	body := *openapiclient.NewDeleteInfectedObjectsParameters([]openapiclient.InfectedObject{*openapiclient.NewInfectedObject("BucketName_example", "ObjectName_example")}) // DeleteInfectedObjectsParameters | Specifies the parameters of infected objects to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AntivirusServiceAPI.DeleteInfectedObjects(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AntivirusServiceAPI.DeleteInfectedObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInfectedObjects`: DeleteInfectedObjects
	fmt.Fprintf(os.Stdout, "Response from `AntivirusServiceAPI.DeleteInfectedObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInfectedObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteInfectedObjectsParameters**](DeleteInfectedObjectsParameters.md) | Specifies the parameters of infected objects to be deleted. | 

### Return type

[**DeleteInfectedObjects**](DeleteInfectedObjects.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAntivirusServiceGroups

> AntivirusServiceGroups GetAntivirusServiceGroups(ctx).Execute()

Get Antivirus Service groups.



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
	resp, r, err := apiClient.AntivirusServiceAPI.GetAntivirusServiceGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AntivirusServiceAPI.GetAntivirusServiceGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAntivirusServiceGroups`: AntivirusServiceGroups
	fmt.Fprintf(os.Stdout, "Response from `AntivirusServiceAPI.GetAntivirusServiceGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAntivirusServiceGroupsRequest struct via the builder pattern


### Return type

[**AntivirusServiceGroups**](AntivirusServiceGroups.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIcapUriConnectionStatus

> IcapUriConnectionStatusList GetIcapUriConnectionStatus(ctx).Uris(uris).Execute()

Get ICAP Uri connection status.



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
	uris := []string{"Inner_example"} // []string | Specifies a list of URIs to check connection status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AntivirusServiceAPI.GetIcapUriConnectionStatus(context.Background()).Uris(uris).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AntivirusServiceAPI.GetIcapUriConnectionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIcapUriConnectionStatus`: IcapUriConnectionStatusList
	fmt.Fprintf(os.Stdout, "Response from `AntivirusServiceAPI.GetIcapUriConnectionStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIcapUriConnectionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uris** | **[]string** | Specifies a list of URIs to check connection status. | 

### Return type

[**IcapUriConnectionStatusList**](IcapUriConnectionStatusList.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfectedFiles

> InfectedFiles GetInfectedFiles(ctx).ViewIds(viewIds).Path(path).States(states).MaxCount(maxCount).Cookie(cookie).Execute()

Get infected entities.



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
	viewIds := []int64{int64(123)} // []int64 | Specifies a list of view ids. Only infected entities from these views will be returned. (optional)
	path := "path_example" // string | Specifies the file path. (optional)
	states := []string{"States_example"} // []string | Specifies the file states. (optional)
	maxCount := int64(789) // int64 | Specifies the max number of files to be returned. (optional)
	cookie := "cookie_example" // string | Specifies the pagination cookie. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AntivirusServiceAPI.GetInfectedFiles(context.Background()).ViewIds(viewIds).Path(path).States(states).MaxCount(maxCount).Cookie(cookie).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AntivirusServiceAPI.GetInfectedFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfectedFiles`: InfectedFiles
	fmt.Fprintf(os.Stdout, "Response from `AntivirusServiceAPI.GetInfectedFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInfectedFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewIds** | **[]int64** | Specifies a list of view ids. Only infected entities from these views will be returned. | 
 **path** | **string** | Specifies the file path. | 
 **states** | **[]string** | Specifies the file states. | 
 **maxCount** | **int64** | Specifies the max number of files to be returned. | 
 **cookie** | **string** | Specifies the pagination cookie. | 

### Return type

[**InfectedFiles**](InfectedFiles.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAntivirusGroup

> AntivirusServiceGroup UpdateAntivirusGroup(ctx, id).Body(body).Execute()

Update an Antivirus Service group with given parameters or if state is specified, enable or disable given group.



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
	id := int64(789) // int64 | Specifies a unique id of the Antivirus Group to update.
	body := *openapiclient.NewAntivirusServiceGroup([]openapiclient.AntivirusService{*openapiclient.NewAntivirusService("IcapUri_example")}, "Name_example") // AntivirusServiceGroup | Specifies the parameters to update antivirus service group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AntivirusServiceAPI.UpdateAntivirusGroup(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AntivirusServiceAPI.UpdateAntivirusGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAntivirusGroup`: AntivirusServiceGroup
	fmt.Fprintf(os.Stdout, "Response from `AntivirusServiceAPI.UpdateAntivirusGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the Antivirus Group to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAntivirusGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AntivirusServiceGroup**](AntivirusServiceGroup.md) | Specifies the parameters to update antivirus service group. | 

### Return type

[**AntivirusServiceGroup**](AntivirusServiceGroup.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInfectedFiles

> UpdateInfectedFilesList UpdateInfectedFiles(ctx).Body(body).Execute()

Update infected entities state.



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
	body := *openapiclient.NewUpdateInfectedFilesParameters([]openapiclient.InfectedFile{*openapiclient.NewInfectedFile(NullableInt64(123), NullableInt64(123), NullableInt64(123))}) // UpdateInfectedFilesParameters | Specifies the parameters of infected entities to be updated.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AntivirusServiceAPI.UpdateInfectedFiles(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AntivirusServiceAPI.UpdateInfectedFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInfectedFiles`: UpdateInfectedFilesList
	fmt.Fprintf(os.Stdout, "Response from `AntivirusServiceAPI.UpdateInfectedFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInfectedFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateInfectedFilesParameters**](UpdateInfectedFilesParameters.md) | Specifies the parameters of infected entities to be updated. | 

### Return type

[**UpdateInfectedFilesList**](UpdateInfectedFilesList.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

