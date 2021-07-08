# \ClustersApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeServiceState**](ClustersApi.md#ChangeServiceState) | **Post** /public/clusters/services/states | Change the state of one or more services on a Cohesity Cluster.
[**CreateCloudCluster**](ClustersApi.md#CreateCloudCluster) | **Post** /public/clusters/cloudEdition | Create a new Cloud Edition Cohesity Cluster.
[**CreatePhysicalCluster**](ClustersApi.md#CreatePhysicalCluster) | **Post** /public/clusters/physicalEdition | Create a new Physical Edition Cohesity Cluster.
[**CreateVirtualCluster**](ClustersApi.md#CreateVirtualCluster) | **Post** /public/clusters/virtualEdition | Create a new Virtual Edition Cohesity Cluster.
[**DestroyCluster**](ClustersApi.md#DestroyCluster) | **Delete** /public/clusters | Destroy a Cohesity Cluster.
[**ExpandCloudCluster**](ClustersApi.md#ExpandCloudCluster) | **Post** /public/clusters/cloudEdition/nodes | Expand a Cloud Edition Cohesity Cluster.
[**ExpandPhysicalCluster**](ClustersApi.md#ExpandPhysicalCluster) | **Post** /public/clusters/physicalEdition/nodes | Expand a Physical Edition Cohesity Cluster.
[**GetBackgroundActivitySchedule**](ClustersApi.md#GetBackgroundActivitySchedule) | **Get** /public/cluster/backgroundActivitySchedule | Get the Apollo Throttling schedule.
[**GetClusterCreationProgress**](ClustersApi.md#GetClusterCreationProgress) | **Get** /public/clusters/creationProgress | Check the progress of the creation of a new Cohesity Cluster.
[**GetClusterKeys**](ClustersApi.md#GetClusterKeys) | **Get** /public/cluster/keys | List the Public Keys for the cluster.
[**GetExternalClientSubnets**](ClustersApi.md#GetExternalClientSubnets) | **Get** /public/externalClientSubnets | List the external Client Subnets for the cluster.
[**GetIoPreferentialTier**](ClustersApi.md#GetIoPreferentialTier) | **Get** /public/clusters/ioPreferentialTier | Return the IO preferential tiers of the cluster.
[**ListServiceStates**](ClustersApi.md#ListServiceStates) | **Get** /public/clusters/services/states | List the states of the services on the Cluster.
[**PutIoPreferentialTier**](ClustersApi.md#PutIoPreferentialTier) | **Put** /public/clusters/ioPreferentialTier | Update the IO preferential tiers and return the updated IO preferential tiers of the cluster.
[**RemoveNode**](ClustersApi.md#RemoveNode) | **Delete** /public/clusters/nodes/{id} | Remove a Node from a Cohesity Cluster.
[**UpdateBackgroundActivitySchedule**](ClustersApi.md#UpdateBackgroundActivitySchedule) | **Put** /public/cluster/backgroundActivitySchedule | Update the Apollo Throttling schedule.
[**UpdateExternalClientSubnets**](ClustersApi.md#UpdateExternalClientSubnets) | **Put** /public/externalClientSubnets | Update the external Client Subnet of the Cluster.
[**UpgradeCluster**](ClustersApi.md#UpgradeCluster) | **Put** /public/clusters/software | Upgrade a Cohesity Cluster.



## ChangeServiceState

> ChangeServiceStateResult ChangeServiceState(ctx).Body(body).Execute()

Change the state of one or more services on a Cohesity Cluster.



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
    body := *openapiclient.NewChangeServiceStateParameters() // ChangeServiceStateParameters | 

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

    request := cohesitysdk.ApiChangeServiceStateRequest{
        Body: &body
    }

    resp, r, err := api_client.ClustersApi.ChangeServiceState(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ChangeServiceState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ChangeServiceState`: ChangeServiceStateResult
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ChangeServiceState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeServiceStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ChangeServiceStateParameters**](ChangeServiceStateParameters.md) |  | 

### Return type

[**ChangeServiceStateResult**](ChangeServiceStateResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCloudCluster

> CreateClusterResult CreateCloudCluster(ctx).Body(body).Execute()

Create a new Cloud Edition Cohesity Cluster.



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
    body := *openapiclient.NewCreateCloudClusterParameters("ClusterName_example", *openapiclient.NewCloudNetworkConfiguration(), []string{"NodeIps_example"}) // CreateCloudClusterParameters | 

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

    request := cohesitysdk.ApiCreateCloudClusterRequest{
        Body: &body
    }

    resp, r, err := api_client.ClustersApi.CreateCloudCluster(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.CreateCloudCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCloudCluster`: CreateClusterResult
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.CreateCloudCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloudClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateCloudClusterParameters**](CreateCloudClusterParameters.md) |  | 

### Return type

[**CreateClusterResult**](CreateClusterResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePhysicalCluster

> CreateClusterResult CreatePhysicalCluster(ctx).Body(body).Execute()

Create a new Physical Edition Cohesity Cluster.



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
    body := *openapiclient.NewCreatePhysicalClusterParameters("ClusterName_example", *openapiclient.NewIpmiConfiguration(), *openapiclient.NewNetworkConfiguration(), []openapiclient.PhysicalNodeConfiguration{*openapiclient.NewPhysicalNodeConfiguration()}) // CreatePhysicalClusterParameters | 

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

    request := cohesitysdk.ApiCreatePhysicalClusterRequest{
        Body: &body
    }

    resp, r, err := api_client.ClustersApi.CreatePhysicalCluster(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.CreatePhysicalCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePhysicalCluster`: CreateClusterResult
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.CreatePhysicalCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePhysicalClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreatePhysicalClusterParameters**](CreatePhysicalClusterParameters.md) |  | 

### Return type

[**CreateClusterResult**](CreateClusterResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVirtualCluster

> CreateClusterResult CreateVirtualCluster(ctx).Body(body).Execute()

Create a new Virtual Edition Cohesity Cluster.



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
    body := *openapiclient.NewCreateVirtualClusterParameters("ClusterName_example", *openapiclient.NewNetworkConfiguration(), []openapiclient.VirtualNodeConfiguration{*openapiclient.NewVirtualNodeConfiguration()}) // CreateVirtualClusterParameters | 

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

    request := cohesitysdk.ApiCreateVirtualClusterRequest{
        Body: &body
    }

    resp, r, err := api_client.ClustersApi.CreateVirtualCluster(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.CreateVirtualCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVirtualCluster`: CreateClusterResult
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.CreateVirtualCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVirtualClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateVirtualClusterParameters**](CreateVirtualClusterParameters.md) |  | 

### Return type

[**CreateClusterResult**](CreateClusterResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyCluster

> DestroyCluster(ctx).Execute()

Destroy a Cohesity Cluster.



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

    request := cohesitysdk.ApiDestroyClusterRequest{
    }

    resp, r, err := api_client.ClustersApi.DestroyCluster(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.DestroyCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyClusterRequest struct via the builder pattern


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


## ExpandCloudCluster

> CreateClusterResult ExpandCloudCluster(ctx).Body(body).Execute()

Expand a Cloud Edition Cohesity Cluster.



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
    body := *openapiclient.NewExpandCloudClusterParameters([]string{"NodeIps_example"}) // ExpandCloudClusterParameters | 

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

    request := cohesitysdk.ApiExpandCloudClusterRequest{
        Body: &body
    }

    resp, r, err := api_client.ClustersApi.ExpandCloudCluster(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ExpandCloudCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpandCloudCluster`: CreateClusterResult
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ExpandCloudCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExpandCloudClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ExpandCloudClusterParameters**](ExpandCloudClusterParameters.md) |  | 

### Return type

[**CreateClusterResult**](CreateClusterResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpandPhysicalCluster

> CreateClusterResult ExpandPhysicalCluster(ctx).Body(body).Execute()

Expand a Physical Edition Cohesity Cluster.



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
    body := *openapiclient.NewExpandPhysicalClusterParameters([]openapiclient.PhysicalNodeConfiguration{*openapiclient.NewPhysicalNodeConfiguration()}) // ExpandPhysicalClusterParameters | 

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

    request := cohesitysdk.ApiExpandPhysicalClusterRequest{
        Body: &body
    }

    resp, r, err := api_client.ClustersApi.ExpandPhysicalCluster(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ExpandPhysicalCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpandPhysicalCluster`: CreateClusterResult
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ExpandPhysicalCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExpandPhysicalClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ExpandPhysicalClusterParameters**](ExpandPhysicalClusterParameters.md) |  | 

### Return type

[**CreateClusterResult**](CreateClusterResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackgroundActivitySchedule

> BandwidthLimit GetBackgroundActivitySchedule(ctx).Execute()

Get the Apollo Throttling schedule.



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

    request := cohesitysdk.ApiGetBackgroundActivityScheduleRequest{
    }

    resp, r, err := api_client.ClustersApi.GetBackgroundActivitySchedule(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetBackgroundActivitySchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBackgroundActivitySchedule`: BandwidthLimit
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetBackgroundActivitySchedule`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackgroundActivityScheduleRequest struct via the builder pattern


### Return type

[**BandwidthLimit**](BandwidthLimit.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterCreationProgress

> ClusterCreationProgressResult GetClusterCreationProgress(ctx).Execute()

Check the progress of the creation of a new Cohesity Cluster.



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

    request := cohesitysdk.ApiGetClusterCreationProgressRequest{
    }

    resp, r, err := api_client.ClustersApi.GetClusterCreationProgress(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusterCreationProgress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterCreationProgress`: ClusterCreationProgressResult
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusterCreationProgress`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterCreationProgressRequest struct via the builder pattern


### Return type

[**ClusterCreationProgressResult**](ClusterCreationProgressResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterKeys

> ClusterPublicKeys GetClusterKeys(ctx).Execute()

List the Public Keys for the cluster.



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

    request := cohesitysdk.ApiGetClusterKeysRequest{
    }

    resp, r, err := api_client.ClustersApi.GetClusterKeys(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetClusterKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterKeys`: ClusterPublicKeys
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetClusterKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterKeysRequest struct via the builder pattern


### Return type

[**ClusterPublicKeys**](ClusterPublicKeys.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalClientSubnets

> ExternalClientSubnets GetExternalClientSubnets(ctx).Execute()

List the external Client Subnets for the cluster.



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

    request := cohesitysdk.ApiGetExternalClientSubnetsRequest{
    }

    resp, r, err := api_client.ClustersApi.GetExternalClientSubnets(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetExternalClientSubnets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExternalClientSubnets`: ExternalClientSubnets
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetExternalClientSubnets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalClientSubnetsRequest struct via the builder pattern


### Return type

[**ExternalClientSubnets**](ExternalClientSubnets.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIoPreferentialTier

> IoPreferentialTier GetIoPreferentialTier(ctx).Execute()

Return the IO preferential tiers of the cluster.



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

    request := cohesitysdk.ApiGetIoPreferentialTierRequest{
    }

    resp, r, err := api_client.ClustersApi.GetIoPreferentialTier(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetIoPreferentialTier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIoPreferentialTier`: IoPreferentialTier
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetIoPreferentialTier`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIoPreferentialTierRequest struct via the builder pattern


### Return type

[**IoPreferentialTier**](IoPreferentialTier.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceStates

> []ServiceStateResult ListServiceStates(ctx).Execute()

List the states of the services on the Cluster.



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

    request := cohesitysdk.ApiListServiceStatesRequest{
    }

    resp, r, err := api_client.ClustersApi.ListServiceStates(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ListServiceStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceStates`: []ServiceStateResult
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ListServiceStates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceStatesRequest struct via the builder pattern


### Return type

[**[]ServiceStateResult**](ServiceStateResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIoPreferentialTier

> IoPreferentialTier PutIoPreferentialTier(ctx).Body(body).Execute()

Update the IO preferential tiers and return the updated IO preferential tiers of the cluster.



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
    body := *openapiclient.NewIoPreferentialTier() // IoPreferentialTier | 

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

    request := cohesitysdk.ApiPutIoPreferentialTierRequest{
        Body: &body
    }

    resp, r, err := api_client.ClustersApi.PutIoPreferentialTier(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.PutIoPreferentialTier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutIoPreferentialTier`: IoPreferentialTier
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.PutIoPreferentialTier`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutIoPreferentialTierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IoPreferentialTier**](IoPreferentialTier.md) |  | 

### Return type

[**IoPreferentialTier**](IoPreferentialTier.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveNode

> RemoveNode(ctx, id).Execute()

Remove a Node from a Cohesity Cluster.



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
    id := int64(789) // int64 | Specifies the ID of the node being removed.

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

    request := cohesitysdk.ApiRemoveNodeRequest{
        Id: &id
    }

    resp, r, err := api_client.ClustersApi.RemoveNode(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.RemoveNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the ID of the node being removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNodeRequest struct via the builder pattern


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


## UpdateBackgroundActivitySchedule

> BandwidthLimit UpdateBackgroundActivitySchedule(ctx).Execute()

Update the Apollo Throttling schedule.



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

    request := cohesitysdk.ApiUpdateBackgroundActivityScheduleRequest{
    }

    resp, r, err := api_client.ClustersApi.UpdateBackgroundActivitySchedule(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpdateBackgroundActivitySchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBackgroundActivitySchedule`: BandwidthLimit
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpdateBackgroundActivitySchedule`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBackgroundActivityScheduleRequest struct via the builder pattern


### Return type

[**BandwidthLimit**](BandwidthLimit.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExternalClientSubnets

> ExternalClientSubnets UpdateExternalClientSubnets(ctx).Body(body).Execute()

Update the external Client Subnet of the Cluster.



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
    body := *openapiclient.NewExternalClientSubnets() // ExternalClientSubnets |  (optional)

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

    request := cohesitysdk.ApiUpdateExternalClientSubnetsRequest{
        Body: &body
    }

    resp, r, err := api_client.ClustersApi.UpdateExternalClientSubnets(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpdateExternalClientSubnets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExternalClientSubnets`: ExternalClientSubnets
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpdateExternalClientSubnets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExternalClientSubnetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ExternalClientSubnets**](ExternalClientSubnets.md) |  | 

### Return type

[**ExternalClientSubnets**](ExternalClientSubnets.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeCluster

> UpgradeClusterResult UpgradeCluster(ctx).Body(body).Execute()

Upgrade a Cohesity Cluster.



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
    body := *openapiclient.NewUpgradeClusterParameters("TargetSwVersion_example") // UpgradeClusterParameters | 

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

    request := cohesitysdk.ApiUpgradeClusterRequest{
        Body: &body
    }

    resp, r, err := api_client.ClustersApi.UpgradeCluster(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpgradeCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeCluster`: UpgradeClusterResult
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpgradeCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpgradeClusterParameters**](UpgradeClusterParameters.md) |  | 

### Return type

[**UpgradeClusterResult**](UpgradeClusterResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

