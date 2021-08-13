# \Privilege

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPrivileges**](Privilege.md#GetPrivileges) | **Get** /privileges | Get Privileges.



## GetPrivileges

> Privileges GetPrivileges(ctx).Names(names).Execute()

Get Privileges.



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

    request := onprem.ApiGetPrivilegesRequest{
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

