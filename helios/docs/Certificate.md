# \Certificate

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHeliosSslCertificate**](Certificate.md#GetHeliosSslCertificate) | **Get** /mcm/ssl-certificate | Get the Helios SSL Certificate.



## GetHeliosSslCertificate

> SslCertificateParams GetHeliosSslCertificate(ctx).RegionId(regionId).Execute()

Get the Helios SSL Certificate.



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

    request := helios.ApiGetHeliosSslCertificateRequest{
        RegionId: &regionId
    }

    resp, r, err := api_client.Certificate.GetHeliosSslCertificate(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Certificate.GetHeliosSslCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHeliosSslCertificate`: SslCertificateParams
    fmt.Fprintf(os.Stdout, "Response from `Certificate.GetHeliosSslCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHeliosSslCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**SslCertificateParams**](SslCertificateParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

