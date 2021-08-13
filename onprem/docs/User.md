# \User

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSession**](User.md#CreateSession) | **Post** /users/sessions | Create a user session
[**DeleteSession**](User.md#DeleteSession) | **Delete** /users/sessions | Delete user sessions
[**GetGroups**](User.md#GetGroups) | **Get** /groups | Get Groups.
[**GetSecurityPrincipals**](User.md#GetSecurityPrincipals) | **Get** /security-principals | Get Security Principals.
[**GetUsers**](User.md#GetUsers) | **Get** /users | Get Users.



## CreateSession

> UserSession CreateSession(ctx).Body(body).Execute()

Create a user session



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
    body := *openapiclient.NewCreateUserSessionRequestParams() // CreateUserSessionRequestParams | Specifies the parameters to create a user session

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

    request := onprem.ApiCreateSessionRequest{
        Body: &body
    }

    resp, r, err := api_client.User.CreateSession(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `User.CreateSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSession`: UserSession
    fmt.Fprintf(os.Stdout, "Response from `User.CreateSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateUserSessionRequestParams**](CreateUserSessionRequestParams.md) | Specifies the parameters to create a user session | 

### Return type

[**UserSession**](UserSession.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSession

> DeleteSession(ctx).Sid(sid).Execute()

Delete user sessions



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
    sid := "sid_example" // string | Specifies a user sid. If sid is not given system wide sessions are deleted. (optional)

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

    request := onprem.ApiDeleteSessionRequest{
        Sid: &sid
    }

    resp, r, err := api_client.User.DeleteSession(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `User.DeleteSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sid** | **string** | Specifies a user sid. If sid is not given system wide sessions are deleted. | 

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


## GetGroups

> Groups GetGroups(ctx).Domain(domain).Names(names).Roles(roles).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()

Get Groups.



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
    domain := "domain_example" // string | Specifies the Group domain to filter. (optional)
    names := []string{"Inner_example"} // []string | Specifies a list of names to filter. (optional)
    roles := []string{"Inner_example"} // []string | Specifies a list of roles to filter. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which groups are to be returned. (optional)
    includeTenants := true // bool | IncludeTenants specifies if groups of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)

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

    request := onprem.ApiGetGroupsRequest{
        Domain: &domain
        Names: &names
        Roles: &roles
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
    }

    resp, r, err := api_client.User.GetGroups(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `User.GetGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroups`: Groups
    fmt.Fprintf(os.Stdout, "Response from `User.GetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | Specifies the Group domain to filter. | 
 **names** | **[]string** | Specifies a list of names to filter. | 
 **roles** | **[]string** | Specifies a list of roles to filter. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which groups are to be returned. | 
 **includeTenants** | **bool** | IncludeTenants specifies if groups of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 

### Return type

[**Groups**](Groups.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityPrincipals

> SecurityPrincipals GetSecurityPrincipals(ctx).Sids(sids).Execute()

Get Security Principals.



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
    sids := []string{"Inner_example"} // []string | Specifies a list of SIDs.

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

    request := onprem.ApiGetSecurityPrincipalsRequest{
        Sids: &sids
    }

    resp, r, err := api_client.User.GetSecurityPrincipals(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `User.GetSecurityPrincipals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityPrincipals`: SecurityPrincipals
    fmt.Fprintf(os.Stdout, "Response from `User.GetSecurityPrincipals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityPrincipalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sids** | **[]string** | Specifies a list of SIDs. | 

### Return type

[**SecurityPrincipals**](SecurityPrincipals.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> UsersList GetUsers(ctx).Domain(domain).Usernames(usernames).MatchPartialNames(matchPartialNames).EmailAddresses(emailAddresses).Roles(roles).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()

Get Users.



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
    domain := "domain_example" // string | Specifies the user domain to filter. (optional)
    usernames := []string{"Inner_example"} // []string | Specifies a list of usernames to filter. (optional)
    matchPartialNames := true // bool | If true, the names in usernames are matched by any partial rather than exactly matched. (optional)
    emailAddresses := []string{"Inner_example"} // []string | Specifies a list of email addresses to filter. (optional)
    roles := []string{"Inner_example"} // []string | Specifies a list of roles to filter. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which users are to be returned. (optional)
    includeTenants := true // bool | IncludeTenants specifies if users of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)

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

    request := onprem.ApiGetUsersRequest{
        Domain: &domain
        Usernames: &usernames
        MatchPartialNames: &matchPartialNames
        EmailAddresses: &emailAddresses
        Roles: &roles
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
    }

    resp, r, err := api_client.User.GetUsers(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `User.GetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsers`: UsersList
    fmt.Fprintf(os.Stdout, "Response from `User.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | Specifies the user domain to filter. | 
 **usernames** | **[]string** | Specifies a list of usernames to filter. | 
 **matchPartialNames** | **bool** | If true, the names in usernames are matched by any partial rather than exactly matched. | 
 **emailAddresses** | **[]string** | Specifies a list of email addresses to filter. | 
 **roles** | **[]string** | Specifies a list of roles to filter. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which users are to be returned. | 
 **includeTenants** | **bool** | IncludeTenants specifies if users of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 

### Return type

[**UsersList**](UsersList.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

