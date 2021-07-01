# \NetworkApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppendHosts**](NetworkApi.md#AppendHosts) | **Post** /public/network/hosts | Add new entries to the /etc/hosts file.
[**CreateBond**](NetworkApi.md#CreateBond) | **Post** /public/network/bonds | Create a new network bond.
[**DeleteBond**](NetworkApi.md#DeleteBond) | **Delete** /public/network/bonds/{name} | Delete a network bond.
[**DeleteHosts**](NetworkApi.md#DeleteHosts) | **Delete** /public/network/hosts | Remove entries from the /etc/hosts file.
[**EditHosts**](NetworkApi.md#EditHosts) | **Put** /public/network/hosts | Edit entries in the Cluster&#39;s /etc/hosts file.
[**ListHosts**](NetworkApi.md#ListHosts) | **Get** /public/network/hosts | Get the current entries in the hosts file on the Cluster.



## AppendHosts

> HostResult AppendHosts(ctx).Body(body).Execute()

Add new entries to the /etc/hosts file.



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
    body := *openapiclient.NewAppendHostsParameters() // AppendHostsParameters | 

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

    request := cohesitysdk.ApiAppendHostsRequest{
        body: &Body
    }

    resp, r, err := api_client.NetworkApi.AppendHosts(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkApi.AppendHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppendHosts`: HostResult
    fmt.Fprintf(os.Stdout, "Response from `NetworkApi.AppendHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppendHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AppendHostsParameters**](AppendHostsParameters.md) |  | 

### Return type

[**HostResult**](HostResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBond

> CreateBondResult CreateBond(ctx).Body(body).Execute()

Create a new network bond.



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
    body := *openapiclient.NewCreateBondParameters("Name_example", []string{"Slaves_example"}) // CreateBondParameters | 

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

    request := cohesitysdk.ApiCreateBondRequest{
        body: &Body
    }

    resp, r, err := api_client.NetworkApi.CreateBond(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkApi.CreateBond``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBond`: CreateBondResult
    fmt.Fprintf(os.Stdout, "Response from `NetworkApi.CreateBond`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBondRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateBondParameters**](CreateBondParameters.md) |  | 

### Return type

[**CreateBondResult**](CreateBondResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBond

> DeleteBond(ctx, name).Execute()

Delete a network bond.



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
    name := "name_example" // string | Specifies the name of the bond being deleted.

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

    request := cohesitysdk.ApiDeleteBondRequest{
        name: &Name
    }

    resp, r, err := api_client.NetworkApi.DeleteBond(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkApi.DeleteBond``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Specifies the name of the bond being deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBondRequest struct via the builder pattern


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


## DeleteHosts

> HostResult DeleteHosts(ctx).Ips(ips).Execute()

Remove entries from the /etc/hosts file.



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
    ips := []string{"Inner_example"} // []string | Specifies a list of the IP addresses of entries to remove from the Cluster's /etc/hosts file.

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

    request := cohesitysdk.ApiDeleteHostsRequest{
        ips: &Ips
    }

    resp, r, err := api_client.NetworkApi.DeleteHosts(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkApi.DeleteHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteHosts`: HostResult
    fmt.Fprintf(os.Stdout, "Response from `NetworkApi.DeleteHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ips** | **[]string** | Specifies a list of the IP addresses of entries to remove from the Cluster&#39;s /etc/hosts file. | 

### Return type

[**HostResult**](HostResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditHosts

> HostResult EditHosts(ctx).Body(body).Execute()

Edit entries in the Cluster's /etc/hosts file.



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
    body := *openapiclient.NewEditHostsParameters() // EditHostsParameters | 

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

    request := cohesitysdk.ApiEditHostsRequest{
        body: &Body
    }

    resp, r, err := api_client.NetworkApi.EditHosts(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkApi.EditHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditHosts`: HostResult
    fmt.Fprintf(os.Stdout, "Response from `NetworkApi.EditHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EditHostsParameters**](EditHostsParameters.md) |  | 

### Return type

[**HostResult**](HostResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHosts

> []HostEntry ListHosts(ctx).Execute()

Get the current entries in the hosts file on the Cluster.



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

    request := cohesitysdk.ApiListHostsRequest{
    }

    resp, r, err := api_client.NetworkApi.ListHosts(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkApi.ListHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHosts`: []HostEntry
    fmt.Fprintf(os.Stdout, "Response from `NetworkApi.ListHosts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListHostsRequest struct via the builder pattern


### Return type

[**[]HostEntry**](HostEntry.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

