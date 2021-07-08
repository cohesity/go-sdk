# \IpApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigureIp**](IpApi.md#ConfigureIp) | **Put** /public/ip | Configure the specfied IP settings on the Cohesity Cluster.
[**UnconfigureIp**](IpApi.md#UnconfigureIp) | **Delete** /public/ip | Delete the specified IP settings on the Cohesity Cluster.



## ConfigureIp

> IpConfig ConfigureIp(ctx).Body(body).Execute()

Configure the specfied IP settings on the Cohesity Cluster.



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
    body := *openapiclient.NewIpConfig() // IpConfig |  (optional)

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

    request := cohesitysdk.ApiConfigureIpRequest{
        Body: &body
    }

    resp, r, err := api_client.IpApi.ConfigureIp(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IpApi.ConfigureIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfigureIp`: IpConfig
    fmt.Fprintf(os.Stdout, "Response from `IpApi.ConfigureIp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigureIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IpConfig**](IpConfig.md) |  | 

### Return type

[**IpConfig**](IpConfig.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnconfigureIp

> UnconfigureIp(ctx).Body(body).Execute()

Delete the specified IP settings on the Cohesity Cluster.



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
    body := *openapiclient.NewIpUnconfig() // IpUnconfig |  (optional)

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

    request := cohesitysdk.ApiUnconfigureIpRequest{
        Body: &body
    }

    resp, r, err := api_client.IpApi.UnconfigureIp(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IpApi.UnconfigureIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnconfigureIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IpUnconfig**](IpUnconfig.md) |  | 

### Return type

 (empty response body)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

