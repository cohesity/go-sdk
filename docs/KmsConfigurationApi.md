# \KmsConfigurationApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKmsConfig**](KmsConfigurationApi.md#CreateKmsConfig) | **Post** /public/kmsConfig | Create a KMS config.



## CreateKmsConfig

> KmsConfigurationResponse CreateKmsConfig(ctx).Body(body).Execute()

Create a KMS config.



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
    body := *openapiclient.NewKmsCreateRequestParameters() // KmsCreateRequestParameters |  (optional)

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

    request := cohesitysdk.ApiCreateKmsConfigRequest{
        body: &Body
    }

    resp, r, err := api_client.KmsConfigurationApi.CreateKmsConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KmsConfigurationApi.CreateKmsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKmsConfig`: KmsConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `KmsConfigurationApi.CreateKmsConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKmsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**KmsCreateRequestParameters**](KmsCreateRequestParameters.md) |  | 

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

