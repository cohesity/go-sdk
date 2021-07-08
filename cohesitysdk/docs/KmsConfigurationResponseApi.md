# \KmsConfigurationResponseApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKmsConfig**](KmsConfigurationResponseApi.md#GetKmsConfig) | **Get** /public/kmsConfig | List KMS configurations in the cluster.



## GetKmsConfig

> []KmsConfigurationResponse GetKmsConfig(ctx).Id(id).Execute()

List KMS configurations in the cluster.

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
    id := int64(789) // int64 | The Id of a KMS server. (optional)

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

    request := cohesitysdk.ApiGetKmsConfigRequest{
        Id: &id
    }

    resp, r, err := api_client.KmsConfigurationResponseApi.GetKmsConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KmsConfigurationResponseApi.GetKmsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKmsConfig`: []KmsConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `KmsConfigurationResponseApi.GetKmsConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKmsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64** | The Id of a KMS server. | 

### Return type

[**[]KmsConfigurationResponse**](KmsConfigurationResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

