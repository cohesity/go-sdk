# \VlanApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVlan**](VlanApi.md#CreateVlan) | **Post** /public/vlans | Creates a VLAN of the Cohesity Cluster.
[**GetVlanById**](VlanApi.md#GetVlanById) | **Get** /public/vlans/{id} | List the details about a single VLAN.
[**GetVlans**](VlanApi.md#GetVlans) | **Get** /public/vlans | List the VLANs for the Cohesity Cluster.
[**RemoveVlan**](VlanApi.md#RemoveVlan) | **Delete** /public/vlans/{id} | Remove the specified VLAN from the Cohesity Cluster.
[**UpdateVlan**](VlanApi.md#UpdateVlan) | **Put** /public/vlans/{id} | Creates or updates a VLAN of the Cohesity Cluster.



## CreateVlan

> Vlan CreateVlan(ctx).Body(body).Execute()

Creates a VLAN of the Cohesity Cluster.



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
    body := *openapiclient.NewVlan() // Vlan | 

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

    request := cohesitysdk.ApiCreateVlanRequest{
        body: &Body
    }

    resp, r, err := api_client.VlanApi.CreateVlan(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.CreateVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVlan`: Vlan
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.CreateVlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Vlan**](Vlan.md) |  | 

### Return type

[**Vlan**](Vlan.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVlanById

> Vlan GetVlanById(ctx, id).Execute()

List the details about a single VLAN.



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
    id := int32(56) // int32 | Specifies the id of the VLAN.

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

    request := cohesitysdk.ApiGetVlanByIdRequest{
        id: &Id
    }

    resp, r, err := api_client.VlanApi.GetVlanById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.GetVlanById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVlanById`: Vlan
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.GetVlanById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Specifies the id of the VLAN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVlanByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Vlan**](Vlan.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVlans

> []Vlan GetVlans(ctx).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).SkipPrimaryAndBondIface(skipPrimaryAndBondIface).Execute()

List the VLANs for the Cohesity Cluster.



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
    skipPrimaryAndBondIface := true // bool | SkipPrimaryAndBondIface is to filter interfaces entries which are primary interface or bond interfaces. (optional)

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

    request := cohesitysdk.ApiGetVlansRequest{
        tenantIds: &TenantIds
        allUnderHierarchy: &AllUnderHierarchy
        skipPrimaryAndBondIface: &SkipPrimaryAndBondIface
    }

    resp, r, err := api_client.VlanApi.GetVlans(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.GetVlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVlans`: []Vlan
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.GetVlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 
 **skipPrimaryAndBondIface** | **bool** | SkipPrimaryAndBondIface is to filter interfaces entries which are primary interface or bond interfaces. | 

### Return type

[**[]Vlan**](Vlan.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveVlan

> RemoveVlan(ctx, id).Execute()

Remove the specified VLAN from the Cohesity Cluster.



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
    id := int32(56) // int32 | Specifies the id of the VLAN.

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

    request := cohesitysdk.ApiRemoveVlanRequest{
        id: &Id
    }

    resp, r, err := api_client.VlanApi.RemoveVlan(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.RemoveVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Specifies the id of the VLAN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveVlanRequest struct via the builder pattern


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


## UpdateVlan

> Vlan UpdateVlan(ctx, id).Body(body).Execute()

Creates or updates a VLAN of the Cohesity Cluster.



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
    id := int32(56) // int32 | Specifies the id of the VLAN.
    body := *openapiclient.NewVlan() // Vlan |  (optional)

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

    request := cohesitysdk.ApiUpdateVlanRequest{
        id: &Id
        body: &Body
    }

    resp, r, err := api_client.VlanApi.UpdateVlan(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.UpdateVlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVlan`: Vlan
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.UpdateVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Specifies the id of the VLAN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Vlan**](Vlan.md) |  | 

### Return type

[**Vlan**](Vlan.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

