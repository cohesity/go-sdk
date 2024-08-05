# \StatsAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterStorageStats**](StatsAPI.md#GetClusterStorageStats) | **Get** /stats/cluster-storage | Get Cluster Storage Stats.
[**GetFilesStats**](StatsAPI.md#GetFilesStats) | **Get** /stats/files | Get Stats of Files.
[**GetProtectionRunsStats**](StatsAPI.md#GetProtectionRunsStats) | **Get** /stats/protection-runs | Get statistics of protection runs.
[**GetTimeSeriesStats**](StatsAPI.md#GetTimeSeriesStats) | **Get** /stats/time-series-stats | Get Time Series Stats.
[**GetTopViewsStats**](StatsAPI.md#GetTopViewsStats) | **Get** /stats/top-views | Get stats for the top views, which are the views with largest value of &#39;stats.valueInLastHours&#39; for a given combination of &#39;metric&#39;, &#39;protocol&#39; &amp; &#39;lastHours&#39; params. The API uses suitable defaults if any of the parameters are not specified.
[**GetViewClientStats**](StatsAPI.md#GetViewClientStats) | **Get** /stats/view-clients | Get Stats of View Clients
[**GetViewsStats**](StatsAPI.md#GetViewsStats) | **Get** /stats/views | Get stats for the top views, which are the views with largest value of &#39;stats.valueInLastHours&#39; for a given combination of &#39;metric&#39;, &#39;protocol&#39; &amp; &#39;lastHours&#39; params. The API uses suitable defaults if any of the parameters are not specified.
[**GetWorkloadStats**](StatsAPI.md#GetWorkloadStats) | **Get** /stats/workload-stats | Get Workload Stats Schema.



## GetClusterStorageStats

> ClusterStorageStats GetClusterStorageStats(ctx).Execute()

Get Cluster Storage Stats.



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
	resp, r, err := apiClient.StatsAPI.GetClusterStorageStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.GetClusterStorageStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterStorageStats`: ClusterStorageStats
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.GetClusterStorageStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterStorageStatsRequest struct via the builder pattern


### Return type

[**ClusterStorageStats**](ClusterStorageStats.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilesStats

> FilesStats GetFilesStats(ctx).EntityType(entityType).Execute()

Get Stats of Files.



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
	entityType := "entityType_example" // string | Specifies the entity type based on which the files stats are calculated. By default stats are calculated based on Cluster (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.GetFilesStats(context.Background()).EntityType(entityType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.GetFilesStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilesStats`: FilesStats
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.GetFilesStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityType** | **string** | Specifies the entity type based on which the files stats are calculated. By default stats are calculated based on Cluster | 

### Return type

[**FilesStats**](FilesStats.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionRunsStats

> GetProtectionRunsStatusResponseBody GetProtectionRunsStats(ctx).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).RunStatus(runStatus).Execute()

Get statistics of protection runs.



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
	startTimeUsecs := int64(789) // int64 | Specify the start time as a Unix epoch Timestamp (in microseconds), only runs executing after this time will be counted. By default it is current time minus a day. (optional)
	endTimeUsecs := int64(789) // int64 | Specify the end time as a Unix epoch Timestamp (in microseconds), only runs executing before this time will be counted. By default it is current time. (optional)
	runStatus := []string{"RunStatus_example"} // []string | Specifies a list of status, runs matching the status will be returned. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Failed' indicates that the run has failed. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.GetProtectionRunsStats(context.Background()).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).RunStatus(runStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.GetProtectionRunsStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtectionRunsStats`: GetProtectionRunsStatusResponseBody
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.GetProtectionRunsStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionRunsStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimeUsecs** | **int64** | Specify the start time as a Unix epoch Timestamp (in microseconds), only runs executing after this time will be counted. By default it is current time minus a day. | 
 **endTimeUsecs** | **int64** | Specify the end time as a Unix epoch Timestamp (in microseconds), only runs executing before this time will be counted. By default it is current time. | 
 **runStatus** | **[]string** | Specifies a list of status, runs matching the status will be returned. &#39;Running&#39; indicates that the run is still running. &#39;Canceled&#39; indicates that the run has been canceled. &#39;Failed&#39; indicates that the run has failed. &#39;Succeeded&#39; indicates that the run has finished successfully. &#39;SucceededWithWarning&#39; indicates that the run finished successfully, but there were some warning messages. | 

### Return type

[**GetProtectionRunsStatusResponseBody**](GetProtectionRunsStatusResponseBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeSeriesStats

> TimeSeriesStats GetTimeSeriesStats(ctx).SchemaName(schemaName).MetricNames(metricNames).StartTimeMsecs(startTimeMsecs).EntityId(entityId).EntityIdList(entityIdList).ProrateDataPoints(prorateDataPoints).IncludeGrowthChange(includeGrowthChange).EndTimeMsecs(endTimeMsecs).RollupFunction(rollupFunction).RollupIntervalSecs(rollupIntervalSecs).Execute()

Get Time Series Stats.



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
	schemaName := "schemaName_example" // string | Specifies the schema name.
	metricNames := []string{"Inner_example"} // []string | Specifies a list of metric names.
	startTimeMsecs := int64(789) // int64 | Specifies the start time of series stats.
	entityId := "entityId_example" // string | Specifies the entity id. (optional)
	entityIdList := []string{"Inner_example"} // []string | Specifies an entity id list represented as a string. The stats result will be the sum over all these entities. Duplicate id's will be ignored. If both EntityIdList and EntityId are specified, EntityId will be ignored. (optional)
	prorateDataPoints := true // bool | Specifies to create pro rated data point for every rollup interval instead of returning the actual raw data points. This should be used only when rollup function is provided. (optional)
	includeGrowthChange := true // bool | Specifies if the response should return the difference of a data point with the previous datapoint. Used for determining the change in growth rate. Datapoint could be +x, 0, -x showing the growth is up, no change or down respectively. (optional)
	endTimeMsecs := int64(789) // int64 | Specifies the end time of series stats, by default it is current time. (optional)
	rollupFunction := "rollupFunction_example" // string | Specifies the rollup function to apply to the data points for the time interval specified by rollupInternalSecs. (optional)
	rollupIntervalSecs := int32(56) // int32 | Specifies the time interval granularity for the specified rollup function. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.GetTimeSeriesStats(context.Background()).SchemaName(schemaName).MetricNames(metricNames).StartTimeMsecs(startTimeMsecs).EntityId(entityId).EntityIdList(entityIdList).ProrateDataPoints(prorateDataPoints).IncludeGrowthChange(includeGrowthChange).EndTimeMsecs(endTimeMsecs).RollupFunction(rollupFunction).RollupIntervalSecs(rollupIntervalSecs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.GetTimeSeriesStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeSeriesStats`: TimeSeriesStats
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.GetTimeSeriesStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeSeriesStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaName** | **string** | Specifies the schema name. | 
 **metricNames** | **[]string** | Specifies a list of metric names. | 
 **startTimeMsecs** | **int64** | Specifies the start time of series stats. | 
 **entityId** | **string** | Specifies the entity id. | 
 **entityIdList** | **[]string** | Specifies an entity id list represented as a string. The stats result will be the sum over all these entities. Duplicate id&#39;s will be ignored. If both EntityIdList and EntityId are specified, EntityId will be ignored. | 
 **prorateDataPoints** | **bool** | Specifies to create pro rated data point for every rollup interval instead of returning the actual raw data points. This should be used only when rollup function is provided. | 
 **includeGrowthChange** | **bool** | Specifies if the response should return the difference of a data point with the previous datapoint. Used for determining the change in growth rate. Datapoint could be +x, 0, -x showing the growth is up, no change or down respectively. | 
 **endTimeMsecs** | **int64** | Specifies the end time of series stats, by default it is current time. | 
 **rollupFunction** | **string** | Specifies the rollup function to apply to the data points for the time interval specified by rollupInternalSecs. | 
 **rollupIntervalSecs** | **int32** | Specifies the time interval granularity for the specified rollup function. | 

### Return type

[**TimeSeriesStats**](TimeSeriesStats.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopViewsStats

> ViewsStats GetTopViewsStats(ctx).Metric(metric).Protocol(protocol).NumTopViews(numTopViews).LastHours(lastHours).Execute()

Get stats for the top views, which are the views with largest value of 'stats.valueInLastHours' for a given combination of 'metric', 'protocol' & 'lastHours' params. The API uses suitable defaults if any of the parameters are not specified.



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
	metric := "metric_example" // string | Specifies the metric to which stats has to be sorted. Defaults to kNumBytesRead. (optional) (default to "kNumBytesRead")
	protocol := "protocol_example" // string | Specifies the protocol to sort. Defaults to kAny. (optional) (default to "kAny")
	numTopViews := int64(789) // int64 | Specifies the number of view for which stats has to be computed. Returned Views will be sorted in descending order based on the 'metric' param. Minimum value has to be 1. Defaults to 100. (optional) (default to 100)
	lastHours := int64(789) // int64 | Specifies the last hours of stats to sort. Defaults to 24. (optional) (default to 24)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.GetTopViewsStats(context.Background()).Metric(metric).Protocol(protocol).NumTopViews(numTopViews).LastHours(lastHours).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.GetTopViewsStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTopViewsStats`: ViewsStats
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.GetTopViewsStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTopViewsStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metric** | **string** | Specifies the metric to which stats has to be sorted. Defaults to kNumBytesRead. | [default to &quot;kNumBytesRead&quot;]
 **protocol** | **string** | Specifies the protocol to sort. Defaults to kAny. | [default to &quot;kAny&quot;]
 **numTopViews** | **int64** | Specifies the number of view for which stats has to be computed. Returned Views will be sorted in descending order based on the &#39;metric&#39; param. Minimum value has to be 1. Defaults to 100. | [default to 100]
 **lastHours** | **int64** | Specifies the last hours of stats to sort. Defaults to 24. | [default to 24]

### Return type

[**ViewsStats**](ViewsStats.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewClientStats

> ViewClientsStats GetViewClientStats(ctx).Metric(metric).NumTopViewClients(numTopViewClients).LastHours(lastHours).Execute()

Get Stats of View Clients



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
	metric := "metric_example" // string | Specifies the metric to which stats has to be sorted. (optional)
	numTopViewClients := int64(789) // int64 | Specifies the number of view clients for which stats has to be computed. Specifying this field will return the Views sorted in the descending order on the metric specified. If specified, minimum value is 1. If not specified, all view clients will be returned. If metric is not specified, this parameter must also not be specified. (optional)
	lastHours := int64(789) // int64 | Specifies the last hours of stats to sort. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.GetViewClientStats(context.Background()).Metric(metric).NumTopViewClients(numTopViewClients).LastHours(lastHours).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.GetViewClientStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetViewClientStats`: ViewClientsStats
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.GetViewClientStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetViewClientStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metric** | **string** | Specifies the metric to which stats has to be sorted. | 
 **numTopViewClients** | **int64** | Specifies the number of view clients for which stats has to be computed. Specifying this field will return the Views sorted in the descending order on the metric specified. If specified, minimum value is 1. If not specified, all view clients will be returned. If metric is not specified, this parameter must also not be specified. | 
 **lastHours** | **int64** | Specifies the last hours of stats to sort. | 

### Return type

[**ViewClientsStats**](ViewClientsStats.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewsStats

> ViewsStats GetViewsStats(ctx).Metric(metric).Protocol(protocol).NumTopViews(numTopViews).LastHours(lastHours).Execute()

Get stats for the top views, which are the views with largest value of 'stats.valueInLastHours' for a given combination of 'metric', 'protocol' & 'lastHours' params. The API uses suitable defaults if any of the parameters are not specified.



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
	metric := "metric_example" // string | Specifies the metric to which stats has to be sorted. Defaults to kNumBytesRead. (optional) (default to "kNumBytesRead")
	protocol := "protocol_example" // string | Specifies the protocol to sort. Defaults to kAny. (optional) (default to "kAny")
	numTopViews := int64(789) // int64 | Specifies the number of view for which stats has to be computed. Returned Views will be sorted in descending order based on the 'metric' param. Minimum value has to be 1. Defaults to 100. (optional) (default to 100)
	lastHours := int64(789) // int64 | Specifies the last hours of stats to sort. Defaults to 24. (optional) (default to 24)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.GetViewsStats(context.Background()).Metric(metric).Protocol(protocol).NumTopViews(numTopViews).LastHours(lastHours).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.GetViewsStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetViewsStats`: ViewsStats
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.GetViewsStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetViewsStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metric** | **string** | Specifies the metric to which stats has to be sorted. Defaults to kNumBytesRead. | [default to &quot;kNumBytesRead&quot;]
 **protocol** | **string** | Specifies the protocol to sort. Defaults to kAny. | [default to &quot;kAny&quot;]
 **numTopViews** | **int64** | Specifies the number of view for which stats has to be computed. Returned Views will be sorted in descending order based on the &#39;metric&#39; param. Minimum value has to be 1. Defaults to 100. | [default to 100]
 **lastHours** | **int64** | Specifies the last hours of stats to sort. Defaults to 24. | [default to 24]

### Return type

[**ViewsStats**](ViewsStats.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkloadStats

> WorkloadStatsSummary GetWorkloadStats(ctx).Execute()

Get Workload Stats Schema.



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
	resp, r, err := apiClient.StatsAPI.GetWorkloadStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.GetWorkloadStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkloadStats`: WorkloadStatsSummary
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.GetWorkloadStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkloadStatsRequest struct via the builder pattern


### Return type

[**WorkloadStatsSummary**](WorkloadStatsSummary.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

