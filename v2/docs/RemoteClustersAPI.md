# \RemoteClustersAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRemoteCluster**](RemoteClustersAPI.md#DeleteRemoteCluster) | **Delete** /remote-clusters/{clusterId} | Unregister a Remote Cluster.
[**GetRemoteClusterById**](RemoteClustersAPI.md#GetRemoteClusterById) | **Get** /remote-clusters/{clusterId} | Get Remote Cluster config by id.
[**GetRemoteClusters**](RemoteClustersAPI.md#GetRemoteClusters) | **Get** /remote-clusters | Get all registered Remote Clusters.
[**RegisterRemoteCluster**](RemoteClustersAPI.md#RegisterRemoteCluster) | **Post** /remote-clusters | Register a Remote Cluster.
[**UpdateRemoteCluster**](RemoteClustersAPI.md#UpdateRemoteCluster) | **Put** /remote-clusters/{clusterId} | Update a Remote Cluster config.
[**ValidateRemoteCluster**](RemoteClustersAPI.md#ValidateRemoteCluster) | **Post** /remote-clusters/validate | Validate Remote Cluster config.



## DeleteRemoteCluster

> DeleteRemoteCluster(ctx, clusterId).Execute()

Unregister a Remote Cluster.



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
	clusterId := int64(789) // int64 | Specifies the cluster id of the Remote Cluster to unregister.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RemoteClustersAPI.DeleteRemoteCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteClustersAPI.DeleteRemoteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int64** | Specifies the cluster id of the Remote Cluster to unregister. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRemoteClusterRequest struct via the builder pattern


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


## GetRemoteClusterById

> RemoteCluster GetRemoteClusterById(ctx, clusterId).Execute()

Get Remote Cluster config by id.



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
	clusterId := int64(789) // int64 | Specifies the cluster id of Remote Cluster to fetch.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteClustersAPI.GetRemoteClusterById(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteClustersAPI.GetRemoteClusterById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteClusterById`: RemoteCluster
	fmt.Fprintf(os.Stdout, "Response from `RemoteClustersAPI.GetRemoteClusterById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int64** | Specifies the cluster id of Remote Cluster to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteClusterByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemoteCluster**](RemoteCluster.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteClusters

> RemoteClusters GetRemoteClusters(ctx).ClusterIds(clusterIds).ClusterNames(clusterNames).NodeAddresses(nodeAddresses).Purpose(purpose).IncludeEncryptedCredentials(includeEncryptedCredentials).Execute()

Get all registered Remote Clusters.



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
	clusterIds := []int64{int64(123)} // []int64 | Specifies a list of Remote Cluster ids to filter. (optional)
	clusterNames := []string{"Inner_example"} // []string | Specifies a list of Remote Cluster names to filter. (optional)
	nodeAddresses := []string{"Inner_example"} // []string | Specifies a list of Remote Cluster IPs to filter. (optional)
	purpose := []string{"Purpose_example"} // []string | Specifies the purpose for which the remote cluster is being registered. (optional)
	includeEncryptedCredentials := true // bool | If true, the response will include encrypted password. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteClustersAPI.GetRemoteClusters(context.Background()).ClusterIds(clusterIds).ClusterNames(clusterNames).NodeAddresses(nodeAddresses).Purpose(purpose).IncludeEncryptedCredentials(includeEncryptedCredentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteClustersAPI.GetRemoteClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteClusters`: RemoteClusters
	fmt.Fprintf(os.Stdout, "Response from `RemoteClustersAPI.GetRemoteClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterIds** | **[]int64** | Specifies a list of Remote Cluster ids to filter. | 
 **clusterNames** | **[]string** | Specifies a list of Remote Cluster names to filter. | 
 **nodeAddresses** | **[]string** | Specifies a list of Remote Cluster IPs to filter. | 
 **purpose** | **[]string** | Specifies the purpose for which the remote cluster is being registered. | 
 **includeEncryptedCredentials** | **bool** | If true, the response will include encrypted password. | 

### Return type

[**RemoteClusters**](RemoteClusters.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterRemoteCluster

> RemoteCluster RegisterRemoteCluster(ctx).Body(body).Execute()

Register a Remote Cluster.



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
	body := *openapiclient.NewRegisterRemoteClusterParameters([]string{"NodeAddresses_example"}, "Password_example", "Username_example") // RegisterRemoteClusterParameters | Specifies the request to register Remote Cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteClustersAPI.RegisterRemoteCluster(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteClustersAPI.RegisterRemoteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterRemoteCluster`: RemoteCluster
	fmt.Fprintf(os.Stdout, "Response from `RemoteClustersAPI.RegisterRemoteCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterRemoteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RegisterRemoteClusterParameters**](RegisterRemoteClusterParameters.md) | Specifies the request to register Remote Cluster. | 

### Return type

[**RemoteCluster**](RemoteCluster.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRemoteCluster

> RemoteCluster UpdateRemoteCluster(ctx, clusterId).Body(body).Execute()

Update a Remote Cluster config.



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
	clusterId := int64(789) // int64 | Specifies the cluster id of the Remote Cluster to update.
	body := *openapiclient.NewRemoteCluster() // RemoteCluster | Specifies the request to update Remote Cluster config.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteClustersAPI.UpdateRemoteCluster(context.Background(), clusterId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteClustersAPI.UpdateRemoteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRemoteCluster`: RemoteCluster
	fmt.Fprintf(os.Stdout, "Response from `RemoteClustersAPI.UpdateRemoteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int64** | Specifies the cluster id of the Remote Cluster to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRemoteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RemoteCluster**](RemoteCluster.md) | Specifies the request to update Remote Cluster config. | 

### Return type

[**RemoteCluster**](RemoteCluster.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateRemoteCluster

> RemoteClusterParams ValidateRemoteCluster(ctx).Body(body).IncludeMetadata(includeMetadata).Execute()

Validate Remote Cluster config.



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
	body := *openapiclient.NewValidateRemoteClusterConnectionParams([]string{"NodeAddresses_example"}, "Password_example", "Username_example") // ValidateRemoteClusterConnectionParams | Specifies the request to validate Remote Cluster.
	includeMetadata := true // bool | Specifies if Remote Cluster metadata should be included in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteClustersAPI.ValidateRemoteCluster(context.Background()).Body(body).IncludeMetadata(includeMetadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteClustersAPI.ValidateRemoteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateRemoteCluster`: RemoteClusterParams
	fmt.Fprintf(os.Stdout, "Response from `RemoteClustersAPI.ValidateRemoteCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateRemoteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ValidateRemoteClusterConnectionParams**](ValidateRemoteClusterConnectionParams.md) | Specifies the request to validate Remote Cluster. | 
 **includeMetadata** | **bool** | Specifies if Remote Cluster metadata should be included in the response. | 

### Return type

[**RemoteClusterParams**](RemoteClusterParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

