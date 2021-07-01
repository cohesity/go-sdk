# \ActiveDirectoryApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddActiveDirectoryPrincipals**](ActiveDirectoryApi.md#AddActiveDirectoryPrincipals) | **Post** /public/activeDirectory/principals | Add multiple groups or users on the Cohesity Cluster for the specified Active Directory principals. In addition, assign Cohesity roles to the users or groups to define their Cohesity privileges.
[**CreateActiveDirectoryEntry**](ActiveDirectoryApi.md#CreateActiveDirectoryEntry) | **Post** /public/activeDirectory | Join the Cohesity Cluster to the specified Active Directory.
[**DeleteActiveDirectoryEntry**](ActiveDirectoryApi.md#DeleteActiveDirectoryEntry) | **Delete** /public/activeDirectory | Deletes the join with the Active Directory.
[**EnableTrustedDomainDiscovery**](ActiveDirectoryApi.md#EnableTrustedDomainDiscovery) | **Post** /public/activeDirectory/{name}/enableTrustedDomainState | Updates the states of trusted domains discovery.
[**GetActiveDirectoryDomainControllers**](ActiveDirectoryApi.md#GetActiveDirectoryDomainControllers) | **Get** /public/activeDirectory/domainControllers | List the domain controllers for a domain.
[**GetActiveDirectoryEntry**](ActiveDirectoryApi.md#GetActiveDirectoryEntry) | **Get** /public/activeDirectory | Lists the Active Directories that the Cohesity Cluster has joined.
[**ListCentrifyZones**](ActiveDirectoryApi.md#ListCentrifyZones) | **Get** /public/activeDirectory/centrifyZones | Fetches the list centrify zones of an active directory domain.
[**SearchActiveDirectoryPrincipals**](ActiveDirectoryApi.md#SearchActiveDirectoryPrincipals) | **Get** /public/activeDirectory/principals | List the user and group principals in the Active Directory that match the filter criteria specified using parameters.
[**UpdateActiveDirectoryIdMapping**](ActiveDirectoryApi.md#UpdateActiveDirectoryIdMapping) | **Put** /public/activeDirectory/{name}/idMappingInfo | Updates the user id mapping info of an Active Directory.
[**UpdateActiveDirectoryIgnoredTrustedDomains**](ActiveDirectoryApi.md#UpdateActiveDirectoryIgnoredTrustedDomains) | **Put** /public/activeDirectory/{name}/ignoredTrustedDomains | Updates the list of trusted domains to be ignored during trusted domain discovery of an Active Directory.
[**UpdateActiveDirectoryLdapProvider**](ActiveDirectoryApi.md#UpdateActiveDirectoryLdapProvider) | **Put** /public/activeDirectory/{name}/ldapProvider | Updates the LDAP provide Id for an Active Directory domain.
[**UpdateActiveDirectoryMachineAccounts**](ActiveDirectoryApi.md#UpdateActiveDirectoryMachineAccounts) | **Post** /public/activeDirectory/{name}/machineAccounts | Updates the machine accounts of an Active Directory.
[**UpdatePreferredDomainControllers**](ActiveDirectoryApi.md#UpdatePreferredDomainControllers) | **Put** /public/activeDirectory/{name}/preferredDomainControllers | 



## AddActiveDirectoryPrincipals

> []AddedActiveDirectoryPrincipal AddActiveDirectoryPrincipals(ctx).Body(body).Execute()

Add multiple groups or users on the Cohesity Cluster for the specified Active Directory principals. In addition, assign Cohesity roles to the users or groups to define their Cohesity privileges.



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
    body := []openapiclient.ActiveDirectoryPrincipalsAddParameters{*openapiclient.NewActiveDirectoryPrincipalsAddParameters()} // []ActiveDirectoryPrincipalsAddParameters | Request to add groups or users to the Cohesity Cluster. (optional)

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

    request := cohesitysdk.ApiAddActiveDirectoryPrincipalsRequest{
        body: &Body
    }

    resp, r, err := api_client.ActiveDirectoryApi.AddActiveDirectoryPrincipals(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.AddActiveDirectoryPrincipals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddActiveDirectoryPrincipals`: []AddedActiveDirectoryPrincipal
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.AddActiveDirectoryPrincipals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddActiveDirectoryPrincipalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]ActiveDirectoryPrincipalsAddParameters**](ActiveDirectoryPrincipalsAddParameters.md) | Request to add groups or users to the Cohesity Cluster. | 

### Return type

[**[]AddedActiveDirectoryPrincipal**](AddedActiveDirectoryPrincipal.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateActiveDirectoryEntry

> ActiveDirectoryEntry CreateActiveDirectoryEntry(ctx).Body(body).Execute()

Join the Cohesity Cluster to the specified Active Directory.



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
    body := *openapiclient.NewCreateActiveDirectoryEntryParams() // CreateActiveDirectoryEntryParams | Request to join an Active Directory.

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

    request := cohesitysdk.ApiCreateActiveDirectoryEntryRequest{
        body: &Body
    }

    resp, r, err := api_client.ActiveDirectoryApi.CreateActiveDirectoryEntry(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.CreateActiveDirectoryEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateActiveDirectoryEntry`: ActiveDirectoryEntry
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.CreateActiveDirectoryEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateActiveDirectoryEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateActiveDirectoryEntryParams**](CreateActiveDirectoryEntryParams.md) | Request to join an Active Directory. | 

### Return type

[**ActiveDirectoryEntry**](ActiveDirectoryEntry.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteActiveDirectoryEntry

> DeleteActiveDirectoryEntry(ctx).Body(body).Execute()

Deletes the join with the Active Directory.



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
    body := *openapiclient.NewActiveDirectoryEntry() // ActiveDirectoryEntry | Request to delete a join with an Active Directory.

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

    request := cohesitysdk.ApiDeleteActiveDirectoryEntryRequest{
        body: &Body
    }

    resp, r, err := api_client.ActiveDirectoryApi.DeleteActiveDirectoryEntry(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.DeleteActiveDirectoryEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteActiveDirectoryEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ActiveDirectoryEntry**](ActiveDirectoryEntry.md) | Request to delete a join with an Active Directory. | 

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


## EnableTrustedDomainDiscovery

> ActiveDirectoryEntry EnableTrustedDomainDiscovery(ctx, name).Body(body).Execute()

Updates the states of trusted domains discovery.

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
    name := "name_example" // string | Specifies the Active Directory Domain Name.
    body := *openapiclient.NewUpdateTrustedDomainEnableParams() // UpdateTrustedDomainEnableParams | Request to update enable trusted domains state of an Active Directory.

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

    request := cohesitysdk.ApiEnableTrustedDomainDiscoveryRequest{
        name: &Name
        body: &Body
    }

    resp, r, err := api_client.ActiveDirectoryApi.EnableTrustedDomainDiscovery(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.EnableTrustedDomainDiscovery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableTrustedDomainDiscovery`: ActiveDirectoryEntry
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.EnableTrustedDomainDiscovery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the Active Directory Domain Name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableTrustedDomainDiscoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateTrustedDomainEnableParams**](UpdateTrustedDomainEnableParams.md) | Request to update enable trusted domains state of an Active Directory. | 

### Return type

[**ActiveDirectoryEntry**](ActiveDirectoryEntry.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveDirectoryDomainControllers

> DomainControllers GetActiveDirectoryDomainControllers(ctx).DomainName(domainName).Execute()

List the domain controllers for a domain.

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
    domainName := "domainName_example" // string | Specifies the domain name (optional)

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

    request := cohesitysdk.ApiGetActiveDirectoryDomainControllersRequest{
        domainName: &DomainName
    }

    resp, r, err := api_client.ActiveDirectoryApi.GetActiveDirectoryDomainControllers(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.GetActiveDirectoryDomainControllers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveDirectoryDomainControllers`: DomainControllers
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.GetActiveDirectoryDomainControllers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveDirectoryDomainControllersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainName** | **string** | Specifies the domain name | 

### Return type

[**DomainControllers**](DomainControllers.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveDirectoryEntry

> []ActiveDirectoryEntry GetActiveDirectoryEntry(ctx).Domains(domains).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).Execute()

Lists the Active Directories that the Cohesity Cluster has joined.



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
    domains := []string{"Inner_example"} // []string | Specifies the domains to fetch active directory entries. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    allUnderHierarchy := true // bool | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)

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

    request := cohesitysdk.ApiGetActiveDirectoryEntryRequest{
        domains: &Domains
        tenantIds: &TenantIds
        allUnderHierarchy: &AllUnderHierarchy
    }

    resp, r, err := api_client.ActiveDirectoryApi.GetActiveDirectoryEntry(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.GetActiveDirectoryEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveDirectoryEntry`: []ActiveDirectoryEntry
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.GetActiveDirectoryEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveDirectoryEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domains** | **[]string** | Specifies the domains to fetch active directory entries. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 

### Return type

[**[]ActiveDirectoryEntry**](ActiveDirectoryEntry.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCentrifyZones

> []ListCentrifyZone ListCentrifyZones(ctx).DomainName(domainName).Execute()

Fetches the list centrify zones of an active directory domain.

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
    domainName := "domainName_example" // string | Specifies the fully qualified domain name (FQDN) of an Active Directory. (optional)

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

    request := cohesitysdk.ApiListCentrifyZonesRequest{
        domainName: &DomainName
    }

    resp, r, err := api_client.ActiveDirectoryApi.ListCentrifyZones(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.ListCentrifyZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCentrifyZones`: []ListCentrifyZone
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.ListCentrifyZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCentrifyZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainName** | **string** | Specifies the fully qualified domain name (FQDN) of an Active Directory. | 

### Return type

[**[]ListCentrifyZone**](ListCentrifyZone.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchActiveDirectoryPrincipals

> []ActiveDirectoryPrincipal SearchActiveDirectoryPrincipals(ctx).Domain(domain).ObjectClass(objectClass).Search(search).Sids(sids).IncludeComputers(includeComputers).Execute()

List the user and group principals in the Active Directory that match the filter criteria specified using parameters.



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
    domain := "domain_example" // string | Specifies the domain name of the principals to search. If specified the principals in that domain are searched. Domain could be an Active Directory domain joined by the Cluster or any one of the trusted domains of the Active Directory domain or the LOCAL domain. If not specified, all the domains are searched. (optional)
    objectClass := "objectClass_example" // string | Optionally filter by a principal object class such as 'kGroup' or 'kUser'. If 'kGroup' is specified, only group principals are returned. If 'kUser' is specified, only user principals are returned. If not specified, both group and user principals are returned. 'kUser' specifies a user object class. 'kGroup' specifies a group object class. 'kComputer' specifies a computer object class. 'kWellKnownPrincipal' specifies a well known principal. (optional)
    search := "search_example" // string | Optionally filter by matching a substring. Only principals in the with a name or sAMAccountName that matches part or all of the specified substring are returned. If specified, a 'sids' parameter should not be specified. (optional)
    sids := []string{"Inner_example"} // []string | Optionally filter by a list of security identifiers (SIDs) found in the specified domain. Only principals matching the specified SIDs are returned. If specified, a 'search' parameter should not be specified. (optional)
    includeComputers := true // bool | Specifies if Computer/GMSA accounts need to be included in this search. (optional)

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

    request := cohesitysdk.ApiSearchActiveDirectoryPrincipalsRequest{
        domain: &Domain
        objectClass: &ObjectClass
        search: &Search
        sids: &Sids
        includeComputers: &IncludeComputers
    }

    resp, r, err := api_client.ActiveDirectoryApi.SearchActiveDirectoryPrincipals(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.SearchActiveDirectoryPrincipals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchActiveDirectoryPrincipals`: []ActiveDirectoryPrincipal
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.SearchActiveDirectoryPrincipals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchActiveDirectoryPrincipalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | Specifies the domain name of the principals to search. If specified the principals in that domain are searched. Domain could be an Active Directory domain joined by the Cluster or any one of the trusted domains of the Active Directory domain or the LOCAL domain. If not specified, all the domains are searched. | 
 **objectClass** | **string** | Optionally filter by a principal object class such as &#39;kGroup&#39; or &#39;kUser&#39;. If &#39;kGroup&#39; is specified, only group principals are returned. If &#39;kUser&#39; is specified, only user principals are returned. If not specified, both group and user principals are returned. &#39;kUser&#39; specifies a user object class. &#39;kGroup&#39; specifies a group object class. &#39;kComputer&#39; specifies a computer object class. &#39;kWellKnownPrincipal&#39; specifies a well known principal. | 
 **search** | **string** | Optionally filter by matching a substring. Only principals in the with a name or sAMAccountName that matches part or all of the specified substring are returned. If specified, a &#39;sids&#39; parameter should not be specified. | 
 **sids** | **[]string** | Optionally filter by a list of security identifiers (SIDs) found in the specified domain. Only principals matching the specified SIDs are returned. If specified, a &#39;search&#39; parameter should not be specified. | 
 **includeComputers** | **bool** | Specifies if Computer/GMSA accounts need to be included in this search. | 

### Return type

[**[]ActiveDirectoryPrincipal**](ActiveDirectoryPrincipal.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateActiveDirectoryIdMapping

> ActiveDirectoryEntry UpdateActiveDirectoryIdMapping(ctx, name).Body(body).Execute()

Updates the user id mapping info of an Active Directory.

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
    name := "name_example" // string | Specifies the Active Directory Domain Name.
    body := *openapiclient.NewIdMappingInfo() // IdMappingInfo | Request to update user id mapping of an Active Directory.

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

    request := cohesitysdk.ApiUpdateActiveDirectoryIdMappingRequest{
        name: &Name
        body: &Body
    }

    resp, r, err := api_client.ActiveDirectoryApi.UpdateActiveDirectoryIdMapping(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.UpdateActiveDirectoryIdMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateActiveDirectoryIdMapping`: ActiveDirectoryEntry
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.UpdateActiveDirectoryIdMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the Active Directory Domain Name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateActiveDirectoryIdMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IdMappingInfo**](IdMappingInfo.md) | Request to update user id mapping of an Active Directory. | 

### Return type

[**ActiveDirectoryEntry**](ActiveDirectoryEntry.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateActiveDirectoryIgnoredTrustedDomains

> ActiveDirectoryEntry UpdateActiveDirectoryIgnoredTrustedDomains(ctx, name).Body(body).Execute()

Updates the list of trusted domains to be ignored during trusted domain discovery of an Active Directory.

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
    name := "name_example" // string | Specifies the Active Directory Domain Name.
    body := *openapiclient.NewUpdateIgnoredTrustedDomainsParams() // UpdateIgnoredTrustedDomainsParams | Request to update the list of ignored trusted domains of an AD.

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

    request := cohesitysdk.ApiUpdateActiveDirectoryIgnoredTrustedDomainsRequest{
        name: &Name
        body: &Body
    }

    resp, r, err := api_client.ActiveDirectoryApi.UpdateActiveDirectoryIgnoredTrustedDomains(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.UpdateActiveDirectoryIgnoredTrustedDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateActiveDirectoryIgnoredTrustedDomains`: ActiveDirectoryEntry
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.UpdateActiveDirectoryIgnoredTrustedDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the Active Directory Domain Name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateActiveDirectoryIgnoredTrustedDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateIgnoredTrustedDomainsParams**](UpdateIgnoredTrustedDomainsParams.md) | Request to update the list of ignored trusted domains of an AD. | 

### Return type

[**ActiveDirectoryEntry**](ActiveDirectoryEntry.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateActiveDirectoryLdapProvider

> ActiveDirectoryEntry UpdateActiveDirectoryLdapProvider(ctx, name).Body(body).Execute()

Updates the LDAP provide Id for an Active Directory domain.

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
    name := "name_example" // string | Specifies the Active Directory Domain Name.
    body := *openapiclient.NewUpdateLdapProviderParams() // UpdateLdapProviderParams | Request to update the LDAP provider info.

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

    request := cohesitysdk.ApiUpdateActiveDirectoryLdapProviderRequest{
        name: &Name
        body: &Body
    }

    resp, r, err := api_client.ActiveDirectoryApi.UpdateActiveDirectoryLdapProvider(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.UpdateActiveDirectoryLdapProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateActiveDirectoryLdapProvider`: ActiveDirectoryEntry
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.UpdateActiveDirectoryLdapProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the Active Directory Domain Name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateActiveDirectoryLdapProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateLdapProviderParams**](UpdateLdapProviderParams.md) | Request to update the LDAP provider info. | 

### Return type

[**ActiveDirectoryEntry**](ActiveDirectoryEntry.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateActiveDirectoryMachineAccounts

> ActiveDirectoryEntry UpdateActiveDirectoryMachineAccounts(ctx, name).Body(body).Execute()

Updates the machine accounts of an Active Directory.

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
    name := "name_example" // string | Specifies the Active Directory Domain Name.
    body := *openapiclient.NewUpdateMachineAccountsParams() // UpdateMachineAccountsParams | Request to update machine accounts of an Active Directory.

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

    request := cohesitysdk.ApiUpdateActiveDirectoryMachineAccountsRequest{
        name: &Name
        body: &Body
    }

    resp, r, err := api_client.ActiveDirectoryApi.UpdateActiveDirectoryMachineAccounts(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.UpdateActiveDirectoryMachineAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateActiveDirectoryMachineAccounts`: ActiveDirectoryEntry
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.UpdateActiveDirectoryMachineAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the Active Directory Domain Name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateActiveDirectoryMachineAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateMachineAccountsParams**](UpdateMachineAccountsParams.md) | Request to update machine accounts of an Active Directory. | 

### Return type

[**ActiveDirectoryEntry**](ActiveDirectoryEntry.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePreferredDomainControllers

> ActiveDirectoryEntry UpdatePreferredDomainControllers(ctx, name).Body(body).Execute()





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
    name := "name_example" // string | Specifies the Active Directory Domain Name.
    body := []openapiclient.PreferredDomainController{*openapiclient.NewPreferredDomainController()} // []PreferredDomainController | Request to update preferred domain controllers of an Active Directory.

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

    request := cohesitysdk.ApiUpdatePreferredDomainControllersRequest{
        name: &Name
        body: &Body
    }

    resp, r, err := api_client.ActiveDirectoryApi.UpdatePreferredDomainControllers(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryApi.UpdatePreferredDomainControllers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePreferredDomainControllers`: ActiveDirectoryEntry
    fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryApi.UpdatePreferredDomainControllers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the Active Directory Domain Name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePreferredDomainControllersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**[]PreferredDomainController**](PreferredDomainController.md) | Request to update preferred domain controllers of an Active Directory. | 

### Return type

[**ActiveDirectoryEntry**](ActiveDirectoryEntry.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

