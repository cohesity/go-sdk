# \Platform

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRemoteDisk**](Platform.md#AddRemoteDisk) | **Post** /disks/remote | Add remote disk
[**CreateCluster**](Platform.md#CreateCluster) | **Post** /clusters | Create a cluster.
[**CreateRacks**](Platform.md#CreateRacks) | **Post** /racks | Create racks
[**DeleteAMQPTargetConfig**](Platform.md#DeleteAMQPTargetConfig) | **Delete** /clusters/amqp-target-config | Delete AMQP Target Config
[**DeleteRackById**](Platform.md#DeleteRackById) | **Delete** /racks/{id} | Delete a rack by id.
[**DeleteRacks**](Platform.md#DeleteRacks) | **Delete** /racks | Delete all the racks.
[**DiscoverDisks**](Platform.md#DiscoverDisks) | **Get** /disks/discover | Discover new disks
[**DiskIdentify**](Platform.md#DiskIdentify) | **Post** /disks/identify | Identify a disk
[**DisksAssimilate**](Platform.md#DisksAssimilate) | **Post** /disks/assimilate | Assimilate disks.
[**GetAMQPTargetConfig**](Platform.md#GetAMQPTargetConfig) | **Get** /clusters/amqp-target-config | Get AMQP Target Config
[**GetChassis**](Platform.md#GetChassis) | **Get** /chassis | Get list of chassis
[**GetChassisById**](Platform.md#GetChassisById) | **Get** /chassis/{id} | Get a chassis by chassis id.
[**GetCluster**](Platform.md#GetCluster) | **Get** /clusters | Retrieve Cluster Configuration
[**GetClusterLocalDomainSID**](Platform.md#GetClusterLocalDomainSID) | **Get** /clusters/local-domain-sid | Get Cluster Local Domain SID
[**GetIsDMaaSCluster**](Platform.md#GetIsDMaaSCluster) | **Get** /clusters/is-dmaas | Get whether the cluster is a DMaaS cluster.
[**GetNetworkInterfaces**](Platform.md#GetNetworkInterfaces) | **Get** /network-interfaces | Get list of interfaces
[**GetRackById**](Platform.md#GetRackById) | **Get** /racks/{id} | Get a rack by rack id.
[**GetRacks**](Platform.md#GetRacks) | **Get** /racks | Get list of racks
[**GetRemoteDisks**](Platform.md#GetRemoteDisks) | **Get** /disks/remote | Get remote disks
[**GetSupportChannelConfig**](Platform.md#GetSupportChannelConfig) | **Get** /support-channel-config | Get support channel configuration.
[**IdentifyNode**](Platform.md#IdentifyNode) | **Post** /nodes/{id}/identify | Identify node
[**ListDisks**](Platform.md#ListDisks) | **Get** /disks/local | Get list of disks
[**ListFeatureFlag**](Platform.md#ListFeatureFlag) | **Get** /clusters/feature-flag | Get feature flag overrides list.
[**MarkDiskRemoval**](Platform.md#MarkDiskRemoval) | **Post** /disks/{id}/remove | Mark Disk for removal
[**MarkNodeRemoval**](Platform.md#MarkNodeRemoval) | **Post** /nodes/{id}/remove | Mark Node for removal
[**RemoveRemoteDisk**](Platform.md#RemoveRemoteDisk) | **Delete** /disks/remote/{id} | Remove remote disk
[**SetNodePower**](Platform.md#SetNodePower) | **Post** /node-power | Reboot or shutdown nodes in cluster.
[**UpdateAMQPTargetConfig**](Platform.md#UpdateAMQPTargetConfig) | **Put** /clusters/amqp-target-config | Update AMQP Target Config
[**UpdateChassisById**](Platform.md#UpdateChassisById) | **Patch** /chassis/{id} | Update a chassis by chassis id.
[**UpdateCluster**](Platform.md#UpdateCluster) | **Put** /clusters | Update a cluster.
[**UpdateFeatureFlag**](Platform.md#UpdateFeatureFlag) | **Put** /clusters/feature-flag | Update feature flag override status.
[**UpdateIsDMaaSCluster**](Platform.md#UpdateIsDMaaSCluster) | **Put** /clusters/is-dmaas | Update whether the cluster is a DMaaS cluster.
[**UpdateRackById**](Platform.md#UpdateRackById) | **Patch** /racks/{id} | 
[**UpdateRacks**](Platform.md#UpdateRacks) | **Patch** /racks | Update racks
[**UpdateSupportChannelConfig**](Platform.md#UpdateSupportChannelConfig) | **Put** /support-channel-config | Update support channel configuration.



## AddRemoteDisk

> AddRemoteDiskResponseBody AddRemoteDisk(ctx).Body(body).Execute()

Add remote disk



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewRemoteDisks() // RemoteDisks | Specifies the remote disk configuration.

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

    request := onprem.ApiAddRemoteDiskRequest{
        Body: &body
    }

    resp, r, err := api_client.Platform.AddRemoteDisk(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.AddRemoteDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRemoteDisk`: AddRemoteDiskResponseBody
    fmt.Fprintf(os.Stdout, "Response from `Platform.AddRemoteDisk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRemoteDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RemoteDisks**](RemoteDisks.md) | Specifies the remote disk configuration. | 

### Return type

[**AddRemoteDiskResponseBody**](AddRemoteDiskResponseBody.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCluster

> Cluster CreateCluster(ctx).Body(body).Execute()

Create a cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewCreateClusterParams("Name_example", "Type_example", false, *openapiclient.NewClusterCreateNetworkConfig([]string{"NtpServers_example"}, []string{"DomainNames_example"}, false)) // CreateClusterParams | Specifies the parameters to create cluster.

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

    request := onprem.ApiCreateClusterRequest{
        Body: &body
    }

    resp, r, err := api_client.Platform.CreateCluster(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.CreateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `Platform.CreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateClusterParams**](CreateClusterParams.md) | Specifies the parameters to create cluster. | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRacks

> Racks CreateRacks(ctx).Body(body).Execute()

Create racks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewRacks() // Racks | Specifies the parameters to create racks.

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

    request := onprem.ApiCreateRacksRequest{
        Body: &body
    }

    resp, r, err := api_client.Platform.CreateRacks(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.CreateRacks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRacks`: Racks
    fmt.Fprintf(os.Stdout, "Response from `Platform.CreateRacks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRacksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Racks**](Racks.md) | Specifies the parameters to create racks. | 

### Return type

[**Racks**](Racks.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAMQPTargetConfig

> DeleteAMQPTargetConfig(ctx).Execute()

Delete AMQP Target Config



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
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

    request := onprem.ApiDeleteAMQPTargetConfigRequest{
    }

    resp, r, err := api_client.Platform.DeleteAMQPTargetConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.DeleteAMQPTargetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAMQPTargetConfigRequest struct via the builder pattern


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


## DeleteRackById

> DeleteRackById(ctx, id).Execute()

Delete a rack by id.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the rack.

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

    request := onprem.ApiDeleteRackByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.Platform.DeleteRackById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.DeleteRackById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the rack. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRackByIdRequest struct via the builder pattern


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


## DeleteRacks

> DeleteRacks(ctx).Execute()

Delete all the racks.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
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

    request := onprem.ApiDeleteRacksRequest{
    }

    resp, r, err := api_client.Platform.DeleteRacks(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.DeleteRacks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRacksRequest struct via the builder pattern


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


## DiscoverDisks

> ClusterFreeDisks DiscoverDisks(ctx).Execute()

Discover new disks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
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

    request := onprem.ApiDiscoverDisksRequest{
    }

    resp, r, err := api_client.Platform.DiscoverDisks(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.DiscoverDisks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DiscoverDisks`: ClusterFreeDisks
    fmt.Fprintf(os.Stdout, "Response from `Platform.DiscoverDisks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverDisksRequest struct via the builder pattern


### Return type

[**ClusterFreeDisks**](ClusterFreeDisks.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiskIdentify

> DiskIdentify DiskIdentify(ctx).Body(body).Execute()

Identify a disk



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewDiskIdentify(NullableInt64(123), "SerialNumber_example", false) // DiskIdentify | Specifies the parameter to identify disk.

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

    request := onprem.ApiDiskIdentifyRequest{
        Body: &body
    }

    resp, r, err := api_client.Platform.DiskIdentify(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.DiskIdentify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DiskIdentify`: DiskIdentify
    fmt.Fprintf(os.Stdout, "Response from `Platform.DiskIdentify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiskIdentifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DiskIdentify**](DiskIdentify.md) | Specifies the parameter to identify disk. | 

### Return type

[**DiskIdentify**](DiskIdentify.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisksAssimilate

> ClusterFreeDisks DisksAssimilate(ctx).Body(body).Execute()

Assimilate disks.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewClusterFreeDisks([]openapiclient.NodeFreeDisks{*openapiclient.NewNodeFreeDisks(NullableInt64(123), []openapiclient.FreeDisk{*openapiclient.NewFreeDisk("SerialNumber_example")})}) // ClusterFreeDisks | Specifies the parameter to assimilate disks.

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

    request := onprem.ApiDisksAssimilateRequest{
        Body: &body
    }

    resp, r, err := api_client.Platform.DisksAssimilate(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.DisksAssimilate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisksAssimilate`: ClusterFreeDisks
    fmt.Fprintf(os.Stdout, "Response from `Platform.DisksAssimilate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisksAssimilateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClusterFreeDisks**](ClusterFreeDisks.md) | Specifies the parameter to assimilate disks. | 

### Return type

[**ClusterFreeDisks**](ClusterFreeDisks.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAMQPTargetConfig

> ClusterAMQPTargetConfig GetAMQPTargetConfig(ctx).Execute()

Get AMQP Target Config



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
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

    request := onprem.ApiGetAMQPTargetConfigRequest{
    }

    resp, r, err := api_client.Platform.GetAMQPTargetConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.GetAMQPTargetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAMQPTargetConfig`: ClusterAMQPTargetConfig
    fmt.Fprintf(os.Stdout, "Response from `Platform.GetAMQPTargetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAMQPTargetConfigRequest struct via the builder pattern


### Return type

[**ClusterAMQPTargetConfig**](ClusterAMQPTargetConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChassis

> ChassisList GetChassis(ctx).NoRackAssigned(noRackAssigned).Execute()

Get list of chassis



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    noRackAssigned := true // bool | Filters chassis that have no rack assigned. (optional)

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

    request := onprem.ApiGetChassisRequest{
        NoRackAssigned: &noRackAssigned
    }

    resp, r, err := api_client.Platform.GetChassis(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.GetChassis``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChassis`: ChassisList
    fmt.Fprintf(os.Stdout, "Response from `Platform.GetChassis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChassisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noRackAssigned** | **bool** | Filters chassis that have no rack assigned. | 

### Return type

[**ChassisList**](ChassisList.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChassisById

> Chassis GetChassisById(ctx, id).Execute()

Get a chassis by chassis id.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    id := int64(789) // int64 | Specifies the id of chassis.

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

    request := onprem.ApiGetChassisByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.Platform.GetChassisById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.GetChassisById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChassisById`: Chassis
    fmt.Fprintf(os.Stdout, "Response from `Platform.GetChassisById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of chassis. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChassisByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Chassis**](Chassis.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCluster

> Cluster GetCluster(ctx).Execute()

Retrieve Cluster Configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
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

    request := onprem.ApiGetClusterRequest{
    }

    resp, r, err := api_client.Platform.GetCluster(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.GetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `Platform.GetCluster`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterRequest struct via the builder pattern


### Return type

[**Cluster**](Cluster.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterLocalDomainSID

> ClusterLocalDomainSID GetClusterLocalDomainSID(ctx).Execute()

Get Cluster Local Domain SID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
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

    request := onprem.ApiGetClusterLocalDomainSIDRequest{
    }

    resp, r, err := api_client.Platform.GetClusterLocalDomainSID(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.GetClusterLocalDomainSID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterLocalDomainSID`: ClusterLocalDomainSID
    fmt.Fprintf(os.Stdout, "Response from `Platform.GetClusterLocalDomainSID`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterLocalDomainSIDRequest struct via the builder pattern


### Return type

[**ClusterLocalDomainSID**](ClusterLocalDomainSID.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIsDMaaSCluster

> DMaaSInfo GetIsDMaaSCluster(ctx).Execute()

Get whether the cluster is a DMaaS cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
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

    request := onprem.ApiGetIsDMaaSClusterRequest{
    }

    resp, r, err := api_client.Platform.GetIsDMaaSCluster(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.GetIsDMaaSCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIsDMaaSCluster`: DMaaSInfo
    fmt.Fprintf(os.Stdout, "Response from `Platform.GetIsDMaaSCluster`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIsDMaaSClusterRequest struct via the builder pattern


### Return type

[**DMaaSInfo**](DMaaSInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkInterfaces

> ClusterInterfaces GetNetworkInterfaces(ctx).Execute()

Get list of interfaces



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
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

    request := onprem.ApiGetNetworkInterfacesRequest{
    }

    resp, r, err := api_client.Platform.GetNetworkInterfaces(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.GetNetworkInterfaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkInterfaces`: ClusterInterfaces
    fmt.Fprintf(os.Stdout, "Response from `Platform.GetNetworkInterfaces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkInterfacesRequest struct via the builder pattern


### Return type

[**ClusterInterfaces**](ClusterInterfaces.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRackById

> Rack GetRackById(ctx, id).Execute()

Get a rack by rack id.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    id := int64(789) // int64 | Specifies the id of rack.

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

    request := onprem.ApiGetRackByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.Platform.GetRackById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.GetRackById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRackById`: Rack
    fmt.Fprintf(os.Stdout, "Response from `Platform.GetRackById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of rack. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRackByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Rack**](Rack.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRacks

> Racks GetRacks(ctx).Execute()

Get list of racks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
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

    request := onprem.ApiGetRacksRequest{
    }

    resp, r, err := api_client.Platform.GetRacks(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.GetRacks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRacks`: Racks
    fmt.Fprintf(os.Stdout, "Response from `Platform.GetRacks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRacksRequest struct via the builder pattern


### Return type

[**Racks**](Racks.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteDisks

> RemoteDisks GetRemoteDisks(ctx).DiskIds(diskIds).NodeIds(nodeIds).Tiers(tiers).MountPath(mountPath).FileSystem(fileSystem).Execute()

Get remote disks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    diskIds := []int64{int64(123)} // []int64 | Specifies a list of disk ids, only disks having these ids will be returned. (optional)
    nodeIds := []int64{int64(123)} // []int64 | Specifies a list of node ids, only disks in these nodes will be returned. (optional)
    tiers := []string{"Tiers_example"} // []string | Specifies a list of disk tiers, only disks with given tiers will be returned. (optional)
    mountPath := "mountPath_example" // string | This field is deprecated. Providing this queryparam will not have any impact. Please use fileSystem query param to filter instead. (optional)
    fileSystem := "fileSystem_example" // string | Specified file system name to search. only disks with file system name that partially matches the specified name will be returned. (optional)

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

    request := onprem.ApiGetRemoteDisksRequest{
        DiskIds: &diskIds
        NodeIds: &nodeIds
        Tiers: &tiers
        MountPath: &mountPath
        FileSystem: &fileSystem
    }

    resp, r, err := api_client.Platform.GetRemoteDisks(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.GetRemoteDisks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemoteDisks`: RemoteDisks
    fmt.Fprintf(os.Stdout, "Response from `Platform.GetRemoteDisks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteDisksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **diskIds** | **[]int64** | Specifies a list of disk ids, only disks having these ids will be returned. | 
 **nodeIds** | **[]int64** | Specifies a list of node ids, only disks in these nodes will be returned. | 
 **tiers** | **[]string** | Specifies a list of disk tiers, only disks with given tiers will be returned. | 
 **mountPath** | **string** | This field is deprecated. Providing this queryparam will not have any impact. Please use fileSystem query param to filter instead. | 
 **fileSystem** | **string** | Specified file system name to search. only disks with file system name that partially matches the specified name will be returned. | 

### Return type

[**RemoteDisks**](RemoteDisks.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportChannelConfig

> SupportChannelConfig GetSupportChannelConfig(ctx).Execute()

Get support channel configuration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
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

    request := onprem.ApiGetSupportChannelConfigRequest{
    }

    resp, r, err := api_client.Platform.GetSupportChannelConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.GetSupportChannelConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupportChannelConfig`: SupportChannelConfig
    fmt.Fprintf(os.Stdout, "Response from `Platform.GetSupportChannelConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportChannelConfigRequest struct via the builder pattern


### Return type

[**SupportChannelConfig**](SupportChannelConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentifyNode

> NodeIdentifyParams IdentifyNode(ctx, id).Body(body).Execute()

Identify node



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    id := int64(789) // int64 | Specifies id of node to identify.
    body := *openapiclient.NewNodeIdentifyParams(false) // NodeIdentifyParams | Specifies the parameter to identify node.

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

    request := onprem.ApiIdentifyNodeRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.Platform.IdentifyNode(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.IdentifyNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentifyNode`: NodeIdentifyParams
    fmt.Fprintf(os.Stdout, "Response from `Platform.IdentifyNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies id of node to identify. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentifyNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NodeIdentifyParams**](NodeIdentifyParams.md) | Specifies the parameter to identify node. | 

### Return type

[**NodeIdentifyParams**](NodeIdentifyParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDisks

> DisksList ListDisks(ctx).NodeId(nodeId).Execute()

Get list of disks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    nodeId := int64(789) // int64 | Specifies node id of the node to get list of disks (optional)

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

    request := onprem.ApiListDisksRequest{
        NodeId: &nodeId
    }

    resp, r, err := api_client.Platform.ListDisks(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.ListDisks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDisks`: DisksList
    fmt.Fprintf(os.Stdout, "Response from `Platform.ListDisks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDisksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeId** | **int64** | Specifies node id of the node to get list of disks | 

### Return type

[**DisksList**](DisksList.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFeatureFlag

> ListFeatureFlag(ctx).Execute()

Get feature flag overrides list.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
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

    request := onprem.ApiListFeatureFlagRequest{
    }

    resp, r, err := api_client.Platform.ListFeatureFlag(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.ListFeatureFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFeatureFlagRequest struct via the builder pattern


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


## MarkDiskRemoval

> RemoveDisk MarkDiskRemoval(ctx, id).Body(body).Execute()

Mark Disk for removal



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    id := int64(789) // int64 | Specifies unique id of the disk to mark for removal.
    body := *openapiclient.NewDiskRemovalParams(false) // DiskRemovalParams | Specifies parameters to mark/cancel disk removal.

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

    request := onprem.ApiMarkDiskRemovalRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.Platform.MarkDiskRemoval(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.MarkDiskRemoval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarkDiskRemoval`: RemoveDisk
    fmt.Fprintf(os.Stdout, "Response from `Platform.MarkDiskRemoval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies unique id of the disk to mark for removal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkDiskRemovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DiskRemovalParams**](DiskRemovalParams.md) | Specifies parameters to mark/cancel disk removal. | 

### Return type

[**RemoveDisk**](RemoveDisk.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkNodeRemoval

> RemoveNode MarkNodeRemoval(ctx, id).Body(body).Execute()

Mark Node for removal



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    id := int64(789) // int64 | Specifies id of node to cancel removal.
    body := *openapiclient.NewNodeRemovalParams(false) // NodeRemovalParams | Specifies parameters to initiate/cancel node removal .

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

    request := onprem.ApiMarkNodeRemovalRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.Platform.MarkNodeRemoval(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.MarkNodeRemoval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarkNodeRemoval`: RemoveNode
    fmt.Fprintf(os.Stdout, "Response from `Platform.MarkNodeRemoval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies id of node to cancel removal. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkNodeRemovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NodeRemovalParams**](NodeRemovalParams.md) | Specifies parameters to initiate/cancel node removal . | 

### Return type

[**RemoveNode**](RemoveNode.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRemoteDisk

> RemoveRemoteDisk(ctx, id).Execute()

Remove remote disk



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    id := int64(789) // int64 | Specifies the id of the remote disk to remove.

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

    request := onprem.ApiRemoveRemoteDiskRequest{
        Id: &id
    }

    resp, r, err := api_client.Platform.RemoveRemoteDisk(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.RemoveRemoteDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of the remote disk to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRemoteDiskRequest struct via the builder pattern


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


## SetNodePower

> SetNodePower(ctx).Body(body).Execute()

Reboot or shutdown nodes in cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewNodePowerOperation("Operation_example") // NodePowerOperation | Specifies the reboot or shutdown operation.

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

    request := onprem.ApiSetNodePowerRequest{
        Body: &body
    }

    resp, r, err := api_client.Platform.SetNodePower(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.SetNodePower``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetNodePowerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NodePowerOperation**](NodePowerOperation.md) | Specifies the reboot or shutdown operation. | 

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


## UpdateAMQPTargetConfig

> ClusterAMQPTargetConfig UpdateAMQPTargetConfig(ctx).Body(body).Execute()

Update AMQP Target Config



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewClusterAMQPTargetConfig() // ClusterAMQPTargetConfig | Specifies the parameters to update cluster AMQP target config.

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

    request := onprem.ApiUpdateAMQPTargetConfigRequest{
        Body: &body
    }

    resp, r, err := api_client.Platform.UpdateAMQPTargetConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.UpdateAMQPTargetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAMQPTargetConfig`: ClusterAMQPTargetConfig
    fmt.Fprintf(os.Stdout, "Response from `Platform.UpdateAMQPTargetConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAMQPTargetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClusterAMQPTargetConfig**](ClusterAMQPTargetConfig.md) | Specifies the parameters to update cluster AMQP target config. | 

### Return type

[**ClusterAMQPTargetConfig**](ClusterAMQPTargetConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChassisById

> Chassis UpdateChassisById(ctx, id).Body(body).Execute()

Update a chassis by chassis id.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    id := int64(789) // int64 | Specifies the id of chassis.
    body := *openapiclient.NewChassis() // Chassis | Specifies the parameters to update chassis. (optional)

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

    request := onprem.ApiUpdateChassisByIdRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.Platform.UpdateChassisById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.UpdateChassisById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateChassisById`: Chassis
    fmt.Fprintf(os.Stdout, "Response from `Platform.UpdateChassisById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of chassis. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChassisByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Chassis**](Chassis.md) | Specifies the parameters to update chassis. | 

### Return type

[**Chassis**](Chassis.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCluster

> Cluster UpdateCluster(ctx).Body(body).Execute()

Update a cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewCluster() // Cluster | Specifies the parameters to update cluster.

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

    request := onprem.ApiUpdateClusterRequest{
        Body: &body
    }

    resp, r, err := api_client.Platform.UpdateCluster(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.UpdateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `Platform.UpdateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Cluster**](Cluster.md) | Specifies the parameters to update cluster. | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeatureFlag

> UpdateFeatureFlag(ctx).Body(body).Execute()

Update feature flag override status.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewUpdateFeatureFlagParams() // UpdateFeatureFlagParams | Param for feature flag override request.

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

    request := onprem.ApiUpdateFeatureFlagRequest{
        Body: &body
    }

    resp, r, err := api_client.Platform.UpdateFeatureFlag(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.UpdateFeatureFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeatureFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateFeatureFlagParams**](UpdateFeatureFlagParams.md) | Param for feature flag override request. | 

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


## UpdateIsDMaaSCluster

> DMaaSInfo UpdateIsDMaaSCluster(ctx).Body(body).Execute()

Update whether the cluster is a DMaaS cluster.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewDMaaSInfo(false) // DMaaSInfo | Param to update whether the cluster is a DMaaS cluster.

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

    request := onprem.ApiUpdateIsDMaaSClusterRequest{
        Body: &body
    }

    resp, r, err := api_client.Platform.UpdateIsDMaaSCluster(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.UpdateIsDMaaSCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIsDMaaSCluster`: DMaaSInfo
    fmt.Fprintf(os.Stdout, "Response from `Platform.UpdateIsDMaaSCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIsDMaaSClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DMaaSInfo**](DMaaSInfo.md) | Param to update whether the cluster is a DMaaS cluster. | 

### Return type

[**DMaaSInfo**](DMaaSInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRackById

> Rack UpdateRackById(ctx, id).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    id := int64(789) // int64 | Specifies the id of rack.
    body := *openapiclient.NewRack() // Rack | Specifies the parameters to update rack. (optional)

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

    request := onprem.ApiUpdateRackByIdRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.Platform.UpdateRackById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.UpdateRackById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRackById`: Rack
    fmt.Fprintf(os.Stdout, "Response from `Platform.UpdateRackById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the id of rack. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRackByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Rack**](Rack.md) | Specifies the parameters to update rack. | 

### Return type

[**Rack**](Rack.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRacks

> Racks UpdateRacks(ctx).Body(body).Execute()

Update racks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewRacks() // Racks | Specifies the parameters to update racks.

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

    request := onprem.ApiUpdateRacksRequest{
        Body: &body
    }

    resp, r, err := api_client.Platform.UpdateRacks(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.UpdateRacks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRacks`: Racks
    fmt.Fprintf(os.Stdout, "Response from `Platform.UpdateRacks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRacksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Racks**](Racks.md) | Specifies the parameters to update racks. | 

### Return type

[**Racks**](Racks.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSupportChannelConfig

> SupportChannelConfig UpdateSupportChannelConfig(ctx).Body(body).Execute()

Update support channel configuration.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewSupportChannelConfig(false, NullableInt64(123)) // SupportChannelConfig | Specifies the support channel configuration.

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

    request := onprem.ApiUpdateSupportChannelConfigRequest{
        Body: &body
    }

    resp, r, err := api_client.Platform.UpdateSupportChannelConfig(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Platform.UpdateSupportChannelConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSupportChannelConfig`: SupportChannelConfig
    fmt.Fprintf(os.Stdout, "Response from `Platform.UpdateSupportChannelConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSupportChannelConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SupportChannelConfig**](SupportChannelConfig.md) | Specifies the support channel configuration. | 

### Return type

[**SupportChannelConfig**](SupportChannelConfig.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

