# \LicenseApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LicenseUsage**](LicenseApi.md#LicenseUsage) | **Get** /public/licenseUsage | Get the current license usage of a Cohesity Cluster.



## LicenseUsage

> LicenseAccountUsageRsp LicenseUsage(ctx).Execute()

Get the current license usage of a Cohesity Cluster.



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

    request := cohesitysdk.ApiLicenseUsageRequest{
    }

    resp, r, err := api_client.LicenseApi.LicenseUsage(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseApi.LicenseUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LicenseUsage`: LicenseAccountUsageRsp
    fmt.Fprintf(os.Stdout, "Response from `LicenseApi.LicenseUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLicenseUsageRequest struct via the builder pattern


### Return type

[**LicenseAccountUsageRsp**](LicenseAccountUsageRsp.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

