# \VaultsApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVault**](VaultsApi.md#CreateVault) | **Post** /public/vaults | Create a new Vault (External Target).
[**DeleteVault**](VaultsApi.md#DeleteVault) | **Delete** /public/vaults/{id} | Delete a Vault (External Target).
[**GetArchiveMediaInfo**](VaultsApi.md#GetArchiveMediaInfo) | **Get** /public/vaults/archiveMediaInfo | List the media information for the specified archive service.
[**GetBandwidthSettings**](VaultsApi.md#GetBandwidthSettings) | **Get** /public/vaults/bandwidthSettings | List the upload and download bandwidth limit and bandwidth overrides settings.
[**GetVaultById**](VaultsApi.md#GetVaultById) | **Get** /public/vaults/{id} | List details about a single Vault (External Target).
[**GetVaultEncryptionKey**](VaultsApi.md#GetVaultEncryptionKey) | **Get** /public/vaults/encryptionKey/{id} | Get encryption information for a Vault (External Target). A Vault is equivalent to an External Target in the Cohesity Dashboard.
[**GetVaults**](VaultsApi.md#GetVaults) | **Get** /public/vaults | List the Vaults (External Targets) registered on the Cohesity Cluster filtered by the specified parameters.
[**UpdateBandwidthSettings**](VaultsApi.md#UpdateBandwidthSettings) | **Put** /public/vaults/bandwidthSettings | Updates bandwidth limits.
[**UpdateVault**](VaultsApi.md#UpdateVault) | **Put** /public/vaults/{id} | Update a Vault (External Target).



## CreateVault

> Vault CreateVault(ctx).Body(body).Execute()

Create a new Vault (External Target).



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
    body := *openapiclient.NewVault() // Vault | Request to create a new Vault.

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

    request := cohesitysdk.ApiCreateVaultRequest{
        Body: &body
    }

    resp, r, err := api_client.VaultsApi.CreateVault(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultsApi.CreateVault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVault`: Vault
    fmt.Fprintf(os.Stdout, "Response from `VaultsApi.CreateVault`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Vault**](Vault.md) | Request to create a new Vault. | 

### Return type

[**Vault**](Vault.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVault

> DeleteVault(ctx, id).ForceDelete(forceDelete).Retry(retry).IncludeMarkedForRemoval(includeMarkedForRemoval).Body(body).Execute()

Delete a Vault (External Target).



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
    id := int64(789) // int64 | Specifies a unique id of the Vault.
    forceDelete := true // bool | Specifies whether to force delete the vault. If the flag is set to true, the RemovalState of the vault is changed to 'kMarkedForRemoval' and Eventually vault is removed from cluster config and archived metadata from scribe is removed without necessarily deleting the associated archived data. (optional)
    retry := true // bool | Specifies whether to retry a request after failure. (optional)
    includeMarkedForRemoval := true // bool | Specifies if Vaults that are marked for removal should be returned. (optional)
    body := *openapiclient.NewVaultDeleteParams() // VaultDeleteParams | Request to delete vault. (optional)

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

    request := cohesitysdk.ApiDeleteVaultRequest{
        Id: &id
        ForceDelete: &forceDelete
        Retry: &retry
        IncludeMarkedForRemoval: &includeMarkedForRemoval
        Body: &body
    }

    resp, r, err := api_client.VaultsApi.DeleteVault(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultsApi.DeleteVault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the Vault. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceDelete** | **bool** | Specifies whether to force delete the vault. If the flag is set to true, the RemovalState of the vault is changed to &#39;kMarkedForRemoval&#39; and Eventually vault is removed from cluster config and archived metadata from scribe is removed without necessarily deleting the associated archived data. | 
 **retry** | **bool** | Specifies whether to retry a request after failure. | 
 **includeMarkedForRemoval** | **bool** | Specifies if Vaults that are marked for removal should be returned. | 
 **body** | [**VaultDeleteParams**](VaultDeleteParams.md) | Request to delete vault. | 

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


## GetArchiveMediaInfo

> []TapeMediaInformation GetArchiveMediaInfo(ctx).ClusterId(clusterId).ClusterIncarnationId(clusterIncarnationId).QstarArchiveJobId(qstarArchiveJobId).QstarRestoreTaskId(qstarRestoreTaskId).EntityIds(entityIds).Execute()

List the media information for the specified archive service.



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
    clusterId := int64(789) // int64 | Specifies the id of the Cohesity Cluster that archived to a QStar media Vault.
    clusterIncarnationId := int64(789) // int64 | Specifies the incarnation id of the Cohesity Cluster that archived to a QStar media Vault.
    qstarArchiveJobId := int64(789) // int64 | Specifies the id of the Job that archived to a QStar media Vault.
    qstarRestoreTaskId := int64(789) // int64 | Specifies the id of the restore task to optionally filter by. The restore task that is restoring data from the specified media Vault. (optional)
    entityIds := []int64{int64(123)} // []int64 | Specifies an array of entityIds to optionally filter by. An entityId is a unique id for a VM assigned by the Cohesity Cluster. (optional)

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

    request := cohesitysdk.ApiGetArchiveMediaInfoRequest{
        ClusterId: &clusterId
        ClusterIncarnationId: &clusterIncarnationId
        QstarArchiveJobId: &qstarArchiveJobId
        QstarRestoreTaskId: &qstarRestoreTaskId
        EntityIds: &entityIds
    }

    resp, r, err := api_client.VaultsApi.GetArchiveMediaInfo(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultsApi.GetArchiveMediaInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArchiveMediaInfo`: []TapeMediaInformation
    fmt.Fprintf(os.Stdout, "Response from `VaultsApi.GetArchiveMediaInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetArchiveMediaInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **int64** | Specifies the id of the Cohesity Cluster that archived to a QStar media Vault. | 
 **clusterIncarnationId** | **int64** | Specifies the incarnation id of the Cohesity Cluster that archived to a QStar media Vault. | 
 **qstarArchiveJobId** | **int64** | Specifies the id of the Job that archived to a QStar media Vault. | 
 **qstarRestoreTaskId** | **int64** | Specifies the id of the restore task to optionally filter by. The restore task that is restoring data from the specified media Vault. | 
 **entityIds** | **[]int64** | Specifies an array of entityIds to optionally filter by. An entityId is a unique id for a VM assigned by the Cohesity Cluster. | 

### Return type

[**[]TapeMediaInformation**](TapeMediaInformation.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBandwidthSettings

> VaultBandwidthLimits GetBandwidthSettings(ctx).Execute()

List the upload and download bandwidth limit and bandwidth overrides settings.



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

    request := cohesitysdk.ApiGetBandwidthSettingsRequest{
    }

    resp, r, err := api_client.VaultsApi.GetBandwidthSettings(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultsApi.GetBandwidthSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBandwidthSettings`: VaultBandwidthLimits
    fmt.Fprintf(os.Stdout, "Response from `VaultsApi.GetBandwidthSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBandwidthSettingsRequest struct via the builder pattern


### Return type

[**VaultBandwidthLimits**](VaultBandwidthLimits.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVaultById

> Vault GetVaultById(ctx, id).Execute()

List details about a single Vault (External Target).



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
    id := int64(789) // int64 | Specifies a unique id of the Vault.

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

    request := cohesitysdk.ApiGetVaultByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.VaultsApi.GetVaultById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultsApi.GetVaultById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultById`: Vault
    fmt.Fprintf(os.Stdout, "Response from `VaultsApi.GetVaultById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the Vault. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Vault**](Vault.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVaultEncryptionKey

> VaultEncryptionKey GetVaultEncryptionKey(ctx, id).Execute()

Get encryption information for a Vault (External Target). A Vault is equivalent to an External Target in the Cohesity Dashboard.



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
    id := int64(789) // int64 | Specifies a unique id of the Vault.

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

    request := cohesitysdk.ApiGetVaultEncryptionKeyRequest{
        Id: &id
    }

    resp, r, err := api_client.VaultsApi.GetVaultEncryptionKey(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultsApi.GetVaultEncryptionKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultEncryptionKey`: VaultEncryptionKey
    fmt.Fprintf(os.Stdout, "Response from `VaultsApi.GetVaultEncryptionKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the Vault. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultEncryptionKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VaultEncryptionKey**](VaultEncryptionKey.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVaults

> []Vault GetVaults(ctx).Id(id).Name(name).IncludeMarkedForRemoval(includeMarkedForRemoval).Execute()

List the Vaults (External Targets) registered on the Cohesity Cluster filtered by the specified parameters.



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
    id := int64(789) // int64 | Specifies the id of Vault to return. If empty, all Vaults are returned. (optional)
    name := "name_example" // string | Specifies the name of the Vault to return. If empty, all Vaults are returned. (optional)
    includeMarkedForRemoval := true // bool | Specifies if Vaults that are marked for removal should be returned. (optional)

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

    request := cohesitysdk.ApiGetVaultsRequest{
        Id: &id
        Name: &name
        IncludeMarkedForRemoval: &includeMarkedForRemoval
    }

    resp, r, err := api_client.VaultsApi.GetVaults(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultsApi.GetVaults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaults`: []Vault
    fmt.Fprintf(os.Stdout, "Response from `VaultsApi.GetVaults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64** | Specifies the id of Vault to return. If empty, all Vaults are returned. | 
 **name** | **string** | Specifies the name of the Vault to return. If empty, all Vaults are returned. | 
 **includeMarkedForRemoval** | **bool** | Specifies if Vaults that are marked for removal should be returned. | 

### Return type

[**[]Vault**](Vault.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBandwidthSettings

> VaultBandwidthLimits UpdateBandwidthSettings(ctx).Body(body).Execute()

Updates bandwidth limits.



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
    body := *openapiclient.NewVaultBandwidthLimits() // VaultBandwidthLimits | Request to update global bandwidth limits settings.

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

    request := cohesitysdk.ApiUpdateBandwidthSettingsRequest{
        Body: &body
    }

    resp, r, err := api_client.VaultsApi.UpdateBandwidthSettings(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultsApi.UpdateBandwidthSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBandwidthSettings`: VaultBandwidthLimits
    fmt.Fprintf(os.Stdout, "Response from `VaultsApi.UpdateBandwidthSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBandwidthSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VaultBandwidthLimits**](VaultBandwidthLimits.md) | Request to update global bandwidth limits settings. | 

### Return type

[**VaultBandwidthLimits**](VaultBandwidthLimits.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVault

> Vault UpdateVault(ctx, id).Body(body).Execute()

Update a Vault (External Target).



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
    id := int64(789) // int64 | Specifies a unique id of the Vault.
    body := *openapiclient.NewVault() // Vault | Request to update a Vault's settings.

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

    request := cohesitysdk.ApiUpdateVaultRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.VaultsApi.UpdateVault(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultsApi.UpdateVault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVault`: Vault
    fmt.Fprintf(os.Stdout, "Response from `VaultsApi.UpdateVault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the Vault. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Vault**](Vault.md) | Request to update a Vault&#39;s settings. | 

### Return type

[**Vault**](Vault.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

