# \Security

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClientcsr**](Security.md#CreateClientcsr) | **Post** /client-csr | Create Certificate Signing Requests on the cluster.
[**CreateCsr**](Security.md#CreateCsr) | **Post** /csr | Create a Certificate Signing Request on the cluster.
[**DeleteCsr**](Security.md#DeleteCsr) | **Delete** /csr/{id} | Delete a Certificate Signing Request on the cluster.
[**GetCiphers**](Security.md#GetCiphers) | **Get** /security/ciphers | Gets the list of ciphers enabled on the cluster.
[**GetCsrById**](Security.md#GetCsrById) | **Get** /csr/{id} | List the specified Certificate Signing Request.
[**GetCsrList**](Security.md#GetCsrList) | **Get** /csr | List Certificate Signing Requests on the cluster.
[**GetSecurityConfig**](Security.md#GetSecurityConfig) | **Get** /security-config | Get cluster security settings.
[**ImportCertificateByClientcsr**](Security.md#ImportCertificateByClientcsr) | **Post** /client-csr/certificate | Import the signed certificates on the cluster after the Certificate Signing Requests are created.
[**ListTrustedCaById**](Security.md#ListTrustedCaById) | **Get** /trusted-cas/{id} | List the specified Certificate.
[**ListTrustedCas**](Security.md#ListTrustedCas) | **Get** /trusted-cas | List all Certificates with cluster trust store.
[**ModifyCiphers**](Security.md#ModifyCiphers) | **Post** /security/ciphers | Enable/Disable a list of ciphers on the cluster. Iris must be restarted for the change to take effect.
[**RegisterTrustedCas**](Security.md#RegisterTrustedCas) | **Post** /trusted-cas | Register CA Certificate to the cluster trust store.
[**UnregisterTrustedCa**](Security.md#UnregisterTrustedCa) | **Delete** /trusted-cas/{id} | Unregister CA Certificate from the cluster trust store.
[**UpdateCertificateByCsr**](Security.md#UpdateCertificateByCsr) | **Post** /csr/certificate | Update the signed certificate on the cluster after a Certificate Signing Request is created.
[**UpdateSecurityConfig**](Security.md#UpdateSecurityConfig) | **Put** /security-config | Update cluster security settings.
[**ValidateTrustedCaById**](Security.md#ValidateTrustedCaById) | **Post** /trusted-cas/{id}/validate | Validate CA Certificate.



## CreateClientcsr

> CreateClientcsrResponseBody CreateClientcsr(ctx).Body(body).Execute()

Create Certificate Signing Requests on the cluster.



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
    body := *openapiclient.NewCreateCsrRequest("Organization_example", "OrganizationUnit_example", "CountryCode_example", "State_example", "City_example") // CreateCsrRequest | Specifies the parameters to create the Certificate Signing Requests.

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

    request := onprem.ApiCreateClientcsrRequest{
        Body: &body
    }

    resp, r, err := api_client.Security.CreateClientcsr(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Security.CreateClientcsr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateClientcsr`: CreateClientcsrResponseBody
    fmt.Fprintf(os.Stdout, "Response from `Security.CreateClientcsr`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientcsrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateCsrRequest**](CreateCsrRequest.md) | Specifies the parameters to create the Certificate Signing Requests. | 

### Return type

[**CreateClientcsrResponseBody**](CreateClientcsrResponseBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCsr

> CreateCsrResponseBody CreateCsr(ctx).Body(body).Execute()

Create a Certificate Signing Request on the cluster.



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
    body := *openapiclient.NewCreateCsrRequest("Organization_example", "OrganizationUnit_example", "CountryCode_example", "State_example", "City_example") // CreateCsrRequest | Specifies the parameters to create a Certificate Signing Request.

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

    request := onprem.ApiCreateCsrRequest{
        Body: &body
    }

    resp, r, err := api_client.Security.CreateCsr(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Security.CreateCsr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCsr`: CreateCsrResponseBody
    fmt.Fprintf(os.Stdout, "Response from `Security.CreateCsr`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCsrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateCsrRequest**](CreateCsrRequest.md) | Specifies the parameters to create a Certificate Signing Request. | 

### Return type

[**CreateCsrResponseBody**](CreateCsrResponseBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCsr

> DeleteCsr(ctx, id).Execute()

Delete a Certificate Signing Request on the cluster.



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
    id := "id_example" // string | Specifies the id of the csr to be deleted.

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

    request := onprem.ApiDeleteCsrRequest{
        Id: &id
    }

    resp, r, err := api_client.Security.DeleteCsr(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Security.DeleteCsr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the csr to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCsrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCiphers

> CiphersResp GetCiphers(ctx).Execute()

Gets the list of ciphers enabled on the cluster.



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

    request := onprem.ApiGetCiphersRequest{
    }

    resp, r, err := api_client.Security.GetCiphers(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Security.GetCiphers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCiphers`: CiphersResp
    fmt.Fprintf(os.Stdout, "Response from `Security.GetCiphers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCiphersRequest struct via the builder pattern


### Return type

[**CiphersResp**](CiphersResp.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCsrById

> CommonCsrResponseParams GetCsrById(ctx, id).Execute()

List the specified Certificate Signing Request.



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
    id := "id_example" // string | Specifies the id of the csr.

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

    request := onprem.ApiGetCsrByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.Security.GetCsrById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Security.GetCsrById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCsrById`: CommonCsrResponseParams
    fmt.Fprintf(os.Stdout, "Response from `Security.GetCsrById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the csr. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCsrByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommonCsrResponseParams**](CommonCsrResponseParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCsrList

> []CommonCsrResponseParams GetCsrList(ctx).ServiceName(serviceName).Ids(ids).Execute()

List Certificate Signing Requests on the cluster.



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
    serviceName := "serviceName_example" // string | Specifies the Cohesity service name for which the CSR is generated. If this is not specified, all the csrs on the cluster will be returned. (optional)
    ids := []string{"Inner_example"} // []string | Specifies the ids of the csrs. If this is not specified, all the csrs on the cluster will be returned. (optional)

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

    request := onprem.ApiGetCsrListRequest{
        ServiceName: &serviceName
        Ids: &ids
    }

    resp, r, err := api_client.Security.GetCsrList(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Security.GetCsrList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCsrList`: []CommonCsrResponseParams
    fmt.Fprintf(os.Stdout, "Response from `Security.GetCsrList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCsrListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceName** | **string** | Specifies the Cohesity service name for which the CSR is generated. If this is not specified, all the csrs on the cluster will be returned. | 
 **ids** | **[]string** | Specifies the ids of the csrs. If this is not specified, all the csrs on the cluster will be returned. | 

### Return type

[**[]CommonCsrResponseParams**](CommonCsrResponseParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityConfig

> SecurityConfigResponse GetSecurityConfig(ctx).Execute()

Get cluster security settings.



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

    request := onprem.ApiGetSecurityConfigRequest{
    }

    resp, r, err := api_client.Security.GetSecurityConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Security.GetSecurityConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityConfig`: SecurityConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `Security.GetSecurityConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityConfigRequest struct via the builder pattern


### Return type

[**SecurityConfigResponse**](SecurityConfigResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportCertificateByClientcsr

> ImportCertificateByClientcsrResponseBody ImportCertificateByClientcsr(ctx).Body(body).Execute()

Import the signed certificates on the cluster after the Certificate Signing Requests are created.



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
    body := *openapiclient.NewImportCertificateByClientcsrRequest("CertificateServer_example", "CertificateClient_example") // ImportCertificateByClientcsrRequest | Specifies the parameters to import the certificate.

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

    request := onprem.ApiImportCertificateByClientcsrRequest{
        Body: &body
    }

    resp, r, err := api_client.Security.ImportCertificateByClientcsr(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Security.ImportCertificateByClientcsr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportCertificateByClientcsr`: ImportCertificateByClientcsrResponseBody
    fmt.Fprintf(os.Stdout, "Response from `Security.ImportCertificateByClientcsr`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportCertificateByClientcsrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ImportCertificateByClientcsrRequest**](ImportCertificateByClientcsrRequest.md) | Specifies the parameters to import the certificate. | 

### Return type

[**ImportCertificateByClientcsrResponseBody**](ImportCertificateByClientcsrResponseBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrustedCaById

> TrustedCa ListTrustedCaById(ctx, id).Execute()

List the specified Certificate.



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
    id := "id_example" // string | Specifies the id of the certificate.

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

    request := onprem.ApiListTrustedCaByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.Security.ListTrustedCaById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Security.ListTrustedCaById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTrustedCaById`: TrustedCa
    fmt.Fprintf(os.Stdout, "Response from `Security.ListTrustedCaById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTrustedCaByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrustedCa**](TrustedCa.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrustedCas

> ListTrustedCasResult ListTrustedCas(ctx).Ids(ids).Names(names).Execute()

List all Certificates with cluster trust store.



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
    ids := []string{"Inner_example"} // []string | Specifies the ids of the certificates to be returned. (optional)
    names := []string{"Inner_example"} // []string | Specifies the names of the certificates to be returned. (optional)

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

    request := onprem.ApiListTrustedCasRequest{
        Ids: &ids
        Names: &names
    }

    resp, r, err := api_client.Security.ListTrustedCas(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Security.ListTrustedCas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTrustedCas`: ListTrustedCasResult
    fmt.Fprintf(os.Stdout, "Response from `Security.ListTrustedCas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTrustedCasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | Specifies the ids of the certificates to be returned. | 
 **names** | **[]string** | Specifies the names of the certificates to be returned. | 

### Return type

[**ListTrustedCasResult**](ListTrustedCasResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyCiphers

> CiphersResp ModifyCiphers(ctx).Body(body).Execute()

Enable/Disable a list of ciphers on the cluster. Iris must be restarted for the change to take effect.



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
    body := *openapiclient.NewModifyCiphersRequestBody(false, []string{"Ciphers_example"}) // ModifyCiphersRequestBody | Enable/Disable ciphers.

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

    request := onprem.ApiModifyCiphersRequest{
        Body: &body
    }

    resp, r, err := api_client.Security.ModifyCiphers(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Security.ModifyCiphers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyCiphers`: CiphersResp
    fmt.Fprintf(os.Stdout, "Response from `Security.ModifyCiphers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyCiphersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ModifyCiphersRequestBody**](ModifyCiphersRequestBody.md) | Enable/Disable ciphers. | 

### Return type

[**CiphersResp**](CiphersResp.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterTrustedCas

> ListTrustedCasResult RegisterTrustedCas(ctx).Body(body).Execute()

Register CA Certificate to the cluster trust store.



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
    body := *openapiclient.NewRegisterTrustedCas([]openapiclient.TrustedCaRequest{*openapiclient.NewTrustedCaRequest("Certificate_example", "Name_example")}) // RegisterTrustedCas | Specifies the parameters to register a Certificate.

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

    request := onprem.ApiRegisterTrustedCasRequest{
        Body: &body
    }

    resp, r, err := api_client.Security.RegisterTrustedCas(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Security.RegisterTrustedCas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterTrustedCas`: ListTrustedCasResult
    fmt.Fprintf(os.Stdout, "Response from `Security.RegisterTrustedCas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterTrustedCasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RegisterTrustedCas**](RegisterTrustedCas.md) | Specifies the parameters to register a Certificate. | 

### Return type

[**ListTrustedCasResult**](ListTrustedCasResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterTrustedCa

> UnregisterTrustedCa(ctx, id).Execute()

Unregister CA Certificate from the cluster trust store.



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
    id := "id_example" // string | Specifies the id of the certificate to be unregistered.

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

    request := onprem.ApiUnregisterTrustedCaRequest{
        Id: &id
    }

    resp, r, err := api_client.Security.UnregisterTrustedCa(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Security.UnregisterTrustedCa``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the certificate to be unregistered. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterTrustedCaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificateByCsr

> UpdateCertificateByCsrResponseBody UpdateCertificateByCsr(ctx).Body(body).Execute()

Update the signed certificate on the cluster after a Certificate Signing Request is created.



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
    body := *openapiclient.NewUpdateCertificateByCsrRequest("Certificate_example", "CsrId_example") // UpdateCertificateByCsrRequest | Specifies the parameters to update the certificate.

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

    request := onprem.ApiUpdateCertificateByCsrRequest{
        Body: &body
    }

    resp, r, err := api_client.Security.UpdateCertificateByCsr(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Security.UpdateCertificateByCsr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCertificateByCsr`: UpdateCertificateByCsrResponseBody
    fmt.Fprintf(os.Stdout, "Response from `Security.UpdateCertificateByCsr`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificateByCsrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateCertificateByCsrRequest**](UpdateCertificateByCsrRequest.md) | Specifies the parameters to update the certificate. | 

### Return type

[**UpdateCertificateByCsrResponseBody**](UpdateCertificateByCsrResponseBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecurityConfig

> SecurityConfigResponse UpdateSecurityConfig(ctx).Body(body).Execute()

Update cluster security settings.



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
    body := *openapiclient.NewUpdateSecurityConfigRequest() // UpdateSecurityConfigRequest | Specifies the parameters to update security config.

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

    request := onprem.ApiUpdateSecurityConfigRequest{
        Body: &body
    }

    resp, r, err := api_client.Security.UpdateSecurityConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Security.UpdateSecurityConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSecurityConfig`: SecurityConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `Security.UpdateSecurityConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateSecurityConfigRequest**](UpdateSecurityConfigRequest.md) | Specifies the parameters to update security config. | 

### Return type

[**SecurityConfigResponse**](SecurityConfigResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateTrustedCaById

> TrustedCa ValidateTrustedCaById(ctx, id).Execute()

Validate CA Certificate.



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
    id := "id_example" // string | Specifies the id of the certificate to be validated.

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

    request := onprem.ApiValidateTrustedCaByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.Security.ValidateTrustedCaById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Security.ValidateTrustedCaById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateTrustedCaById`: TrustedCa
    fmt.Fprintf(os.Stdout, "Response from `Security.ValidateTrustedCaById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the certificate to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateTrustedCaByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrustedCa**](TrustedCa.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

