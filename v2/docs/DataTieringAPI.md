# \DataTieringAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelDataTieringAnalysisGroupRun**](DataTieringAPI.md#CancelDataTieringAnalysisGroupRun) | **Post** /data-tiering/analysis-groups/{id}/runs/{runId}/cancel | Cancel data tiering analysis run.
[**CancelDataTieringTaskRun**](DataTieringAPI.md#CancelDataTieringTaskRun) | **Post** /data-tiering/tasks/{id}/runs/{runId}/cancel | Cancel data tiering task.
[**CreateDataTieringAnalysisGroup**](DataTieringAPI.md#CreateDataTieringAnalysisGroup) | **Post** /data-tiering/analysis-groups | Create a data tiering analysis group.
[**CreateDataTieringAnalysisGroupRun**](DataTieringAPI.md#CreateDataTieringAnalysisGroupRun) | **Post** /data-tiering/analysis-groups/{id}/runs | Create a data tiering analysis group run.
[**CreateDataTieringTask**](DataTieringAPI.md#CreateDataTieringTask) | **Post** /data-tiering/tasks | Create a data tiering task.
[**CreateDataTieringTaskRun**](DataTieringAPI.md#CreateDataTieringTaskRun) | **Post** /data-tiering/tasks/{id}/runs | Create a data tiering tasks run.
[**DeleteDataTieringAnalysisGroup**](DataTieringAPI.md#DeleteDataTieringAnalysisGroup) | **Delete** /data-tiering/analysis-groups/{id} | Delete data tiering analysis group.
[**DeleteDataTieringTask**](DataTieringAPI.md#DeleteDataTieringTask) | **Delete** /data-tiering/tasks/{id} | delete the data tiering task.
[**DownloadTieringReports**](DataTieringAPI.md#DownloadTieringReports) | **Get** /data-tiering/tasks/{id}/runs/{runId}/download-report | Download Tiering reports.
[**GetCapacityTrendAnalysis**](DataTieringAPI.md#GetCapacityTrendAnalysis) | **Get** /data-tiering/capacity-trend | Get capacity trend analysis for all sources or a specific source.
[**GetDataTieringAnalysisGroupById**](DataTieringAPI.md#GetDataTieringAnalysisGroupById) | **Get** /data-tiering/analysis-groups/{id} | Get data tiering analysis group by id.
[**GetDataTieringAnalysisGroupRuns**](DataTieringAPI.md#GetDataTieringAnalysisGroupRuns) | **Get** /data-tiering/analysis-groups/{id}/runs | Get data tiering analysis group runs.
[**GetDataTieringAnalysisGroups**](DataTieringAPI.md#GetDataTieringAnalysisGroups) | **Get** /data-tiering/analysis-groups | Get the list of data tiering analysis groups.
[**GetDataTieringAnalysisGroupsDefaultConfig**](DataTieringAPI.md#GetDataTieringAnalysisGroupsDefaultConfig) | **Get** /data-tiering/analysis-groups/config | Get the default config of data tiering analysis groups.
[**GetDataTieringTaskById**](DataTieringAPI.md#GetDataTieringTaskById) | **Get** /data-tiering/tasks/{id} | Get data tiering task by id.
[**GetDataTieringTasks**](DataTieringAPI.md#GetDataTieringTasks) | **Get** /data-tiering/tasks | Get the list of data tiering tasks.
[**UpdateDataTieringAnalysisGroup**](DataTieringAPI.md#UpdateDataTieringAnalysisGroup) | **Put** /data-tiering/analysis-groups/{id} | Update a data tiering analysis group. Currently, it supports updating sources and schedule only.
[**UpdateDataTieringAnalysisGroupTagsConfig**](DataTieringAPI.md#UpdateDataTieringAnalysisGroupTagsConfig) | **Put** /data-tiering/analysis-groups/{id}/config | Update data tiering analysis group config.
[**UpdateDataTieringAnalysisGroupsState**](DataTieringAPI.md#UpdateDataTieringAnalysisGroupsState) | **Post** /data-tiering/analysis-groups/states | Update data tiering analysis groups state.
[**UpdateDataTieringTask**](DataTieringAPI.md#UpdateDataTieringTask) | **Put** /data-tiering/tasks/{id} | Update a data tiering task.
[**UpdateDataTieringTasksState**](DataTieringAPI.md#UpdateDataTieringTasksState) | **Post** /data-tiering/tasks/states | Update data tiering source analysis tasks state.



## CancelDataTieringAnalysisGroupRun

> CancelDataTieringAnalysisGroupRun(ctx, id, runId).Execute()

Cancel data tiering analysis run.



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
	id := "id_example" // string | Specifies a unique id of data tiering group.
	runId := "runId_example" // string | Specifies a unique run id of data tiering group run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataTieringAPI.CancelDataTieringAnalysisGroupRun(context.Background(), id, runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.CancelDataTieringAnalysisGroupRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of data tiering group. | 
**runId** | **string** | Specifies a unique run id of data tiering group run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelDataTieringAnalysisGroupRunRequest struct via the builder pattern


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


## CancelDataTieringTaskRun

> CancelDataTieringTaskRun(ctx, id, runId).Execute()

Cancel data tiering task.



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
	id := "id_example" // string | Specifies a unique id of data tiering task.
	runId := "runId_example" // string | Specifies a unique run id of data tiering task.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataTieringAPI.CancelDataTieringTaskRun(context.Background(), id, runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.CancelDataTieringTaskRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of data tiering task. | 
**runId** | **string** | Specifies a unique run id of data tiering task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelDataTieringTaskRunRequest struct via the builder pattern


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


## CreateDataTieringAnalysisGroup

> DataTieringAnalysisGroup CreateDataTieringAnalysisGroup(ctx).Body(body).Execute()

Create a data tiering analysis group.



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
	body := *openapiclient.NewCommonDataTieringAnalysisGroupParams("Name_example") // CommonDataTieringAnalysisGroupParams | Specifies the data tiering analysis group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTieringAPI.CreateDataTieringAnalysisGroup(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.CreateDataTieringAnalysisGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDataTieringAnalysisGroup`: DataTieringAnalysisGroup
	fmt.Fprintf(os.Stdout, "Response from `DataTieringAPI.CreateDataTieringAnalysisGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataTieringAnalysisGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CommonDataTieringAnalysisGroupParams**](CommonDataTieringAnalysisGroupParams.md) | Specifies the data tiering analysis group. | 

### Return type

[**DataTieringAnalysisGroup**](DataTieringAnalysisGroup.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDataTieringAnalysisGroupRun

> CreateDataTieringAnalysisGroupRun(ctx, id).Body(body).Execute()

Create a data tiering analysis group run.



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
	id := "id_example" // string | Specifies the id of the data tiering analysis group.
	body := *openapiclient.NewDataTieringAnalysisRunRequest() // DataTieringAnalysisRunRequest | Specifies the request to run analysis group once. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataTieringAPI.CreateDataTieringAnalysisGroupRun(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.CreateDataTieringAnalysisGroupRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the data tiering analysis group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataTieringAnalysisGroupRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DataTieringAnalysisRunRequest**](DataTieringAnalysisRunRequest.md) | Specifies the request to run analysis group once. | 

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


## CreateDataTieringTask

> DataTieringTask CreateDataTieringTask(ctx).Body(body).Execute()

Create a data tiering task.



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
	body := *openapiclient.NewCreateOrUpdateDataTieringTaskRequest("Name_example", "Type_example") // CreateOrUpdateDataTieringTaskRequest | Specifies the parameters to create a data tiering task.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTieringAPI.CreateDataTieringTask(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.CreateDataTieringTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDataTieringTask`: DataTieringTask
	fmt.Fprintf(os.Stdout, "Response from `DataTieringAPI.CreateDataTieringTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataTieringTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateOrUpdateDataTieringTaskRequest**](CreateOrUpdateDataTieringTaskRequest.md) | Specifies the parameters to create a data tiering task. | 

### Return type

[**DataTieringTask**](DataTieringTask.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDataTieringTaskRun

> CreateDataTieringTaskRun(ctx, id).Body(body).Execute()

Create a data tiering tasks run.



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
	id := "id_example" // string | Specifies the id of the data tiering tasks.
	body := *openapiclient.NewDataTieringTaskRunRequest() // DataTieringTaskRunRequest | Specifies the request to run tiering task once. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataTieringAPI.CreateDataTieringTaskRun(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.CreateDataTieringTaskRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the data tiering tasks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataTieringTaskRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DataTieringTaskRunRequest**](DataTieringTaskRunRequest.md) | Specifies the request to run tiering task once. | 

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


## DeleteDataTieringAnalysisGroup

> DeleteDataTieringAnalysisGroup(ctx, id).Execute()

Delete data tiering analysis group.



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
	id := "id_example" // string | Specifies a unique id of the data tiering analysis group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataTieringAPI.DeleteDataTieringAnalysisGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.DeleteDataTieringAnalysisGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the data tiering analysis group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataTieringAnalysisGroupRequest struct via the builder pattern


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


## DeleteDataTieringTask

> DeleteDataTieringTask(ctx, id).Execute()

delete the data tiering task.



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
	id := "id_example" // string | Specifies the id of the data tiering task.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataTieringAPI.DeleteDataTieringTask(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.DeleteDataTieringTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the data tiering task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataTieringTaskRequest struct via the builder pattern


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


## DownloadTieringReports

> DownloadTieringReports(ctx, id, runId).TargetViewName(targetViewName).FilePath(filePath).Execute()

Download Tiering reports.



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
	id := "id_example" // string | Specifies a unique id of data tiering task.
	runId := "runId_example" // string | Specifies a unique run id of data tiering task.
	targetViewName := "targetViewName_example" // string | Specifies the View name from which the tiering job report file should be read from.
	filePath := "filePath_example" // string | Specifies the file path in the targetView.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataTieringAPI.DownloadTieringReports(context.Background(), id, runId).TargetViewName(targetViewName).FilePath(filePath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.DownloadTieringReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of data tiering task. | 
**runId** | **string** | Specifies a unique run id of data tiering task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadTieringReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **targetViewName** | **string** | Specifies the View name from which the tiering job report file should be read from. | 
 **filePath** | **string** | Specifies the file path in the targetView. | 

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


## GetCapacityTrendAnalysis

> CapacityTrendAnalysis GetCapacityTrendAnalysis(ctx).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).SourceId(sourceId).TruncateResponse(truncateResponse).Execute()

Get capacity trend analysis for all sources or a specific source.



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
	startTimeUsecs := int64(789) // int64 | Filter by a start time. Specify the start time as a Unix epoch Timestamp (in microseconds). (optional)
	endTimeUsecs := int64(789) // int64 | Filter by a end time. Specify the end time as a Unix epoch Timestamp (in microseconds). (optional)
	sourceId := int64(789) // int64 | Filter by source id. If specified, this will only return the capacity trend analysis of the specific source. (optional)
	truncateResponse := true // bool | If set, magneto will truncate the response if it exceeds max size limit governed by magneto_http_rpc_response_size_limit_bytes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTieringAPI.GetCapacityTrendAnalysis(context.Background()).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).SourceId(sourceId).TruncateResponse(truncateResponse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.GetCapacityTrendAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCapacityTrendAnalysis`: CapacityTrendAnalysis
	fmt.Fprintf(os.Stdout, "Response from `DataTieringAPI.GetCapacityTrendAnalysis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCapacityTrendAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimeUsecs** | **int64** | Filter by a start time. Specify the start time as a Unix epoch Timestamp (in microseconds). | 
 **endTimeUsecs** | **int64** | Filter by a end time. Specify the end time as a Unix epoch Timestamp (in microseconds). | 
 **sourceId** | **int64** | Filter by source id. If specified, this will only return the capacity trend analysis of the specific source. | 
 **truncateResponse** | **bool** | If set, magneto will truncate the response if it exceeds max size limit governed by magneto_http_rpc_response_size_limit_bytes | 

### Return type

[**CapacityTrendAnalysis**](CapacityTrendAnalysis.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTieringAnalysisGroupById

> DataTieringAnalysisGroup GetDataTieringAnalysisGroupById(ctx, id).Execute()

Get data tiering analysis group by id.



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
	id := "id_example" // string | Specifies a unique id of the data tiering analysis group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTieringAPI.GetDataTieringAnalysisGroupById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.GetDataTieringAnalysisGroupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataTieringAnalysisGroupById`: DataTieringAnalysisGroup
	fmt.Fprintf(os.Stdout, "Response from `DataTieringAPI.GetDataTieringAnalysisGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the data tiering analysis group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTieringAnalysisGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataTieringAnalysisGroup**](DataTieringAnalysisGroup.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTieringAnalysisGroupRuns

> DataTieringAnalysisGroupRuns GetDataTieringAnalysisGroupRuns(ctx, id).RunIds(runIds).TruncateResponse(truncateResponse).Execute()

Get data tiering analysis group runs.



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
	id := "id_example" // string | Specifies a unique id of the data tiering analysis group.
	runIds := []string{"Inner_example"} // []string | Filter by a list of analysis group run ids. (optional)
	truncateResponse := true // bool | If set, magneto will truncate the response if it exceeds max size limit governed by magneto_http_rpc_response_size_limit_bytes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTieringAPI.GetDataTieringAnalysisGroupRuns(context.Background(), id).RunIds(runIds).TruncateResponse(truncateResponse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.GetDataTieringAnalysisGroupRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataTieringAnalysisGroupRuns`: DataTieringAnalysisGroupRuns
	fmt.Fprintf(os.Stdout, "Response from `DataTieringAPI.GetDataTieringAnalysisGroupRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the data tiering analysis group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTieringAnalysisGroupRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **runIds** | **[]string** | Filter by a list of analysis group run ids. | 
 **truncateResponse** | **bool** | If set, magneto will truncate the response if it exceeds max size limit governed by magneto_http_rpc_response_size_limit_bytes | 

### Return type

[**DataTieringAnalysisGroupRuns**](DataTieringAnalysisGroupRuns.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTieringAnalysisGroups

> []DataTieringAnalysisGroup GetDataTieringAnalysisGroups(ctx).Ids(ids).IncludeLastRunStats(includeLastRunStats).Execute()

Get the list of data tiering analysis groups.



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
	ids := []string{"Inner_example"} // []string | Filter by a list of Analysis Group IDs. (optional)
	includeLastRunStats := true // bool | If true, the response will include last run info. If it is false or not specified, the last run info won't be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTieringAPI.GetDataTieringAnalysisGroups(context.Background()).Ids(ids).IncludeLastRunStats(includeLastRunStats).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.GetDataTieringAnalysisGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataTieringAnalysisGroups`: []DataTieringAnalysisGroup
	fmt.Fprintf(os.Stdout, "Response from `DataTieringAPI.GetDataTieringAnalysisGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTieringAnalysisGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | Filter by a list of Analysis Group IDs. | 
 **includeLastRunStats** | **bool** | If true, the response will include last run info. If it is false or not specified, the last run info won&#39;t be returned. | 

### Return type

[**[]DataTieringAnalysisGroup**](DataTieringAnalysisGroup.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTieringAnalysisGroupsDefaultConfig

> DataTieringTagConfig GetDataTieringAnalysisGroupsDefaultConfig(ctx).Execute()

Get the default config of data tiering analysis groups.



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
	resp, r, err := apiClient.DataTieringAPI.GetDataTieringAnalysisGroupsDefaultConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.GetDataTieringAnalysisGroupsDefaultConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataTieringAnalysisGroupsDefaultConfig`: DataTieringTagConfig
	fmt.Fprintf(os.Stdout, "Response from `DataTieringAPI.GetDataTieringAnalysisGroupsDefaultConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTieringAnalysisGroupsDefaultConfigRequest struct via the builder pattern


### Return type

[**DataTieringTagConfig**](DataTieringTagConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTieringTaskById

> DataTieringTask GetDataTieringTaskById(ctx, id).Execute()

Get data tiering task by id.



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
	id := "id_example" // string | Specifies the id of the data tiering task.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTieringAPI.GetDataTieringTaskById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.GetDataTieringTaskById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataTieringTaskById`: DataTieringTask
	fmt.Fprintf(os.Stdout, "Response from `DataTieringAPI.GetDataTieringTaskById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the data tiering task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTieringTaskByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataTieringTask**](DataTieringTask.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTieringTasks

> []DataTieringTask GetDataTieringTasks(ctx).Ids(ids).IncludeDowntieredDataLocation(includeDowntieredDataLocation).Execute()

Get the list of data tiering tasks.



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
	ids := []string{"Inner_example"} // []string | Filter by a list of data tiering task ids. (optional)
	includeDowntieredDataLocation := true // bool | If true, it will also return a list of downtiered data locations for downtiered tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTieringAPI.GetDataTieringTasks(context.Background()).Ids(ids).IncludeDowntieredDataLocation(includeDowntieredDataLocation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.GetDataTieringTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataTieringTasks`: []DataTieringTask
	fmt.Fprintf(os.Stdout, "Response from `DataTieringAPI.GetDataTieringTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTieringTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | Filter by a list of data tiering task ids. | 
 **includeDowntieredDataLocation** | **bool** | If true, it will also return a list of downtiered data locations for downtiered tasks. | 

### Return type

[**[]DataTieringTask**](DataTieringTask.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataTieringAnalysisGroup

> DataTieringAnalysisGroup UpdateDataTieringAnalysisGroup(ctx, id).Body(body).Execute()

Update a data tiering analysis group. Currently, it supports updating sources and schedule only.



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
	id := "id_example" // string | Specifies a unique id of the data tiering analysis group.
	body := *openapiclient.NewCommonDataTieringAnalysisGroupParams("Name_example") // CommonDataTieringAnalysisGroupParams | Specifies the data tiering analysis group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTieringAPI.UpdateDataTieringAnalysisGroup(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.UpdateDataTieringAnalysisGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDataTieringAnalysisGroup`: DataTieringAnalysisGroup
	fmt.Fprintf(os.Stdout, "Response from `DataTieringAPI.UpdateDataTieringAnalysisGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the data tiering analysis group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataTieringAnalysisGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CommonDataTieringAnalysisGroupParams**](CommonDataTieringAnalysisGroupParams.md) | Specifies the data tiering analysis group. | 

### Return type

[**DataTieringAnalysisGroup**](DataTieringAnalysisGroup.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataTieringAnalysisGroupTagsConfig

> DataTieringTagConfig UpdateDataTieringAnalysisGroupTagsConfig(ctx, id).Body(body).Execute()

Update data tiering analysis group config.



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
	id := "id_example" // string | Specifies a unique id of the data tiering analysis group.
	body := *openapiclient.NewDataTieringTagConfig([]openapiclient.DataTieringTagObject{*openapiclient.NewDataTieringTagObject()}) // DataTieringTagConfig | Specifies the data tiering analysis Tags Config.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTieringAPI.UpdateDataTieringAnalysisGroupTagsConfig(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.UpdateDataTieringAnalysisGroupTagsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDataTieringAnalysisGroupTagsConfig`: DataTieringTagConfig
	fmt.Fprintf(os.Stdout, "Response from `DataTieringAPI.UpdateDataTieringAnalysisGroupTagsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the data tiering analysis group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataTieringAnalysisGroupTagsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DataTieringTagConfig**](DataTieringTagConfig.md) | Specifies the data tiering analysis Tags Config. | 

### Return type

[**DataTieringTagConfig**](DataTieringTagConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataTieringAnalysisGroupsState

> UpdateDataTieringState UpdateDataTieringAnalysisGroupsState(ctx).Body(body).Execute()

Update data tiering analysis groups state.



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
	body := *openapiclient.NewUpdateDataTieringStateRequest("Action_example", []string{"Ids_example"}) // UpdateDataTieringStateRequest | Specifies the parameters to perform an action of list of data tiering analysis groups.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTieringAPI.UpdateDataTieringAnalysisGroupsState(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.UpdateDataTieringAnalysisGroupsState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDataTieringAnalysisGroupsState`: UpdateDataTieringState
	fmt.Fprintf(os.Stdout, "Response from `DataTieringAPI.UpdateDataTieringAnalysisGroupsState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataTieringAnalysisGroupsStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateDataTieringStateRequest**](UpdateDataTieringStateRequest.md) | Specifies the parameters to perform an action of list of data tiering analysis groups. | 

### Return type

[**UpdateDataTieringState**](UpdateDataTieringState.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataTieringTask

> DataTieringTask UpdateDataTieringTask(ctx, id).Body(body).Execute()

Update a data tiering task.



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
	id := "id_example" // string | Specifies the id of the data tiering task.
	body := *openapiclient.NewCreateOrUpdateDataTieringTaskRequest("Name_example", "Type_example") // CreateOrUpdateDataTieringTaskRequest | Specifies the parameters to update a data tiering task.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTieringAPI.UpdateDataTieringTask(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.UpdateDataTieringTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDataTieringTask`: DataTieringTask
	fmt.Fprintf(os.Stdout, "Response from `DataTieringAPI.UpdateDataTieringTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the data tiering task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataTieringTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateOrUpdateDataTieringTaskRequest**](CreateOrUpdateDataTieringTaskRequest.md) | Specifies the parameters to update a data tiering task. | 

### Return type

[**DataTieringTask**](DataTieringTask.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataTieringTasksState

> UpdateDataTieringState UpdateDataTieringTasksState(ctx).Body(body).Execute()

Update data tiering source analysis tasks state.



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
	body := *openapiclient.NewUpdateDataTieringStateRequest("Action_example", []string{"Ids_example"}) // UpdateDataTieringStateRequest | Specifies the parameters to perform an action of list of data tiering tasks.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataTieringAPI.UpdateDataTieringTasksState(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataTieringAPI.UpdateDataTieringTasksState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDataTieringTasksState`: UpdateDataTieringState
	fmt.Fprintf(os.Stdout, "Response from `DataTieringAPI.UpdateDataTieringTasksState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataTieringTasksStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateDataTieringStateRequest**](UpdateDataTieringStateRequest.md) | Specifies the parameters to perform an action of list of data tiering tasks. | 

### Return type

[**UpdateDataTieringState**](UpdateDataTieringState.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

