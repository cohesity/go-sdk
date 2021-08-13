# \Alert

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlertSummary**](Alert.md#GetAlertSummary) | **Get** /alerts-summary | Get alerts summary.
[**GetHeliosAlertsSummary**](Alert.md#GetHeliosAlertsSummary) | **Get** /mcm/stats/alerts-summary | Get alerts summary on Helios.



## GetAlertSummary

> AlertsSummaryResponse GetAlertSummary(ctx).AccessClusterId(accessClusterId).RegionId(regionId).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).IncludeTenants(includeTenants).TenantIds(tenantIds).StatesList(statesList).Execute()

Get alerts summary.



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
    startTimeUsecs := int64(789) // int64 | Filter by start time. Specify the start time as a Unix epoch Timestamp (in microseconds). By default it is current time minus a day. (optional)
    endTimeUsecs := int64(789) // int64 | Filter by end time. Specify the end time as a Unix epoch Timestamp (in microseconds). By default it is current time. (optional)
    includeTenants := true // bool | IncludeTenants specifies if alerts of all the tenants under the hierarchy of the logged in user's organization should be used to compute summary. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which alerts are to be used to compute summary. (optional)
    statesList := []string{"StatesList_example"} // []string | Specifies list of alert states to filter alerts by. If not specified, only open alerts will be used to get summary. (optional)

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

    request := helios.ApiGetAlertSummaryRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        StartTimeUsecs: &startTimeUsecs
        EndTimeUsecs: &endTimeUsecs
        IncludeTenants: &includeTenants
        TenantIds: &tenantIds
        StatesList: &statesList
    }

    resp, r, err := api_client.Alert.GetAlertSummary(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Alert.GetAlertSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertSummary`: AlertsSummaryResponse
    fmt.Fprintf(os.Stdout, "Response from `Alert.GetAlertSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **startTimeUsecs** | **int64** | Filter by start time. Specify the start time as a Unix epoch Timestamp (in microseconds). By default it is current time minus a day. | 
 **endTimeUsecs** | **int64** | Filter by end time. Specify the end time as a Unix epoch Timestamp (in microseconds). By default it is current time. | 
 **includeTenants** | **bool** | IncludeTenants specifies if alerts of all the tenants under the hierarchy of the logged in user&#39;s organization should be used to compute summary. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which alerts are to be used to compute summary. | 
 **statesList** | **[]string** | Specifies list of alert states to filter alerts by. If not specified, only open alerts will be used to get summary. | 

### Return type

[**AlertsSummaryResponse**](AlertsSummaryResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHeliosAlertsSummary

> AlertsSummaryResponse GetHeliosAlertsSummary(ctx).RegionId(regionId).ClusterIdentifiers(clusterIdentifiers).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).StatesList(statesList).Execute()

Get alerts summary on Helios.



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
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    clusterIdentifiers := []string{"Inner_example"} // []string | Specifies the list of cluster identifiers. Format is clusterId:clusterIncarnationId. (optional)
    startTimeUsecs := int64(789) // int64 | Filter by start time. Specify the start time as a Unix epoch Timestamp (in microseconds). By default it is current time minus a day. (optional)
    endTimeUsecs := int64(789) // int64 | Filter by end time. Specify the end time as a Unix epoch Timestamp (in microseconds). By default it is current time. (optional)
    statesList := []string{"StatesList_example"} // []string | Specifies list of alert states to filter alerts by. If not specified, only open alerts will be used to get summary. (optional)

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

    request := helios.ApiGetHeliosAlertsSummaryRequest{
        RegionId: &regionId
        ClusterIdentifiers: &clusterIdentifiers
        StartTimeUsecs: &startTimeUsecs
        EndTimeUsecs: &endTimeUsecs
        StatesList: &statesList
    }

    resp, r, err := api_client.Alert.GetHeliosAlertsSummary(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Alert.GetHeliosAlertsSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHeliosAlertsSummary`: AlertsSummaryResponse
    fmt.Fprintf(os.Stdout, "Response from `Alert.GetHeliosAlertsSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHeliosAlertsSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **clusterIdentifiers** | **[]string** | Specifies the list of cluster identifiers. Format is clusterId:clusterIncarnationId. | 
 **startTimeUsecs** | **int64** | Filter by start time. Specify the start time as a Unix epoch Timestamp (in microseconds). By default it is current time minus a day. | 
 **endTimeUsecs** | **int64** | Filter by end time. Specify the end time as a Unix epoch Timestamp (in microseconds). By default it is current time. | 
 **statesList** | **[]string** | Specifies list of alert states to filter alerts by. If not specified, only open alerts will be used to get summary. | 

### Return type

[**AlertsSummaryResponse**](AlertsSummaryResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

