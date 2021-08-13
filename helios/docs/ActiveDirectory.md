# \ActiveDirectory

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateActiveDirectory**](ActiveDirectory.md#CreateActiveDirectory) | **Post** /active-directories | Create an Active Directory.
[**DeleteActiveDirectory**](ActiveDirectory.md#DeleteActiveDirectory) | **Delete** /active-directories/{id} | Delete an Active Directory.
[**GetActiveDirectory**](ActiveDirectory.md#GetActiveDirectory) | **Get** /active-directories | Get the list of Active Directories.
[**GetActiveDirectoryById**](ActiveDirectory.md#GetActiveDirectoryById) | **Get** /active-directories/{id} | Get an Active Directory by id.
[**GetDomainControllers**](ActiveDirectory.md#GetDomainControllers) | **Get** /domain-controllers | Get Domain Controllers of specified domains.
[**UpdateActiveDirectory**](ActiveDirectory.md#UpdateActiveDirectory) | **Put** /active-directories/{id} | Update an Active Directory.



## CreateActiveDirectory

> ActiveDirectory CreateActiveDirectory(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Create an Active Directory.



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
    body := *openapiclient.NewCreateActiveDirectoryRequest([]openapiclient.MachineAccount{*openapiclient.NewMachineAccount("Name_example")}, "DomainName_example", "TODO") // CreateActiveDirectoryRequest | Specifies the parameters to create an Active Directory.
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

    request := helios.ApiCreateActiveDirectoryRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.ActiveDirectory.CreateActiveDirectory(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectory.CreateActiveDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateActiveDirectory`: ActiveDirectory
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectory.CreateActiveDirectory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateActiveDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateActiveDirectoryRequest**](CreateActiveDirectoryRequest.md) | Specifies the parameters to create an Active Directory. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**ActiveDirectory**](ActiveDirectory.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteActiveDirectory

> DeleteActiveDirectory(ctx, id).ActiveDirectoryAdminUsername(activeDirectoryAdminUsername).ActiveDirectoryAdminPassword(activeDirectoryAdminPassword).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Delete an Active Directory.



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
    id := int64(789) // int64 | Specifies id of an Active Directory.
    activeDirectoryAdminUsername := "activeDirectoryAdminUsername_example" // string | Specifies the username of the Active Direcotry Admin.
    activeDirectoryAdminPassword := "activeDirectoryAdminPassword_example" // string | Specifies the password of the Active Direcotry Admin.
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

    request := helios.ApiDeleteActiveDirectoryRequest{
        Id: &id
        ActiveDirectoryAdminUsername: &activeDirectoryAdminUsername
        ActiveDirectoryAdminPassword: &activeDirectoryAdminPassword
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.ActiveDirectory.DeleteActiveDirectory(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectory.DeleteActiveDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies id of an Active Directory. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteActiveDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activeDirectoryAdminUsername** | **string** | Specifies the username of the Active Direcotry Admin. | 
 **activeDirectoryAdminPassword** | **string** | Specifies the password of the Active Direcotry Admin. | 
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


## GetActiveDirectory

> ActiveDirectories GetActiveDirectory(ctx).AccessClusterId(accessClusterId).RegionId(regionId).DomainNames(domainNames).Ids(ids).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()

Get the list of Active Directories.



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
    domainNames := []string{"Inner_example"} // []string | Filter by a list of Active Directory domain names. (optional)
    ids := []int64{int64(123)} // []int64 | Filter by a list of Active Directory Ids. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which Active Directories are to be returned. (optional)
    includeTenants := true // bool | If true, the response will include Active Directories which were created by all tenants which the current user has permission to see. If false, then only Active Directories created by the current user will be returned. (optional)

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

    request := helios.ApiGetActiveDirectoryRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        DomainNames: &domainNames
        Ids: &ids
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
    }

    resp, r, err := api_client.ActiveDirectory.GetActiveDirectory(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectory.GetActiveDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveDirectory`: ActiveDirectories
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectory.GetActiveDirectory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **domainNames** | **[]string** | Filter by a list of Active Directory domain names. | 
 **ids** | **[]int64** | Filter by a list of Active Directory Ids. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which Active Directories are to be returned. | 
 **includeTenants** | **bool** | If true, the response will include Active Directories which were created by all tenants which the current user has permission to see. If false, then only Active Directories created by the current user will be returned. | 

### Return type

[**ActiveDirectories**](ActiveDirectories.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveDirectoryById

> ActiveDirectory GetActiveDirectoryById(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).IncludeCentrifyZones(includeCentrifyZones).IncludeDomainControllers(includeDomainControllers).IncludeSecurityPrincipals(includeSecurityPrincipals).Prefix(prefix).ObjectClass(objectClass).Execute()

Get an Active Directory by id.



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
    id := int64(789) // int64 | Specifies id of an Active Directory.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    includeCentrifyZones := true // bool | Specifies whether to include Centrify Zones of the Active Directory in response. (optional)
    includeDomainControllers := true // bool | Specifies whether to include Domain Controllers of the Active Directory in response. (optional)
    includeSecurityPrincipals := true // bool | Specifies whether to include Security Principals of the Active Directory in response. (optional)
    prefix := "prefix_example" // string | Specifies a prefix, only security principals with name or sAMAccountName having this prefix (ignoring cases) will be returned. This field is appliciable and mandatory if 'includeSecurityPrincipals' is set to true. (optional)
    objectClass := []string{"ObjectClass_example"} // []string | Specifies a list of object classes, only security principals with object class in this list will be returned. This field is appliciable if 'includeSecurityPrincipals' is set to true. (optional)

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

    request := helios.ApiGetActiveDirectoryByIdRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        IncludeCentrifyZones: &includeCentrifyZones
        IncludeDomainControllers: &includeDomainControllers
        IncludeSecurityPrincipals: &includeSecurityPrincipals
        Prefix: &prefix
        ObjectClass: &objectClass
    }

    resp, r, err := api_client.ActiveDirectory.GetActiveDirectoryById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectory.GetActiveDirectoryById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveDirectoryById`: ActiveDirectory
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectory.GetActiveDirectoryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies id of an Active Directory. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveDirectoryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **includeCentrifyZones** | **bool** | Specifies whether to include Centrify Zones of the Active Directory in response. | 
 **includeDomainControllers** | **bool** | Specifies whether to include Domain Controllers of the Active Directory in response. | 
 **includeSecurityPrincipals** | **bool** | Specifies whether to include Security Principals of the Active Directory in response. | 
 **prefix** | **string** | Specifies a prefix, only security principals with name or sAMAccountName having this prefix (ignoring cases) will be returned. This field is appliciable and mandatory if &#39;includeSecurityPrincipals&#39; is set to true. | 
 **objectClass** | **[]string** | Specifies a list of object classes, only security principals with object class in this list will be returned. This field is appliciable if &#39;includeSecurityPrincipals&#39; is set to true. | 

### Return type

[**ActiveDirectory**](ActiveDirectory.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainControllers

> DomainControllersResponse GetDomainControllers(ctx).DomainNames(domainNames).AccessClusterId(accessClusterId).RegionId(regionId).ConnectionId(connectionId).Execute()

Get Domain Controllers of specified domains.



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
    domainNames := []string{"Inner_example"} // []string | Specifies a list of domain names.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    connectionId := int64(789) // int64 | Specifies the Id of the connection which the connector belongs to. (optional)

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

    request := helios.ApiGetDomainControllersRequest{
        DomainNames: &domainNames
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        ConnectionId: &connectionId
    }

    resp, r, err := api_client.ActiveDirectory.GetDomainControllers(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectory.GetDomainControllers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomainControllers`: DomainControllersResponse
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectory.GetDomainControllers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainControllersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainNames** | **[]string** | Specifies a list of domain names. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **connectionId** | **int64** | Specifies the Id of the connection which the connector belongs to. | 

### Return type

[**DomainControllersResponse**](DomainControllersResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateActiveDirectory

> ActiveDirectory UpdateActiveDirectory(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update an Active Directory.



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
    id := int64(789) // int64 | Specifies id of an Active Directory.
    body := *openapiclient.NewUpdateActiveDirectoryRequest([]openapiclient.MachineAccount{*openapiclient.NewMachineAccount("Name_example")}) // UpdateActiveDirectoryRequest | Request to update an Active Directory.
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

    request := helios.ApiUpdateActiveDirectoryRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.ActiveDirectory.UpdateActiveDirectory(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectory.UpdateActiveDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateActiveDirectory`: ActiveDirectory
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectory.UpdateActiveDirectory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies id of an Active Directory. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateActiveDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateActiveDirectoryRequest**](UpdateActiveDirectoryRequest.md) | Request to update an Active Directory. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**ActiveDirectory**](ActiveDirectory.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

