# \RemoteClusterApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRemoteCluster**](RemoteClusterApi.md#CreateRemoteCluster) | **Post** /public/remoteClusters | Register a remote Cluster on this local Cluster for replication.
[**DeleteRemoteCluster**](RemoteClusterApi.md#DeleteRemoteCluster) | **Delete** /public/remoteClusters/{id} | Delete a remote Cluster registration connection.
[**GetRemoteClusterById**](RemoteClusterApi.md#GetRemoteClusterById) | **Get** /public/remoteClusters/{id} | List details about a single remote Cluster registered on this local Cluster.
[**GetRemoteClusters**](RemoteClusterApi.md#GetRemoteClusters) | **Get** /public/remoteClusters | List the remote Cohesity Clusters that are registered on this local Cohesity Cluster that match the filter criteria specified using parameters.
[**GetReplicationEncryptionKey**](RemoteClusterApi.md#GetReplicationEncryptionKey) | **Get** /public/replicationEncryptionKey | Get the encryption key on this Cluster.
[**UpdateRemoteCluster**](RemoteClusterApi.md#UpdateRemoteCluster) | **Put** /public/remoteClusters/{id} | Update the connection settings of the registered remote Cluster.



## CreateRemoteCluster

> RemoteCluster CreateRemoteCluster(ctx).Body(body).Execute()

Register a remote Cluster on this local Cluster for replication.



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
    body := *openapiclient.NewRegisterRemoteCluster() // RegisterRemoteCluster | Request to register a remote Cluster.

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

    request := cohesitysdk.ApiCreateRemoteClusterRequest{
        body: &Body
    }

    resp, r, err := api_client.RemoteClusterApi.CreateRemoteCluster(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteClusterApi.CreateRemoteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRemoteCluster`: RemoteCluster
    fmt.Fprintf(os.Stdout, "Response from `RemoteClusterApi.CreateRemoteCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRemoteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RegisterRemoteCluster**](RegisterRemoteCluster.md) | Request to register a remote Cluster. | 

### Return type

[**RemoteCluster**](RemoteCluster.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRemoteCluster

> DeleteRemoteCluster(ctx, id).Execute()

Delete a remote Cluster registration connection.



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
    id := int64(789) // int64 | id of the remote Cluster

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

    request := cohesitysdk.ApiDeleteRemoteClusterRequest{
        id: &Id
    }

    resp, r, err := api_client.RemoteClusterApi.DeleteRemoteCluster(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteClusterApi.DeleteRemoteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of the remote Cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRemoteClusterRequest struct via the builder pattern


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


## GetRemoteClusterById

> []RemoteCluster GetRemoteClusterById(ctx, id).Execute()

List details about a single remote Cluster registered on this local Cluster.



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
    id := int64(789) // int64 | id of the remote Cluster

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

    request := cohesitysdk.ApiGetRemoteClusterByIdRequest{
        id: &Id
    }

    resp, r, err := api_client.RemoteClusterApi.GetRemoteClusterById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteClusterApi.GetRemoteClusterById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemoteClusterById`: []RemoteCluster
    fmt.Fprintf(os.Stdout, "Response from `RemoteClusterApi.GetRemoteClusterById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of the remote Cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteClusterByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]RemoteCluster**](RemoteCluster.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteClusters

> []RemoteCluster GetRemoteClusters(ctx).ClusterIds(clusterIds).ClusterNames(clusterNames).PurposeReplication(purposeReplication).PurposeRemoteAccess(purposeRemoteAccess).Execute()

List the remote Cohesity Clusters that are registered on this local Cohesity Cluster that match the filter criteria specified using parameters.



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
    clusterIds := []int64{int64(123)} // []int64 | Filter by a list of Cluster ids. (optional)
    clusterNames := []string{"Inner_example"} // []string | Filter by a list of Cluster names. (optional)
    purposeReplication := true // bool | Filter for purpose as Replication. (optional)
    purposeRemoteAccess := true // bool | Filter for purpose as Remote Access. (optional)

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

    request := cohesitysdk.ApiGetRemoteClustersRequest{
        clusterIds: &ClusterIds
        clusterNames: &ClusterNames
        purposeReplication: &PurposeReplication
        purposeRemoteAccess: &PurposeRemoteAccess
    }

    resp, r, err := api_client.RemoteClusterApi.GetRemoteClusters(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteClusterApi.GetRemoteClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemoteClusters`: []RemoteCluster
    fmt.Fprintf(os.Stdout, "Response from `RemoteClusterApi.GetRemoteClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterIds** | **[]int64** | Filter by a list of Cluster ids. | 
 **clusterNames** | **[]string** | Filter by a list of Cluster names. | 
 **purposeReplication** | **bool** | Filter for purpose as Replication. | 
 **purposeRemoteAccess** | **bool** | Filter for purpose as Remote Access. | 

### Return type

[**[]RemoteCluster**](RemoteCluster.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReplicationEncryptionKey

> ReplicationEncryptionKeyReponse GetReplicationEncryptionKey(ctx).Execute()

Get the encryption key on this Cluster.



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

    request := cohesitysdk.ApiGetReplicationEncryptionKeyRequest{
    }

    resp, r, err := api_client.RemoteClusterApi.GetReplicationEncryptionKey(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteClusterApi.GetReplicationEncryptionKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReplicationEncryptionKey`: ReplicationEncryptionKeyReponse
    fmt.Fprintf(os.Stdout, "Response from `RemoteClusterApi.GetReplicationEncryptionKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetReplicationEncryptionKeyRequest struct via the builder pattern


### Return type

[**ReplicationEncryptionKeyReponse**](ReplicationEncryptionKeyReponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRemoteCluster

> RemoteCluster UpdateRemoteCluster(ctx, id).Body(body).Execute()

Update the connection settings of the registered remote Cluster.



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
    id := int64(789) // int64 | id of the remote Cluster
    body := *openapiclient.NewRegisterRemoteCluster() // RegisterRemoteCluster | Request to update a remote Cluster.

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

    request := cohesitysdk.ApiUpdateRemoteClusterRequest{
        id: &Id
        body: &Body
    }

    resp, r, err := api_client.RemoteClusterApi.UpdateRemoteCluster(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteClusterApi.UpdateRemoteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRemoteCluster`: RemoteCluster
    fmt.Fprintf(os.Stdout, "Response from `RemoteClusterApi.UpdateRemoteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | id of the remote Cluster | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRemoteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RegisterRemoteCluster**](RegisterRemoteCluster.md) | Request to update a remote Cluster. | 

### Return type

[**RemoteCluster**](RemoteCluster.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

