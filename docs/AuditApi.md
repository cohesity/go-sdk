# \AuditApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditLogsActions**](AuditApi.md#GetAuditLogsActions) | **Get** /public/auditLogs/actions | Return list of audit log actions.
[**GetAuditLogsCategories**](AuditApi.md#GetAuditLogsCategories) | **Get** /public/auditLogs/categories | Return list of audit log categories.
[**SearchClusterAuditLogs**](AuditApi.md#SearchClusterAuditLogs) | **Get** /public/auditLogs/cluster | Lists the cluster audit logs on the Cohesity Cluster that match the filter criteria specified using parameters.



## GetAuditLogsActions

> []string GetAuditLogsActions(ctx).Execute()

Return list of audit log actions.



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

    request := cohesitysdk.ApiGetAuditLogsActionsRequest{
    }

    resp, r, err := api_client.AuditApi.GetAuditLogsActions(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.GetAuditLogsActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuditLogsActions`: []string
    fmt.Fprintf(os.Stdout, "Response from `AuditApi.GetAuditLogsActions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogsActionsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditLogsCategories

> []string GetAuditLogsCategories(ctx).Execute()

Return list of audit log categories.



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

    request := cohesitysdk.ApiGetAuditLogsCategoriesRequest{
    }

    resp, r, err := api_client.AuditApi.GetAuditLogsCategories(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.GetAuditLogsCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuditLogsCategories`: []string
    fmt.Fprintf(os.Stdout, "Response from `AuditApi.GetAuditLogsCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogsCategoriesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchClusterAuditLogs

> ClusterAuditLogsSearchResult SearchClusterAuditLogs(ctx).UserNames(userNames).Domains(domains).EntityTypes(entityTypes).Actions(actions).Search(search).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).StartIndex(startIndex).PageCount(pageCount).OutputFormat(outputFormat).TenantId(tenantId).AllUnderHierarchy(allUnderHierarchy).Execute()

Lists the cluster audit logs on the Cohesity Cluster that match the filter criteria specified using parameters.



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
    userNames := []string{"Inner_example"} // []string | Filter by user names who cause the actions that generate Cluster Audit Logs. (optional)
    domains := []string{"Inner_example"} // []string | Filter by domains of users who cause the actions that trigger Cluster audit logs. (optional)
    entityTypes := []string{"Inner_example"} // []string | Filter by entity types involved in the actions that generate the Cluster audit logs, such as User, Protection Job, View, etc. For a complete list, see the Category drop-down in the Admin > Audit Logs page of the Cohesity Dashboard. (optional)
    actions := []string{"Inner_example"} // []string | Filter by the actions that generate Cluster audit logs such as Activate, Cancel, Clone, Create, etc. For a complete list, see the Actions drop-down in the Admin > Audit Logs page of the Cohesity Dashboard. (optional)
    search := "search_example" // string | Filter by matching a substring in entity name or details of the Cluster audit log. (optional)
    startTimeUsecs := int64(789) // int64 | Filter by a start time. Only Cluster audit logs that were generated after the specified time are returned. Specify the start time as a Unix epoch Timestamp (in microseconds). (optional)
    endTimeUsecs := int64(789) // int64 | Filter by a end time specified as a Unix epoch Timestamp (in microseconds). Only Cluster audit logs that were generated before the specified end time are returned. (optional)
    startIndex := int64(789) // int64 | Specifies an index number that can be used to return subsets of items in multiple requests. Break up the items to return into multiple requests by setting pageCount and startIndex to return a subsets of items in the search result. For example, set startIndex to 0 to get the first set of pageCount items for the first request. Increment startIndex by pageCount to get the next set of pageCount items for a next request. Continue until all items are returned and therefore the total number of returned items is equal to totalCount. Default value is 0. (optional)
    pageCount := int64(789) // int64 | Limit the number of items to return in the response for pagination purposes. Default value is 1000. (optional)
    outputFormat := "outputFormat_example" // string | Specifies the format of the output such as csv and json. If not specified, the json format is returned. If csv is specified, a comma-separated list with a heading row is returned. (optional)
    tenantId := "tenantId_example" // string | TenantId specifies the tenant whose action resulted in the audit log. (optional)
    allUnderHierarchy := true // bool | AllUnderHierarchy specifies if logs of all the tenants under the hierarchy of tenant with id TenantId should be returned. (optional)

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

    request := cohesitysdk.ApiSearchClusterAuditLogsRequest{
        userNames: &UserNames
        domains: &Domains
        entityTypes: &EntityTypes
        actions: &Actions
        search: &Search
        startTimeUsecs: &StartTimeUsecs
        endTimeUsecs: &EndTimeUsecs
        startIndex: &StartIndex
        pageCount: &PageCount
        outputFormat: &OutputFormat
        tenantId: &TenantId
        allUnderHierarchy: &AllUnderHierarchy
    }

    resp, r, err := api_client.AuditApi.SearchClusterAuditLogs(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditApi.SearchClusterAuditLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchClusterAuditLogs`: ClusterAuditLogsSearchResult
    fmt.Fprintf(os.Stdout, "Response from `AuditApi.SearchClusterAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchClusterAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userNames** | **[]string** | Filter by user names who cause the actions that generate Cluster Audit Logs. | 
 **domains** | **[]string** | Filter by domains of users who cause the actions that trigger Cluster audit logs. | 
 **entityTypes** | **[]string** | Filter by entity types involved in the actions that generate the Cluster audit logs, such as User, Protection Job, View, etc. For a complete list, see the Category drop-down in the Admin &gt; Audit Logs page of the Cohesity Dashboard. | 
 **actions** | **[]string** | Filter by the actions that generate Cluster audit logs such as Activate, Cancel, Clone, Create, etc. For a complete list, see the Actions drop-down in the Admin &gt; Audit Logs page of the Cohesity Dashboard. | 
 **search** | **string** | Filter by matching a substring in entity name or details of the Cluster audit log. | 
 **startTimeUsecs** | **int64** | Filter by a start time. Only Cluster audit logs that were generated after the specified time are returned. Specify the start time as a Unix epoch Timestamp (in microseconds). | 
 **endTimeUsecs** | **int64** | Filter by a end time specified as a Unix epoch Timestamp (in microseconds). Only Cluster audit logs that were generated before the specified end time are returned. | 
 **startIndex** | **int64** | Specifies an index number that can be used to return subsets of items in multiple requests. Break up the items to return into multiple requests by setting pageCount and startIndex to return a subsets of items in the search result. For example, set startIndex to 0 to get the first set of pageCount items for the first request. Increment startIndex by pageCount to get the next set of pageCount items for a next request. Continue until all items are returned and therefore the total number of returned items is equal to totalCount. Default value is 0. | 
 **pageCount** | **int64** | Limit the number of items to return in the response for pagination purposes. Default value is 1000. | 
 **outputFormat** | **string** | Specifies the format of the output such as csv and json. If not specified, the json format is returned. If csv is specified, a comma-separated list with a heading row is returned. | 
 **tenantId** | **string** | TenantId specifies the tenant whose action resulted in the audit log. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if logs of all the tenants under the hierarchy of tenant with id TenantId should be returned. | 

### Return type

[**ClusterAuditLogsSearchResult**](ClusterAuditLogsSearchResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

