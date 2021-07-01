# \BannerApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBanner**](BannerApi.md#GetBanner) | **Get** /public/banners | List the banner for a persona.
[**UpdateBanner**](BannerApi.md#UpdateBanner) | **Put** /public/banners | Update an existing banner on the Cohesity Cluster.



## GetBanner

> Banner GetBanner(ctx).Execute()

List the banner for a persona.



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

    request := cohesitysdk.ApiGetBannerRequest{
    }

    resp, r, err := api_client.BannerApi.GetBanner(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BannerApi.GetBanner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBanner`: Banner
    fmt.Fprintf(os.Stdout, "Response from `BannerApi.GetBanner`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBannerRequest struct via the builder pattern


### Return type

[**Banner**](Banner.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBanner

> Banner UpdateBanner(ctx).Body(body).Execute()

Update an existing banner on the Cohesity Cluster.



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
    body := *openapiclient.NewBannerUpdateParameters() // BannerUpdateParameters | Request to update existing banner. (optional)

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

    request := cohesitysdk.ApiUpdateBannerRequest{
        body: &Body
    }

    resp, r, err := api_client.BannerApi.UpdateBanner(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BannerApi.UpdateBanner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBanner`: Banner
    fmt.Fprintf(os.Stdout, "Response from `BannerApi.UpdateBanner`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBannerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**BannerUpdateParameters**](BannerUpdateParameters.md) | Request to update existing banner. | 

### Return type

[**Banner**](Banner.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

