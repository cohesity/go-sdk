# \ExternalTarget

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalTarget**](ExternalTarget.md#CreateExternalTarget) | **Post** /data-protect/external-targets | Create a External Target.
[**DeleteExternalTarget**](ExternalTarget.md#DeleteExternalTarget) | **Delete** /data-protect/external-targets/{id} | Delete a External Target.
[**GetExternalTargetById**](ExternalTarget.md#GetExternalTargetById) | **Get** /data-protect/external-targets/{id} | List details about single External Target.
[**GetExternalTargetSettings**](ExternalTarget.md#GetExternalTargetSettings) | **Get** /data-protect/external-targets/settings | Get the list of External Target Settings.
[**GetExternalTargets**](ExternalTarget.md#GetExternalTargets) | **Get** /data-protect/external-targets | Get the list of External Targets.
[**UpdateExternalTarget**](ExternalTarget.md#UpdateExternalTarget) | **Put** /data-protect/external-targets/{id} | Update a External Target.
[**UpdateExternalTargetSettings**](ExternalTarget.md#UpdateExternalTargetSettings) | **Put** /data-protect/external-targets/settings | Update External Target Settings



## CreateExternalTarget

> ExternalTarget CreateExternalTarget(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Create a External Target.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    body := *openapiclient.NewExternalTarget("Name_example", "PurposeType_example") // ExternalTarget | Specifies the parameters to create a External Target.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiCreateExternalTargetRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.ExternalTarget.CreateExternalTarget(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalTarget.CreateExternalTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateExternalTarget`: ExternalTarget
    fmt.Fprintf(os.Stdout, "Response from `ExternalTarget.CreateExternalTarget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ExternalTarget**](ExternalTarget.md) | Specifies the parameters to create a External Target. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> DeleteExternalTarget(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Delete a External Target.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the External Target.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiDeleteExternalTargetRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.ExternalTarget.DeleteExternalTarget(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalTarget.DeleteExternalTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the External Target. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExternalTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> ExternalTarget GetExternalTargetById(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

List details about single External Target.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := int64(789) // int64 | Specifies a unique id of the External Target.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiGetExternalTargetByIdRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.ExternalTarget.GetExternalTargetById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalTarget.GetExternalTargetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExternalTargetById`: ExternalTarget
    fmt.Fprintf(os.Stdout, "Response from `ExternalTarget.GetExternalTargetById`: %v\n", resp)
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

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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


## GetExternalTargetSettings

> ExternalTarget GetExternalTargetSettings(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get the list of External Target Settings.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiGetExternalTargetSettingsRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.ExternalTarget.GetExternalTargetSettings(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalTarget.GetExternalTargetSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExternalTargetSettings`: ExternalTarget
    fmt.Fprintf(os.Stdout, "Response from `ExternalTarget.GetExternalTargetSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalTargetSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> ExternalTargets GetExternalTargets(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Ids(ids).Names(names).PurposeTypes(purposeTypes).StorageTypes(storageTypes).StorageClasses(storageClasses).Execute()

Get the list of External Targets.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    ids := []int64{int64(123)} // []int64 | Filter by a list of External Target ids. (optional)
    names := []string{"Inner_example"} // []string | Filter by a list of External Target names. (optional)
    purposeTypes := []string{"PurposeTypes_example"} // []string | Filter by a list of External Target purpose types. (optional)
    storageTypes := []string{"StorageTypes_example"} // []string | Filter by a list of External Target storage types. (optional)
    storageClasses := []string{"StorageClasses_example"} // []string | Filter by a list of External Target storage classes. (optional)

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

    request := helios.ApiGetExternalTargetsRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        Ids: &ids
        Names: &names
        PurposeTypes: &purposeTypes
        StorageTypes: &storageTypes
        StorageClasses: &storageClasses
    }

    resp, r, err := api_client.ExternalTarget.GetExternalTargets(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalTarget.GetExternalTargets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExternalTargets`: ExternalTargets
    fmt.Fprintf(os.Stdout, "Response from `ExternalTarget.GetExternalTargets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **ids** | **[]int64** | Filter by a list of External Target ids. | 
 **names** | **[]string** | Filter by a list of External Target names. | 
 **purposeTypes** | **[]string** | Filter by a list of External Target purpose types. | 
 **storageTypes** | **[]string** | Filter by a list of External Target storage types. | 
 **storageClasses** | **[]string** | Filter by a list of External Target storage classes. | 

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

> ExternalTarget UpdateExternalTarget(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update a External Target.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := int64(789) // int64 | Specifies the id of the External Target.
    body := *openapiclient.NewExternalTarget("Name_example", "PurposeType_example") // ExternalTarget | Specifies the parameters to update a External Target.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiUpdateExternalTargetRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.ExternalTarget.UpdateExternalTarget(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalTarget.UpdateExternalTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExternalTarget`: ExternalTarget
    fmt.Fprintf(os.Stdout, "Response from `ExternalTarget.UpdateExternalTarget`: %v\n", resp)
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> GlobalBandwidthSettings UpdateExternalTargetSettings(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update External Target Settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    body := *openapiclient.NewGlobalBandwidthSettings() // GlobalBandwidthSettings | Specifies the parameters to update a External Target Settings.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiUpdateExternalTargetSettingsRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.ExternalTarget.UpdateExternalTargetSettings(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalTarget.UpdateExternalTargetSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExternalTargetSettings`: GlobalBandwidthSettings
    fmt.Fprintf(os.Stdout, "Response from `ExternalTarget.UpdateExternalTargetSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExternalTargetSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GlobalBandwidthSettings**](GlobalBandwidthSettings.md) | Specifies the parameters to update a External Target Settings. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

