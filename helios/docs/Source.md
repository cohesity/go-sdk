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
[**McmDeleteProtectionSourceRegistration**](Source.md#McmDeleteProtectionSourceRegistration) | **Delete** /mcm/data-protect/sources/registrations/{id} | Delete Protection Source Registration.
[**McmGetProtectionSourceRegistration**](Source.md#McmGetProtectionSourceRegistration) | **Get** /mcm/data-protect/sources/registrations/{id} | Get a Protection Source registration.
[**McmGetProtectionSources**](Source.md#McmGetProtectionSources) | **Get** /mcm/data-protect/sources | Get a List of Protection Sources.
[**McmRegisterProtectionSource**](Source.md#McmRegisterProtectionSource) | **Post** /mcm/data-protect/sources/registrations | Register a Protection Source.
[**McmTestSourceConnection**](Source.md#McmTestSourceConnection) | **Post** /mcm/data-protect/sources/test-connection | Test connection to a source.
[**ProtectionSourceById**](Source.md#ProtectionSourceById) | **Get** /data-protect/sources/{id} | Get a Protection Sources.
[**RegisterProtectionSource**](Source.md#RegisterProtectionSource) | **Post** /data-protect/sources/registrations | Register a Protection Source.
[**TestConnectionProtectionSource**](Source.md#TestConnectionProtectionSource) | **Post** /data-protect/sources/test-connection | Test connection to a source.
[**UpdateProtectionSourceRegistration**](Source.md#UpdateProtectionSourceRegistration) | **Put** /data-protect/sources/registrations/{id} | Update Protection Source registration.
[**UpdateProtectionSourceRegistrationMixin1**](Source.md#UpdateProtectionSourceRegistrationMixin1) | **Put** /mcm/data-protect/sources/registrations/{id} | Update Protection Source registration.



## DeleteProtectionSourceRegistration

> DeleteProtectionSourceRegistration(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Delete Protection Source Registration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := int64(789) // int64 | Specifies the ID of the Protection Source Registration.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiDeleteProtectionSourceRegistrationRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> SourceRegistration GetProtectionSourceRegistration(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get a Protection Source registration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := int64(789) // int64 | Specifies the id of the Protection Source registration.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiGetProtectionSourceRegistrationRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> Sources GetProtectionSources(ctx).AccessClusterId(accessClusterId).RegionId(regionId).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeSourceCredentials(includeSourceCredentials).EncryptionKey(encryptionKey).Execute()

Get a List of Protection Sources.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
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

    request := helios.ApiGetProtectionSourcesRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
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

> SourceAttributeFiltersResponseParams GetSourceAttributeFilters(ctx).SourceUuid(sourceUuid).AccessClusterId(accessClusterId).RegionId(regionId).Environment(environment).Execute()

List attribute filters for a source.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    sourceUuid := "sourceUuid_example" // string | Specifies the source UUID of the parent entity.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
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

    request := helios.ApiGetSourceAttributeFiltersRequest{
        SourceUuid: &sourceUuid
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
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

> SourceRegistrations GetSourceRegistrations(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Ids(ids).TenantIds(tenantIds).IncludeTenants(includeTenants).IncludeSourceCredentials(includeSourceCredentials).EncryptionKey(encryptionKey).Execute()

Get the list of Protection Source registrations.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
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

    request := helios.ApiGetSourceRegistrationsRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
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

> VdcObject GetVdcDetails(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get VDC Details.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := int64(789) // int64 | Specifies the ID of the VMware virtual datacenter.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiGetVdcDetailsRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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


## McmDeleteProtectionSourceRegistration

> McmDeleteProtectionSourceRegistration(ctx, id).RegionId(regionId).Execute()

Delete Protection Source Registration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies the ID of the Protection Source Registration.
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiMcmDeleteProtectionSourceRegistrationRequest{
        Id: &id
        RegionId: &regionId
    }

    resp, r, err := api_client.Source.McmDeleteProtectionSourceRegistration(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Source.McmDeleteProtectionSourceRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the ID of the Protection Source Registration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcmDeleteProtectionSourceRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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


## McmGetProtectionSourceRegistration

> McmSourceRegistration McmGetProtectionSourceRegistration(ctx, id).RegionId(regionId).Execute()

Get a Protection Source registration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies the id of the Protection Source registration.
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiMcmGetProtectionSourceRegistrationRequest{
        Id: &id
        RegionId: &regionId
    }

    resp, r, err := api_client.Source.McmGetProtectionSourceRegistration(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Source.McmGetProtectionSourceRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `McmGetProtectionSourceRegistration`: McmSourceRegistration
    fmt.Fprintf(os.Stdout, "Response from `Source.McmGetProtectionSourceRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the Protection Source registration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMcmGetProtectionSourceRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**McmSourceRegistration**](McmSourceRegistration.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McmGetProtectionSources

> McmSources McmGetProtectionSources(ctx).RegionId(regionId).Environments(environments).Ids(ids).RegionIds(regionIds).ExcludeProtectionStats(excludeProtectionStats).Execute()

Get a List of Protection Sources.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    environments := []string{"Environments_example"} // []string | Specifies the list of environment type of the Protection Source. (optional)
    ids := []string{"Inner_example"} // []string | Specifies the list of ids to filter Protection Sources. (optional)
    regionIds := []string{"Inner_example"} // []string | Filter by a list of region ids. (optional)
    excludeProtectionStats := true // bool | Whether to exclude Protection Sources protection stats in response. (optional)

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

    request := helios.ApiMcmGetProtectionSourcesRequest{
        RegionId: &regionId
        Environments: &environments
        Ids: &ids
        RegionIds: &regionIds
        ExcludeProtectionStats: &excludeProtectionStats
    }

    resp, r, err := api_client.Source.McmGetProtectionSources(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Source.McmGetProtectionSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `McmGetProtectionSources`: McmSources
    fmt.Fprintf(os.Stdout, "Response from `Source.McmGetProtectionSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMcmGetProtectionSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **environments** | **[]string** | Specifies the list of environment type of the Protection Source. | 
 **ids** | **[]string** | Specifies the list of ids to filter Protection Sources. | 
 **regionIds** | **[]string** | Filter by a list of region ids. | 
 **excludeProtectionStats** | **bool** | Whether to exclude Protection Sources protection stats in response. | 

### Return type

[**McmSources**](McmSources.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McmRegisterProtectionSource

> McmSourceRegistration McmRegisterProtectionSource(ctx).Body(body).RegionId(regionId).AccessClusterId(accessClusterId).Execute()

Register a Protection Source.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    body := *openapiclient.NewMcmSourceRegistrationRequestParams("Environment_example") // McmSourceRegistrationRequestParams | Specifies the parameters to register a Protection Source.
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    accessClusterId := int64(789) // int64 | Specifies the destination cluster id on which this Source needs to be registered. (optional)

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

    request := helios.ApiMcmRegisterProtectionSourceRequest{
        Body: &body
        RegionId: &regionId
        AccessClusterId: &accessClusterId
    }

    resp, r, err := api_client.Source.McmRegisterProtectionSource(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Source.McmRegisterProtectionSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `McmRegisterProtectionSource`: McmSourceRegistration
    fmt.Fprintf(os.Stdout, "Response from `Source.McmRegisterProtectionSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMcmRegisterProtectionSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**McmSourceRegistrationRequestParams**](McmSourceRegistrationRequestParams.md) | Specifies the parameters to register a Protection Source. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **accessClusterId** | **int64** | Specifies the destination cluster id on which this Source needs to be registered. | 

### Return type

[**McmSourceRegistration**](McmSourceRegistration.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## McmTestSourceConnection

> SourceConnectionResponseParams McmTestSourceConnection(ctx).Body(body).RegionId(regionId).AccessClusterId(accessClusterId).Execute()

Test connection to a source.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    body := *openapiclient.NewSourceConnectionRequestParams("Environment_example") // SourceConnectionRequestParams | Specifies the parameters to test connectivity of a Protection Source.
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    accessClusterId := int64(789) // int64 | Specifies the destination cluster id on which this Source needs to be registered. (optional)

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

    request := helios.ApiMcmTestSourceConnectionRequest{
        Body: &body
        RegionId: &regionId
        AccessClusterId: &accessClusterId
    }

    resp, r, err := api_client.Source.McmTestSourceConnection(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Source.McmTestSourceConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `McmTestSourceConnection`: SourceConnectionResponseParams
    fmt.Fprintf(os.Stdout, "Response from `Source.McmTestSourceConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMcmTestSourceConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SourceConnectionRequestParams**](SourceConnectionRequestParams.md) | Specifies the parameters to test connectivity of a Protection Source. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **accessClusterId** | **int64** | Specifies the destination cluster id on which this Source needs to be registered. | 

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


## ProtectionSourceById

> Source ProtectionSourceById(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Get a Protection Sources.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := int64(789) // int64 | Specifies the id of the Protection Source.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiProtectionSourceByIdRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> SourceRegistration RegisterProtectionSource(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Register a Protection Source.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    body := *openapiclient.NewSourceRegistrationRequestParams("Environment_example") // SourceRegistrationRequestParams | Specifies the parameters to register a Protection Source.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiRegisterProtectionSourceRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> SourceConnectionResponseParams TestConnectionProtectionSource(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Test connection to a source.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    body := *openapiclient.NewSourceConnectionRequestParams("Environment_example") // SourceConnectionRequestParams | Specifies the parameters to test connectivity with a source.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiTestConnectionProtectionSourceRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> SourceRegistration UpdateProtectionSourceRegistration(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update Protection Source registration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := int64(789) // int64 | Specifies the id of the Protection Source registration.
    body := *openapiclient.NewSourceRegistrationUpdateRequestParams("Environment_example") // SourceRegistrationUpdateRequestParams | Specifies the parameters to update the registration.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiUpdateProtectionSourceRegistrationRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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


## UpdateProtectionSourceRegistrationMixin1

> McmSourceRegistration UpdateProtectionSourceRegistrationMixin1(ctx, id).Body(body).RegionId(regionId).Execute()

Update Protection Source registration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies the id of the Protection Source registration.
    body := *openapiclient.NewMcmSourceRegistrationUpdateRequestParams("Environment_example") // McmSourceRegistrationUpdateRequestParams | Specifies the parameters to update the registration.
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiUpdateProtectionSourceRegistrationMixin1Request{
        Id: &id
        Body: &body
        RegionId: &regionId
    }

    resp, r, err := api_client.Source.UpdateProtectionSourceRegistrationMixin1(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Source.UpdateProtectionSourceRegistrationMixin1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProtectionSourceRegistrationMixin1`: McmSourceRegistration
    fmt.Fprintf(os.Stdout, "Response from `Source.UpdateProtectionSourceRegistrationMixin1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the Protection Source registration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProtectionSourceRegistrationMixin1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**McmSourceRegistrationUpdateRequestParams**](McmSourceRegistrationUpdateRequestParams.md) | Specifies the parameters to update the registration. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**McmSourceRegistration**](McmSourceRegistration.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

