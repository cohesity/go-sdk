# \Privilege

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPrivileges**](Privilege.md#GetPrivileges) | **Get** /privileges | Get Privileges.



## GetPrivileges

> Privileges GetPrivileges(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Names(names).Execute()

Get Privileges.



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
    names := []string{"Inner_example"} // []string | Filter by a list of Privilege names. (optional)

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

    request := helios.ApiGetPrivilegesRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        Names: &names
    }

    resp, r, err := api_client.Privilege.GetPrivileges(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Privilege.GetPrivileges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivileges`: Privileges
    fmt.Fprintf(os.Stdout, "Response from `Privilege.GetPrivileges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivilegesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **names** | **[]string** | Filter by a list of Privilege names. | 

### Return type

[**Privileges**](Privileges.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

