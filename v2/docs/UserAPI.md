# \UserAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroup**](UserAPI.md#CreateGroup) | **Post** /groups | Create Groups
[**CreateSession**](UserAPI.md#CreateSession) | **Post** /users/sessions | Create a user session
[**CreateUserAPIKey**](UserAPI.md#CreateUserAPIKey) | **Post** /users/{userSid}/api-keys | Create a new user API key.
[**CreateUsers**](UserAPI.md#CreateUsers) | **Post** /users | Add one or more users to Cohesity Cluster.
[**DeleteGroup**](UserAPI.md#DeleteGroup) | **Delete** /groups/{sid} | Delete Group
[**DeleteGroups**](UserAPI.md#DeleteGroups) | **Post** /groups/delete | Delete Groups
[**DeleteSession**](UserAPI.md#DeleteSession) | **Delete** /users/sessions | Delete user sessions
[**DeleteUser**](UserAPI.md#DeleteUser) | **Delete** /users/{sid} | Delete a Cohesity (LOCAL/IdP/AD) user.
[**DeleteUserAPIKeyById**](UserAPI.md#DeleteUserAPIKeyById) | **Delete** /users/{userSid}/api-keys/{id} | Delete a user API key.
[**DeleteUsers**](UserAPI.md#DeleteUsers) | **Post** /users/delete | Delete one or more Cohesity users.
[**GetActiveSessionsCount**](UserAPI.md#GetActiveSessionsCount) | **Get** /users/sessions | Get sessions count
[**GetAllAPIKeys**](UserAPI.md#GetAllAPIKeys) | **Get** /api-keys | Get the list of all API keys which are created or owned by the user.
[**GetGroupBySID**](UserAPI.md#GetGroupBySID) | **Get** /groups/{sid} | Get Group by SID
[**GetGroups**](UserAPI.md#GetGroups) | **Get** /groups | Get Groups.
[**GetPrincipalSources**](UserAPI.md#GetPrincipalSources) | **Get** /security-principals/{sid}/sources | Fetch sources &amp; views assigned to a user/group.
[**GetSecurityPrincipals**](UserAPI.md#GetSecurityPrincipals) | **Get** /security-principals | Get Security Principals.
[**GetUserAPIKeyById**](UserAPI.md#GetUserAPIKeyById) | **Get** /users/{userSid}/api-keys/{id} | Get the API key by id.
[**GetUserAPIKeys**](UserAPI.md#GetUserAPIKeys) | **Get** /users/{userSid}/api-keys | Get the list of API keys owned by the user.
[**GetUserBySID**](UserAPI.md#GetUserBySID) | **Get** /users/{sid} | Get User by SID.
[**GetUsers**](UserAPI.md#GetUsers) | **Get** /users | Get Users.
[**RegenerateS3Key**](UserAPI.md#RegenerateS3Key) | **Post** /users/{sid}/s3-secret-key | Reset S3 secret access key
[**RotateUserAPIKey**](UserAPI.md#RotateUserAPIKey) | **Post** /users/{userSid}/api-keys/{id}/rotate | Refresh an existing user API key.
[**UpdateGroup**](UserAPI.md#UpdateGroup) | **Put** /groups/{sid} | Update Group
[**UpdateLinuxCredentialsV2**](UserAPI.md#UpdateLinuxCredentialsV2) | **Put** /users/linux-password | Update or validate linux user password.
[**UpdatePrincipalSources**](UserAPI.md#UpdatePrincipalSources) | **Put** /security-principals/{sid}/sources | Update protection sources assigned to a user/group.
[**UpdateUser**](UserAPI.md#UpdateUser) | **Put** /users/{sid} | Update User information.
[**UpdateUserAPIKeyById**](UserAPI.md#UpdateUserAPIKeyById) | **Put** /users/{userSid}/api-keys/{id} | Update a user API key.
[**UpdateUserS3Keys**](UserAPI.md#UpdateUserS3Keys) | **Post** /users/{sid}/update-s3-keys | Update S3 keys for a User



## CreateGroup

> Groups CreateGroup(ctx).Body(body).Execute()

Create Groups



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
	body := []openapiclient.CreateGroupParams{*openapiclient.NewCreateGroupParams("Domain_example", "Name_example")} // []CreateGroupParams | Specifies the new group parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.CreateGroup(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.CreateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroup`: Groups
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]CreateGroupParams**](CreateGroupParams.md) | Specifies the new group parameters. | 

### Return type

[**Groups**](Groups.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewCreateUserSessionRequestParams() // CreateUserSessionRequestParams | Specifies the parameters to create a user session

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.CreateSession(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.CreateSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSession`: UserSession
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.CreateSession`: %v\n", resp)
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


## CreateUserAPIKey

> CreatedUserAPIKey CreateUserAPIKey(ctx, userSid).Body(body).Execute()

Create a new user API key.



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
	userSid := "userSid_example" // string | Specify the SID of the API key owner.
	body := *openapiclient.NewCreateOrUpdateAPIKeyRequest("Name_example") // CreateOrUpdateAPIKeyRequest | Request to create a new API Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.CreateUserAPIKey(context.Background(), userSid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.CreateUserAPIKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserAPIKey`: CreatedUserAPIKey
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.CreateUserAPIKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userSid** | **string** | Specify the SID of the API key owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateOrUpdateAPIKeyRequest**](CreateOrUpdateAPIKeyRequest.md) | Request to create a new API Key | 

### Return type

[**CreatedUserAPIKey**](CreatedUserAPIKey.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUsers

> UsersList CreateUsers(ctx).Body(body).Execute()

Add one or more users to Cohesity Cluster.



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
	body := []openapiclient.CreateUserParameters{*openapiclient.NewCreateUserParameters("Domain_example", "Username_example")} // []CreateUserParameters | If an Active Directory or an IdP domain is specified, a new user is added to the Cohesity Cluster against the specified Active Directory/IdP user principal. If the LOCAL domain is specified, a new user is created directly in the default LOCAL domain on the Cohesity Cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.CreateUsers(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.CreateUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUsers`: UsersList
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.CreateUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]CreateUserParameters**](CreateUserParameters.md) | If an Active Directory or an IdP domain is specified, a new user is added to the Cohesity Cluster against the specified Active Directory/IdP user principal. If the LOCAL domain is specified, a new user is created directly in the default LOCAL domain on the Cohesity Cluster. | 

### Return type

[**UsersList**](UsersList.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> DeleteGroup(ctx, sid).Execute()

Delete Group



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
	sid := "sid_example" // string | Specify the SID of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.DeleteGroup(context.Background(), sid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | Specify the SID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


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


## DeleteGroups

> DeleteGroups(ctx).Body(body).Execute()

Delete Groups



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
	body := *openapiclient.NewDeleteGroupsRequest([]string{"Sids_example"}) // DeleteGroupsRequest | Specifies the delete request payload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.DeleteGroups(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteGroupsRequest**](DeleteGroupsRequest.md) | Specifies the delete request payload. | 

### Return type

 (empty response body)

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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	sid := "sid_example" // string | Specifies a user sid. If sid is not given system wide sessions are deleted. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.DeleteSession(context.Background()).Sid(sid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteSession``: %v\n", err)
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


## DeleteUser

> DeleteUser(ctx, sid).Execute()

Delete a Cohesity (LOCAL/IdP/AD) user.



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
	sid := "sid_example" // string | Specify the SID of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.DeleteUser(context.Background(), sid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | Specify the SID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## DeleteUserAPIKeyById

> DeleteUserAPIKeyById(ctx, userSid, id).Execute()

Delete a user API key.



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
	userSid := "userSid_example" // string | Specify the SID of the API key owner.
	id := "id_example" // string | Specify the id of the API key.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.DeleteUserAPIKeyById(context.Background(), userSid, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteUserAPIKeyById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userSid** | **string** | Specify the SID of the API key owner. | 
**id** | **string** | Specify the id of the API key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserAPIKeyByIdRequest struct via the builder pattern


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


## DeleteUsers

> DeleteUsers(ctx).Body(body).Execute()

Delete one or more Cohesity users.



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
	body := *openapiclient.NewDeleteUsersRequest([]string{"Sids_example"}) // DeleteUsersRequest | If the Cohesity user was created against an Active Directory/IdP user, the referenced principal user on the Active Directory/IdP domain is NOT deleted. Only the user on the Cohesity Cluster is deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.DeleteUsers(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteUsersRequest**](DeleteUsersRequest.md) | If the Cohesity user was created against an Active Directory/IdP user, the referenced principal user on the Active Directory/IdP domain is NOT deleted. Only the user on the Cohesity Cluster is deleted. | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveSessionsCount

> ActiveSessionsCountParams GetActiveSessionsCount(ctx).Sids(sids).Execute()

Get sessions count



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
	sids := []string{"Inner_example"} // []string | Filter sessions based on user sids. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetActiveSessionsCount(context.Background()).Sids(sids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetActiveSessionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveSessionsCount`: ActiveSessionsCountParams
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetActiveSessionsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveSessionsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sids** | **[]string** | Filter sessions based on user sids. | 

### Return type

[**ActiveSessionsCountParams**](ActiveSessionsCountParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAPIKeys

> UserAPIKeys GetAllAPIKeys(ctx).Ids(ids).OwnerSids(ownerSids).IsActive(isActive).IsExpired(isExpired).Execute()

Get the list of all API keys which are created or owned by the user.



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
	ids := []string{"Inner_example"} // []string | Filter by API Key Ids (optional)
	ownerSids := []string{"Inner_example"} // []string | Filter by list of owner (user) SIDs (optional)
	isActive := true // bool | If true, the response will only include API keys which are active. Returns all keys if the query param is not set. (optional)
	isExpired := true // bool | If true, the response will only include API keys which has been expired. Returns all keys if the query param is not set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetAllAPIKeys(context.Background()).Ids(ids).OwnerSids(ownerSids).IsActive(isActive).IsExpired(isExpired).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetAllAPIKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllAPIKeys`: UserAPIKeys
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetAllAPIKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAPIKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | Filter by API Key Ids | 
 **ownerSids** | **[]string** | Filter by list of owner (user) SIDs | 
 **isActive** | **bool** | If true, the response will only include API keys which are active. Returns all keys if the query param is not set. | 
 **isExpired** | **bool** | If true, the response will only include API keys which has been expired. Returns all keys if the query param is not set. | 

### Return type

[**UserAPIKeys**](UserAPIKeys.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupBySID

> GroupParams GetGroupBySID(ctx, sid).Execute()

Get Group by SID



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
	sid := "sid_example" // string | Specify the SID of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetGroupBySID(context.Background(), sid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetGroupBySID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupBySID`: GroupParams
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetGroupBySID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | Specify the SID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupBySIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupParams**](GroupParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroups

> Groups GetGroups(ctx).Names(names).Roles(roles).Domain(domain).Sids(sids).MatchPartialNames(matchPartialNames).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()

Get Groups.



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
	names := []string{"Inner_example"} // []string | Specifies a list of group names to filter. (optional)
	roles := []string{"Inner_example"} // []string | Specifies a list of roles to filter. (optional)
	domain := "domain_example" // string | Specifies the group domain to filter. (optional)
	sids := []string{"Inner_example"} // []string | Specifies a list of sids to filter. (optional)
	matchPartialNames := true // bool | If true, the names in group names are matched by any partial rather than exactly matched. Default is false. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which groups are to be returned. (optional)
	includeTenants := true // bool | IncludeTenants specifies if groups of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetGroups(context.Background()).Names(names).Roles(roles).Domain(domain).Sids(sids).MatchPartialNames(matchPartialNames).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroups`: Groups
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | **[]string** | Specifies a list of group names to filter. | 
 **roles** | **[]string** | Specifies a list of roles to filter. | 
 **domain** | **string** | Specifies the group domain to filter. | 
 **sids** | **[]string** | Specifies a list of sids to filter. | 
 **matchPartialNames** | **bool** | If true, the names in group names are matched by any partial rather than exactly matched. Default is false. | 
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


## GetPrincipalSources

> AssignedSources GetPrincipalSources(ctx, sid).Execute()

Fetch sources & views assigned to a user/group.



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
	sid := "sid_example" // string | Specify the SID of the principal.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetPrincipalSources(context.Background(), sid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetPrincipalSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrincipalSources`: AssignedSources
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetPrincipalSources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | Specify the SID of the principal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrincipalSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssignedSources**](AssignedSources.md)

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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	sids := []string{"Inner_example"} // []string | Specifies a list of SIDs.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetSecurityPrincipals(context.Background()).Sids(sids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetSecurityPrincipals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityPrincipals`: SecurityPrincipals
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetSecurityPrincipals`: %v\n", resp)
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


## GetUserAPIKeyById

> UserAPIKey GetUserAPIKeyById(ctx, id, userSid).Execute()

Get the API key by id.



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
	id := "id_example" // string | Specify the id of the API key.
	userSid := "userSid_example" // string | Specify the SID of the API key owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUserAPIKeyById(context.Background(), id, userSid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserAPIKeyById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAPIKeyById`: UserAPIKey
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserAPIKeyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specify the id of the API key. | 
**userSid** | **string** | Specify the SID of the API key owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAPIKeyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserAPIKey**](UserAPIKey.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAPIKeys

> UserAPIKeys GetUserAPIKeys(ctx, userSid).Ids(ids).IsActive(isActive).IsExpired(isExpired).Execute()

Get the list of API keys owned by the user.



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
	userSid := "userSid_example" // string | Specify the SID of the API key owner.
	ids := []string{"Inner_example"} // []string | Filter by API Key Ids (optional)
	isActive := true // bool | If true, the response will only include API keys which are active. Returns all keys if the query param is not set. (optional)
	isExpired := true // bool | If true, the response will only include API keys which has been expired. Returns all keys if the query param is not set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUserAPIKeys(context.Background(), userSid).Ids(ids).IsActive(isActive).IsExpired(isExpired).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserAPIKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAPIKeys`: UserAPIKeys
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserAPIKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userSid** | **string** | Specify the SID of the API key owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAPIKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ids** | **[]string** | Filter by API Key Ids | 
 **isActive** | **bool** | If true, the response will only include API keys which are active. Returns all keys if the query param is not set. | 
 **isExpired** | **bool** | If true, the response will only include API keys which has been expired. Returns all keys if the query param is not set. | 

### Return type

[**UserAPIKeys**](UserAPIKeys.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBySID

> UserParams GetUserBySID(ctx, sid).Execute()

Get User by SID.



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
	sid := "sid_example" // string | Specify the SID of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUserBySID(context.Background(), sid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserBySID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserBySID`: UserParams
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserBySID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | Specify the SID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBySIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserParams**](UserParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> UsersList GetUsers(ctx).Domain(domain).Sids(sids).Usernames(usernames).MatchPartialNames(matchPartialNames).EmailAddresses(emailAddresses).Roles(roles).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()

Get Users.



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
	domain := "domain_example" // string | Specifies the user domain to filter. (optional)
	sids := []string{"Inner_example"} // []string | Specifies a list of sids to filter. (optional)
	usernames := []string{"Inner_example"} // []string | Specifies a list of usernames to filter. (optional)
	matchPartialNames := true // bool | If true, the names in usernames are matched by any partial rather than exactly matched. (optional)
	emailAddresses := []string{"Inner_example"} // []string | Specifies a list of email addresses to filter. (optional)
	roles := []string{"Inner_example"} // []string | Specifies a list of roles to filter. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which users are to be returned. (optional)
	includeTenants := true // bool | IncludeTenants specifies if users of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUsers(context.Background()).Domain(domain).Sids(sids).Usernames(usernames).MatchPartialNames(matchPartialNames).EmailAddresses(emailAddresses).Roles(roles).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsers`: UsersList
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | Specifies the user domain to filter. | 
 **sids** | **[]string** | Specifies a list of sids to filter. | 
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


## RegenerateS3Key

> SecretKeyEntity RegenerateS3Key(ctx, sid).Execute()

Reset S3 secret access key



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
	sid := "sid_example" // string | Specify the SID of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.RegenerateS3Key(context.Background(), sid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RegenerateS3Key``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegenerateS3Key`: SecretKeyEntity
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.RegenerateS3Key`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | Specify the SID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateS3KeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecretKeyEntity**](SecretKeyEntity.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateUserAPIKey

> CreatedUserAPIKey RotateUserAPIKey(ctx, userSid, id).Execute()

Refresh an existing user API key.



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
	userSid := "userSid_example" // string | Specify the SID of the API key owner.
	id := "id_example" // string | Specify the id of the API key.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.RotateUserAPIKey(context.Background(), userSid, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RotateUserAPIKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateUserAPIKey`: CreatedUserAPIKey
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.RotateUserAPIKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userSid** | **string** | Specify the SID of the API key owner. | 
**id** | **string** | Specify the id of the API key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateUserAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreatedUserAPIKey**](CreatedUserAPIKey.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> GroupParams UpdateGroup(ctx, sid).Body(body).Execute()

Update Group



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
	sid := "sid_example" // string | Specify the SID of the group.
	body := *openapiclient.NewUpdateGroupParameters() // UpdateGroupParameters | Specifies the group information.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UpdateGroup(context.Background(), sid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroup`: GroupParams
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UpdateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | Specify the SID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateGroupParameters**](UpdateGroupParameters.md) | Specifies the group information. | 

### Return type

[**GroupParams**](GroupParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLinuxCredentialsV2

> SuccessResp UpdateLinuxCredentialsV2(ctx).Body(body).Execute()

Update or validate linux user password.



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
	body := *openapiclient.NewUpdateLinuxPasswordRequest("Username_example") // UpdateLinuxPasswordRequest | Specifies the linux user parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UpdateLinuxCredentialsV2(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateLinuxCredentialsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLinuxCredentialsV2`: SuccessResp
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UpdateLinuxCredentialsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLinuxCredentialsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateLinuxPasswordRequest**](UpdateLinuxPasswordRequest.md) | Specifies the linux user parameters. | 

### Return type

[**SuccessResp**](SuccessResp.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrincipalSources

> AssignedSources UpdatePrincipalSources(ctx, sid).Body(body).Execute()

Update protection sources assigned to a user/group.



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
	sid := "sid_example" // string | Specify the SID of the principal.
	body := *openapiclient.NewAssignedSources() // AssignedSources | Specify the sources to be assigned to a principal.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UpdatePrincipalSources(context.Background(), sid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdatePrincipalSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePrincipalSources`: AssignedSources
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UpdatePrincipalSources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | Specify the SID of the principal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrincipalSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AssignedSources**](AssignedSources.md) | Specify the sources to be assigned to a principal. | 

### Return type

[**AssignedSources**](AssignedSources.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UserParams UpdateUser(ctx, sid).Body(body).Execute()

Update User information.



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
	sid := "sid_example" // string | Specify the SID of the user.
	body := *openapiclient.NewUpdateUserParameters() // UpdateUserParameters | Specifies the user information.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UpdateUser(context.Background(), sid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUser`: UserParams
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | Specify the SID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateUserParameters**](UpdateUserParameters.md) | Specifies the user information. | 

### Return type

[**UserParams**](UserParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserAPIKeyById

> UserAPIKey UpdateUserAPIKeyById(ctx, id, userSid).Body(body).Execute()

Update a user API key.



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
	id := "id_example" // string | Specify the id of the API key.
	userSid := "userSid_example" // string | Specify the SID of the API key owner.
	body := *openapiclient.NewCreateOrUpdateAPIKeyRequest("Name_example") // CreateOrUpdateAPIKeyRequest | Request to update a user API key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UpdateUserAPIKeyById(context.Background(), id, userSid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUserAPIKeyById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserAPIKeyById`: UserAPIKey
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UpdateUserAPIKeyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specify the id of the API key. | 
**userSid** | **string** | Specify the SID of the API key owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserAPIKeyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**CreateOrUpdateAPIKeyRequest**](CreateOrUpdateAPIKeyRequest.md) | Request to update a user API key | 

### Return type

[**UserAPIKey**](UserAPIKey.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserS3Keys

> UpdateUserS3Keys(ctx, sid).Body(body).Execute()

Update S3 keys for a User



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
	sid := "sid_example" // string | Specify the SID of the user.
	body := *openapiclient.NewS3Keys() // S3Keys | Specifies the body to update the User S3 Keys

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.UpdateUserS3Keys(context.Background(), sid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUserS3Keys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | Specify the SID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserS3KeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**S3Keys**](S3Keys.md) | Specifies the body to update the User S3 Keys | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

