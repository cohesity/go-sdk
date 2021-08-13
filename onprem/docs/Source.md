# \Source

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProtectionSourceRegistration**](Source.md#DeleteProtectionSourceRegistration) | **Delete** /data-protect/sources/registrations/{id} | Delete Protection Source Registration.
[**GetProtectionSourceRegistration**](Source.md#GetProtectionSourceRegistration) | **Get** /data-protect/sources/registrations/{id} | Get a Protection Source registration.
[**GetProtectionSources**](Source.md#GetProtectionSources) | **Get** /data-protect/sources | Get a List of Protection Sources.
[**GetSourceAttributeFilters**](Source.md#GetSourceAttributeFilters) | **Get** /data-protect/sources/filters | List attribute filters for a source.
[**GetSourceRegistrations**](Source.md#GetSourceRegistrations) | **Get** /data-protect/sources/registrations | Get the list of Protection Source registrations.
[**GetVdcDetails**](Source.md#GetVdcDetails) | **Get** /data-protect/sources/virtual-datacenter/{id} | Get VDC Details.
[**ProtectionSourceById**](Source.md#ProtectionSourceById) | **Get** /data-protect/sources/{id} | Get a Protection Sources.
[**RegisterProtectionSource**](Source.md#RegisterProtectionSource) | **Post** /data-protect/sources/registrations | Register a Protection Source.
[**TestConnectionProtectionSource**](Source.md#TestConnectionProtectionSource) | **Post** /data-protect/sources/test-connection | Test connection to a source.
[**UpdateProtectionSourceRegistration**](Source.md#UpdateProtectionSourceRegistration) | **Put** /data-protect/sources/registrations/{id} | Update Protection Source registration.



## DeleteProtectionSourceRegistration

> DeleteProtectionSourceRegistration(ctx, id).Execute()

Delete Protection Source Registration.



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
    id := int64(789) // int64 | Specifies the ID of the Protection Source Registration.

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

    request := onprem.ApiDeleteProtectionSourceRegistrationRequest{
        Id: &id
    }

    resp, r, err := api_client.Source.DeleteProtectionSourceRegistration(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Source.DeleteProtectionSourceRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the ID of the Protection Source Registration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProtectionSourceRegistrationRequest struct via the builder pattern


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


## GetProtectionSourceRegistration

> SourceRegistration GetProtectionSourceRegistration(ctx, id).Execute()

Get a Protection Source registration.



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
    id := int64(789) // int64 | Specifies the id of the Protection Source registration.

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

    request := onprem.ApiGetProtectionSourceRegistrationRequest{
        Id: &id
    }

    resp, r, err := api_client.Source.GetProtectionSourceRegistration(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Source.GetProtectionSourceRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionSourceRegistration`: SourceRegistration
    fmt.Fprintf(os.Stdout, "Response from `Source.GetProtectionSourceRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the Protection Source registration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionSourceRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SourceRegistration**](SourceRegistration.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionSources

> Sources GetProtectionSources(ctx).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeSourceCredentials(includeSourceCredentials).EncryptionKey(encryptionKey).Execute()

Get a List of Protection Sources.



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
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which Sources are to be returned. (optional)
    includeTenants := true // bool | If true, the response will include Sources which belong belong to all tenants which the current user has permission to see. If false, then only Sources for the current user will be returned. (optional)
    includeSourceCredentials := true // bool | If true, the encrypted crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied encryption key. (optional)
    encryptionKey := "encryptionKey_example" // string | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. (optional)

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

    request := onprem.ApiGetProtectionSourcesRequest{
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
        IncludeSourceCredentials: &includeSourceCredentials
        EncryptionKey: &encryptionKey
    }

    resp, r, err := api_client.Source.GetProtectionSources(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Source.GetProtectionSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionSources`: Sources
    fmt.Fprintf(os.Stdout, "Response from `Source.GetProtectionSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which Sources are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Sources which belong belong to all tenants which the current user has permission to see. If false, then only Sources for the current user will be returned. | 
 **includeSourceCredentials** | **bool** | If true, the encrypted crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied encryption key. | 
 **encryptionKey** | **string** | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | 

### Return type

[**Sources**](Sources.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceAttributeFilters

> SourceAttributeFiltersResponseParams GetSourceAttributeFilters(ctx).SourceUuid(sourceUuid).Environment(environment).Execute()

List attribute filters for a source.



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
    sourceUuid := "sourceUuid_example" // string | Specifies the source UUID of the parent entity.
    environment := "environment_example" // string | Specifies the environment type of the Protection Source. (optional)

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

    request := onprem.ApiGetSourceAttributeFiltersRequest{
        SourceUuid: &sourceUuid
        Environment: &environment
    }

    resp, r, err := api_client.Source.GetSourceAttributeFilters(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Source.GetSourceAttributeFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceAttributeFilters`: SourceAttributeFiltersResponseParams
    fmt.Fprintf(os.Stdout, "Response from `Source.GetSourceAttributeFilters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceAttributeFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceUuid** | **string** | Specifies the source UUID of the parent entity. | 
 **environment** | **string** | Specifies the environment type of the Protection Source. | 

### Return type

[**SourceAttributeFiltersResponseParams**](SourceAttributeFiltersResponseParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceRegistrations

> SourceRegistrations GetSourceRegistrations(ctx).Ids(ids).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeSourceCredentials(includeSourceCredentials).EncryptionKey(encryptionKey).Execute()

Get the list of Protection Source registrations.



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
    ids := []int32{int32(123)} // []int32 | Ids specifies the list of source registration ids to return. If left empty, every source registration will be returned by default. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    includeTenants := true // bool | If true, the response will include Registrations which were created by all tenants which the current user has permission to see. If false, then only Registrations created by the current user will be returned. (optional)
    includeSourceCredentials := true // bool | If true, the encrypted crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied encryption key. (optional)
    encryptionKey := "encryptionKey_example" // string | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. (optional)

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

    request := onprem.ApiGetSourceRegistrationsRequest{
        Ids: &ids
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
        IncludeSourceCredentials: &includeSourceCredentials
        EncryptionKey: &encryptionKey
    }

    resp, r, err := api_client.Source.GetSourceRegistrations(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Source.GetSourceRegistrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceRegistrations`: SourceRegistrations
    fmt.Fprintf(os.Stdout, "Response from `Source.GetSourceRegistrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceRegistrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int32** | Ids specifies the list of source registration ids to return. If left empty, every source registration will be returned by default. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Registrations which were created by all tenants which the current user has permission to see. If false, then only Registrations created by the current user will be returned. | 
 **includeSourceCredentials** | **bool** | If true, the encrypted crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied encryption key. | 
 **encryptionKey** | **string** | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | 

### Return type

[**SourceRegistrations**](SourceRegistrations.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVdcDetails

> VdcObject GetVdcDetails(ctx, id).Execute()

Get VDC Details.



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
    id := int64(789) // int64 | Specifies the ID of the VMware virtual datacenter.

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

    request := onprem.ApiGetVdcDetailsRequest{
        Id: &id
    }

    resp, r, err := api_client.Source.GetVdcDetails(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Source.GetVdcDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVdcDetails`: VdcObject
    fmt.Fprintf(os.Stdout, "Response from `Source.GetVdcDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the ID of the VMware virtual datacenter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVdcDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VdcObject**](VdcObject.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProtectionSourceById

> Source ProtectionSourceById(ctx, id).Execute()

Get a Protection Sources.



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
    id := int64(789) // int64 | Specifies the id of the Protection Source.

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

    request := onprem.ApiProtectionSourceByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.Source.ProtectionSourceById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Source.ProtectionSourceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProtectionSourceById`: Source
    fmt.Fprintf(os.Stdout, "Response from `Source.ProtectionSourceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the Protection Source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProtectionSourceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Source**](Source.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterProtectionSource

> SourceRegistration RegisterProtectionSource(ctx).Body(body).Execute()

Register a Protection Source.



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
    body := *openapiclient.NewSourceRegistrationRequestParams("Environment_example") // SourceRegistrationRequestParams | Specifies the parameters to register a Protection Source.

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

    request := onprem.ApiRegisterProtectionSourceRequest{
        Body: &body
    }

    resp, r, err := api_client.Source.RegisterProtectionSource(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Source.RegisterProtectionSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterProtectionSource`: SourceRegistration
    fmt.Fprintf(os.Stdout, "Response from `Source.RegisterProtectionSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterProtectionSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SourceRegistrationRequestParams**](SourceRegistrationRequestParams.md) | Specifies the parameters to register a Protection Source. | 

### Return type

[**SourceRegistration**](SourceRegistration.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestConnectionProtectionSource

> SourceConnectionResponseParams TestConnectionProtectionSource(ctx).Body(body).Execute()

Test connection to a source.



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
    body := *openapiclient.NewSourceConnectionRequestParams("Environment_example") // SourceConnectionRequestParams | Specifies the parameters to test connectivity with a source.

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

    request := onprem.ApiTestConnectionProtectionSourceRequest{
        Body: &body
    }

    resp, r, err := api_client.Source.TestConnectionProtectionSource(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Source.TestConnectionProtectionSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestConnectionProtectionSource`: SourceConnectionResponseParams
    fmt.Fprintf(os.Stdout, "Response from `Source.TestConnectionProtectionSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestConnectionProtectionSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SourceConnectionRequestParams**](SourceConnectionRequestParams.md) | Specifies the parameters to test connectivity with a source. | 

### Return type

[**SourceConnectionResponseParams**](SourceConnectionResponseParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProtectionSourceRegistration

> SourceRegistration UpdateProtectionSourceRegistration(ctx, id).Body(body).Execute()

Update Protection Source registration.



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
    id := int64(789) // int64 | Specifies the id of the Protection Source registration.
    body := *openapiclient.NewSourceRegistrationUpdateRequestParams("Environment_example") // SourceRegistrationUpdateRequestParams | Specifies the parameters to update the registration.

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

    request := onprem.ApiUpdateProtectionSourceRegistrationRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.Source.UpdateProtectionSourceRegistration(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Source.UpdateProtectionSourceRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProtectionSourceRegistration`: SourceRegistration
    fmt.Fprintf(os.Stdout, "Response from `Source.UpdateProtectionSourceRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the Protection Source registration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProtectionSourceRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SourceRegistrationUpdateRequestParams**](SourceRegistrationUpdateRequestParams.md) | Specifies the parameters to update the registration. | 

### Return type

[**SourceRegistration**](SourceRegistration.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

