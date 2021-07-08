# \PrincipalsApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](PrincipalsApi.md#CreateUser) | **Post** /public/users | Create or add a new user to the Cohesity Cluster.
[**CreateUserApiKey**](PrincipalsApi.md#CreateUserApiKey) | **Post** /public/users/{sid}/apiKeys | Create an API key for user.
[**DeleteUserApiKey**](PrincipalsApi.md#DeleteUserApiKey) | **Delete** /public/users/{sid}/apiKeys/{id} | Delete an API key for user.
[**DeleteUsers**](PrincipalsApi.md#DeleteUsers) | **Delete** /public/users | Delete one or more users on the Cohesity Cluster.
[**GetAllUserApiKeys**](PrincipalsApi.md#GetAllUserApiKeys) | **Get** /public/usersApiKeys | Fetch API keys across all users.
[**GetSessionUser**](PrincipalsApi.md#GetSessionUser) | **Get** /public/sessionUser | Get the information of the logged in user.
[**GetUserApiKeyById**](PrincipalsApi.md#GetUserApiKeyById) | **Get** /public/users/{sid}/apiKeys/{id} | Fetch an API key for user by its id.
[**GetUserApiKeys**](PrincipalsApi.md#GetUserApiKeys) | **Get** /public/users/{sid}/apiKeys | Fetch API keys for user.
[**GetUserPrivileges**](PrincipalsApi.md#GetUserPrivileges) | **Get** /public/users/privileges | List the privileges of the session user.
[**GetUsers**](PrincipalsApi.md#GetUsers) | **Get** /public/users | List the users on the Cohesity Cluster that match the filter criteria specified using parameters.
[**ListSourcesForPrincipals**](PrincipalsApi.md#ListSourcesForPrincipals) | **Get** /public/principals/protectionSources | Returns the Protection Sources objects and View names that the principals have permissions to access.
[**ResetS3SecretKey**](PrincipalsApi.md#ResetS3SecretKey) | **Post** /public/users/s3SecretKey | Reset the S3 secret access key for the specified user on the Cohesity Cluster.
[**RotateUserApiKey**](PrincipalsApi.md#RotateUserApiKey) | **Post** /public/users/{sid}/apiKeys/{id}/rotate | Fetch an API key for user by its id.
[**SearchPrincipals**](PrincipalsApi.md#SearchPrincipals) | **Get** /public/principals/searchPrincipals | List the user and group principals that match the filter criteria specified using parameters.
[**UpdateSourcesForPrincipals**](PrincipalsApi.md#UpdateSourcesForPrincipals) | **Put** /public/principals/protectionSources | Set the Protection Sources and Views that the specified principal has permissions to access.
[**UpdateUser**](PrincipalsApi.md#UpdateUser) | **Put** /public/users | Update an existing user on the Cohesity Cluster. Only user settings on the Cohesity Cluster are updated. No changes are made to the referenced user principal on the Active Directory.
[**UpdateUserApiKey**](PrincipalsApi.md#UpdateUserApiKey) | **Put** /public/users/{sid}/apiKeys/{id} | Update an API key.



## CreateUser

> User CreateUser(ctx).Body(body).Execute()

Create or add a new user to the Cohesity Cluster.



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
    body := *openapiclient.NewUserParameters() // UserParameters | Request to add or create a new user. (optional)

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

    request := cohesitysdk.ApiCreateUserRequest{
        Body: &body
    }

    resp, r, err := api_client.PrincipalsApi.CreateUser(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalsApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `PrincipalsApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserParameters**](UserParameters.md) | Request to add or create a new user. | 

### Return type

[**User**](User.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserApiKey

> CreatedApiKey CreateUserApiKey(ctx, sid).Body(body).Execute()

Create an API key for user.

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
    sid := "sid_example" // string | Specifies the user sid.
    body := *openapiclient.NewCreateApiKeyParams("Name_example") // CreateApiKeyParams | Request to create an API key.

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

    request := cohesitysdk.ApiCreateUserApiKeyRequest{
        Sid: &sid
        Body: &body
    }

    resp, r, err := api_client.PrincipalsApi.CreateUserApiKey(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalsApi.CreateUserApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserApiKey`: CreatedApiKey
    fmt.Fprintf(os.Stdout, "Response from `PrincipalsApi.CreateUserApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | Specifies the user sid. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateApiKeyParams**](CreateApiKeyParams.md) | Request to create an API key. | 

### Return type

[**CreatedApiKey**](CreatedApiKey.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserApiKey

> DeleteUserApiKey(ctx, sid, id).Execute()

Delete an API key for user.

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
    sid := "sid_example" // string | Specifies the user sid.
    id := "id_example" // string | Specifies the API key id.

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

    request := cohesitysdk.ApiDeleteUserApiKeyRequest{
        Sid: &sid
        Id: &id
    }

    resp, r, err := api_client.PrincipalsApi.DeleteUserApiKey(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalsApi.DeleteUserApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | Specifies the user sid. | 
**id** | **string** | Specifies the API key id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserApiKeyRequest struct via the builder pattern


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


## DeleteUsers

> DeleteUsers(ctx).Body(body).Execute()

Delete one or more users on the Cohesity Cluster.



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
    body := *openapiclient.NewUserDeleteParameters() // UserDeleteParameters | Request to delete one or more users on the Cohesity Cluster. (optional)

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

    request := cohesitysdk.ApiDeleteUsersRequest{
        Body: &body
    }

    resp, r, err := api_client.PrincipalsApi.DeleteUsers(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalsApi.DeleteUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserDeleteParameters**](UserDeleteParameters.md) | Request to delete one or more users on the Cohesity Cluster. | 

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


## GetAllUserApiKeys

> []ApiKey GetAllUserApiKeys(ctx).UserSids(userSids).Execute()

Fetch API keys across all users.

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
    userSids := []string{"Inner_example"} // []string | Specifies a list of user sids. (optional)

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

    request := cohesitysdk.ApiGetAllUserApiKeysRequest{
        UserSids: &userSids
    }

    resp, r, err := api_client.PrincipalsApi.GetAllUserApiKeys(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalsApi.GetAllUserApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllUserApiKeys`: []ApiKey
    fmt.Fprintf(os.Stdout, "Response from `PrincipalsApi.GetAllUserApiKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllUserApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userSids** | **[]string** | Specifies a list of user sids. | 

### Return type

[**[]ApiKey**](ApiKey.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionUser

> User GetSessionUser(ctx).Execute()

Get the information of the logged in user.

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

    request := cohesitysdk.ApiGetSessionUserRequest{
    }

    resp, r, err := api_client.PrincipalsApi.GetSessionUser(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalsApi.GetSessionUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionUser`: User
    fmt.Fprintf(os.Stdout, "Response from `PrincipalsApi.GetSessionUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionUserRequest struct via the builder pattern


### Return type

[**User**](User.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserApiKeyById

> ApiKey GetUserApiKeyById(ctx, sid, id).Execute()

Fetch an API key for user by its id.

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
    sid := "sid_example" // string | Specifies the user sid.
    id := "id_example" // string | Specifies the API key id.

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

    request := cohesitysdk.ApiGetUserApiKeyByIdRequest{
        Sid: &sid
        Id: &id
    }

    resp, r, err := api_client.PrincipalsApi.GetUserApiKeyById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalsApi.GetUserApiKeyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserApiKeyById`: ApiKey
    fmt.Fprintf(os.Stdout, "Response from `PrincipalsApi.GetUserApiKeyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | Specifies the user sid. | 
**id** | **string** | Specifies the API key id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserApiKeyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiKey**](ApiKey.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserApiKeys

> []ApiKey GetUserApiKeys(ctx, sid).Ids(ids).Execute()

Fetch API keys for user.

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
    sid := "sid_example" // string | Specifies the user sid.
    ids := []string{"Inner_example"} // []string | Specifies a list of API key ids. (optional)

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

    request := cohesitysdk.ApiGetUserApiKeysRequest{
        Sid: &sid
        Ids: &ids
    }

    resp, r, err := api_client.PrincipalsApi.GetUserApiKeys(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalsApi.GetUserApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserApiKeys`: []ApiKey
    fmt.Fprintf(os.Stdout, "Response from `PrincipalsApi.GetUserApiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | Specifies the user sid. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ids** | **[]string** | Specifies a list of API key ids. | 

### Return type

[**[]ApiKey**](ApiKey.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPrivileges

> []string GetUserPrivileges(ctx).Execute()

List the privileges of the session user.

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

    request := cohesitysdk.ApiGetUserPrivilegesRequest{
    }

    resp, r, err := api_client.PrincipalsApi.GetUserPrivileges(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalsApi.GetUserPrivileges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserPrivileges`: []string
    fmt.Fprintf(os.Stdout, "Response from `PrincipalsApi.GetUserPrivileges`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPrivilegesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> []User GetUsers(ctx).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).Usernames(usernames).EmailAddresses(emailAddresses).Domain(domain).PartialMatch(partialMatch).Execute()

List the users on the Cohesity Cluster that match the filter criteria specified using parameters.



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
    usernames := []string{"Inner_example"} // []string | Optionally specify a list of usernames to filter by. All users containing username will be returned. (optional)
    emailAddresses := []string{"Inner_example"} // []string | Optionally specify a list of email addresses to filter by. (optional)
    domain := "domain_example" // string | Optionally specify a domain to filter by. If no domain is specified, all users on the Cohesity Cluster are searched. If a domain is specified, only users on the Cohesity Cluster associated with that domain are searched. (optional)
    partialMatch := true // bool | Optionally specify whether to enable partial match. If set, all users with name containing Usernames will be returned. If set to false, only users with exact the same name as Usernames will be returned. By default this parameter is set to true. (optional)

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

    request := cohesitysdk.ApiGetUsersRequest{
        TenantIds: &tenantIds
        AllUnderHierarchy: &allUnderHierarchy
        Usernames: &usernames
        EmailAddresses: &emailAddresses
        Domain: &domain
        PartialMatch: &partialMatch
    }

    resp, r, err := api_client.PrincipalsApi.GetUsers(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalsApi.GetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsers`: []User
    fmt.Fprintf(os.Stdout, "Response from `PrincipalsApi.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 
 **usernames** | **[]string** | Optionally specify a list of usernames to filter by. All users containing username will be returned. | 
 **emailAddresses** | **[]string** | Optionally specify a list of email addresses to filter by. | 
 **domain** | **string** | Optionally specify a domain to filter by. If no domain is specified, all users on the Cohesity Cluster are searched. If a domain is specified, only users on the Cohesity Cluster associated with that domain are searched. | 
 **partialMatch** | **bool** | Optionally specify whether to enable partial match. If set, all users with name containing Usernames will be returned. If set to false, only users with exact the same name as Usernames will be returned. By default this parameter is set to true. | 

### Return type

[**[]User**](User.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSourcesForPrincipals

> []SourcesForSid ListSourcesForPrincipals(ctx).Sids(sids).Execute()

Returns the Protection Sources objects and View names that the principals have permissions to access.



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
    sids := []string{"Inner_example"} // []string | Specifies a list of security identifiers (SIDs) that specify user or group principals. (optional)

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

    request := cohesitysdk.ApiListSourcesForPrincipalsRequest{
        Sids: &sids
    }

    resp, r, err := api_client.PrincipalsApi.ListSourcesForPrincipals(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalsApi.ListSourcesForPrincipals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSourcesForPrincipals`: []SourcesForSid
    fmt.Fprintf(os.Stdout, "Response from `PrincipalsApi.ListSourcesForPrincipals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSourcesForPrincipalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sids** | **[]string** | Specifies a list of security identifiers (SIDs) that specify user or group principals. | 

### Return type

[**[]SourcesForSid**](SourcesForSid.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetS3SecretKey

> NewS3SecretAccessKey ResetS3SecretKey(ctx).Body(body).Execute()

Reset the S3 secret access key for the specified user on the Cohesity Cluster.



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
    body := *openapiclient.NewResetS3SecretKeyParameters() // ResetS3SecretKeyParameters | Request to reset the S3 secret access key for the specified Cohesity user. (optional)

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

    request := cohesitysdk.ApiResetS3SecretKeyRequest{
        Body: &body
    }

    resp, r, err := api_client.PrincipalsApi.ResetS3SecretKey(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalsApi.ResetS3SecretKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetS3SecretKey`: NewS3SecretAccessKey
    fmt.Fprintf(os.Stdout, "Response from `PrincipalsApi.ResetS3SecretKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetS3SecretKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ResetS3SecretKeyParameters**](ResetS3SecretKeyParameters.md) | Request to reset the S3 secret access key for the specified Cohesity user. | 

### Return type

[**NewS3SecretAccessKey**](NewS3SecretAccessKey.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateUserApiKey

> CreatedApiKey RotateUserApiKey(ctx, sid, id).Execute()

Fetch an API key for user by its id.

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
    sid := "sid_example" // string | Specifies the user sid.
    id := "id_example" // string | Specifies the API key id.

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

    request := cohesitysdk.ApiRotateUserApiKeyRequest{
        Sid: &sid
        Id: &id
    }

    resp, r, err := api_client.PrincipalsApi.RotateUserApiKey(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalsApi.RotateUserApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RotateUserApiKey`: CreatedApiKey
    fmt.Fprintf(os.Stdout, "Response from `PrincipalsApi.RotateUserApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | Specifies the user sid. | 
**id** | **string** | Specifies the API key id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateUserApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreatedApiKey**](CreatedApiKey.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPrincipals

> []Principal SearchPrincipals(ctx).Domain(domain).ObjectClass(objectClass).Search(search).Sids(sids).IncludeComputers(includeComputers).Execute()

List the user and group principals that match the filter criteria specified using parameters.



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

    request := cohesitysdk.ApiSearchPrincipalsRequest{
        Domain: &domain
        ObjectClass: &objectClass
        Search: &search
        Sids: &sids
        IncludeComputers: &includeComputers
    }

    resp, r, err := api_client.PrincipalsApi.SearchPrincipals(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalsApi.SearchPrincipals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchPrincipals`: []Principal
    fmt.Fprintf(os.Stdout, "Response from `PrincipalsApi.SearchPrincipals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPrincipalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | Specifies the domain name of the principals to search. If specified the principals in that domain are searched. Domain could be an Active Directory domain joined by the Cluster or any one of the trusted domains of the Active Directory domain or the LOCAL domain. If not specified, all the domains are searched. | 
 **objectClass** | **string** | Optionally filter by a principal object class such as &#39;kGroup&#39; or &#39;kUser&#39;. If &#39;kGroup&#39; is specified, only group principals are returned. If &#39;kUser&#39; is specified, only user principals are returned. If not specified, both group and user principals are returned. &#39;kUser&#39; specifies a user object class. &#39;kGroup&#39; specifies a group object class. &#39;kComputer&#39; specifies a computer object class. &#39;kWellKnownPrincipal&#39; specifies a well known principal. | 
 **search** | **string** | Optionally filter by matching a substring. Only principals in the with a name or sAMAccountName that matches part or all of the specified substring are returned. If specified, a &#39;sids&#39; parameter should not be specified. | 
 **sids** | **[]string** | Optionally filter by a list of security identifiers (SIDs) found in the specified domain. Only principals matching the specified SIDs are returned. If specified, a &#39;search&#39; parameter should not be specified. | 
 **includeComputers** | **bool** | Specifies if Computer/GMSA accounts need to be included in this search. | 

### Return type

[**[]Principal**](Principal.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSourcesForPrincipals

> UpdateSourcesForPrincipals(ctx).Body(body).Execute()

Set the Protection Sources and Views that the specified principal has permissions to access.



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
    body := *openapiclient.NewUpdateSourcesForPrincipalsParams() // UpdateSourcesForPrincipalsParams | Request to set access permissions to Protection Sources and Views for a principal.

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

    request := cohesitysdk.ApiUpdateSourcesForPrincipalsRequest{
        Body: &body
    }

    resp, r, err := api_client.PrincipalsApi.UpdateSourcesForPrincipals(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalsApi.UpdateSourcesForPrincipals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSourcesForPrincipalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateSourcesForPrincipalsParams**](UpdateSourcesForPrincipalsParams.md) | Request to set access permissions to Protection Sources and Views for a principal. | 

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


## UpdateUser

> User UpdateUser(ctx).Body(body).Execute()

Update an existing user on the Cohesity Cluster. Only user settings on the Cohesity Cluster are updated. No changes are made to the referenced user principal on the Active Directory.



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
    body := *openapiclient.NewUser() // User | Request to update an existing user. (optional)

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

    request := cohesitysdk.ApiUpdateUserRequest{
        Body: &body
    }

    resp, r, err := api_client.PrincipalsApi.UpdateUser(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalsApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `PrincipalsApi.UpdateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**User**](User.md) | Request to update an existing user. | 

### Return type

[**User**](User.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserApiKey

> CreatedApiKey UpdateUserApiKey(ctx, sid, id).Body(body).Execute()

Update an API key.

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
    sid := "sid_example" // string | Specifies the user sid.
    id := "id_example" // string | Specifies the API key id.
    body := *openapiclient.NewUpdateApiKeyParams("Name_example") // UpdateApiKeyParams | Request to update an API key.

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

    request := cohesitysdk.ApiUpdateUserApiKeyRequest{
        Sid: &sid
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.PrincipalsApi.UpdateUserApiKey(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalsApi.UpdateUserApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserApiKey`: CreatedApiKey
    fmt.Fprintf(os.Stdout, "Response from `PrincipalsApi.UpdateUserApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | Specifies the user sid. | 
**id** | **string** | Specifies the API key id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateApiKeyParams**](UpdateApiKeyParams.md) | Request to update an API key. | 

### Return type

[**CreatedApiKey**](CreatedApiKey.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

