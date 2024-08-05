# \ExternalTargetAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalTarget**](ExternalTargetAPI.md#CreateExternalTarget) | **Post** /data-protect/external-targets | Create a External Target.
[**DeleteExternalTarget**](ExternalTargetAPI.md#DeleteExternalTarget) | **Delete** /data-protect/external-targets/{id} | Delete a External Target.
[**GetExternalTargetById**](ExternalTargetAPI.md#GetExternalTargetById) | **Get** /data-protect/external-targets/{id} | List details about single External Target.
[**GetExternalTargetEncryptionKeyInfo**](ExternalTargetAPI.md#GetExternalTargetEncryptionKeyInfo) | **Get** /data-protect/external-targets/{id}/encryption-key | Get the encryption key info for an external target
[**GetExternalTargetMediaInfo**](ExternalTargetAPI.md#GetExternalTargetMediaInfo) | **Get** /data-protect/external-targets/media-info | List archive media information
[**GetExternalTargetSettings**](ExternalTargetAPI.md#GetExternalTargetSettings) | **Get** /data-protect/external-targets/settings | Get the list of External Target Settings.
[**GetExternalTargets**](ExternalTargetAPI.md#GetExternalTargets) | **Get** /data-protect/external-targets | Get the list of External Targets.
[**UpdateExternalTarget**](ExternalTargetAPI.md#UpdateExternalTarget) | **Put** /data-protect/external-targets/{id} | Update a External Target.
[**UpdateExternalTargetSettings**](ExternalTargetAPI.md#UpdateExternalTargetSettings) | **Put** /data-protect/external-targets/settings | Update External Target Settings



## CreateExternalTarget

> ExternalTarget CreateExternalTarget(ctx).Body(body).Execute()

Create a External Target.



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
	body := *openapiclient.NewExternalTarget("Name_example", "PurposeType_example") // ExternalTarget | Specifies the parameters to create a External Target.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalTargetAPI.CreateExternalTarget(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTargetAPI.CreateExternalTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExternalTarget`: ExternalTarget
	fmt.Fprintf(os.Stdout, "Response from `ExternalTargetAPI.CreateExternalTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ExternalTarget**](ExternalTarget.md) | Specifies the parameters to create a External Target. | 

### Return type

[**ExternalTarget**](ExternalTarget.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExternalTarget

> DeleteExternalTarget(ctx, id).ForceDelete(forceDelete).Execute()

Delete a External Target.



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
	id := int64(789) // int64 | Specifies a unique id of the External Target.
	forceDelete := true // bool | Specifies whether to force delete the External target. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExternalTargetAPI.DeleteExternalTarget(context.Background(), id).ForceDelete(forceDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTargetAPI.DeleteExternalTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the External Target. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExternalTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceDelete** | **bool** | Specifies whether to force delete the External target. | 

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


## GetExternalTargetById

> ExternalTarget GetExternalTargetById(ctx, id).Execute()

List details about single External Target.



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
	id := int64(789) // int64 | Specifies a unique id of the External Target.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalTargetAPI.GetExternalTargetById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTargetAPI.GetExternalTargetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalTargetById`: ExternalTarget
	fmt.Fprintf(os.Stdout, "Response from `ExternalTargetAPI.GetExternalTargetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the External Target. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalTargetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalTarget**](ExternalTarget.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalTargetEncryptionKeyInfo

> *os.File GetExternalTargetEncryptionKeyInfo(ctx, id).Execute()

Get the encryption key info for an external target



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
	id := int64(789) // int64 | Specifies the id of the External Target.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalTargetAPI.GetExternalTargetEncryptionKeyInfo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTargetAPI.GetExternalTargetEncryptionKeyInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalTargetEncryptionKeyInfo`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ExternalTargetAPI.GetExternalTargetEncryptionKeyInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the External Target. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalTargetEncryptionKeyInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalTargetMediaInfo

> ArchivalMediaInfo GetExternalTargetMediaInfo(ctx).ClusterId(clusterId).ClusterIncarnationId(clusterIncarnationId).ArchivalJobId(archivalJobId).RestoreTaskId(restoreTaskId).EntityIds(entityIds).Execute()

List archive media information



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
	clusterId := int64(789) // int64 | Specifies the id of the Cohesity cluster which archived to a QStart media target.
	clusterIncarnationId := int64(789) // int64 | Specifies the incarnation Id of the Cohesity cluster which archived to a QStart media target.
	archivalJobId := int64(789) // int64 | Specifies the id of the Job that archived to a QStar media Vault.
	restoreTaskId := int64(789) // int64 | Specifies the id of the restore task to optionally filter by. (optional)
	entityIds := []int64{int64(123)} // []int64 | Specifies an array of entityIds to optionally filter by. An entityId is a unique id for a VM assigned by the Cohesity Cluster. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalTargetAPI.GetExternalTargetMediaInfo(context.Background()).ClusterId(clusterId).ClusterIncarnationId(clusterIncarnationId).ArchivalJobId(archivalJobId).RestoreTaskId(restoreTaskId).EntityIds(entityIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTargetAPI.GetExternalTargetMediaInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalTargetMediaInfo`: ArchivalMediaInfo
	fmt.Fprintf(os.Stdout, "Response from `ExternalTargetAPI.GetExternalTargetMediaInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalTargetMediaInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **int64** | Specifies the id of the Cohesity cluster which archived to a QStart media target. | 
 **clusterIncarnationId** | **int64** | Specifies the incarnation Id of the Cohesity cluster which archived to a QStart media target. | 
 **archivalJobId** | **int64** | Specifies the id of the Job that archived to a QStar media Vault. | 
 **restoreTaskId** | **int64** | Specifies the id of the restore task to optionally filter by. | 
 **entityIds** | **[]int64** | Specifies an array of entityIds to optionally filter by. An entityId is a unique id for a VM assigned by the Cohesity Cluster. | 

### Return type

[**ArchivalMediaInfo**](ArchivalMediaInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalTargetSettings

> ExternalTarget GetExternalTargetSettings(ctx).Execute()

Get the list of External Target Settings.



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
	resp, r, err := apiClient.ExternalTargetAPI.GetExternalTargetSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTargetAPI.GetExternalTargetSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalTargetSettings`: ExternalTarget
	fmt.Fprintf(os.Stdout, "Response from `ExternalTargetAPI.GetExternalTargetSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalTargetSettingsRequest struct via the builder pattern


### Return type

[**ExternalTarget**](ExternalTarget.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalTargets

> ExternalTargets GetExternalTargets(ctx).Ids(ids).GlobalIds(globalIds).Names(names).PurposeTypes(purposeTypes).StorageTypes(storageTypes).StorageClasses(storageClasses).OwnershipContexts(ownershipContexts).Execute()

Get the list of External Targets.



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
	ids := []int64{int64(123)} // []int64 | Filter by a list of External Target ids. (optional)
	globalIds := []string{"Inner_example"} // []string | Filter by a list of External Target global ids. (optional)
	names := []string{"Inner_example"} // []string | Filter by a list of External Target names. (optional)
	purposeTypes := []string{"PurposeTypes_example"} // []string | Filter by a list of External Target purpose types. (optional)
	storageTypes := []string{"StorageTypes_example"} // []string | Filter by a list of External Target storage types. Nas option in archival_target_storage_type will soon be deprecated. Please use NAS instead. (optional)
	storageClasses := []string{"StorageClasses_example"} // []string | Filter by a list of External Target storage classes. (optional)
	ownershipContexts := []string{"OwnershipContexts_example"} // []string | Specifies whether how this external target is being consumed either Local or FortKnox. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalTargetAPI.GetExternalTargets(context.Background()).Ids(ids).GlobalIds(globalIds).Names(names).PurposeTypes(purposeTypes).StorageTypes(storageTypes).StorageClasses(storageClasses).OwnershipContexts(ownershipContexts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTargetAPI.GetExternalTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalTargets`: ExternalTargets
	fmt.Fprintf(os.Stdout, "Response from `ExternalTargetAPI.GetExternalTargets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | Filter by a list of External Target ids. | 
 **globalIds** | **[]string** | Filter by a list of External Target global ids. | 
 **names** | **[]string** | Filter by a list of External Target names. | 
 **purposeTypes** | **[]string** | Filter by a list of External Target purpose types. | 
 **storageTypes** | **[]string** | Filter by a list of External Target storage types. Nas option in archival_target_storage_type will soon be deprecated. Please use NAS instead. | 
 **storageClasses** | **[]string** | Filter by a list of External Target storage classes. | 
 **ownershipContexts** | **[]string** | Specifies whether how this external target is being consumed either Local or FortKnox. | 

### Return type

[**ExternalTargets**](ExternalTargets.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExternalTarget

> ExternalTarget UpdateExternalTarget(ctx, id).Body(body).Execute()

Update a External Target.



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
	id := int64(789) // int64 | Specifies the id of the External Target.
	body := *openapiclient.NewExternalTarget("Name_example", "PurposeType_example") // ExternalTarget | Specifies the parameters to update a External Target.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalTargetAPI.UpdateExternalTarget(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTargetAPI.UpdateExternalTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExternalTarget`: ExternalTarget
	fmt.Fprintf(os.Stdout, "Response from `ExternalTargetAPI.UpdateExternalTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the External Target. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExternalTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ExternalTarget**](ExternalTarget.md) | Specifies the parameters to update a External Target. | 

### Return type

[**ExternalTarget**](ExternalTarget.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExternalTargetSettings

> GlobalBandwidthSettings UpdateExternalTargetSettings(ctx).Body(body).Execute()

Update External Target Settings



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
	body := *openapiclient.NewGlobalBandwidthSettings() // GlobalBandwidthSettings | Specifies the parameters to update a External Target Settings.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalTargetAPI.UpdateExternalTargetSettings(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalTargetAPI.UpdateExternalTargetSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExternalTargetSettings`: GlobalBandwidthSettings
	fmt.Fprintf(os.Stdout, "Response from `ExternalTargetAPI.UpdateExternalTargetSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExternalTargetSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GlobalBandwidthSettings**](GlobalBandwidthSettings.md) | Specifies the parameters to update a External Target Settings. | 

### Return type

[**GlobalBandwidthSettings**](GlobalBandwidthSettings.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

