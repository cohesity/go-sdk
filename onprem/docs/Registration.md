# \Registration

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHeliosRegConfig**](Registration.md#GetHeliosRegConfig) | **Get** /helios-registration-config | Lists the Helios Registration Config.
[**HeliosClaim**](Registration.md#HeliosClaim) | **Post** /helios-registration | Register to Helios.



## GetHeliosRegConfig

> HeliosRegConfig GetHeliosRegConfig(ctx).Execute()

Lists the Helios Registration Config.



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

    request := onprem.ApiGetHeliosRegConfigRequest{
    }

    resp, r, err := api_client.Registration.GetHeliosRegConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Registration.GetHeliosRegConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHeliosRegConfig`: HeliosRegConfig
    fmt.Fprintf(os.Stdout, "Response from `Registration.GetHeliosRegConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHeliosRegConfigRequest struct via the builder pattern


### Return type

[**HeliosRegConfig**](HeliosRegConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeliosClaim

> HeliosClaim(ctx).Body(body).Execute()

Register to Helios.



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
    body := *openapiclient.NewHeliosClaimRequest("RegistrationToken_example") // HeliosClaimRequest | Specifies the parameters to claim to Helios.

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

    request := onprem.ApiHeliosClaimRequest{
        Body: &body
    }

    resp, r, err := api_client.Registration.HeliosClaim(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Registration.HeliosClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHeliosClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HeliosClaimRequest**](HeliosClaimRequest.md) | Specifies the parameters to claim to Helios. | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

