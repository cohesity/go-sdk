# \ActiveDirectoryAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddActiveDirectoryPrincipals**](ActiveDirectoryAPI.md#AddActiveDirectoryPrincipals) | **Post** /active-directory-principals | Add multiple groups or users on the Cohesity Cluster for the specified Active Directory principals. In addition, assign Cohesity roles to the users or groups to define their Cohesity privileges.
[**CreateActiveDirectory**](ActiveDirectoryAPI.md#CreateActiveDirectory) | **Post** /active-directories | Create an Active Directory.
[**DeleteActiveDirectory**](ActiveDirectoryAPI.md#DeleteActiveDirectory) | **Delete** /active-directories/{id} | Delete an Active Directory.
[**GetActiveDirectory**](ActiveDirectoryAPI.md#GetActiveDirectory) | **Get** /active-directories | Get the list of Active Directories.
[**GetActiveDirectoryById**](ActiveDirectoryAPI.md#GetActiveDirectoryById) | **Get** /active-directories/{id} | Get an Active Directory by id.
[**GetActiveDirectoryPrincipals**](ActiveDirectoryAPI.md#GetActiveDirectoryPrincipals) | **Get** /active-directory-principals | Get the list of user and group principals from the Active Directory that match the specified filter criteria.
[**GetCentrifyZones**](ActiveDirectoryAPI.md#GetCentrifyZones) | **Get** /centrify-zones | Get Centrify Zones.
[**GetDomainControllers**](ActiveDirectoryAPI.md#GetDomainControllers) | **Get** /domain-controllers | Get Domain Controllers of specified domains.
[**GetTrustedDomains**](ActiveDirectoryAPI.md#GetTrustedDomains) | **Get** /trusted-domains | Get Trusted Domains.
[**TriggerTrustedDomainsDiscovery**](ActiveDirectoryAPI.md#TriggerTrustedDomainsDiscovery) | **Put** /trusted-domains | Rediscover trusted domains.
[**UpdateActiveDirectory**](ActiveDirectoryAPI.md#UpdateActiveDirectory) | **Put** /active-directories/{id} | Update an Active Directory.
[**UpdateTrustedDomains**](ActiveDirectoryAPI.md#UpdateTrustedDomains) | **Post** /trusted-domains | Update trusted domains.



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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := []openapiclient.AddActiveDirectoryPrincipalsParameters{*openapiclient.NewAddActiveDirectoryPrincipalsParameters("DomainName_example", "Name_example", "ObjectClass_example")} // []AddActiveDirectoryPrincipalsParameters | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActiveDirectoryAPI.AddActiveDirectoryPrincipals(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryAPI.AddActiveDirectoryPrincipals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddActiveDirectoryPrincipals`: []AddedActiveDirectoryPrincipal
	fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryAPI.AddActiveDirectoryPrincipals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddActiveDirectoryPrincipalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]AddActiveDirectoryPrincipalsParameters**](AddActiveDirectoryPrincipalsParameters.md) |  | 

### Return type

[**[]AddedActiveDirectoryPrincipal**](AddedActiveDirectoryPrincipal.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateActiveDirectory

> ActiveDirectory CreateActiveDirectory(ctx).Body(body).Execute()

Create an Active Directory.



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
	body := *openapiclient.NewCreateActiveDirectoryRequest([]openapiclient.MachineAccount{*openapiclient.NewMachineAccount("Name_example")}, map[string]interface{}(123), "DomainName_example") // CreateActiveDirectoryRequest | Specifies the parameters to create an Active Directory.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActiveDirectoryAPI.CreateActiveDirectory(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryAPI.CreateActiveDirectory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateActiveDirectory`: ActiveDirectory
	fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryAPI.CreateActiveDirectory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateActiveDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateActiveDirectoryRequest**](CreateActiveDirectoryRequest.md) | Specifies the parameters to create an Active Directory. | 

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

> DeleteActiveDirectory(ctx, id).ActiveDirectoryAdminUsername(activeDirectoryAdminUsername).ActiveDirectoryAdminPassword(activeDirectoryAdminPassword).ForceRemove(forceRemove).Execute()

Delete an Active Directory.



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
	id := int64(789) // int64 | Specifies id of an Active Directory.
	activeDirectoryAdminUsername := "activeDirectoryAdminUsername_example" // string | Specifies the username of the Active Directory Admin.
	activeDirectoryAdminPassword := "activeDirectoryAdminPassword_example" // string | Specifies the password of the Active Directory Admin.
	forceRemove := true // bool | To force delete the Active directory from cluster. This will skip all the checks that prevents cluster from leaving an AD domain. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ActiveDirectoryAPI.DeleteActiveDirectory(context.Background(), id).ActiveDirectoryAdminUsername(activeDirectoryAdminUsername).ActiveDirectoryAdminPassword(activeDirectoryAdminPassword).ForceRemove(forceRemove).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryAPI.DeleteActiveDirectory``: %v\n", err)
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

 **activeDirectoryAdminUsername** | **string** | Specifies the username of the Active Directory Admin. | 
 **activeDirectoryAdminPassword** | **string** | Specifies the password of the Active Directory Admin. | 
 **forceRemove** | **bool** | To force delete the Active directory from cluster. This will skip all the checks that prevents cluster from leaving an AD domain. | 

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

> ActiveDirectories GetActiveDirectory(ctx).DomainNames(domainNames).Ids(ids).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()

Get the list of Active Directories.



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
	domainNames := []string{"Inner_example"} // []string | Filter by a list of Active Directory domain names. (optional)
	ids := []int64{int64(123)} // []int64 | Filter by a list of Active Directory Ids. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which Active Directories are to be returned. (optional)
	includeTenants := true // bool | If true, the response will include Active Directories which were created by all tenants which the current user has permission to see. If false, then only Active Directories created by the current user will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActiveDirectoryAPI.GetActiveDirectory(context.Background()).DomainNames(domainNames).Ids(ids).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryAPI.GetActiveDirectory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveDirectory`: ActiveDirectories
	fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryAPI.GetActiveDirectory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

> ActiveDirectory GetActiveDirectoryById(ctx, id).IncludeCentrifyZones(includeCentrifyZones).IncludeDomainControllers(includeDomainControllers).IncludeSecurityPrincipals(includeSecurityPrincipals).Prefix(prefix).ObjectClass(objectClass).Execute()

Get an Active Directory by id.



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
	id := int64(789) // int64 | Specifies id of an Active Directory.
	includeCentrifyZones := true // bool | Specifies whether to include Centrify Zones of the Active Directory in response. (optional)
	includeDomainControllers := true // bool | Specifies whether to include Domain Controllers of the Active Directory in response. (optional)
	includeSecurityPrincipals := true // bool | Specifies whether to include Security Principals of the Active Directory in response. (optional)
	prefix := "prefix_example" // string | Specifies a prefix, only security principals with name or sAMAccountName having this prefix (ignoring cases) will be returned. This field is appliciable and mandatory if 'includeSecurityPrincipals' is set to true. (optional)
	objectClass := []string{"ObjectClass_example"} // []string | Specifies a list of object classes, only security principals with object class in this list will be returned. This field is appliciable if 'includeSecurityPrincipals' is set to true. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActiveDirectoryAPI.GetActiveDirectoryById(context.Background(), id).IncludeCentrifyZones(includeCentrifyZones).IncludeDomainControllers(includeDomainControllers).IncludeSecurityPrincipals(includeSecurityPrincipals).Prefix(prefix).ObjectClass(objectClass).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryAPI.GetActiveDirectoryById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveDirectoryById`: ActiveDirectory
	fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryAPI.GetActiveDirectoryById`: %v\n", resp)
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


## GetActiveDirectoryPrincipals

> ActiveDirectoryPrincipals GetActiveDirectoryPrincipals(ctx).DomainName(domainName).Sids(sids).SearchTerm(searchTerm).IncludeComputers(includeComputers).IncludeServiceAccounts(includeServiceAccounts).ObjectClass(objectClass).Execute()

Get the list of user and group principals from the Active Directory that match the specified filter criteria.



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
	domainName := "domainName_example" // string | Specifies the domain name of the principals to search. If specified the principals in that domain are searched. Domain could be an Active Directory domain joined by the Cluster or any one of the trusted domains of the Active Directory domain or the LOCAL domain. If not specified, all the domains are searched. (optional)
	sids := []string{"Inner_example"} // []string | Optionally filter by a list of security identifiers (SIDs) found in the specified domain. Only principals matching the specified SIDs are returned. If specified, a 'searchTerm' parameter should not be specified. (optional)
	searchTerm := "searchTerm_example" // string | Optionally filter by matching a substring. Only principals with a name or sAMAccountName that matches part or all of the specified substring are returned. If specified, a 'sids' parameter should not be specified (optional)
	includeComputers := true // bool | Specifies if Computer/GMSA accounts need to be included in this search. (optional)
	includeServiceAccounts := true // bool | Specifies if service accounts should be included in the search result. (optional)
	objectClass := "objectClass_example" // string | Specifies the type of principal, a user or a group. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActiveDirectoryAPI.GetActiveDirectoryPrincipals(context.Background()).DomainName(domainName).Sids(sids).SearchTerm(searchTerm).IncludeComputers(includeComputers).IncludeServiceAccounts(includeServiceAccounts).ObjectClass(objectClass).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryAPI.GetActiveDirectoryPrincipals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveDirectoryPrincipals`: ActiveDirectoryPrincipals
	fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryAPI.GetActiveDirectoryPrincipals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveDirectoryPrincipalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainName** | **string** | Specifies the domain name of the principals to search. If specified the principals in that domain are searched. Domain could be an Active Directory domain joined by the Cluster or any one of the trusted domains of the Active Directory domain or the LOCAL domain. If not specified, all the domains are searched. | 
 **sids** | **[]string** | Optionally filter by a list of security identifiers (SIDs) found in the specified domain. Only principals matching the specified SIDs are returned. If specified, a &#39;searchTerm&#39; parameter should not be specified. | 
 **searchTerm** | **string** | Optionally filter by matching a substring. Only principals with a name or sAMAccountName that matches part or all of the specified substring are returned. If specified, a &#39;sids&#39; parameter should not be specified | 
 **includeComputers** | **bool** | Specifies if Computer/GMSA accounts need to be included in this search. | 
 **includeServiceAccounts** | **bool** | Specifies if service accounts should be included in the search result. | 
 **objectClass** | **string** | Specifies the type of principal, a user or a group. | 

### Return type

[**ActiveDirectoryPrincipals**](ActiveDirectoryPrincipals.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCentrifyZones

> CentrifyZones GetCentrifyZones(ctx).DomainName(domainName).Execute()

Get Centrify Zones.



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
	domainName := "domainName_example" // string | Specifies the FQDN of the domain name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActiveDirectoryAPI.GetCentrifyZones(context.Background()).DomainName(domainName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryAPI.GetCentrifyZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCentrifyZones`: CentrifyZones
	fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryAPI.GetCentrifyZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCentrifyZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainName** | **string** | Specifies the FQDN of the domain name. | 

### Return type

[**CentrifyZones**](CentrifyZones.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainControllers

> DomainControllersResponse GetDomainControllers(ctx).DomainNames(domainNames).ConnectionId(connectionId).Execute()

Get Domain Controllers of specified domains.



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
	domainNames := []string{"Inner_example"} // []string | Specifies a list of domain names.
	connectionId := int64(789) // int64 | Specifies the Id of the connection which the connector belongs to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActiveDirectoryAPI.GetDomainControllers(context.Background()).DomainNames(domainNames).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryAPI.GetDomainControllers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomainControllers`: DomainControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryAPI.GetDomainControllers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainControllersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainNames** | **[]string** | Specifies a list of domain names. | 
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


## GetTrustedDomains

> TrustedDomainParams GetTrustedDomains(ctx).DomainName(domainName).Execute()

Get Trusted Domains.



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
	domainName := "domainName_example" // string | Specifies the FQDN of an Active directory domain.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActiveDirectoryAPI.GetTrustedDomains(context.Background()).DomainName(domainName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryAPI.GetTrustedDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustedDomains`: TrustedDomainParams
	fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryAPI.GetTrustedDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustedDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainName** | **string** | Specifies the FQDN of an Active directory domain. | 

### Return type

[**TrustedDomainParams**](TrustedDomainParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerTrustedDomainsDiscovery

> TriggerTrustedDomainsDiscovery(ctx).DomainName(domainName).Rediscover(rediscover).Execute()

Rediscover trusted domains.



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
	domainName := "domainName_example" // string | Specifies the FQDN of an Active directory domain.
	rediscover := true // bool | Specifies if trusted domains should be rediscovered.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ActiveDirectoryAPI.TriggerTrustedDomainsDiscovery(context.Background()).DomainName(domainName).Rediscover(rediscover).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryAPI.TriggerTrustedDomainsDiscovery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerTrustedDomainsDiscoveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainName** | **string** | Specifies the FQDN of an Active directory domain. | 
 **rediscover** | **bool** | Specifies if trusted domains should be rediscovered. | 

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


## UpdateActiveDirectory

> ActiveDirectory UpdateActiveDirectory(ctx, id).Body(body).Execute()

Update an Active Directory.



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
	id := int64(789) // int64 | Specifies id of an Active Directory.
	body := *openapiclient.NewUpdateActiveDirectoryRequest([]openapiclient.MachineAccount{*openapiclient.NewMachineAccount("Name_example")}) // UpdateActiveDirectoryRequest | Request to update an Active Directory.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActiveDirectoryAPI.UpdateActiveDirectory(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryAPI.UpdateActiveDirectory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateActiveDirectory`: ActiveDirectory
	fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryAPI.UpdateActiveDirectory`: %v\n", resp)
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


## UpdateTrustedDomains

> TrustedDomainParams UpdateTrustedDomains(ctx).DomainName(domainName).Body(body).Execute()

Update trusted domains.



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
	domainName := "domainName_example" // string | Specifies the FQDN of an Active directory domain.
	body := *openapiclient.NewTrustedDomainParams(false) // TrustedDomainParams | Specifies the trusted domains params.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActiveDirectoryAPI.UpdateTrustedDomains(context.Background()).DomainName(domainName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActiveDirectoryAPI.UpdateTrustedDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTrustedDomains`: TrustedDomainParams
	fmt.Fprintf(os.Stdout, "Response from `ActiveDirectoryAPI.UpdateTrustedDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTrustedDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainName** | **string** | Specifies the FQDN of an Active directory domain. | 
 **body** | [**TrustedDomainParams**](TrustedDomainParams.md) | Specifies the trusted domains params. | 

### Return type

[**TrustedDomainParams**](TrustedDomainParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

