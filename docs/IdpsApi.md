# \IdpsApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddActiveIdpPrincipals**](IdpsApi.md#AddActiveIdpPrincipals) | **Post** /public/idp/principals | Add multiple groups or users on the Cohesity Cluster for the specified Idp principals. In addition, assign Cohesity roles to the users or groups to define their Cohesity privileges.
[**CreateIdp**](IdpsApi.md#CreateIdp) | **Post** /public/idps | Create an IdP configuration.
[**DeleteIdp**](IdpsApi.md#DeleteIdp) | **Delete** /public/idps/{id} | Delete one IdP configuration.
[**GetIdps**](IdpsApi.md#GetIdps) | **Get** /public/idps | List the IdPs configured on the Cluster.
[**IdpLogin**](IdpsApi.md#IdpLogin) | **Get** /public/idps/login | Login to Cohesity Cluster using an IdP.
[**UpdateIdp**](IdpsApi.md#UpdateIdp) | **Put** /public/idps/{id} | Update an IdP configuration.



## AddActiveIdpPrincipals

> []AddedIdpPrincipal AddActiveIdpPrincipals(ctx).Execute()

Add multiple groups or users on the Cohesity Cluster for the specified Idp principals. In addition, assign Cohesity roles to the users or groups to define their Cohesity privileges.



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

    request := cohesitysdk.ApiAddActiveIdpPrincipalsRequest{
    }

    resp, r, err := api_client.IdpsApi.AddActiveIdpPrincipals(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpsApi.AddActiveIdpPrincipals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddActiveIdpPrincipals`: []AddedIdpPrincipal
    fmt.Fprintf(os.Stdout, "Response from `IdpsApi.AddActiveIdpPrincipals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAddActiveIdpPrincipalsRequest struct via the builder pattern


### Return type

[**[]AddedIdpPrincipal**](AddedIdpPrincipal.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIdp

> IdpServiceConfiguration CreateIdp(ctx).Body(body).Execute()

Create an IdP configuration.



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
    body := *openapiclient.NewCreateIdpConfigurationRequest() // CreateIdpConfigurationRequest | Request to create a new IdP Configuration. (optional)

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

    request := cohesitysdk.ApiCreateIdpRequest{
        body: &Body
    }

    resp, r, err := api_client.IdpsApi.CreateIdp(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpsApi.CreateIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdp`: IdpServiceConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IdpsApi.CreateIdp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateIdpConfigurationRequest**](CreateIdpConfigurationRequest.md) | Request to create a new IdP Configuration. | 

### Return type

[**IdpServiceConfiguration**](IdpServiceConfiguration.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdp

> DeleteIdp(ctx, id).Execute()

Delete one IdP configuration.



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
    id := int64(789) // int64 | Specifies the Id assigned for the IdP Service by the Cluster.

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

    request := cohesitysdk.ApiDeleteIdpRequest{
        id: &Id
    }

    resp, r, err := api_client.IdpsApi.DeleteIdp(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpsApi.DeleteIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the Id assigned for the IdP Service by the Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetIdps

> []IdpServiceConfiguration GetIdps(ctx).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).Names(names).Ids(ids).Domains(domains).Execute()

List the IdPs configured on the Cluster.



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
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    allUnderHierarchy := true // bool | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)
    names := []string{"Inner_example"} // []string | Specifies the names of the IdP vendors like Okta. If specified, returns IdP configurations of the vendors matching the names in the parameters. (optional)
    ids := []int64{int64(123)} // []int64 | Specifies the Ids of the IdP configuration. If specified, returns IdP configurations of the matching Ids in the IdP configuration. (optional)
    domains := []string{"Inner_example"} // []string | Specifies the domains of the IdP configurations. If specified, returns IdP configurations matching the domains in the parameters. (optional)

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

    request := cohesitysdk.ApiGetIdpsRequest{
        tenantIds: &TenantIds
        allUnderHierarchy: &AllUnderHierarchy
        names: &Names
        ids: &Ids
        domains: &Domains
    }

    resp, r, err := api_client.IdpsApi.GetIdps(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpsApi.GetIdps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdps`: []IdpServiceConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IdpsApi.GetIdps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 
 **names** | **[]string** | Specifies the names of the IdP vendors like Okta. If specified, returns IdP configurations of the vendors matching the names in the parameters. | 
 **ids** | **[]int64** | Specifies the Ids of the IdP configuration. If specified, returns IdP configurations of the matching Ids in the IdP configuration. | 
 **domains** | **[]string** | Specifies the domains of the IdP configurations. If specified, returns IdP configurations matching the domains in the parameters. | 

### Return type

[**[]IdpServiceConfiguration**](IdpServiceConfiguration.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdpLogin

> RequestError IdpLogin(ctx).TenantId(tenantId).Execute()

Login to Cohesity Cluster using an IdP.



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
    tenantId := "tenantId_example" // string | Specifies an optional tenantId for which the SSO login should be done. If this is not specified, Cluster SSO login is done. (optional)

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

    request := cohesitysdk.ApiIdpLoginRequest{
        tenantId: &TenantId
    }

    resp, r, err := api_client.IdpsApi.IdpLogin(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpsApi.IdpLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdpLogin`: RequestError
    fmt.Fprintf(os.Stdout, "Response from `IdpsApi.IdpLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdpLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | Specifies an optional tenantId for which the SSO login should be done. If this is not specified, Cluster SSO login is done. | 

### Return type

[**RequestError**](RequestError.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdp

> IdpServiceConfiguration UpdateIdp(ctx, id).Body(body).Execute()

Update an IdP configuration.



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
    id := int64(789) // int64 | Specifies the Id assigned for the IdP Service by the Cluster.
    body := *openapiclient.NewUpdateIdpConfigurationRequest() // UpdateIdpConfigurationRequest | Request to update an Idp Configuration. (optional)

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

    request := cohesitysdk.ApiUpdateIdpRequest{
        id: &Id
        body: &Body
    }

    resp, r, err := api_client.IdpsApi.UpdateIdp(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdpsApi.UpdateIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdp`: IdpServiceConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IdpsApi.UpdateIdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the Id assigned for the IdP Service by the Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateIdpConfigurationRequest**](UpdateIdpConfigurationRequest.md) | Request to update an Idp Configuration. | 

### Return type

[**IdpServiceConfiguration**](IdpServiceConfiguration.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

