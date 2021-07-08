# \ClusterPartitionsApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterPartitionById**](ClusterPartitionsApi.md#GetClusterPartitionById) | **Get** /public/clusterPartitions/{id} | List details about a single Cluster Partition.
[**GetClusterPartitions**](ClusterPartitionsApi.md#GetClusterPartitions) | **Get** /public/clusterPartitions | List Cluster Partitions filtered by the specified parameters.



## GetClusterPartitionById

> ClusterPartition GetClusterPartitionById(ctx, id).Execute()

List details about a single Cluster Partition.



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
    id := int64(789) // int64 | Specifies a unique id of the Cluster Partition to return.

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

    request := cohesitysdk.ApiGetClusterPartitionByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.ClusterPartitionsApi.GetClusterPartitionById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterPartitionsApi.GetClusterPartitionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterPartitionById`: ClusterPartition
    fmt.Fprintf(os.Stdout, "Response from `ClusterPartitionsApi.GetClusterPartitionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the Cluster Partition to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterPartitionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterPartition**](ClusterPartition.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterPartitions

> []ClusterPartition GetClusterPartitions(ctx).Ids(ids).Names(names).Execute()

List Cluster Partitions filtered by the specified parameters.



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
    ids := []int64{int64(123)} // []int64 | Array of Cluster Partition Ids.  Filter by a list of Cluster Partition ids. If empty, the Cluster Partitions are not filtered by id. (optional)
    names := []string{"Inner_example"} // []string | Array of Cluster Partition Names.  Filter by a list of Cluster Partition Names. If empty, the Cluster Partitions are not filtered by names. (optional)

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

    request := cohesitysdk.ApiGetClusterPartitionsRequest{
        Ids: &ids
        Names: &names
    }

    resp, r, err := api_client.ClusterPartitionsApi.GetClusterPartitions(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterPartitionsApi.GetClusterPartitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterPartitions`: []ClusterPartition
    fmt.Fprintf(os.Stdout, "Response from `ClusterPartitionsApi.GetClusterPartitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | Array of Cluster Partition Ids.  Filter by a list of Cluster Partition ids. If empty, the Cluster Partitions are not filtered by id. | 
 **names** | **[]string** | Array of Cluster Partition Names.  Filter by a list of Cluster Partition Names. If empty, the Cluster Partitions are not filtered by names. | 

### Return type

[**[]ClusterPartition**](ClusterPartition.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

