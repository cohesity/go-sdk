# \DashboardApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDashboard**](DashboardApi.md#GetDashboard) | **Get** /public/dashboard | Returns the dashboard that match the filter criteria specified using parameters.



## GetDashboard

> DashboardResponse GetDashboard(ctx).ClusterId(clusterId).AllClusters(allClusters).Refresh(refresh).TileTypes(tileTypes).Execute()

Returns the dashboard that match the filter criteria specified using parameters.



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
    clusterId := int64(789) // int64 | Id of the remote cluster for which to fetch the data. If value is not specified, it is assumed to be local cluster. (optional)
    allClusters := true // bool | Summary data for all clusters. If this is set to true, all other parameters will be ignored. (optional)
    refresh := true // bool | Specifies whether to refresh the tiles selected. (optional)
    tileTypes := []string{"TileTypes_example"} // []string | Specifies the types of the tiles to be returned. If this is not specified, all the tiles are returned. This is ignored when allClusters is set to true. 'kHealthTile' is the tile that shows health of the cluster and the alerts in the past 24 hours. 'kJobRunsTile' is the tile that shows job runs in the past 24 hours. 'kRecoveriesTile' is the tile that shows recoveries done in the past 30 days. 'kProtectedObjectsTile' is the tile that shows the protected objects details. 'kProtectionTile' is the tile that shows the protection information in the past 24 hours. 'kAuditLogsTile' is the tile that shows the recent audit logs. 'kIopsTile' is the tile that shows IP performance in the past 24 hours. 'kThroughputTile' is the tile that shows job runs in the past 24 hours. 'kStorageEfficiencyTile' is the tile that shows job runs in the past 7 days. (optional)

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

    request := cohesitysdk.ApiGetDashboardRequest{
        ClusterId: &clusterId
        AllClusters: &allClusters
        Refresh: &refresh
        TileTypes: &tileTypes
    }

    resp, r, err := api_client.DashboardApi.GetDashboard(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardApi.GetDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboard`: DashboardResponse
    fmt.Fprintf(os.Stdout, "Response from `DashboardApi.GetDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **int64** | Id of the remote cluster for which to fetch the data. If value is not specified, it is assumed to be local cluster. | 
 **allClusters** | **bool** | Summary data for all clusters. If this is set to true, all other parameters will be ignored. | 
 **refresh** | **bool** | Specifies whether to refresh the tiles selected. | 
 **tileTypes** | **[]string** | Specifies the types of the tiles to be returned. If this is not specified, all the tiles are returned. This is ignored when allClusters is set to true. &#39;kHealthTile&#39; is the tile that shows health of the cluster and the alerts in the past 24 hours. &#39;kJobRunsTile&#39; is the tile that shows job runs in the past 24 hours. &#39;kRecoveriesTile&#39; is the tile that shows recoveries done in the past 30 days. &#39;kProtectedObjectsTile&#39; is the tile that shows the protected objects details. &#39;kProtectionTile&#39; is the tile that shows the protection information in the past 24 hours. &#39;kAuditLogsTile&#39; is the tile that shows the recent audit logs. &#39;kIopsTile&#39; is the tile that shows IP performance in the past 24 hours. &#39;kThroughputTile&#39; is the tile that shows job runs in the past 24 hours. &#39;kStorageEfficiencyTile&#39; is the tile that shows job runs in the past 7 days. | 

### Return type

[**DashboardResponse**](DashboardResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

