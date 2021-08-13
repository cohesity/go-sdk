# \HeliosOnPrem

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHeliosOnPremConfig**](HeliosOnPrem.md#GetHeliosOnPremConfig) | **Get** /helios-onprem/config | Retreive Helios OnPrem Configuration
[**UpdateHeliosOnPremConfig**](HeliosOnPrem.md#UpdateHeliosOnPremConfig) | **Put** /helios-onprem/config | Update Helios OnPrem Configuration



## GetHeliosOnPremConfig

> HeliosOnPremConfig GetHeliosOnPremConfig(ctx).Execute()

Retreive Helios OnPrem Configuration



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

    request := onprem.ApiGetHeliosOnPremConfigRequest{
    }

    resp, r, err := api_client.HeliosOnPrem.GetHeliosOnPremConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HeliosOnPrem.GetHeliosOnPremConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHeliosOnPremConfig`: HeliosOnPremConfig
    fmt.Fprintf(os.Stdout, "Response from `HeliosOnPrem.GetHeliosOnPremConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHeliosOnPremConfigRequest struct via the builder pattern


### Return type

[**HeliosOnPremConfig**](HeliosOnPremConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHeliosOnPremConfig

> HeliosOnPremConfig UpdateHeliosOnPremConfig(ctx).Body(body).Execute()

Update Helios OnPrem Configuration



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
    body := *openapiclient.NewHeliosOnPremConfig("Name_example", "KubernetesSubnetCidr_example") // HeliosOnPremConfig | Specifies the parameters for config update.

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

    request := onprem.ApiUpdateHeliosOnPremConfigRequest{
        Body: &body
    }

    resp, r, err := api_client.HeliosOnPrem.UpdateHeliosOnPremConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HeliosOnPrem.UpdateHeliosOnPremConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHeliosOnPremConfig`: HeliosOnPremConfig
    fmt.Fprintf(os.Stdout, "Response from `HeliosOnPrem.UpdateHeliosOnPremConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHeliosOnPremConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HeliosOnPremConfig**](HeliosOnPremConfig.md) | Specifies the parameters for config update. | 

### Return type

[**HeliosOnPremConfig**](HeliosOnPremConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

