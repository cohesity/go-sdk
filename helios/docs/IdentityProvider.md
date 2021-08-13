# \IdentityProvider

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticateIdp**](IdentityProvider.md#AuthenticateIdp) | **Post** /mcm/idp/authenticate | Create an Identity Provider Configuration.
[**CreateIdp**](IdentityProvider.md#CreateIdp) | **Post** /mcm/idps | Create an Identity Provider Configuration.
[**CreateIdpPrincipal**](IdentityProvider.md#CreateIdpPrincipal) | **Post** /mcm/idps/principals | Create an Identity Provider Configuration.
[**DeleteIdp**](IdentityProvider.md#DeleteIdp) | **Delete** /mcm/idps/{id} | Delete a IDP configuration.
[**DeleteIdpPrincipal**](IdentityProvider.md#DeleteIdpPrincipal) | **Delete** /mcm/idps/principals/{sid} | Delete an IDP Principal.
[**GetIdpById**](IdentityProvider.md#GetIdpById) | **Get** /mcm/idps/{id} | List details about single Identity Provider configuration.
[**GetIdpPrincipalBySid**](IdentityProvider.md#GetIdpPrincipalBySid) | **Get** /mcm/idps/principals/{sid} | Get IDP Principal by SID
[**GetIdps**](IdentityProvider.md#GetIdps) | **Get** /mcm/idps | Get the list of IDP configurations.
[**ListIdpPrincipals**](IdentityProvider.md#ListIdpPrincipals) | **Get** /mcm/idps/principals | List IDP Principals
[**UpdateIdp**](IdentityProvider.md#UpdateIdp) | **Put** /mcm/idps/{id} | Update IDP Configuration.
[**UpdateIdpPrincipal**](IdentityProvider.md#UpdateIdpPrincipal) | **Put** /mcm/idps/principals/{sid} | Update IDP Principal.



## AuthenticateIdp

> Error AuthenticateIdp(ctx).SAMLResponse(sAMLResponse).RegionId(regionId).Execute()

Create an Identity Provider Configuration.



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
    sAMLResponse := "sAMLResponse_example" // string | Specifies the parameters to authenticate an Identity Provider.
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

    request := helios.ApiAuthenticateIdpRequest{
        SAMLResponse: &sAMLResponse
        RegionId: &regionId
    }

    resp, r, err := api_client.IdentityProvider.AuthenticateIdp(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvider.AuthenticateIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticateIdp`: Error
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvider.AuthenticateIdp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sAMLResponse** | **string** | Specifies the parameters to authenticate an Identity Provider. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**Error**](Error.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIdp

> Idp CreateIdp(ctx).Body(body).RegionId(regionId).Execute()

Create an Identity Provider Configuration.



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
    body := *openapiclient.NewCreateOrUpdateIdpRequest("Name_example", "Domain_example", "SsoUrl_example", "IssuerId_example", "Certificate_example") // CreateOrUpdateIdpRequest | Specifies the parameters to create an Identity Provider.
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

    request := helios.ApiCreateIdpRequest{
        Body: &body
        RegionId: &regionId
    }

    resp, r, err := api_client.IdentityProvider.CreateIdp(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvider.CreateIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdp`: Idp
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvider.CreateIdp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateOrUpdateIdpRequest**](CreateOrUpdateIdpRequest.md) | Specifies the parameters to create an Identity Provider. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**Idp**](Idp.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIdpPrincipal

> IdpPrincipal CreateIdpPrincipal(ctx).Body(body).RegionId(regionId).Execute()

Create an Identity Provider Configuration.



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
    body := *openapiclient.NewIdpPrincipal("Name_example", NullableInt64(123), "PrincipalType_example") // IdpPrincipal | Specifies the parameters to create an IDP Principal.
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

    request := helios.ApiCreateIdpPrincipalRequest{
        Body: &body
        RegionId: &regionId
    }

    resp, r, err := api_client.IdentityProvider.CreateIdpPrincipal(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvider.CreateIdpPrincipal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdpPrincipal`: IdpPrincipal
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvider.CreateIdpPrincipal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdpPrincipalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IdpPrincipal**](IdpPrincipal.md) | Specifies the parameters to create an IDP Principal. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**IdpPrincipal**](IdpPrincipal.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdp

> DeleteIdp(ctx, id).RegionId(regionId).Execute()

Delete a IDP configuration.



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
    id := int64(789) // int64 | Specifies a unique id of the IDP.
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

    request := helios.ApiDeleteIdpRequest{
        Id: &id
        RegionId: &regionId
    }

    resp, r, err := api_client.IdentityProvider.DeleteIdp(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvider.DeleteIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the IDP. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdpRequest struct via the builder pattern


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


## DeleteIdpPrincipal

> DeleteIdpPrincipal(ctx, sid).RegionId(regionId).Execute()

Delete an IDP Principal.



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
    sid := "sid_example" // string | Specifies a unique SID of the Principal.
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

    request := helios.ApiDeleteIdpPrincipalRequest{
        Sid: &sid
        RegionId: &regionId
    }

    resp, r, err := api_client.IdentityProvider.DeleteIdpPrincipal(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvider.DeleteIdpPrincipal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | Specifies a unique SID of the Principal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdpPrincipalRequest struct via the builder pattern


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


## GetIdpById

> Idp GetIdpById(ctx, id).RegionId(regionId).Execute()

List details about single Identity Provider configuration.



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
    id := int64(789) // int64 | Specifies a unique id of the IDP.
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

    request := helios.ApiGetIdpByIdRequest{
        Id: &id
        RegionId: &regionId
    }

    resp, r, err := api_client.IdentityProvider.GetIdpById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvider.GetIdpById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdpById`: Idp
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvider.GetIdpById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the IDP. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**Idp**](Idp.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdpPrincipalBySid

> IdpPrincipal GetIdpPrincipalBySid(ctx, sid).RegionId(regionId).Execute()

Get IDP Principal by SID



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
    sid := "sid_example" // string | Specifies the ID of the IDP Principal.
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

    request := helios.ApiGetIdpPrincipalBySidRequest{
        Sid: &sid
        RegionId: &regionId
    }

    resp, r, err := api_client.IdentityProvider.GetIdpPrincipalBySid(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvider.GetIdpPrincipalBySid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdpPrincipalBySid`: IdpPrincipal
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvider.GetIdpPrincipalBySid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | Specifies the ID of the IDP Principal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpPrincipalBySidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**IdpPrincipal**](IdpPrincipal.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdps

> Idps GetIdps(ctx).RegionId(regionId).Ids(ids).Execute()

Get the list of IDP configurations.



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
    ids := []int64{int64(123)} // []int64 | Filter by a list of IDP ids. (optional)

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

    request := helios.ApiGetIdpsRequest{
        RegionId: &regionId
        Ids: &ids
    }

    resp, r, err := api_client.IdentityProvider.GetIdps(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvider.GetIdps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdps`: Idps
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvider.GetIdps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **ids** | **[]int64** | Filter by a list of IDP ids. | 

### Return type

[**Idps**](Idps.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdpPrincipals

> IdpPrincipals ListIdpPrincipals(ctx).RegionId(regionId).Execute()

List IDP Principals



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

    request := helios.ApiListIdpPrincipalsRequest{
        RegionId: &regionId
    }

    resp, r, err := api_client.IdentityProvider.ListIdpPrincipals(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvider.ListIdpPrincipals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdpPrincipals`: IdpPrincipals
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvider.ListIdpPrincipals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdpPrincipalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**IdpPrincipals**](IdpPrincipals.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdp

> Idp UpdateIdp(ctx, id).Body(body).RegionId(regionId).Execute()

Update IDP Configuration.



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
    id := int64(789) // int64 | Specifies the id of the IDP configuration.
    body := *openapiclient.NewCreateOrUpdateIdpRequest("Name_example", "Domain_example", "SsoUrl_example", "IssuerId_example", "Certificate_example") // CreateOrUpdateIdpRequest | Specifies the parameters to update IDP configuration.
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

    request := helios.ApiUpdateIdpRequest{
        Id: &id
        Body: &body
        RegionId: &regionId
    }

    resp, r, err := api_client.IdentityProvider.UpdateIdp(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvider.UpdateIdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdp`: Idp
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvider.UpdateIdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the IDP configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateOrUpdateIdpRequest**](CreateOrUpdateIdpRequest.md) | Specifies the parameters to update IDP configuration. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**Idp**](Idp.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdpPrincipal

> IdpPrincipal UpdateIdpPrincipal(ctx, sid).Body(body).RegionId(regionId).Execute()

Update IDP Principal.



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
    sid := "sid_example" // string | Specifies the SID of the IDP Principal.
    body := *openapiclient.NewIdpPrincipal("Name_example", NullableInt64(123), "PrincipalType_example") // IdpPrincipal | Specifies the parameters to update IDP Principal.
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

    request := helios.ApiUpdateIdpPrincipalRequest{
        Sid: &sid
        Body: &body
        RegionId: &regionId
    }

    resp, r, err := api_client.IdentityProvider.UpdateIdpPrincipal(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvider.UpdateIdpPrincipal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdpPrincipal`: IdpPrincipal
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvider.UpdateIdpPrincipal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** | Specifies the SID of the IDP Principal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdpPrincipalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IdpPrincipal**](IdpPrincipal.md) | Specifies the parameters to update IDP Principal. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**IdpPrincipal**](IdpPrincipal.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

