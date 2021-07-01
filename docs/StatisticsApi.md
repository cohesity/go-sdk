# \StatisticsApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEntities**](StatisticsApi.md#GetEntities) | **Get** /public/statistics/entities | Lists the entities for the specified schema.
[**GetEntitiesSchema**](StatisticsApi.md#GetEntitiesSchema) | **Get** /public/statistics/entitiesSchema | List the entity schemas filtered by the specified parameters.
[**GetEntitySchemaByName**](StatisticsApi.md#GetEntitySchemaByName) | **Get** /public/statistics/entitiesSchema/{schemaName} | Get the entity schema for the specified schema.
[**GetTasks**](StatisticsApi.md#GetTasks) | **Get** /public/tasks/status | Gets the progress and status of tasks.
[**GetTimeSeriesSchema**](StatisticsApi.md#GetTimeSeriesSchema) | **Get** /public/statistics/timeSeriesSchema | 
[**GetTimeSeriesStats**](StatisticsApi.md#GetTimeSeriesStats) | **Get** /public/statistics/timeSeriesStats | List a series of data points for an entity of a metric in a schema, during the specified time period.



## GetEntities

> []EntityProto GetEntities(ctx).SchemaName(schemaName).IncludeAggrMetricSources(includeAggrMetricSources).MetricNames(metricNames).MaxEntities(maxEntities).Execute()

Lists the entities for the specified schema.



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
    schemaName := "schemaName_example" // string | Specifies the entity schema to search for entities.
    includeAggrMetricSources := true // bool | Specifies whether to include the sources of aggregate metrics of an entity. (optional)
    metricNames := []string{"Inner_example"} // []string | Specifies the list of metric names to return such as 'kRandomIos' which corresponds to 'Random IOs' in Advanced Diagnostics of the Cohesity Dashboard. (optional)
    maxEntities := int32(56) // int32 | Specifies the maximum entities returned in the result. By default this field is 500. (optional)

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

    request := cohesitysdk.ApiGetEntitiesRequest{
        schemaName: &SchemaName
        includeAggrMetricSources: &IncludeAggrMetricSources
        metricNames: &MetricNames
        maxEntities: &MaxEntities
    }

    resp, r, err := api_client.StatisticsApi.GetEntities(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticsApi.GetEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntities`: []EntityProto
    fmt.Fprintf(os.Stdout, "Response from `StatisticsApi.GetEntities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaName** | **string** | Specifies the entity schema to search for entities. | 
 **includeAggrMetricSources** | **bool** | Specifies whether to include the sources of aggregate metrics of an entity. | 
 **metricNames** | **[]string** | Specifies the list of metric names to return such as &#39;kRandomIos&#39; which corresponds to &#39;Random IOs&#39; in Advanced Diagnostics of the Cohesity Dashboard. | 
 **maxEntities** | **int32** | Specifies the maximum entities returned in the result. By default this field is 500. | 

### Return type

[**[]EntityProto**](EntityProto.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitiesSchema

> []EntitySchemaProto GetEntitiesSchema(ctx).SchemaNames(schemaNames).MetricNames(metricNames).Execute()

List the entity schemas filtered by the specified parameters.



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
    schemaNames := []string{"Inner_example"} // []string | Specifies the list of schema names to filter by such as 'kIceboxJobVaultStats' which corresponds to 'External Target Job Stats' in Advanced Diagnostics of the Cohesity Dashboard. (optional)
    metricNames := []string{"Inner_example"} // []string | Specifies the list of metric names to filter by such as 'kRandomIos' which corresponds to 'Random IOs' in Advanced Diagnostics of the Cohesity Dashboard. (optional)

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

    request := cohesitysdk.ApiGetEntitiesSchemaRequest{
        schemaNames: &SchemaNames
        metricNames: &MetricNames
    }

    resp, r, err := api_client.StatisticsApi.GetEntitiesSchema(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticsApi.GetEntitiesSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntitiesSchema`: []EntitySchemaProto
    fmt.Fprintf(os.Stdout, "Response from `StatisticsApi.GetEntitiesSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitiesSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaNames** | **[]string** | Specifies the list of schema names to filter by such as &#39;kIceboxJobVaultStats&#39; which corresponds to &#39;External Target Job Stats&#39; in Advanced Diagnostics of the Cohesity Dashboard. | 
 **metricNames** | **[]string** | Specifies the list of metric names to filter by such as &#39;kRandomIos&#39; which corresponds to &#39;Random IOs&#39; in Advanced Diagnostics of the Cohesity Dashboard. | 

### Return type

[**[]EntitySchemaProto**](EntitySchemaProto.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitySchemaByName

> []EntitySchemaProto GetEntitySchemaByName(ctx, schemaName).Execute()

Get the entity schema for the specified schema.



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
    schemaName := "schemaName_example" // string | Name of the Schema

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

    request := cohesitysdk.ApiGetEntitySchemaByNameRequest{
        schemaName: &SchemaName
    }

    resp, r, err := api_client.StatisticsApi.GetEntitySchemaByName(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticsApi.GetEntitySchemaByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntitySchemaByName`: []EntitySchemaProto
    fmt.Fprintf(os.Stdout, "Response from `StatisticsApi.GetEntitySchemaByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaName** | **string** | Name of the Schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitySchemaByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]EntitySchemaProto**](EntitySchemaProto.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasks

> []Task GetTasks(ctx).TaskPaths(taskPaths).IncludeFinishedTasks(includeFinishedTasks).StartTimeSeconds(startTimeSeconds).EndTimeSeconds(endTimeSeconds).MaxTasks(maxTasks).ExcludeSubTasks(excludeSubTasks).Attributes(attributes).Execute()

Gets the progress and status of tasks.

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
    taskPaths := []string{"Inner_example"} // []string | Specifies a list of tasks path by which to filter the results. Otherwise all of the tasks will be returned. (optional)
    includeFinishedTasks := true // bool | Specifies whether or not to include finished tasks. By default, Pulse will only include unfinished tasks. (optional)
    startTimeSeconds := int64(789) // int64 | Specifies a start time by which to filter tasks. Including a value here will result in tasks only being included if they started after the time specified. (optional)
    endTimeSeconds := int64(789) // int64 | Specifies an end time by which to filter tasks. Including a value here will result in tasks only being included if the ended before this time. (optional)
    maxTasks := int32(56) // int32 | Specifies the maximum number of tasks to display. Defaults to 100. (optional)
    excludeSubTasks := true // bool | Specifies whether or not to include subtasks. By default all subtasks of any task matching a query will be returned. (optional)
    attributes := []string{"Inner_example"} // []string | If specified, tasks matching the current query are further filtered by these KeyValuePairs. This gives client an ability to search by custom attributes that they specified during the task creation. Only the Tasks having 'all' of the specified key=value pairs will be returned. Attributes passed in should be separated by commas and each must follow the pattern \"name:type:value\". Valid types are \"kInt64\", \"kDouble\", \"kString\", and \"kBytes\". (optional)

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

    request := cohesitysdk.ApiGetTasksRequest{
        taskPaths: &TaskPaths
        includeFinishedTasks: &IncludeFinishedTasks
        startTimeSeconds: &StartTimeSeconds
        endTimeSeconds: &EndTimeSeconds
        maxTasks: &MaxTasks
        excludeSubTasks: &ExcludeSubTasks
        attributes: &Attributes
    }

    resp, r, err := api_client.StatisticsApi.GetTasks(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticsApi.GetTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTasks`: []Task
    fmt.Fprintf(os.Stdout, "Response from `StatisticsApi.GetTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskPaths** | **[]string** | Specifies a list of tasks path by which to filter the results. Otherwise all of the tasks will be returned. | 
 **includeFinishedTasks** | **bool** | Specifies whether or not to include finished tasks. By default, Pulse will only include unfinished tasks. | 
 **startTimeSeconds** | **int64** | Specifies a start time by which to filter tasks. Including a value here will result in tasks only being included if they started after the time specified. | 
 **endTimeSeconds** | **int64** | Specifies an end time by which to filter tasks. Including a value here will result in tasks only being included if the ended before this time. | 
 **maxTasks** | **int32** | Specifies the maximum number of tasks to display. Defaults to 100. | 
 **excludeSubTasks** | **bool** | Specifies whether or not to include subtasks. By default all subtasks of any task matching a query will be returned. | 
 **attributes** | **[]string** | If specified, tasks matching the current query are further filtered by these KeyValuePairs. This gives client an ability to search by custom attributes that they specified during the task creation. Only the Tasks having &#39;all&#39; of the specified key&#x3D;value pairs will be returned. Attributes passed in should be separated by commas and each must follow the pattern \&quot;name:type:value\&quot;. Valid types are \&quot;kInt64\&quot;, \&quot;kDouble\&quot;, \&quot;kString\&quot;, and \&quot;kBytes\&quot;. | 

### Return type

[**[]Task**](Task.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeSeriesSchema

> TimeSeriesSchemaResponse GetTimeSeriesSchema(ctx).EntityType(entityType).EntityId(entityId).EntityName(entityName).Execute()





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
    entityType := "entityType_example" // string | Specifies the type of the entity. The following entity types are available: cluster, viewbox. EntityType represents the various values for the entity type. 'Cluster' indicates entity type of Cluster. 'StorageDomain' indicates entity type of Storage Domain.
    entityId := int64(789) // int64 | Specifies the id of the entity.
    entityName := "entityName_example" // string | Specifies the name of the entity.

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

    request := cohesitysdk.ApiGetTimeSeriesSchemaRequest{
        entityType: &EntityType
        entityId: &EntityId
        entityName: &EntityName
    }

    resp, r, err := api_client.StatisticsApi.GetTimeSeriesSchema(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticsApi.GetTimeSeriesSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTimeSeriesSchema`: TimeSeriesSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `StatisticsApi.GetTimeSeriesSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeSeriesSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityType** | **string** | Specifies the type of the entity. The following entity types are available: cluster, viewbox. EntityType represents the various values for the entity type. &#39;Cluster&#39; indicates entity type of Cluster. &#39;StorageDomain&#39; indicates entity type of Storage Domain. | 
 **entityId** | **int64** | Specifies the id of the entity. | 
 **entityName** | **string** | Specifies the name of the entity. | 

### Return type

[**TimeSeriesSchemaResponse**](TimeSeriesSchemaResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeSeriesStats

> MetricDataBlock GetTimeSeriesStats(ctx).SchemaName(schemaName).MetricName(metricName).StartTimeMsecs(startTimeMsecs).EntityId(entityId).EntityIdList(entityIdList).EndTimeMsecs(endTimeMsecs).RollupFunction(rollupFunction).RollupIntervalSecs(rollupIntervalSecs).Execute()

List a series of data points for an entity of a metric in a schema, during the specified time period.



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
    schemaName := "schemaName_example" // string | Specifies the name of entity schema such as 'kIceboxJobVaultStats'.
    metricName := "metricName_example" // string | Specifies the name of the metric such as the 'kDiskAwaitTimeMsecs' which is displayed as 'Disk Health' in Advanced Diagnostics of the Cohesity Dashboard.
    startTimeMsecs := int64(789) // int64 | Specifies the start time for the series of metric data. Specify the start time as a Unix epoch Timestamp (in milliseconds).
    entityId := "entityId_example" // string | Specifies the id of the entity represented as a string. Get the entityId by running the GET public/statistics/entities operation. (optional)
    entityIdList := []string{"Inner_example"} // []string | Specifies an entity id list represented as a string. The stats result will be the sum over all these entities. If both EntityIdList and EntityId are specified, EntityId will be ignored. (optional)
    endTimeMsecs := int64(789) // int64 | Specifies the end time for the series of metric data. Specify the end time as a Unix epoch Timestamp (in milliseconds). If not specified, the data points up to the current time are returned. (optional)
    rollupFunction := "rollupFunction_example" // string | Specifies the rollup function to apply to the data points for the time interval specified by rollupInternalSecs. The following rollup functions are available: sum, average, count, max, min, median, percentile95, latest, and rate. For more information about the functions, see the Advanced Diagnostics in the Cohesity online help. If not specified, raw data is returned. (optional)
    rollupIntervalSecs := int32(56) // int32 | Specifies the time interval granularity (in seconds) for the specified rollup function. Only specify a value if Rollup function is also specified. (optional)

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

    request := cohesitysdk.ApiGetTimeSeriesStatsRequest{
        schemaName: &SchemaName
        metricName: &MetricName
        startTimeMsecs: &StartTimeMsecs
        entityId: &EntityId
        entityIdList: &EntityIdList
        endTimeMsecs: &EndTimeMsecs
        rollupFunction: &RollupFunction
        rollupIntervalSecs: &RollupIntervalSecs
    }

    resp, r, err := api_client.StatisticsApi.GetTimeSeriesStats(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticsApi.GetTimeSeriesStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTimeSeriesStats`: MetricDataBlock
    fmt.Fprintf(os.Stdout, "Response from `StatisticsApi.GetTimeSeriesStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeSeriesStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaName** | **string** | Specifies the name of entity schema such as &#39;kIceboxJobVaultStats&#39;. | 
 **metricName** | **string** | Specifies the name of the metric such as the &#39;kDiskAwaitTimeMsecs&#39; which is displayed as &#39;Disk Health&#39; in Advanced Diagnostics of the Cohesity Dashboard. | 
 **startTimeMsecs** | **int64** | Specifies the start time for the series of metric data. Specify the start time as a Unix epoch Timestamp (in milliseconds). | 
 **entityId** | **string** | Specifies the id of the entity represented as a string. Get the entityId by running the GET public/statistics/entities operation. | 
 **entityIdList** | **[]string** | Specifies an entity id list represented as a string. The stats result will be the sum over all these entities. If both EntityIdList and EntityId are specified, EntityId will be ignored. | 
 **endTimeMsecs** | **int64** | Specifies the end time for the series of metric data. Specify the end time as a Unix epoch Timestamp (in milliseconds). If not specified, the data points up to the current time are returned. | 
 **rollupFunction** | **string** | Specifies the rollup function to apply to the data points for the time interval specified by rollupInternalSecs. The following rollup functions are available: sum, average, count, max, min, median, percentile95, latest, and rate. For more information about the functions, see the Advanced Diagnostics in the Cohesity online help. If not specified, raw data is returned. | 
 **rollupIntervalSecs** | **int32** | Specifies the time interval granularity (in seconds) for the specified rollup function. Only specify a value if Rollup function is also specified. | 

### Return type

[**MetricDataBlock**](MetricDataBlock.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

