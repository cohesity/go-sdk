# \StatsApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActiveAlertsStats**](StatsApi.md#GetActiveAlertsStats) | **Get** /public/stats/alerts | Compute the statistics on the active Alerts generated on the cluster.
[**GetConsumerStats**](StatsApi.md#GetConsumerStats) | **Get** /public/stats/consumers | Gets the statistics of consumers.
[**GetFileDistributionStats**](StatsApi.md#GetFileDistributionStats) | **Get** /public/stats/files | Compute the file distribution statistics on a given entity like cluster or storage domain.
[**GetLastProtectionRunStats**](StatsApi.md#GetLastProtectionRunStats) | **Get** /public/stats/protectionRuns/lastRun | Compute stats on last Protection Run for every job.
[**GetProtectedObjectsSummary**](StatsApi.md#GetProtectedObjectsSummary) | **Get** /public/stats/protectionSummary | Computes the protected objects summary on the cluster.
[**GetProtectionRunsStats**](StatsApi.md#GetProtectionRunsStats) | **Get** /public/stats/protectionRuns | Compute the statistics on the Protection Runs for the cluster.
[**GetRestoreStats**](StatsApi.md#GetRestoreStats) | **Get** /public/stats/restores | Compute the statistics on the Restore tasks on the cluster.
[**GetStorageStats**](StatsApi.md#GetStorageStats) | **Get** /public/stats/storage | Computes the storage stats on the cluster.
[**GetTenantStats**](StatsApi.md#GetTenantStats) | **Get** /public/stats/tenants | Gets the statistics of organizations (tenant).
[**GetVaultProviderStats**](StatsApi.md#GetVaultProviderStats) | **Get** /public/stats/vaults/providers | Compute the size and count of entities on vaults.
[**GetVaultRunStats**](StatsApi.md#GetVaultRunStats) | **Get** /public/stats/vaults/runs | Compute the statistics on protection runs to or from a vault.
[**GetVaultStats**](StatsApi.md#GetVaultStats) | **Get** /public/stats/vaults | Compute the statistics on vaults.
[**GetViewBoxStats**](StatsApi.md#GetViewBoxStats) | **Get** /public/stats/viewBoxes | Gets the statistics of view boxes (storage domain).
[**GetViewProtocolStats**](StatsApi.md#GetViewProtocolStats) | **Get** /public/stats/views/protocols | Compute the protocol statistics on Views.
[**GetViewStats**](StatsApi.md#GetViewStats) | **Get** /public/stats/views | Compute the statistics on Views.



## GetActiveAlertsStats

> ActiveAlertsStats GetActiveAlertsStats(ctx).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).Execute()

Compute the statistics on the active Alerts generated on the cluster.



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
    startTimeUsecs := int64(789) // int64 | Specifies the start time Unix time epoch in microseconds from which the active alerts stats are computed.
    endTimeUsecs := int64(789) // int64 | Specifies the end time Unix time epoch in microseconds to which the active alerts stats are computed.

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

    request := cohesitysdk.ApiGetActiveAlertsStatsRequest{
        StartTimeUsecs: &startTimeUsecs
        EndTimeUsecs: &endTimeUsecs
    }

    resp, r, err := api_client.StatsApi.GetActiveAlertsStats(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetActiveAlertsStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveAlertsStats`: ActiveAlertsStats
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetActiveAlertsStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveAlertsStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimeUsecs** | **int64** | Specifies the start time Unix time epoch in microseconds from which the active alerts stats are computed. | 
 **endTimeUsecs** | **int64** | Specifies the end time Unix time epoch in microseconds to which the active alerts stats are computed. | 

### Return type

[**ActiveAlertsStats**](ActiveAlertsStats.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsumerStats

> GetConsumerStatsResult GetConsumerStats(ctx).ConsumerType(consumerType).MaxCount(maxCount).Cookie(cookie).ConsumerIdList(consumerIdList).ConsumerEntityIdList(consumerEntityIdList).FetchViewBoxName(fetchViewBoxName).FetchTenantName(fetchTenantName).FetchProtectionPolicy(fetchProtectionPolicy).FetchProtectionEnvironment(fetchProtectionEnvironment).ViewBoxesIdList(viewBoxesIdList).OrganizationsIdList(organizationsIdList).TenantIds(tenantIds).IncludeServiceProvider(includeServiceProvider).Execute()

Gets the statistics of consumers.

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
    consumerType := "consumerType_example" // string | Specifies the consumer type. Type of the consumer can be one of the following three,  'kViews', indicates the stats info of Views used per organization (tenant) per view box (storage domain). 'kProtectionRuns', indicates the stats info of Protection Runs used per organization (tenant) per view box (storage domain). 'kReplicationRuns', indicates the stats info of Replication In used per organization (tenant) per view box (storage domain). 'kViewProtectionRuns', indicates the stats info of View Protection Runs used per organization (tenant) per view box (storage domain). (optional)
    maxCount := int64(789) // int64 | Specifies a limit on the number of stats groups returned. (optional)
    cookie := "cookie_example" // string | Specifies the opaque string returned in the previous response. If this is set, next set of active opens just after the previous response are returned. If this is not set, first set of active opens are returned. (optional)
    consumerIdList := []int64{int64(123)} // []int64 | Specifies a list of consumer ids. (optional)
    consumerEntityIdList := []string{"Inner_example"} // []string | Specifies a list of consumer entity ids. If this field is specified, each entity id must corresponds to the id in 'ConsumerIdList' in the same index, and the length of 'ConsumerEntityIdList' and 'ConsumerIdList' must be the same. (optional)
    fetchViewBoxName := true // bool | Specifies whether to fetch view box (storage domain) name for each consumer. (optional)
    fetchTenantName := true // bool | Specifies whether to fetch tenant (organization) name for each consumer. (optional)
    fetchProtectionPolicy := true // bool | Specifies whether to fetch protection policy for each consumer. This field is applicable only if 'consumerType' is 'kProtectionRuns' or 'kReplicationRuns'. (optional)
    fetchProtectionEnvironment := true // bool | Specifies whether to fetch protection environment for each consumer. This field is applicable only if 'consumerType' is 'kProtectionRuns' or 'kReplicationRuns'. (optional)
    viewBoxesIdList := []int64{int64(123)} // []int64 | Specifies a list of view boxes (storage domain) id. (optional)
    organizationsIdList := []string{"Inner_example"} // []string | Specifies a list of organizations (tenant) id. (optional)
    tenantIds := []string{"Inner_example"} // []string | Specifies a list of organizations (tenant) id. This field is added to allow tenantIds json tag. This list will be concatenated with TenantsIdList to form full list of tenantsIdList. (optional)
    includeServiceProvider := true // bool | Specifies whether to fetch the consumption of external service providers. These information will be listed as a unique organization (tenant) in response. By default it is false. (optional)

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

    request := cohesitysdk.ApiGetConsumerStatsRequest{
        ConsumerType: &consumerType
        MaxCount: &maxCount
        Cookie: &cookie
        ConsumerIdList: &consumerIdList
        ConsumerEntityIdList: &consumerEntityIdList
        FetchViewBoxName: &fetchViewBoxName
        FetchTenantName: &fetchTenantName
        FetchProtectionPolicy: &fetchProtectionPolicy
        FetchProtectionEnvironment: &fetchProtectionEnvironment
        ViewBoxesIdList: &viewBoxesIdList
        OrganizationsIdList: &organizationsIdList
        TenantIds: &tenantIds
        IncludeServiceProvider: &includeServiceProvider
    }

    resp, r, err := api_client.StatsApi.GetConsumerStats(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetConsumerStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsumerStats`: GetConsumerStatsResult
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetConsumerStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConsumerStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consumerType** | **string** | Specifies the consumer type. Type of the consumer can be one of the following three,  &#39;kViews&#39;, indicates the stats info of Views used per organization (tenant) per view box (storage domain). &#39;kProtectionRuns&#39;, indicates the stats info of Protection Runs used per organization (tenant) per view box (storage domain). &#39;kReplicationRuns&#39;, indicates the stats info of Replication In used per organization (tenant) per view box (storage domain). &#39;kViewProtectionRuns&#39;, indicates the stats info of View Protection Runs used per organization (tenant) per view box (storage domain). | 
 **maxCount** | **int64** | Specifies a limit on the number of stats groups returned. | 
 **cookie** | **string** | Specifies the opaque string returned in the previous response. If this is set, next set of active opens just after the previous response are returned. If this is not set, first set of active opens are returned. | 
 **consumerIdList** | **[]int64** | Specifies a list of consumer ids. | 
 **consumerEntityIdList** | **[]string** | Specifies a list of consumer entity ids. If this field is specified, each entity id must corresponds to the id in &#39;ConsumerIdList&#39; in the same index, and the length of &#39;ConsumerEntityIdList&#39; and &#39;ConsumerIdList&#39; must be the same. | 
 **fetchViewBoxName** | **bool** | Specifies whether to fetch view box (storage domain) name for each consumer. | 
 **fetchTenantName** | **bool** | Specifies whether to fetch tenant (organization) name for each consumer. | 
 **fetchProtectionPolicy** | **bool** | Specifies whether to fetch protection policy for each consumer. This field is applicable only if &#39;consumerType&#39; is &#39;kProtectionRuns&#39; or &#39;kReplicationRuns&#39;. | 
 **fetchProtectionEnvironment** | **bool** | Specifies whether to fetch protection environment for each consumer. This field is applicable only if &#39;consumerType&#39; is &#39;kProtectionRuns&#39; or &#39;kReplicationRuns&#39;. | 
 **viewBoxesIdList** | **[]int64** | Specifies a list of view boxes (storage domain) id. | 
 **organizationsIdList** | **[]string** | Specifies a list of organizations (tenant) id. | 
 **tenantIds** | **[]string** | Specifies a list of organizations (tenant) id. This field is added to allow tenantIds json tag. This list will be concatenated with TenantsIdList to form full list of tenantsIdList. | 
 **includeServiceProvider** | **bool** | Specifies whether to fetch the consumption of external service providers. These information will be listed as a unique organization (tenant) in response. By default it is false. | 

### Return type

[**GetConsumerStatsResult**](GetConsumerStatsResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileDistributionStats

> []FileDistributionStats GetFileDistributionStats(ctx).EntityType(entityType).Execute()

Compute the file distribution statistics on a given entity like cluster or storage domain.



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
    entityType := "entityType_example" // string | Specifies the entity type on which file distribution stats are computed.

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

    request := cohesitysdk.ApiGetFileDistributionStatsRequest{
        EntityType: &entityType
    }

    resp, r, err := api_client.StatsApi.GetFileDistributionStats(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetFileDistributionStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFileDistributionStats`: []FileDistributionStats
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetFileDistributionStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFileDistributionStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityType** | **string** | Specifies the entity type on which file distribution stats are computed. | 

### Return type

[**[]FileDistributionStats**](FileDistributionStats.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastProtectionRunStats

> LastProtectionRunStats GetLastProtectionRunStats(ctx).Execute()

Compute stats on last Protection Run for every job.



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

    request := cohesitysdk.ApiGetLastProtectionRunStatsRequest{
    }

    resp, r, err := api_client.StatsApi.GetLastProtectionRunStats(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetLastProtectionRunStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLastProtectionRunStats`: LastProtectionRunStats
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetLastProtectionRunStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLastProtectionRunStatsRequest struct via the builder pattern


### Return type

[**LastProtectionRunStats**](LastProtectionRunStats.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectedObjectsSummary

> ProtectedObjectsSummary GetProtectedObjectsSummary(ctx).ExcludeTypes(excludeTypes).Execute()

Computes the protected objects summary on the cluster.



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
    excludeTypes := []string{"ExcludeTypes_example"} // []string | Specifies the environment types to exclude. (optional)

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

    request := cohesitysdk.ApiGetProtectedObjectsSummaryRequest{
        ExcludeTypes: &excludeTypes
    }

    resp, r, err := api_client.StatsApi.GetProtectedObjectsSummary(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetProtectedObjectsSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectedObjectsSummary`: ProtectedObjectsSummary
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetProtectedObjectsSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectedObjectsSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **excludeTypes** | **[]string** | Specifies the environment types to exclude. | 

### Return type

[**ProtectedObjectsSummary**](ProtectedObjectsSummary.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionRunsStats

> ProtectionRunsStats GetProtectionRunsStats(ctx).Status(status).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).Execute()

Compute the statistics on the Protection Runs for the cluster.



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
    status := "status_example" // string | Specifies the Protection Run status for which stats has to be computed.
    startTimeUsecs := int64(789) // int64 | Specifies the start time in Unix timestamp epoch in microseconds where the end time of the Protection Run is greater than the specified value.
    endTimeUsecs := int64(789) // int64 | Specifies the end time in Unix timestamp epoch in microseconds where the end time of the Protection Run is lesser than the specified value.

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

    request := cohesitysdk.ApiGetProtectionRunsStatsRequest{
        Status: &status
        StartTimeUsecs: &startTimeUsecs
        EndTimeUsecs: &endTimeUsecs
    }

    resp, r, err := api_client.StatsApi.GetProtectionRunsStats(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetProtectionRunsStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionRunsStats`: ProtectionRunsStats
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetProtectionRunsStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionRunsStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Specifies the Protection Run status for which stats has to be computed. | 
 **startTimeUsecs** | **int64** | Specifies the start time in Unix timestamp epoch in microseconds where the end time of the Protection Run is greater than the specified value. | 
 **endTimeUsecs** | **int64** | Specifies the end time in Unix timestamp epoch in microseconds where the end time of the Protection Run is lesser than the specified value. | 

### Return type

[**ProtectionRunsStats**](ProtectionRunsStats.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRestoreStats

> RestoreStats GetRestoreStats(ctx).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).Execute()

Compute the statistics on the Restore tasks on the cluster.



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
    startTimeUsecs := int64(789) // int64 | Specifies the start time Unix time epoch in microseconds from which the restore stats are computed.
    endTimeUsecs := int64(789) // int64 | Specifies the end time Unix time epoch in microseconds to which the restore stats are computed.

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

    request := cohesitysdk.ApiGetRestoreStatsRequest{
        StartTimeUsecs: &startTimeUsecs
        EndTimeUsecs: &endTimeUsecs
    }

    resp, r, err := api_client.StatsApi.GetRestoreStats(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetRestoreStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRestoreStats`: RestoreStats
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetRestoreStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRestoreStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimeUsecs** | **int64** | Specifies the start time Unix time epoch in microseconds from which the restore stats are computed. | 
 **endTimeUsecs** | **int64** | Specifies the end time Unix time epoch in microseconds to which the restore stats are computed. | 

### Return type

[**RestoreStats**](RestoreStats.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageStats

> StorageStats GetStorageStats(ctx).Execute()

Computes the storage stats on the cluster.



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

    request := cohesitysdk.ApiGetStorageStatsRequest{
    }

    resp, r, err := api_client.StatsApi.GetStorageStats(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetStorageStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorageStats`: StorageStats
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetStorageStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageStatsRequest struct via the builder pattern


### Return type

[**StorageStats**](StorageStats.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantStats

> GetTenantStatsResult GetTenantStats(ctx).ConsumerType(consumerType).SkipGroupByTenant(skipGroupByTenant).MaxCount(maxCount).Cookie(cookie).OutputFormat(outputFormat).ViewBoxesIdList(viewBoxesIdList).OrganizationsIdList(organizationsIdList).TenantIds(tenantIds).IncludeServiceProvider(includeServiceProvider).Execute()

Gets the statistics of organizations (tenant).

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
    consumerType := "consumerType_example" // string | Specifies the consumer type. Type of the consumer can be one of the following three,  'kViews', indicates the stats info of Views used per organization (tenant) per view box (storage domain). 'kProtectionRuns', indicates the stats info of Protection Runs used per organization (tenant) per view box (storage domain). 'kReplicationRuns', indicates the stats info of Replication In used per organization (tenant) per view box (storage domain). 'kViewProtectionRuns', indicates the stats info of View Protection Runs used per organization (tenant) per view box (storage domain). (optional)
    skipGroupByTenant := true // bool | Specifies if we should skip grouping by tenant. If false, and tenant has multiple storage domains, then the stats for the storage domains will be aggregated. If true, then the response will return each storage domain cross tenant independently. (optional)
    maxCount := int64(789) // int64 | Specifies a limit on the number of stats groups returned. (optional)
    cookie := "cookie_example" // string | Specifies the opaque string returned in the previous response. If this is set, next set of active opens just after the previous response are returned. If this is not set, first set of active opens are returned. (optional)
    outputFormat := "outputFormat_example" // string | Specifies what format to return the data in. Defaults to proto, but can select other options like csv. (optional)
    viewBoxesIdList := []int64{int64(123)} // []int64 | Specifies a list of view boxes (storage domain) id. (optional)
    organizationsIdList := []string{"Inner_example"} // []string | Specifies a list of organizations (tenant) id. (optional)
    tenantIds := []string{"Inner_example"} // []string | Specifies a list of organizations (tenant) id. This field is added to allow tenantIds json tag. This list will be concatenated with TenantsIdList to form full list of tenantsIdList. (optional)
    includeServiceProvider := true // bool | Specifies whether to fetch the consumption of external service providers. These information will be listed as a unique organization (tenant) in response. By default it is false. (optional)

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

    request := cohesitysdk.ApiGetTenantStatsRequest{
        ConsumerType: &consumerType
        SkipGroupByTenant: &skipGroupByTenant
        MaxCount: &maxCount
        Cookie: &cookie
        OutputFormat: &outputFormat
        ViewBoxesIdList: &viewBoxesIdList
        OrganizationsIdList: &organizationsIdList
        TenantIds: &tenantIds
        IncludeServiceProvider: &includeServiceProvider
    }

    resp, r, err := api_client.StatsApi.GetTenantStats(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetTenantStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantStats`: GetTenantStatsResult
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetTenantStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consumerType** | **string** | Specifies the consumer type. Type of the consumer can be one of the following three,  &#39;kViews&#39;, indicates the stats info of Views used per organization (tenant) per view box (storage domain). &#39;kProtectionRuns&#39;, indicates the stats info of Protection Runs used per organization (tenant) per view box (storage domain). &#39;kReplicationRuns&#39;, indicates the stats info of Replication In used per organization (tenant) per view box (storage domain). &#39;kViewProtectionRuns&#39;, indicates the stats info of View Protection Runs used per organization (tenant) per view box (storage domain). | 
 **skipGroupByTenant** | **bool** | Specifies if we should skip grouping by tenant. If false, and tenant has multiple storage domains, then the stats for the storage domains will be aggregated. If true, then the response will return each storage domain cross tenant independently. | 
 **maxCount** | **int64** | Specifies a limit on the number of stats groups returned. | 
 **cookie** | **string** | Specifies the opaque string returned in the previous response. If this is set, next set of active opens just after the previous response are returned. If this is not set, first set of active opens are returned. | 
 **outputFormat** | **string** | Specifies what format to return the data in. Defaults to proto, but can select other options like csv. | 
 **viewBoxesIdList** | **[]int64** | Specifies a list of view boxes (storage domain) id. | 
 **organizationsIdList** | **[]string** | Specifies a list of organizations (tenant) id. | 
 **tenantIds** | **[]string** | Specifies a list of organizations (tenant) id. This field is added to allow tenantIds json tag. This list will be concatenated with TenantsIdList to form full list of tenantsIdList. | 
 **includeServiceProvider** | **bool** | Specifies whether to fetch the consumption of external service providers. These information will be listed as a unique organization (tenant) in response. By default it is false. | 

### Return type

[**GetTenantStatsResult**](GetTenantStatsResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVaultProviderStats

> []VaultProviderStatsInfo GetVaultProviderStats(ctx).RunType(runType).Execute()

Compute the size and count of entities on vaults.



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
    runType := "runType_example" // string | Specifies the type of the runs.

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

    request := cohesitysdk.ApiGetVaultProviderStatsRequest{
        RunType: &runType
    }

    resp, r, err := api_client.StatsApi.GetVaultProviderStats(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetVaultProviderStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultProviderStats`: []VaultProviderStatsInfo
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetVaultProviderStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultProviderStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runType** | **string** | Specifies the type of the runs. | 

### Return type

[**[]VaultProviderStatsInfo**](VaultProviderStatsInfo.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVaultRunStats

> VaultRunStatsSummary GetVaultRunStats(ctx).RunType(runType).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).Interval(interval).Execute()

Compute the statistics on protection runs to or from a vault.



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
    runType := "runType_example" // string | Specifies the type of the run.
    startTimeUsecs := int64(789) // int64 | Specifies the start time Unix time epoch in microseconds from which the vault run stats are computed.
    endTimeUsecs := int64(789) // int64 | Specifies the end time Unix time epoch in microseconds to which the vault run stats are computed.
    interval := "interval_example" // string | Specifies the interval by which runs will be grouped together in the returned trend line.

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

    request := cohesitysdk.ApiGetVaultRunStatsRequest{
        RunType: &runType
        StartTimeUsecs: &startTimeUsecs
        EndTimeUsecs: &endTimeUsecs
        Interval: &interval
    }

    resp, r, err := api_client.StatsApi.GetVaultRunStats(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetVaultRunStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultRunStats`: VaultRunStatsSummary
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetVaultRunStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultRunStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runType** | **string** | Specifies the type of the run. | 
 **startTimeUsecs** | **int64** | Specifies the start time Unix time epoch in microseconds from which the vault run stats are computed. | 
 **endTimeUsecs** | **int64** | Specifies the end time Unix time epoch in microseconds to which the vault run stats are computed. | 
 **interval** | **string** | Specifies the interval by which runs will be grouped together in the returned trend line. | 

### Return type

[**VaultRunStatsSummary**](VaultRunStatsSummary.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVaultStats

> VaultStats GetVaultStats(ctx).Execute()

Compute the statistics on vaults.



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

    request := cohesitysdk.ApiGetVaultStatsRequest{
    }

    resp, r, err := api_client.StatsApi.GetVaultStats(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetVaultStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultStats`: VaultStats
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetVaultStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultStatsRequest struct via the builder pattern


### Return type

[**VaultStats**](VaultStats.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewBoxStats

> GetViewBoxStatsResult GetViewBoxStats(ctx).ViewBoxesIdList(viewBoxesIdList).OrganizationsIdList(organizationsIdList).TenantIds(tenantIds).IncludeServiceProvider(includeServiceProvider).Execute()

Gets the statistics of view boxes (storage domain).

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
    viewBoxesIdList := []int64{int64(123)} // []int64 | Specifies a list of view boxes (storage domain) id. (optional)
    organizationsIdList := []string{"Inner_example"} // []string | Specifies a list of organizations (tenant) id. (optional)
    tenantIds := []string{"Inner_example"} // []string | Specifies a list of organizations (tenant) id. This field is added to allow tenantIds json tag. This list will be concatenated with TenantsIdList to form full list of tenantsIdList. (optional)
    includeServiceProvider := true // bool | Specifies whether to fetch the consumption of external service providers. These information will be listed as a unique organization (tenant) in response. By default it is false. (optional)

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

    request := cohesitysdk.ApiGetViewBoxStatsRequest{
        ViewBoxesIdList: &viewBoxesIdList
        OrganizationsIdList: &organizationsIdList
        TenantIds: &tenantIds
        IncludeServiceProvider: &includeServiceProvider
    }

    resp, r, err := api_client.StatsApi.GetViewBoxStats(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetViewBoxStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetViewBoxStats`: GetViewBoxStatsResult
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetViewBoxStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetViewBoxStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewBoxesIdList** | **[]int64** | Specifies a list of view boxes (storage domain) id. | 
 **organizationsIdList** | **[]string** | Specifies a list of organizations (tenant) id. | 
 **tenantIds** | **[]string** | Specifies a list of organizations (tenant) id. This field is added to allow tenantIds json tag. This list will be concatenated with TenantsIdList to form full list of tenantsIdList. | 
 **includeServiceProvider** | **bool** | Specifies whether to fetch the consumption of external service providers. These information will be listed as a unique organization (tenant) in response. By default it is false. | 

### Return type

[**GetViewBoxStatsResult**](GetViewBoxStatsResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewProtocolStats

> []ViewProtocolStats GetViewProtocolStats(ctx).Execute()

Compute the protocol statistics on Views.



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

    request := cohesitysdk.ApiGetViewProtocolStatsRequest{
    }

    resp, r, err := api_client.StatsApi.GetViewProtocolStats(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetViewProtocolStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetViewProtocolStats`: []ViewProtocolStats
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetViewProtocolStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetViewProtocolStatsRequest struct via the builder pattern


### Return type

[**[]ViewProtocolStats**](ViewProtocolStats.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewStats

> ViewStatsSnapshot GetViewStats(ctx).Metric(metric).NumTopViews(numTopViews).Execute()

Compute the statistics on Views.



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
    metric := "metric_example" // string | Specifies the metric to which stats has to be sorted. (optional)
    numTopViews := int64(789) // int64 | Specifies the number of views for which stats has to be computed. Specifying this field will return the Views sorted in the descending order on the metric specified. If specified, minimum value is 1. If not specified, all views will be returned. If metric is not specified, this parameter must also not be specified. (optional)

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

    request := cohesitysdk.ApiGetViewStatsRequest{
        Metric: &metric
        NumTopViews: &numTopViews
    }

    resp, r, err := api_client.StatsApi.GetViewStats(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetViewStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetViewStats`: ViewStatsSnapshot
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetViewStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetViewStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metric** | **string** | Specifies the metric to which stats has to be sorted. | 
 **numTopViews** | **int64** | Specifies the number of views for which stats has to be computed. Specifying this field will return the Views sorted in the descending order on the metric specified. If specified, minimum value is 1. If not specified, all views will be returned. If metric is not specified, this parameter must also not be specified. | 

### Return type

[**ViewStatsSnapshot**](ViewStatsSnapshot.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

