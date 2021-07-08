# \CertificatesApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebServerCertificate**](CertificatesApi.md#DeleteWebServerCertificate) | **Delete** /public/certificates/webServer | Delete the SSL Certificate in the Cluster.
[**DeployHostCertificate**](CertificatesApi.md#DeployHostCertificate) | **Post** /public/certificates/global | Generate and deploy certificate for a single or multiple hosts.
[**GetCertificateList**](CertificatesApi.md#GetCertificateList) | **Get** /public/certificates/global | List the certificates generated and deployed on hosts.
[**GetWebServerCertificate**](CertificatesApi.md#GetWebServerCertificate) | **Get** /public/certificates/webServer | Get the Server Certificate configured on the Cluster.
[**UpdateWebServerCertificate**](CertificatesApi.md#UpdateWebServerCertificate) | **Put** /public/certificates/webServer | Update the Web Server Certificate on the Cluster.



## DeleteWebServerCertificate

> DeleteWebServerCertificate(ctx).Execute()

Delete the SSL Certificate in the Cluster.



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

    request := cohesitysdk.ApiDeleteWebServerCertificateRequest{
    }

    resp, r, err := api_client.CertificatesApi.DeleteWebServerCertificate(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.DeleteWebServerCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebServerCertificateRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeployHostCertificate

> CertificateDetails DeployHostCertificate(ctx).Body(body).Execute()

Generate and deploy certificate for a single or multiple hosts.



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
    body := *openapiclient.NewDeployCertParameters() // DeployCertParameters | Request to generate and deploy a new certificate. (optional)

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

    request := cohesitysdk.ApiDeployHostCertificateRequest{
        Body: &body
    }

    resp, r, err := api_client.CertificatesApi.DeployHostCertificate(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.DeployHostCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployHostCertificate`: CertificateDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.DeployHostCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeployHostCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeployCertParameters**](DeployCertParameters.md) | Request to generate and deploy a new certificate. | 

### Return type

[**CertificateDetails**](CertificateDetails.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateList

> ListCertResponse GetCertificateList(ctx).Execute()

List the certificates generated and deployed on hosts.



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

    request := cohesitysdk.ApiGetCertificateListRequest{
    }

    resp, r, err := api_client.CertificatesApi.GetCertificateList(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.GetCertificateList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateList`: ListCertResponse
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.GetCertificateList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateListRequest struct via the builder pattern


### Return type

[**ListCertResponse**](ListCertResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebServerCertificate

> SslCertificateConfig GetWebServerCertificate(ctx).Execute()

Get the Server Certificate configured on the Cluster.



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

    request := cohesitysdk.ApiGetWebServerCertificateRequest{
    }

    resp, r, err := api_client.CertificatesApi.GetWebServerCertificate(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.GetWebServerCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebServerCertificate`: SslCertificateConfig
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.GetWebServerCertificate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebServerCertificateRequest struct via the builder pattern


### Return type

[**SslCertificateConfig**](SslCertificateConfig.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebServerCertificate

> SslCertificateConfig UpdateWebServerCertificate(ctx).Body(body).Execute()

Update the Web Server Certificate on the Cluster.



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
    body := *openapiclient.NewSslCertificateConfig() // SslCertificateConfig |  (optional)

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

    request := cohesitysdk.ApiUpdateWebServerCertificateRequest{
        Body: &body
    }

    resp, r, err := api_client.CertificatesApi.UpdateWebServerCertificate(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.UpdateWebServerCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebServerCertificate`: SslCertificateConfig
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.UpdateWebServerCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebServerCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SslCertificateConfig**](SslCertificateConfig.md) |  | 

### Return type

[**SslCertificateConfig**](SslCertificateConfig.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

