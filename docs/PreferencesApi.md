# \PreferencesApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserPreferences**](PreferencesApi.md#GetUserPreferences) | **Get** /public/sessionUser/preferences | List the preferences of the session user.
[**PatchUserPreferences**](PreferencesApi.md#PatchUserPreferences) | **Patch** /public/sessionUser/preferences | Update specific preferences of the session user
[**UpdateUserPreferences**](PreferencesApi.md#UpdateUserPreferences) | **Put** /public/sessionUser/preferences | Update the preferences of the session user



## GetUserPreferences

> UserPreferencesResult GetUserPreferences(ctx).Execute()

List the preferences of the session user.

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

    request := cohesitysdk.ApiGetUserPreferencesRequest{
    }

    resp, r, err := api_client.PreferencesApi.GetUserPreferences(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreferencesApi.GetUserPreferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserPreferences`: UserPreferencesResult
    fmt.Fprintf(os.Stdout, "Response from `PreferencesApi.GetUserPreferences`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPreferencesRequest struct via the builder pattern


### Return type

[**UserPreferencesResult**](UserPreferencesResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchUserPreferences

> UserPreferencesResult PatchUserPreferences(ctx).Preferences(preferences).Execute()

Update specific preferences of the session user



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
    preferences := map[string]string{"key": "Inner_example"} // map[string]string | Request to create or update User Preferences. (optional)

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

    request := cohesitysdk.ApiPatchUserPreferencesRequest{
        preferences: &Preferences
    }

    resp, r, err := api_client.PreferencesApi.PatchUserPreferences(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreferencesApi.PatchUserPreferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchUserPreferences`: UserPreferencesResult
    fmt.Fprintf(os.Stdout, "Response from `PreferencesApi.PatchUserPreferences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchUserPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **preferences** | **map[string]string** | Request to create or update User Preferences. | 

### Return type

[**UserPreferencesResult**](UserPreferencesResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPreferences

> UserPreferencesResult UpdateUserPreferences(ctx).Preferences(preferences).Execute()

Update the preferences of the session user



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
    preferences := map[string]string{"key": "Inner_example"} // map[string]string | Request to create or update User Preferences. (optional)

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

    request := cohesitysdk.ApiUpdateUserPreferencesRequest{
        preferences: &Preferences
    }

    resp, r, err := api_client.PreferencesApi.UpdateUserPreferences(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreferencesApi.UpdateUserPreferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserPreferences`: UserPreferencesResult
    fmt.Fprintf(os.Stdout, "Response from `PreferencesApi.UpdateUserPreferences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **preferences** | **map[string]string** | Request to create or update User Preferences. | 

### Return type

[**UserPreferencesResult**](UserPreferencesResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

