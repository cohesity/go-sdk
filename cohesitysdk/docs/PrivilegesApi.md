# \PrivilegesApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPrivileges**](PrivilegesApi.md#GetPrivileges) | **Get** /public/privileges | List the privileges defined on the Cohesity Cluster.



## GetPrivileges

> []PrivilegeInfo GetPrivileges(ctx).Name(name).Execute()

List the privileges defined on the Cohesity Cluster.



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
    name := "name_example" // string | Specifies the name of the privilege. (optional)

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

    request := cohesitysdk.ApiGetPrivilegesRequest{
        Name: &name
    }

    resp, r, err := api_client.PrivilegesApi.GetPrivileges(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivilegesApi.GetPrivileges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivileges`: []PrivilegeInfo
    fmt.Fprintf(os.Stdout, "Response from `PrivilegesApi.GetPrivileges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivilegesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Specifies the name of the privilege. | 

### Return type

[**[]PrivilegeInfo**](PrivilegeInfo.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

