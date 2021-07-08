# \PackagesApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadPackage**](PackagesApi.md#DownloadPackage) | **Post** /public/packages/url | Download a package to the Cluster by providing a URL where the package is hosted.
[**ListPackages**](PackagesApi.md#ListPackages) | **Get** /public/packages | List all currently installed packages on a Cohesity Cluster.



## DownloadPackage

> DownloadPackageResult DownloadPackage(ctx).Body(body).Execute()

Download a package to the Cluster by providing a URL where the package is hosted.



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
    body := *openapiclient.NewDownloadPackageParameters("Url_example") // DownloadPackageParameters | 

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

    request := cohesitysdk.ApiDownloadPackageRequest{
        Body: &body
    }

    resp, r, err := api_client.PackagesApi.DownloadPackage(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.DownloadPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadPackage`: DownloadPackageResult
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.DownloadPackage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DownloadPackageParameters**](DownloadPackageParameters.md) |  | 

### Return type

[**DownloadPackageResult**](DownloadPackageResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPackages

> []PackageDetails ListPackages(ctx).Execute()

List all currently installed packages on a Cohesity Cluster.



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

    request := cohesitysdk.ApiListPackagesRequest{
    }

    resp, r, err := api_client.PackagesApi.ListPackages(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesApi.ListPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPackages`: []PackageDetails
    fmt.Fprintf(os.Stdout, "Response from `PackagesApi.ListPackages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPackagesRequest struct via the builder pattern


### Return type

[**[]PackageDetails**](PackageDetails.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

