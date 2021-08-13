# \AuditLog

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditLogs**](AuditLog.md#GetAuditLogs) | **Get** /audit-logs | Get cluster audit logs.
[**GetAuditLogsActions**](AuditLog.md#GetAuditLogsActions) | **Get** /audit-logs/actions | Get cluster audit logs actions.
[**GetAuditLogsEntityTypes**](AuditLog.md#GetAuditLogsEntityTypes) | **Get** /audit-logs/entity-types | Get cluster audit logs entity types.
[**GetFilerAuditLogConfigs**](AuditLog.md#GetFilerAuditLogConfigs) | **Get** /audit-logs/filer-configs | Get filer audit log configs.
[**UpdateFilerAuditLogConfigs**](AuditLog.md#UpdateFilerAuditLogConfigs) | **Put** /audit-logs/filer-configs | Update filer audit log configs.



## GetAuditLogs

> AuditLogs GetAuditLogs(ctx).AccessClusterId(accessClusterId).RegionId(regionId).SearchString(searchString).Usernames(usernames).Domains(domains).EntityTypes(entityTypes).Actions(actions).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).TenantIds(tenantIds).IncludeTenants(includeTenants).StartIndex(startIndex).Count(count).Execute()

Get cluster audit logs.



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
    searchString := "searchString_example" // string | Search audit logs by 'entityName' or 'details'. (optional)
    usernames := []string{"Inner_example"} // []string | Specifies a list of usernames, only audit logs made by these users will be returned. (optional)
    domains := []string{"Inner_example"} // []string | Specifies a list of domains, only audit logs made by user in these domains will be returned. (optional)
    entityTypes := []string{"EntityTypes_example"} // []string | Specifies a list of entity types, only audit logs containing these entity types will be returned. (optional)
    actions := []string{"Actions_example"} // []string | Specifies a list of actions, only audit logs containing these actions will be returned. (optional)
    startTimeUsecs := int64(789) // int64 | Specifies a unix timestamp in microseconds, only audit logs made after this time will be returned. (optional)
    endTimeUsecs := int64(789) // int64 | Specifies a unix timestamp in microseconds, only audit logs made before this time will be returned. (optional)
    tenantIds := []string{"Inner_example"} // []string | Specifies a list of tenant ids, only audit logs made by these tenants will be returned. (optional)
    includeTenants := true // bool | If true, the response will include Protection Groups which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. (optional)
    startIndex := int64(789) // int64 | Specifies a start index. The oldest logs before this index will skipped, only audit logs from this index will be fetched. (optional)
    count := int64(789) // int64 | Specifies the number of indexed obejcts to be fetched from the specified start index. (optional)

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

    request := helios.ApiGetAuditLogsRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        SearchString: &searchString
        Usernames: &usernames
        Domains: &domains
        EntityTypes: &entityTypes
        Actions: &actions
        StartTimeUsecs: &startTimeUsecs
        EndTimeUsecs: &endTimeUsecs
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
        StartIndex: &startIndex
        Count: &count
    }

    resp, r, err := api_client.AuditLog.GetAuditLogs(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLog.GetAuditLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuditLogs`: AuditLogs
    fmt.Fprintf(os.Stdout, "Response from `AuditLog.GetAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **searchString** | **string** | Search audit logs by &#39;entityName&#39; or &#39;details&#39;. | 
 **usernames** | **[]string** | Specifies a list of usernames, only audit logs made by these users will be returned. | 
 **domains** | **[]string** | Specifies a list of domains, only audit logs made by user in these domains will be returned. | 
 **entityTypes** | **[]string** | Specifies a list of entity types, only audit logs containing these entity types will be returned. | 
 **actions** | **[]string** | Specifies a list of actions, only audit logs containing these actions will be returned. | 
 **startTimeUsecs** | **int64** | Specifies a unix timestamp in microseconds, only audit logs made after this time will be returned. | 
 **endTimeUsecs** | **int64** | Specifies a unix timestamp in microseconds, only audit logs made before this time will be returned. | 
 **tenantIds** | **[]string** | Specifies a list of tenant ids, only audit logs made by these tenants will be returned. | 
 **includeTenants** | **bool** | If true, the response will include Protection Groups which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. | 
 **startIndex** | **int64** | Specifies a start index. The oldest logs before this index will skipped, only audit logs from this index will be fetched. | 
 **count** | **int64** | Specifies the number of indexed obejcts to be fetched from the specified start index. | 

### Return type

[**AuditLogs**](AuditLogs.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditLogsActions

> AuditLogsActions GetAuditLogsActions(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get cluster audit logs actions.



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

    request := helios.ApiGetAuditLogsActionsRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.AuditLog.GetAuditLogsActions(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLog.GetAuditLogsActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuditLogsActions`: AuditLogsActions
    fmt.Fprintf(os.Stdout, "Response from `AuditLog.GetAuditLogsActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogsActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**AuditLogsActions**](AuditLogsActions.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditLogsEntityTypes

> AuditLogsEntityTypes GetAuditLogsEntityTypes(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get cluster audit logs entity types.



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

    request := helios.ApiGetAuditLogsEntityTypesRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.AuditLog.GetAuditLogsEntityTypes(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLog.GetAuditLogsEntityTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuditLogsEntityTypes`: AuditLogsEntityTypes
    fmt.Fprintf(os.Stdout, "Response from `AuditLog.GetAuditLogsEntityTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogsEntityTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**AuditLogsEntityTypes**](AuditLogsEntityTypes.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilerAuditLogConfigs

> FilerAuditLogConfigs GetFilerAuditLogConfigs(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get filer audit log configs.



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

    request := helios.ApiGetFilerAuditLogConfigsRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.AuditLog.GetFilerAuditLogConfigs(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLog.GetFilerAuditLogConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFilerAuditLogConfigs`: FilerAuditLogConfigs
    fmt.Fprintf(os.Stdout, "Response from `AuditLog.GetFilerAuditLogConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilerAuditLogConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**FilerAuditLogConfigs**](FilerAuditLogConfigs.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFilerAuditLogConfigs

> FilerAuditLogConfigs UpdateFilerAuditLogConfigs(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update filer audit log configs.



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
    body := *openapiclient.NewFilerAuditLogConfigs() // FilerAuditLogConfigs | Specifies the filer audit log config to update.
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

    request := helios.ApiUpdateFilerAuditLogConfigsRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.AuditLog.UpdateFilerAuditLogConfigs(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLog.UpdateFilerAuditLogConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFilerAuditLogConfigs`: FilerAuditLogConfigs
    fmt.Fprintf(os.Stdout, "Response from `AuditLog.UpdateFilerAuditLogConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFilerAuditLogConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FilerAuditLogConfigs**](FilerAuditLogConfigs.md) | Specifies the filer audit log config to update. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**FilerAuditLogConfigs**](FilerAuditLogConfigs.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

