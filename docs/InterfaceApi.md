# \InterfaceApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListInterface**](InterfaceApi.md#ListInterface) | **Get** /public/interface | Show network interfaces.
[**UpdateInterface**](InterfaceApi.md#UpdateInterface) | **Put** /public/interface | Update an interface.



## ListInterface

> []NodeNetworkInterfaces ListInterface(ctx).NodeId(nodeId).Cache(cache).BondInterfaceOnly(bondInterfaceOnly).IfaceGroupAssignedOnly(ifaceGroupAssignedOnly).Execute()

Show network interfaces.

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
    nodeId := int64(789) // int64 | Specifies the id of the ndde. (optional)
    cache := true // bool | Specifies if interface is cached info. (optional)
    bondInterfaceOnly := true // bool | Specifies if only show bond interface info. (optional)
    ifaceGroupAssignedOnly := true // bool | Specifies if only show interface group assigned interface info. (optional)

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

    request := cohesitysdk.ApiListInterfaceRequest{
        nodeId: &NodeId
        cache: &Cache
        bondInterfaceOnly: &BondInterfaceOnly
        ifaceGroupAssignedOnly: &IfaceGroupAssignedOnly
    }

    resp, r, err := api_client.InterfaceApi.ListInterface(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfaceApi.ListInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInterface`: []NodeNetworkInterfaces
    fmt.Fprintf(os.Stdout, "Response from `InterfaceApi.ListInterface`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeId** | **int64** | Specifies the id of the ndde. | 
 **cache** | **bool** | Specifies if interface is cached info. | 
 **bondInterfaceOnly** | **bool** | Specifies if only show bond interface info. | 
 **ifaceGroupAssignedOnly** | **bool** | Specifies if only show interface group assigned interface info. | 

### Return type

[**[]NodeNetworkInterfaces**](NodeNetworkInterfaces.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInterface

> Interface UpdateInterface(ctx).Body(body).Execute()

Update an interface.



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
    body := *openapiclient.NewInterface() // Interface |  (optional)

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

    request := cohesitysdk.ApiUpdateInterfaceRequest{
        body: &Body
    }

    resp, r, err := api_client.InterfaceApi.UpdateInterface(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfaceApi.UpdateInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInterface`: Interface
    fmt.Fprintf(os.Stdout, "Response from `InterfaceApi.UpdateInterface`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Interface**](Interface.md) |  | 

### Return type

[**Interface**](Interface.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

