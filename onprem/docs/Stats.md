# \Stats

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProtectionRunsStats**](Stats.md#GetProtectionRunsStats) | **Get** /stats/protection-runs | Get statistics of protection runs.



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
    onprem "onprem"
)

func main() {
    startTimeUsecs := int64(789) // int64 | Specify the start time as a Unix epoch Timestamp (in microseconds), only runs executing after this time will be counted. By default it is current time minus a day. (optional)
    endTimeUsecs := int64(789) // int64 | Specify the end time as a Unix epoch Timestamp (in microseconds), only runs executing before this time will be counted. By default it is current time. (optional)
    runStatus := []string{"RunStatus_example"} // []string | Specifies a list of status, runs matching the status will be returned. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Failed' indicates that the run has failed. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. (optional)

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

    request := onprem.ApiGetProtectionRunsStatsRequest{
        StartTimeUsecs: &startTimeUsecs
        EndTimeUsecs: &endTimeUsecs
        RunStatus: &runStatus
    }

    resp, r, err := api_client.Stats.GetProtectionRunsStats(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Stats.GetProtectionRunsStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionRunsStats`: GetProtectionRunsStatusResponseBody
    fmt.Fprintf(os.Stdout, "Response from `Stats.GetProtectionRunsStats`: %v\n", resp)
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

