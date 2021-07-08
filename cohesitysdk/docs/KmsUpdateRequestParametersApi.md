# \KmsUpdateRequestParametersApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateKmsConfig**](KmsUpdateRequestParametersApi.md#UpdateKmsConfig) | **Put** /public/kmsConfig | Update KMS configurations in the cluster.



## UpdateKmsConfig

> KmsConfigurationResponse UpdateKmsConfig(ctx).Body(body).Execute()

Update KMS configurations in the cluster.

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
    body := *openapiclient.NewKmsUpdateRequestParameters() // KmsUpdateRequestParameters |  (optional)

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

    request := cohesitysdk.ApiUpdateKmsConfigRequest{
        Body: &body
    }

    resp, r, err := api_client.KmsUpdateRequestParametersApi.UpdateKmsConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KmsUpdateRequestParametersApi.UpdateKmsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKmsConfig`: KmsConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `KmsUpdateRequestParametersApi.UpdateKmsConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKmsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**KmsUpdateRequestParameters**](KmsUpdateRequestParameters.md) |  | 

### Return type

[**KmsConfigurationResponse**](KmsConfigurationResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

