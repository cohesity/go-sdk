# \AnalyticsApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyzeJar**](AnalyticsApi.md#AnalyzeJar) | **Post** /public/analytics/analyzeJar | Analyze the uploaded jar.
[**CancelMapReduceInstanceRun**](AnalyticsApi.md#CancelMapReduceInstanceRun) | **Put** /public/analytics/cancelAppInstanceRun/{id} | Cancel a specific map reduce instance run.
[**CreateApplication**](AnalyticsApi.md#CreateApplication) | **Post** /public/analytics/apps | Create an Application.
[**CreateMapper**](AnalyticsApi.md#CreateMapper) | **Post** /public/analytics/mappers | Create a Mapper.
[**CreateReducer**](AnalyticsApi.md#CreateReducer) | **Post** /public/analytics/reducers | Create a Reducer.
[**DeleteApplication**](AnalyticsApi.md#DeleteApplication) | **Delete** /public/analytics/apps/{id} | Delete an Application.
[**DeleteMapReduceInstanceRun**](AnalyticsApi.md#DeleteMapReduceInstanceRun) | **Delete** /public/analytics/mrAppRun/{id} | Delete a Map-Reduce Application Instance Run.
[**DeleteMapper**](AnalyticsApi.md#DeleteMapper) | **Delete** /public/analytics/mappers/{id} | Delete a Mapper.
[**DeleteReducer**](AnalyticsApi.md#DeleteReducer) | **Delete** /public/analytics/reducers/{id} | Delete a Reducer.
[**DeleteUploadedJar**](AnalyticsApi.md#DeleteUploadedJar) | **Delete** /public/analytics/uploadJar | Delete the uploaded jar from temporary location.
[**DownloadMRBaseJar**](AnalyticsApi.md#DownloadMRBaseJar) | **Get** /public/analytics/mrBaseJar | Downloads the map reduce base jar.
[**DownloadMROutputFiles**](AnalyticsApi.md#DownloadMROutputFiles) | **Get** /public/analytics/mrOutputfiles | Download map reduce base instance run output files from Yoda.
[**GetApplicationById**](AnalyticsApi.md#GetApplicationById) | **Get** /public/analytics/apps/{id} | List details about a single Application.
[**GetApplications**](AnalyticsApi.md#GetApplications) | **Get** /public/analytics/apps | List Applications filtered by the specified parameters.
[**GetMRUploadJarPath**](AnalyticsApi.md#GetMRUploadJarPath) | **Get** /public/analytics/uploadJarPath | Get details about the mounted path to upload Jars.
[**GetMapReduceAppRuns**](AnalyticsApi.md#GetMapReduceAppRuns) | **Get** /public/analytics/mrAppRuns | List map reduce application runs filtered by the specified parameters.
[**GetMapReduceFileFormats**](AnalyticsApi.md#GetMapReduceFileFormats) | **Get** /public/analytics/mrFileFormats | Used to retrieve supported output file formats from a map reduce instance.
[**GetMapperById**](AnalyticsApi.md#GetMapperById) | **Get** /public/analytics/mappers/{id} | List details about a single Mapper.
[**GetMappers**](AnalyticsApi.md#GetMappers) | **Get** /public/analytics/mappers | List Mappers filtered by the specified parameters.
[**GetReducerById**](AnalyticsApi.md#GetReducerById) | **Get** /public/analytics/reducers/{id} | List details about a single Reducer.
[**GetReducers**](AnalyticsApi.md#GetReducers) | **Get** /public/analytics/reducers | List Reducers filtered by the specified parameters.
[**GetSupportedPatterns**](AnalyticsApi.md#GetSupportedPatterns) | **Get** /public/analytics/supportedPatterns | Used to retrieve supported patterns available for search in an application.
[**RunMapReduceAppInstance**](AnalyticsApi.md#RunMapReduceAppInstance) | **Put** /public/analytics/runAppInstance | Run a map-reduce application instance.
[**SavePattern**](AnalyticsApi.md#SavePattern) | **Put** /public/analytics/supportedPatterns | Used to save user patterns for search in an application.
[**UpdateApplication**](AnalyticsApi.md#UpdateApplication) | **Put** /public/analytics/apps/{id} | Update an Application.
[**UpdateMapper**](AnalyticsApi.md#UpdateMapper) | **Put** /public/analytics/mappers/{id} | Update a Mapper.
[**UpdateReducer**](AnalyticsApi.md#UpdateReducer) | **Put** /public/analytics/reducers/{id} | Update a Reducer.
[**UploadJar**](AnalyticsApi.md#UploadJar) | **Post** /public/analytics/uploadJar | Upload a jar to pre-specified Yoda internal view.



## AnalyzeJar

> AnalyseJarResult AnalyzeJar(ctx).Body(body).Execute()

Analyze the uploaded jar.



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
    body := *openapiclient.NewAnalyseJarArg() // AnalyseJarArg |  (optional)

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

    request := cohesitysdk.ApiAnalyzeJarRequest{
        Body: &body
    }

    resp, r, err := api_client.AnalyticsApi.AnalyzeJar(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.AnalyzeJar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AnalyzeJar`: AnalyseJarResult
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.AnalyzeJar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyzeJarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AnalyseJarArg**](AnalyseJarArg.md) |  | 

### Return type

[**AnalyseJarResult**](AnalyseJarResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelMapReduceInstanceRun

> KillMapReduceInstanceResult CancelMapReduceInstanceRun(ctx, id).Execute()

Cancel a specific map reduce instance run.



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
    id := int64(789) // int64 | 

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

    request := cohesitysdk.ApiCancelMapReduceInstanceRunRequest{
        Id: &id
    }

    resp, r, err := api_client.AnalyticsApi.CancelMapReduceInstanceRun(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.CancelMapReduceInstanceRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelMapReduceInstanceRun`: KillMapReduceInstanceResult
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.CancelMapReduceInstanceRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelMapReduceInstanceRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KillMapReduceInstanceResult**](KillMapReduceInstanceResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplication

> MapReduceInfo CreateApplication(ctx).Body(body).Execute()

Create an Application.



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
    body := *openapiclient.NewMapReduceInfo() // MapReduceInfo |  (optional)

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

    request := cohesitysdk.ApiCreateApplicationRequest{
        Body: &body
    }

    resp, r, err := api_client.AnalyticsApi.CreateApplication(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.CreateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplication`: MapReduceInfo
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.CreateApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MapReduceInfo**](MapReduceInfo.md) |  | 

### Return type

[**MapReduceInfo**](MapReduceInfo.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMapper

> MapperInfo CreateMapper(ctx).Body(body).Execute()

Create a Mapper.



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
    body := *openapiclient.NewMapperInfo() // MapperInfo |  (optional)

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

    request := cohesitysdk.ApiCreateMapperRequest{
        Body: &body
    }

    resp, r, err := api_client.AnalyticsApi.CreateMapper(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.CreateMapper``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMapper`: MapperInfo
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.CreateMapper`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMapperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MapperInfo**](MapperInfo.md) |  | 

### Return type

[**MapperInfo**](MapperInfo.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReducer

> ReducerInfo CreateReducer(ctx).Body(body).Execute()

Create a Reducer.



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
    body := *openapiclient.NewReducerInfo() // ReducerInfo |  (optional)

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

    request := cohesitysdk.ApiCreateReducerRequest{
        Body: &body
    }

    resp, r, err := api_client.AnalyticsApi.CreateReducer(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.CreateReducer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReducer`: ReducerInfo
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.CreateReducer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReducerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ReducerInfo**](ReducerInfo.md) |  | 

### Return type

[**ReducerInfo**](ReducerInfo.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplication

> DeleteApplication(ctx, id).Execute()

Delete an Application.



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
    id := int64(789) // int64 | 

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

    request := cohesitysdk.ApiDeleteApplicationRequest{
        Id: &id
    }

    resp, r, err := api_client.AnalyticsApi.DeleteApplication(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.DeleteApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


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


## DeleteMapReduceInstanceRun

> DeleteMapReduceInstanceRun(ctx, id).Execute()

Delete a Map-Reduce Application Instance Run.



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
    id := int64(789) // int64 | 

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

    request := cohesitysdk.ApiDeleteMapReduceInstanceRunRequest{
        Id: &id
    }

    resp, r, err := api_client.AnalyticsApi.DeleteMapReduceInstanceRun(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.DeleteMapReduceInstanceRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMapReduceInstanceRunRequest struct via the builder pattern


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


## DeleteMapper

> DeleteMapper(ctx, id).Execute()

Delete a Mapper.



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
    id := int64(789) // int64 | 

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

    request := cohesitysdk.ApiDeleteMapperRequest{
        Id: &id
    }

    resp, r, err := api_client.AnalyticsApi.DeleteMapper(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.DeleteMapper``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMapperRequest struct via the builder pattern


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


## DeleteReducer

> DeleteReducer(ctx, id).Execute()

Delete a Reducer.



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
    id := int64(789) // int64 | 

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

    request := cohesitysdk.ApiDeleteReducerRequest{
        Id: &id
    }

    resp, r, err := api_client.AnalyticsApi.DeleteReducer(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.DeleteReducer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReducerRequest struct via the builder pattern


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


## DeleteUploadedJar

> DeleteUploadedJar(ctx).Body(body).Execute()

Delete the uploaded jar from temporary location.



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
    body := *openapiclient.NewUploadMRJarViewPathWrapper() // UploadMRJarViewPathWrapper |  (optional)

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

    request := cohesitysdk.ApiDeleteUploadedJarRequest{
        Body: &body
    }

    resp, r, err := api_client.AnalyticsApi.DeleteUploadedJar(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.DeleteUploadedJar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUploadedJarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UploadMRJarViewPathWrapper**](UploadMRJarViewPathWrapper.md) |  | 

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


## DownloadMRBaseJar

> ExtractFileRangeResult DownloadMRBaseJar(ctx).Execute()

Downloads the map reduce base jar.



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

    request := cohesitysdk.ApiDownloadMRBaseJarRequest{
    }

    resp, r, err := api_client.AnalyticsApi.DownloadMRBaseJar(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.DownloadMRBaseJar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadMRBaseJar`: ExtractFileRangeResult
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.DownloadMRBaseJar`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadMRBaseJarRequest struct via the builder pattern


### Return type

[**ExtractFileRangeResult**](ExtractFileRangeResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadMROutputFiles

> ExtractFileRangeResult DownloadMROutputFiles(ctx).IsNfsFile(isNfsFile).PartitionId(partitionId).ViewBoxId(viewBoxId).ViewName(viewName).FilePath(filePath).StartOffset(startOffset).LengthBytes(lengthBytes).Execute()

Download map reduce base instance run output files from Yoda.



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
    isNfsFile := true // bool | If true, then extracts file from NFS, else from local file system. (optional)
    partitionId := int64(789) // int64 | Cluster partition id for the file to be read. (optional)
    viewBoxId := int64(789) // int64 | View box id for the file to be read. Required for nfs files only. (optional)
    viewName := "viewName_example" // string | View name for the file to be read. Required for nfs files only. (optional)
    filePath := "filePath_example" // string | filepath of the file on NFS. (optional)
    startOffset := int64(789) // int64 | start offset from where bytes will be read. (optional)
    lengthBytes := int64(789) // int64 | Number of bytes to be read from start_offset. (optional)

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

    request := cohesitysdk.ApiDownloadMROutputFilesRequest{
        IsNfsFile: &isNfsFile
        PartitionId: &partitionId
        ViewBoxId: &viewBoxId
        ViewName: &viewName
        FilePath: &filePath
        StartOffset: &startOffset
        LengthBytes: &lengthBytes
    }

    resp, r, err := api_client.AnalyticsApi.DownloadMROutputFiles(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.DownloadMROutputFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadMROutputFiles`: ExtractFileRangeResult
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.DownloadMROutputFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadMROutputFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isNfsFile** | **bool** | If true, then extracts file from NFS, else from local file system. | 
 **partitionId** | **int64** | Cluster partition id for the file to be read. | 
 **viewBoxId** | **int64** | View box id for the file to be read. Required for nfs files only. | 
 **viewName** | **string** | View name for the file to be read. Required for nfs files only. | 
 **filePath** | **string** | filepath of the file on NFS. | 
 **startOffset** | **int64** | start offset from where bytes will be read. | 
 **lengthBytes** | **int64** | Number of bytes to be read from start_offset. | 

### Return type

[**ExtractFileRangeResult**](ExtractFileRangeResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationById

> MapReduceInfo GetApplicationById(ctx, id).Execute()

List details about a single Application.



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
    id := int64(789) // int64 | 

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

    request := cohesitysdk.ApiGetApplicationByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.AnalyticsApi.GetApplicationById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetApplicationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationById`: MapReduceInfo
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetApplicationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MapReduceInfo**](MapReduceInfo.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> ApplicationsWrapper GetApplications(ctx).Execute()

List Applications filtered by the specified parameters.



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

    request := cohesitysdk.ApiGetApplicationsRequest{
    }

    resp, r, err := api_client.AnalyticsApi.GetApplications(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplications`: ApplicationsWrapper
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetApplications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsRequest struct via the builder pattern


### Return type

[**ApplicationsWrapper**](ApplicationsWrapper.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMRUploadJarPath

> GetMRJarUploadPathResult GetMRUploadJarPath(ctx).Execute()

Get details about the mounted path to upload Jars.



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

    request := cohesitysdk.ApiGetMRUploadJarPathRequest{
    }

    resp, r, err := api_client.AnalyticsApi.GetMRUploadJarPath(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetMRUploadJarPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMRUploadJarPath`: GetMRJarUploadPathResult
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetMRUploadJarPath`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMRUploadJarPathRequest struct via the builder pattern


### Return type

[**GetMRJarUploadPathResult**](GetMRJarUploadPathResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMapReduceAppRuns

> []AppRunHistory GetMapReduceAppRuns(ctx).Body(body).Execute()

List map reduce application runs filtered by the specified parameters.



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
    body := *openapiclient.NewGetMapReduceAppRunsParams() // GetMapReduceAppRunsParams |  (optional)

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

    request := cohesitysdk.ApiGetMapReduceAppRunsRequest{
        Body: &body
    }

    resp, r, err := api_client.AnalyticsApi.GetMapReduceAppRuns(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetMapReduceAppRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMapReduceAppRuns`: []AppRunHistory
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetMapReduceAppRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMapReduceAppRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GetMapReduceAppRunsParams**](GetMapReduceAppRunsParams.md) |  | 

### Return type

[**[]AppRunHistory**](AppRunHistory.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMapReduceFileFormats

> MapReduceFileFormats GetMapReduceFileFormats(ctx).ResultFilePath(resultFilePath).StorageDomainId(storageDomainId).ViewName(viewName).Execute()

Used to retrieve supported output file formats from a map reduce instance.



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
    resultFilePath := "resultFilePath_example" // string | Specifies the path where the map reduce instance has the result. file generated.
    storageDomainId := int64(789) // int64 | Specifies the ID of the storage domain.
    viewName := "viewName_example" // string | Specifies the name of the view inside view box where map reduce instance. is run.

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

    request := cohesitysdk.ApiGetMapReduceFileFormatsRequest{
        ResultFilePath: &resultFilePath
        StorageDomainId: &storageDomainId
        ViewName: &viewName
    }

    resp, r, err := api_client.AnalyticsApi.GetMapReduceFileFormats(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetMapReduceFileFormats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMapReduceFileFormats`: MapReduceFileFormats
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetMapReduceFileFormats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMapReduceFileFormatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resultFilePath** | **string** | Specifies the path where the map reduce instance has the result. file generated. | 
 **storageDomainId** | **int64** | Specifies the ID of the storage domain. | 
 **viewName** | **string** | Specifies the name of the view inside view box where map reduce instance. is run. | 

### Return type

[**MapReduceFileFormats**](MapReduceFileFormats.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMapperById

> MapperInfo GetMapperById(ctx, id).Execute()

List details about a single Mapper.



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
    id := int64(789) // int64 | 

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

    request := cohesitysdk.ApiGetMapperByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.AnalyticsApi.GetMapperById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetMapperById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMapperById`: MapperInfo
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetMapperById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMapperByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MapperInfo**](MapperInfo.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMappers

> MappersWrapper GetMappers(ctx).Execute()

List Mappers filtered by the specified parameters.



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

    request := cohesitysdk.ApiGetMappersRequest{
    }

    resp, r, err := api_client.AnalyticsApi.GetMappers(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetMappers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMappers`: MappersWrapper
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetMappers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMappersRequest struct via the builder pattern


### Return type

[**MappersWrapper**](MappersWrapper.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReducerById

> ReducerInfo GetReducerById(ctx, id).Execute()

List details about a single Reducer.



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
    id := int64(789) // int64 | 

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

    request := cohesitysdk.ApiGetReducerByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.AnalyticsApi.GetReducerById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetReducerById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReducerById`: ReducerInfo
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetReducerById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReducerByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReducerInfo**](ReducerInfo.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReducers

> ReducersWrapper GetReducers(ctx).Execute()

List Reducers filtered by the specified parameters.



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

    request := cohesitysdk.ApiGetReducersRequest{
    }

    resp, r, err := api_client.AnalyticsApi.GetReducers(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetReducers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReducers`: ReducersWrapper
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetReducers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetReducersRequest struct via the builder pattern


### Return type

[**ReducersWrapper**](ReducersWrapper.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportedPatterns

> []SupportedPattern GetSupportedPatterns(ctx).ApplicationId(applicationId).ApplicationDataType(applicationDataType).Execute()

Used to retrieve supported patterns available for search in an application.



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
    applicationId := int64(789) // int64 | Specifies the application Id.
    applicationDataType := int32(56) // int32 | Specifies the data type for which supported patterns can be fetched. (optional)

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

    request := cohesitysdk.ApiGetSupportedPatternsRequest{
        ApplicationId: &applicationId
        ApplicationDataType: &applicationDataType
    }

    resp, r, err := api_client.AnalyticsApi.GetSupportedPatterns(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetSupportedPatterns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupportedPatterns`: []SupportedPattern
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetSupportedPatterns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportedPatternsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | **int64** | Specifies the application Id. | 
 **applicationDataType** | **int32** | Specifies the data type for which supported patterns can be fetched. | 

### Return type

[**[]SupportedPattern**](SupportedPattern.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunMapReduceAppInstance

> RunMapReduceInstanceResult RunMapReduceAppInstance(ctx).Body(body).Execute()

Run a map-reduce application instance.



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
    body := *openapiclient.NewRunMapReduceParams() // RunMapReduceParams |  (optional)

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

    request := cohesitysdk.ApiRunMapReduceAppInstanceRequest{
        Body: &body
    }

    resp, r, err := api_client.AnalyticsApi.RunMapReduceAppInstance(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.RunMapReduceAppInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunMapReduceAppInstance`: RunMapReduceInstanceResult
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.RunMapReduceAppInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunMapReduceAppInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RunMapReduceParams**](RunMapReduceParams.md) |  | 

### Return type

[**RunMapReduceInstanceResult**](RunMapReduceInstanceResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SavePattern

> SavePattern(ctx).Body(body).Execute()

Used to save user patterns for search in an application.



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
    body := *openapiclient.NewPatternRequestBody(NullableInt32(123), NullableInt64(123), *openapiclient.NewSupportedPattern()) // PatternRequestBody |  (optional)

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

    request := cohesitysdk.ApiSavePatternRequest{
        Body: &body
    }

    resp, r, err := api_client.AnalyticsApi.SavePattern(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.SavePattern``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSavePatternRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PatternRequestBody**](PatternRequestBody.md) |  | 

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


## UpdateApplication

> MapReduceInfo UpdateApplication(ctx, id).Body(body).Execute()

Update an Application.



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
    id := int64(789) // int64 | 
    body := *openapiclient.NewMapReduceInfo() // MapReduceInfo |  (optional)

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

    request := cohesitysdk.ApiUpdateApplicationRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.AnalyticsApi.UpdateApplication(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.UpdateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplication`: MapReduceInfo
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.UpdateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MapReduceInfo**](MapReduceInfo.md) |  | 

### Return type

[**MapReduceInfo**](MapReduceInfo.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMapper

> MapperInfo UpdateMapper(ctx, id).Body(body).Execute()

Update a Mapper.



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
    id := int64(789) // int64 | 
    body := *openapiclient.NewMapperInfo() // MapperInfo |  (optional)

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

    request := cohesitysdk.ApiUpdateMapperRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.AnalyticsApi.UpdateMapper(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.UpdateMapper``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMapper`: MapperInfo
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.UpdateMapper`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMapperRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MapperInfo**](MapperInfo.md) |  | 

### Return type

[**MapperInfo**](MapperInfo.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReducer

> ReducerInfo UpdateReducer(ctx, id).Body(body).Execute()

Update a Reducer.



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
    id := int64(789) // int64 | 
    body := *openapiclient.NewReducerInfo() // ReducerInfo |  (optional)

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

    request := cohesitysdk.ApiUpdateReducerRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.AnalyticsApi.UpdateReducer(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.UpdateReducer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReducer`: ReducerInfo
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.UpdateReducer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReducerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ReducerInfo**](ReducerInfo.md) |  | 

### Return type

[**ReducerInfo**](ReducerInfo.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadJar

> UploadMRJarViewPathWrapper UploadJar(ctx).Execute()

Upload a jar to pre-specified Yoda internal view.



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

    request := cohesitysdk.ApiUploadJarRequest{
    }

    resp, r, err := api_client.AnalyticsApi.UploadJar(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.UploadJar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadJar`: UploadMRJarViewPathWrapper
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.UploadJar`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUploadJarRequest struct via the builder pattern


### Return type

[**UploadMRJarViewPathWrapper**](UploadMRJarViewPathWrapper.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

