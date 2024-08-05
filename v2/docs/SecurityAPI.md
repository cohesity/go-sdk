# \SecurityAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClientcsr**](SecurityAPI.md#CreateClientcsr) | **Post** /client-csr | Create Certificate Signing Requests on the cluster.
[**CreateCsr**](SecurityAPI.md#CreateCsr) | **Post** /csr | Create a Certificate Signing Request on the cluster.
[**DeleteCsr**](SecurityAPI.md#DeleteCsr) | **Delete** /csr/{id} | Delete a Certificate Signing Request on the cluster.
[**GetCiphers**](SecurityAPI.md#GetCiphers) | **Get** /security/ciphers | Gets the list of ciphers enabled on the cluster.
[**GetCsrById**](SecurityAPI.md#GetCsrById) | **Get** /csr/{id} | List the specified Certificate Signing Request.
[**GetCsrList**](SecurityAPI.md#GetCsrList) | **Get** /csr | List Certificate Signing Requests on the cluster.
[**GetObjectStoreCiphers**](SecurityAPI.md#GetObjectStoreCiphers) | **Get** /security/object-store-ciphers | Gets the list of object store ciphers enabled on the cluster.
[**GetSecurityConfig**](SecurityAPI.md#GetSecurityConfig) | **Get** /security-config | Get cluster security settings.
[**ImportCertificateByClientcsr**](SecurityAPI.md#ImportCertificateByClientcsr) | **Post** /client-csr/certificate | Import the signed certificates on the cluster after the Certificate Signing Requests are created.
[**ListTrustedCaById**](SecurityAPI.md#ListTrustedCaById) | **Get** /trusted-cas/{id} | List the specified Certificate.
[**ListTrustedCas**](SecurityAPI.md#ListTrustedCas) | **Get** /trusted-cas | List all Certificates with cluster trust store.
[**ModifyCiphers**](SecurityAPI.md#ModifyCiphers) | **Post** /security/ciphers | Enable/Disable a list of ciphers on the cluster. Iris must be restarted for the change to take effect.
[**ModifyObjectStoreCiphers**](SecurityAPI.md#ModifyObjectStoreCiphers) | **Post** /security/object-store-ciphers | Enable/Disable a list of object store ciphers on the cluster. Bridge must be restarted for the change to take effect.
[**RegisterTrustedCas**](SecurityAPI.md#RegisterTrustedCas) | **Post** /trusted-cas | Register CA Certificate to the cluster trust store.
[**UnregisterTrustedCa**](SecurityAPI.md#UnregisterTrustedCa) | **Delete** /trusted-cas/{id} | Unregister CA Certificate from the cluster trust store.
[**UpdateCertificateByCsr**](SecurityAPI.md#UpdateCertificateByCsr) | **Post** /csr/certificate | Update the signed certificate on the cluster after a Certificate Signing Request is created.
[**UpdateSecurityConfig**](SecurityAPI.md#UpdateSecurityConfig) | **Put** /security-config | Update cluster security settings.
[**ValidateTrustedCaById**](SecurityAPI.md#ValidateTrustedCaById) | **Post** /trusted-cas/{id}/validate | Validate CA Certificate.



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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewCreateCsrRequest("City_example", "CountryCode_example", "Organization_example", "OrganizationUnit_example", "State_example") // CreateCsrRequest | Specifies the parameters to create the Certificate Signing Requests.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.CreateClientcsr(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.CreateClientcsr``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateClientcsr`: CreateClientcsrResponseBody
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.CreateClientcsr`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewCreateCsrRequest("City_example", "CountryCode_example", "Organization_example", "OrganizationUnit_example", "State_example") // CreateCsrRequest | Specifies the parameters to create a Certificate Signing Request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.CreateCsr(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.CreateCsr``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCsr`: CreateCsrResponseBody
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.CreateCsr`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies the id of the csr to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.DeleteCsr(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.DeleteCsr``: %v\n", err)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetCiphers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetCiphers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCiphers`: CiphersResp
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetCiphers`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies the id of the csr.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetCsrById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetCsrById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCsrById`: CommonCsrResponseParams
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetCsrById`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	serviceName := "serviceName_example" // string | Specifies the Cohesity service name for which the CSR is generated. If this is not specified, all the csrs on the cluster will be returned. (optional)
	ids := []string{"Inner_example"} // []string | Specifies the ids of the csrs. If this is not specified, all the csrs on the cluster will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetCsrList(context.Background()).ServiceName(serviceName).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetCsrList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCsrList`: []CommonCsrResponseParams
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetCsrList`: %v\n", resp)
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


## GetObjectStoreCiphers

> ObjectStoreCiphersResp GetObjectStoreCiphers(ctx).Execute()

Gets the list of object store ciphers enabled on the cluster.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetObjectStoreCiphers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetObjectStoreCiphers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStoreCiphers`: ObjectStoreCiphersResp
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetObjectStoreCiphers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStoreCiphersRequest struct via the builder pattern


### Return type

[**ObjectStoreCiphersResp**](ObjectStoreCiphersResp.md)

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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetSecurityConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetSecurityConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityConfig`: SecurityConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetSecurityConfig`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewImportCertificateByClientcsrRequest("CertificateClient_example", "CertificateServer_example") // ImportCertificateByClientcsrRequest | Specifies the parameters to import the certificate.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ImportCertificateByClientcsr(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ImportCertificateByClientcsr``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportCertificateByClientcsr`: ImportCertificateByClientcsrResponseBody
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ImportCertificateByClientcsr`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies the id of the certificate.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ListTrustedCaById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ListTrustedCaById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTrustedCaById`: TrustedCa
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ListTrustedCaById`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	ids := []string{"Inner_example"} // []string | Specifies the ids of the certificates to be returned. (optional)
	names := []string{"Inner_example"} // []string | Specifies the names of the certificates to be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ListTrustedCas(context.Background()).Ids(ids).Names(names).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ListTrustedCas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTrustedCas`: ListTrustedCasResult
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ListTrustedCas`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewModifyCiphersRequestBody([]string{"Ciphers_example"}, false) // ModifyCiphersRequestBody | Enable/Disable ciphers.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ModifyCiphers(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ModifyCiphers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyCiphers`: CiphersResp
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ModifyCiphers`: %v\n", resp)
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


## ModifyObjectStoreCiphers

> ObjectStoreCiphersResp ModifyObjectStoreCiphers(ctx).Body(body).Execute()

Enable/Disable a list of object store ciphers on the cluster. Bridge must be restarted for the change to take effect.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewModifyObjectStoreCiphersRequestBody([]string{"Ciphers_example"}, false) // ModifyObjectStoreCiphersRequestBody | Enable/Disable object store ciphers.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ModifyObjectStoreCiphers(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ModifyObjectStoreCiphers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyObjectStoreCiphers`: ObjectStoreCiphersResp
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ModifyObjectStoreCiphers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyObjectStoreCiphersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ModifyObjectStoreCiphersRequestBody**](ModifyObjectStoreCiphersRequestBody.md) | Enable/Disable object store ciphers. | 

### Return type

[**ObjectStoreCiphersResp**](ObjectStoreCiphersResp.md)

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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewRegisterTrustedCas([]openapiclient.TrustedCaRequest{*openapiclient.NewTrustedCaRequest("Certificate_example", "Name_example")}) // RegisterTrustedCas | Specifies the parameters to register a Certificate.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.RegisterTrustedCas(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.RegisterTrustedCas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterTrustedCas`: ListTrustedCasResult
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.RegisterTrustedCas`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies the id of the certificate to be unregistered.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.UnregisterTrustedCa(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.UnregisterTrustedCa``: %v\n", err)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewUpdateCertificateByCsrRequest("Certificate_example", "CsrId_example") // UpdateCertificateByCsrRequest | Specifies the parameters to update the certificate.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.UpdateCertificateByCsr(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.UpdateCertificateByCsr``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCertificateByCsr`: UpdateCertificateByCsrResponseBody
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.UpdateCertificateByCsr`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewUpdateSecurityConfigRequest() // UpdateSecurityConfigRequest | Specifies the parameters to update security config.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.UpdateSecurityConfig(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.UpdateSecurityConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSecurityConfig`: SecurityConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.UpdateSecurityConfig`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies the id of the certificate to be validated.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ValidateTrustedCaById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ValidateTrustedCaById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateTrustedCaById`: TrustedCa
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ValidateTrustedCaById`: %v\n", resp)
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

