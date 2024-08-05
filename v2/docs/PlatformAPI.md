# \PlatformAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddHosts**](PlatformAPI.md#AddHosts) | **Post** /clusters/host-mappings | Create Cluster Host Mappings
[**AddRemoteDisk**](PlatformAPI.md#AddRemoteDisk) | **Post** /disks/remote | Add remote disk
[**ClearSMTPConfiguration**](PlatformAPI.md#ClearSMTPConfiguration) | **Delete** /clusters/smtp | Clear SMTP configuration.
[**ClusterUpdateIpmiUsers**](PlatformAPI.md#ClusterUpdateIpmiUsers) | **Put** /ipmi/cluster-users | To update IPMI Users for cluster
[**CreateCluster**](PlatformAPI.md#CreateCluster) | **Post** /clusters | Create a cluster.
[**CreateClusterVlan**](PlatformAPI.md#CreateClusterVlan) | **Post** /network/vlans | Create vlan
[**CreateInterfaceGroup**](PlatformAPI.md#CreateInterfaceGroup) | **Post** /network/interface-groups | Create interface group
[**CreateRacks**](PlatformAPI.md#CreateRacks) | **Post** /racks | Create racks
[**DeleteAMQPTargetConfig**](PlatformAPI.md#DeleteAMQPTargetConfig) | **Delete** /clusters/amqp-target-config | Delete AMQP Target Config
[**DeleteClusterSnapshotPolicy**](PlatformAPI.md#DeleteClusterSnapshotPolicy) | **Delete** /clusters/snapshot-policy | Delete cluster snapshot policy.
[**DeleteHosts**](PlatformAPI.md#DeleteHosts) | **Post** /clusters/host-mappings/delete | Deletes multiple Host Mappings within the cluster
[**DiscoverDisks**](PlatformAPI.md#DiscoverDisks) | **Get** /disks/discover | Discover new disks
[**DiskIdentify**](PlatformAPI.md#DiskIdentify) | **Post** /disks/identify | Identify a disk
[**DisksAssimilate**](PlatformAPI.md#DisksAssimilate) | **Post** /disks/assimilate | Assimilate disks.
[**GetAMQPTargetConfig**](PlatformAPI.md#GetAMQPTargetConfig) | **Get** /clusters/amqp-target-config | Get AMQP Target Config
[**GetChassis**](PlatformAPI.md#GetChassis) | **Get** /chassis | Get list of chassis
[**GetChassisById**](PlatformAPI.md#GetChassisById) | **Get** /chassis/{id} | Get a chassis by chassis id.
[**GetCluster**](PlatformAPI.md#GetCluster) | **Get** /clusters | Retrieve Cluster Configuration
[**GetClusterIpmiLanInfo**](PlatformAPI.md#GetClusterIpmiLanInfo) | **Get** /ipmi/cluster-get-lan-info | To get IPMI LAN info for the cluster
[**GetClusterIpmiUsers**](PlatformAPI.md#GetClusterIpmiUsers) | **Get** /ipmi/cluster-users | To get IPMI users info for the cluster
[**GetClusterLocalDomainSID**](PlatformAPI.md#GetClusterLocalDomainSID) | **Get** /clusters/local-domain-sid | Get Cluster Local Domain SID
[**GetClusterOperationStatusList**](PlatformAPI.md#GetClusterOperationStatusList) | **Get** /clusters/operation-status | Get cluster operations status.
[**GetClusterPackages**](PlatformAPI.md#GetClusterPackages) | **Get** /clusters/packages | Get packages
[**GetClusterSnapshotPolicy**](PlatformAPI.md#GetClusterSnapshotPolicy) | **Get** /clusters/snapshot-policy | Get cluster snapshot policy.
[**GetClusterState**](PlatformAPI.md#GetClusterState) | **Get** /clusters/state | Get cluster state
[**GetClusterStatus**](PlatformAPI.md#GetClusterStatus) | **Get** /clusters/status | Get cluster status.
[**GetClusterSubnetsInfo**](PlatformAPI.md#GetClusterSubnetsInfo) | **Get** /clusters/subnets | Get cluster subnets info.
[**GetHardwareInfo**](PlatformAPI.md#GetHardwareInfo) | **Get** /node/hardware-info | Fetch Node Hardware Information
[**GetIpmiFruInfo**](PlatformAPI.md#GetIpmiFruInfo) | **Get** /ipmi/get-fru-info | To get IPMI FRU info
[**GetIpmiLanInfo**](PlatformAPI.md#GetIpmiLanInfo) | **Get** /ipmi/get-lan-info | To get IPMI LAN info
[**GetIpmiSdrInfo**](PlatformAPI.md#GetIpmiSdrInfo) | **Get** /ipmi/get-sdr-info | To get IPMI SDR Info
[**GetIpmiSel**](PlatformAPI.md#GetIpmiSel) | **Get** /ipmi/get-sel | To get IPMI SEL
[**GetIpmiSelInfo**](PlatformAPI.md#GetIpmiSelInfo) | **Get** /ipmi/get-sel-info | To get IPMI SEL Info
[**GetIpmiUsers**](PlatformAPI.md#GetIpmiUsers) | **Get** /ipmi/users | To get IPMI User Info for node
[**GetIsDMaaSCluster**](PlatformAPI.md#GetIsDMaaSCluster) | **Get** /clusters/is-dmaas | Get whether the cluster is a DMaaS cluster.
[**GetKubernetesInfraHealthStatus**](PlatformAPI.md#GetKubernetesInfraHealthStatus) | **Get** /kubernetes/status | Get Kubernetes Infra Health Status
[**GetNetworkInterfaces**](PlatformAPI.md#GetNetworkInterfaces) | **Get** /network-interfaces | Get list of interfaces
[**GetNodes**](PlatformAPI.md#GetNodes) | **Get** /clusters/nodes | List Nodes of the cluster.
[**GetRackById**](PlatformAPI.md#GetRackById) | **Get** /racks/{id} | Get a rack by rack id.
[**GetRacks**](PlatformAPI.md#GetRacks) | **Get** /racks | Get list of racks
[**GetRemoteDisks**](PlatformAPI.md#GetRemoteDisks) | **Get** /disks/remote | Get remote disks
[**GetSMTPConfiguration**](PlatformAPI.md#GetSMTPConfiguration) | **Get** /clusters/smtp | Get SMTP configuration.
[**GetSWUpdateHistory**](PlatformAPI.md#GetSWUpdateHistory) | **Get** /clusters/softwares | Get cluster software history
[**GetServiceGflags**](PlatformAPI.md#GetServiceGflags) | **Get** /clusters/gflag | Gets cluster gflags for a service.
[**GetSoftwareComponents**](PlatformAPI.md#GetSoftwareComponents) | **Get** /clusters/software-components | Get Software Components
[**GetSupportChannelConfig**](PlatformAPI.md#GetSupportChannelConfig) | **Get** /support-channel-config | Get support channel configuration.
[**IdentifyNode**](PlatformAPI.md#IdentifyNode) | **Post** /nodes/{id}/identify | Identify node
[**ImportCrlFile**](PlatformAPI.md#ImportCrlFile) | **Put** /clusters/import-crl-file | Import Crl File
[**ListDisks**](PlatformAPI.md#ListDisks) | **Get** /disks/local | Get list of disks
[**ListFreeNodes**](PlatformAPI.md#ListFreeNodes) | **Get** /clusters/nodes/free | List the free Cohesity Nodes present on a network.
[**ListHosts**](PlatformAPI.md#ListHosts) | **Get** /clusters/host-mappings | List Host Mappings
[**MarkBaseosUpgrade**](PlatformAPI.md#MarkBaseosUpgrade) | **Put** /clusters/baseos-upgrade | Sets/clears the BaseOS upgrade cluster operation.
[**MarkDiskRemoval**](PlatformAPI.md#MarkDiskRemoval) | **Post** /disks/{id}/remove | Mark Disk for removal
[**MarkNodeRemoval**](PlatformAPI.md#MarkNodeRemoval) | **Post** /nodes/{id}/remove | Mark Node for removal
[**NodeInformation**](PlatformAPI.md#NodeInformation) | **Get** /nodes | Fetch Node General Information
[**NodeStatus**](PlatformAPI.md#NodeStatus) | **Get** /node/status | Fetch Node status Information
[**PublicKeyRequest**](PlatformAPI.md#PublicKeyRequest) | **Post** /clusters/ssh-public-key | Get the SSH public key.
[**RemoveRemoteDisk**](PlatformAPI.md#RemoveRemoteDisk) | **Delete** /disks/remote/{id} | Remove remote disk
[**ResetIpmiBmc**](PlatformAPI.md#ResetIpmiBmc) | **Post** /ipmi/reset-bmc | To reset IPMI BMC for given node
[**SetNodePower**](PlatformAPI.md#SetNodePower) | **Post** /node-power | Reboot or shutdown nodes in cluster.
[**UpdateAMQPTargetConfig**](PlatformAPI.md#UpdateAMQPTargetConfig) | **Put** /clusters/amqp-target-config | Update AMQP Target Config
[**UpdateAirgapConfig**](PlatformAPI.md#UpdateAirgapConfig) | **Put** /clusters/airgap | Update Airgap config
[**UpdateChassisById**](PlatformAPI.md#UpdateChassisById) | **Patch** /chassis/{id} | Update a chassis by chassis id.
[**UpdateCluster**](PlatformAPI.md#UpdateCluster) | **Put** /clusters | Update a cluster.
[**UpdateClusterSnapshotPolicy**](PlatformAPI.md#UpdateClusterSnapshotPolicy) | **Put** /clusters/snapshot-policy | Update cluster snapshot policy.
[**UpdateClusterSoftware**](PlatformAPI.md#UpdateClusterSoftware) | **Put** /clusters/softwares | Update cluster software
[**UpdateFeatureFlag**](PlatformAPI.md#UpdateFeatureFlag) | **Put** /clusters/feature-flag | Update feature flag override status.
[**UpdateHosts**](PlatformAPI.md#UpdateHosts) | **Put** /clusters/host-mappings | Update Host Mappings
[**UpdateIpmiUser**](PlatformAPI.md#UpdateIpmiUser) | **Post** /ipmi/users | To update IPMI User Info for node
[**UpdateIsDMaaSCluster**](PlatformAPI.md#UpdateIsDMaaSCluster) | **Put** /clusters/is-dmaas | Update whether the cluster is a DMaaS cluster.
[**UpdateRackById**](PlatformAPI.md#UpdateRackById) | **Patch** /racks/{id} | 
[**UpdateRacks**](PlatformAPI.md#UpdateRacks) | **Patch** /racks | Update racks
[**UpdateSMTPConfiguration**](PlatformAPI.md#UpdateSMTPConfiguration) | **Put** /clusters/smtp | Update SMTP configuration.
[**UpdateServiceGflags**](PlatformAPI.md#UpdateServiceGflags) | **Put** /clusters/gflag | Update the gflags
[**UpdateSupportChannelConfig**](PlatformAPI.md#UpdateSupportChannelConfig) | **Put** /support-channel-config | Update support channel configuration.
[**UpgradeCheckGetResults**](PlatformAPI.md#UpgradeCheckGetResults) | **Get** /clusters/upgrade-checks/{testRunInstanceId} | Get upgrade checks results.
[**UpgradeCheckRunTests**](PlatformAPI.md#UpgradeCheckRunTests) | **Put** /clusters/upgrade-checks | Run upgrade checks on cluster.
[**UploadFilePackage**](PlatformAPI.md#UploadFilePackage) | **Post** /clusters/packages/file | Upload package by files
[**ValidateSMTPConfiguration**](PlatformAPI.md#ValidateSMTPConfiguration) | **Post** /clusters/smtp/validate | Validate SMTP configuration.



## AddHosts

> HostMappings AddHosts(ctx).Body(body).Execute()

Create Cluster Host Mappings



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
	body := []openapiclient.HostEntry{*openapiclient.NewHostEntry([]string{"DomainNames_example"}, "Ip_example")} // []HostEntry | Specifies the request to add entries to /etc/hosts

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.AddHosts(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.AddHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddHosts`: HostMappings
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.AddHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]HostEntry**](HostEntry.md) | Specifies the request to add entries to /etc/hosts | 

### Return type

[**HostMappings**](HostMappings.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewRemoteDisks() // RemoteDisks | Specifies the remote disk configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.AddRemoteDisk(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.AddRemoteDisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddRemoteDisk`: AddRemoteDiskResponseBody
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.AddRemoteDisk`: %v\n", resp)
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


## ClearSMTPConfiguration

> ClearSMTPConfiguration(ctx).Execute()

Clear SMTP configuration.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.ClearSMTPConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.ClearSMTPConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiClearSMTPConfigurationRequest struct via the builder pattern


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


## ClusterUpdateIpmiUsers

> IpmiTextResponse ClusterUpdateIpmiUsers(ctx).Body(body).Execute()

To update IPMI Users for cluster



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
	body := *openapiclient.NewClusterUpdateIpmiUsers() // ClusterUpdateIpmiUsers | Specifies the parameters to update cluster ipmi users.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.ClusterUpdateIpmiUsers(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.ClusterUpdateIpmiUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClusterUpdateIpmiUsers`: IpmiTextResponse
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.ClusterUpdateIpmiUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClusterUpdateIpmiUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClusterUpdateIpmiUsers**](ClusterUpdateIpmiUsers.md) | Specifies the parameters to update cluster ipmi users. | 

### Return type

[**IpmiTextResponse**](IpmiTextResponse.md)

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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewCreateClusterParams(false, "Name_example", *openapiclient.NewClusterCreateNetworkConfig([]string{"DomainNames_example"}, []string{"NtpServers_example"}, false), "Type_example") // CreateClusterParams | Specifies the parameters to create cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.CreateCluster(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CreateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCluster`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.CreateCluster`: %v\n", resp)
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


## CreateClusterVlan

> ClusterVlanParams CreateClusterVlan(ctx).Body(body).Execute()

Create vlan



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
	body := *openapiclient.NewCreateClusterVlanParams("InterfaceName_example") // CreateClusterVlanParams | Parameters to create a vlan.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.CreateClusterVlan(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CreateClusterVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateClusterVlan`: ClusterVlanParams
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.CreateClusterVlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateClusterVlanParams**](CreateClusterVlanParams.md) | Parameters to create a vlan. | 

### Return type

[**ClusterVlanParams**](ClusterVlanParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInterfaceGroup

> InterfaceGroup CreateInterfaceGroup(ctx).Body(body).Execute()

Create interface group



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
	body := *openapiclient.NewInterfaceGroupParams("Name_example", []openapiclient.NodeInterfaceParams{*openapiclient.NewNodeInterfaceParams(int64(123))}, "Type_example") // InterfaceGroupParams | Parameters to create an interface group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.CreateInterfaceGroup(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CreateInterfaceGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInterfaceGroup`: InterfaceGroup
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.CreateInterfaceGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInterfaceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InterfaceGroupParams**](InterfaceGroupParams.md) | Parameters to create an interface group. | 

### Return type

[**InterfaceGroup**](InterfaceGroup.md)

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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewRacks() // Racks | Specifies the parameters to create racks.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.CreateRacks(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CreateRacks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRacks`: Racks
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.CreateRacks`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.DeleteAMQPTargetConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.DeleteAMQPTargetConfig``: %v\n", err)
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


## DeleteClusterSnapshotPolicy

> DeleteClusterSnapshotPolicy(ctx).Execute()

Delete cluster snapshot policy.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.DeleteClusterSnapshotPolicy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.DeleteClusterSnapshotPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterSnapshotPolicyRequest struct via the builder pattern


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


## DeleteHosts

> DeleteHosts(ctx).Body(body).Execute()

Deletes multiple Host Mappings within the cluster



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
	body := *openapiclient.NewDeleteHostsParameters() // DeleteHostsParameters | Specifies the params to delete host mappings

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.DeleteHosts(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.DeleteHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteHostsParameters**](DeleteHostsParameters.md) | Specifies the params to delete host mappings | 

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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.DiscoverDisks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.DiscoverDisks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverDisks`: ClusterFreeDisks
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.DiscoverDisks`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewDiskIdentify(false) // DiskIdentify | Specifies the parameter to identify disk.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.DiskIdentify(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.DiskIdentify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiskIdentify`: DiskIdentify
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.DiskIdentify`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewClusterFreeDisks([]openapiclient.NodeFreeDisks{*openapiclient.NewNodeFreeDisks([]openapiclient.FreeDisk{*openapiclient.NewFreeDisk("SerialNumber_example")}, NullableInt64(123))}) // ClusterFreeDisks | Specifies the parameter to assimilate disks.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.DisksAssimilate(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.DisksAssimilate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisksAssimilate`: ClusterFreeDisks
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.DisksAssimilate`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetAMQPTargetConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetAMQPTargetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAMQPTargetConfig`: ClusterAMQPTargetConfig
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetAMQPTargetConfig`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	noRackAssigned := true // bool | Filters chassis that have no rack assigned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetChassis(context.Background()).NoRackAssigned(noRackAssigned).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetChassis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChassis`: ChassisList
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetChassis`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := int64(789) // int64 | Specifies the id of chassis.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetChassisById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetChassisById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChassisById`: Chassis
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetChassisById`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetCluster(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCluster`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetCluster`: %v\n", resp)
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


## GetClusterIpmiLanInfo

> ClusterIpmiLanInfo GetClusterIpmiLanInfo(ctx).Execute()

To get IPMI LAN info for the cluster



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetClusterIpmiLanInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetClusterIpmiLanInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterIpmiLanInfo`: ClusterIpmiLanInfo
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetClusterIpmiLanInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterIpmiLanInfoRequest struct via the builder pattern


### Return type

[**ClusterIpmiLanInfo**](ClusterIpmiLanInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterIpmiUsers

> ClusterIpmiUsers GetClusterIpmiUsers(ctx).Execute()

To get IPMI users info for the cluster



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetClusterIpmiUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetClusterIpmiUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterIpmiUsers`: ClusterIpmiUsers
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetClusterIpmiUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterIpmiUsersRequest struct via the builder pattern


### Return type

[**ClusterIpmiUsers**](ClusterIpmiUsers.md)

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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetClusterLocalDomainSID(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetClusterLocalDomainSID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterLocalDomainSID`: ClusterLocalDomainSID
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetClusterLocalDomainSID`: %v\n", resp)
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


## GetClusterOperationStatusList

> []ClusterOperationStatus GetClusterOperationStatusList(ctx).OperationTypes(operationTypes).OperationIds(operationIds).IncludeFinishedOperations(includeFinishedOperations).IncludeEventLogs(includeEventLogs).StartTime(startTime).EndTime(endTime).Execute()

Get cluster operations status.



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
	operationTypes := []string{"OperationTypes_example"} // []string | One or more operation types to query for. (optional)
	operationIds := []string{"Inner_example"} // []string | One or more operation ids to query for. (optional)
	includeFinishedOperations := true // bool | Controls whether finished operations should be included in the query results. The default value is false. Applicable only for patch apply, revert, and upgrade operations (optional)
	includeEventLogs := true // bool | Controls whether event logs should be included in the query results. If set to true, 'operationIds' becomes mandatory. The default value is false. Applicable only for patch apply, revert, and upgrade operations (optional)
	startTime := int64(789) // int64 | Filters operations that started after the specified time. Applicable only for patch apply, revert, and upgrade operations (optional)
	endTime := int64(789) // int64 | Filters operations that ended before the specified time. Applicable only for patch apply, revert, and upgrade operations (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetClusterOperationStatusList(context.Background()).OperationTypes(operationTypes).OperationIds(operationIds).IncludeFinishedOperations(includeFinishedOperations).IncludeEventLogs(includeEventLogs).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetClusterOperationStatusList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterOperationStatusList`: []ClusterOperationStatus
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetClusterOperationStatusList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterOperationStatusListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operationTypes** | **[]string** | One or more operation types to query for. | 
 **operationIds** | **[]string** | One or more operation ids to query for. | 
 **includeFinishedOperations** | **bool** | Controls whether finished operations should be included in the query results. The default value is false. Applicable only for patch apply, revert, and upgrade operations | 
 **includeEventLogs** | **bool** | Controls whether event logs should be included in the query results. If set to true, &#39;operationIds&#39; becomes mandatory. The default value is false. Applicable only for patch apply, revert, and upgrade operations | 
 **startTime** | **int64** | Filters operations that started after the specified time. Applicable only for patch apply, revert, and upgrade operations | 
 **endTime** | **int64** | Filters operations that ended before the specified time. Applicable only for patch apply, revert, and upgrade operations | 

### Return type

[**[]ClusterOperationStatus**](ClusterOperationStatus.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterPackages

> ClusterPackages GetClusterPackages(ctx).Execute()

Get packages



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetClusterPackages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetClusterPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterPackages`: ClusterPackages
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetClusterPackages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterPackagesRequest struct via the builder pattern


### Return type

[**ClusterPackages**](ClusterPackages.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterSnapshotPolicy

> ClusterSnapshotPolicy GetClusterSnapshotPolicy(ctx).Execute()

Get cluster snapshot policy.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetClusterSnapshotPolicy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetClusterSnapshotPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterSnapshotPolicy`: ClusterSnapshotPolicy
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetClusterSnapshotPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterSnapshotPolicyRequest struct via the builder pattern


### Return type

[**ClusterSnapshotPolicy**](ClusterSnapshotPolicy.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterState

> ClusterStateParams GetClusterState(ctx).SystemApps(systemApps).Execute()

Get cluster state



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
	systemApps := true // bool | The filter whether or not to get the system apps state details. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetClusterState(context.Background()).SystemApps(systemApps).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetClusterState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterState`: ClusterStateParams
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetClusterState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemApps** | **bool** | The filter whether or not to get the system apps state details. | 

### Return type

[**ClusterStateParams**](ClusterStateParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterStatus

> ClusterStatus GetClusterStatus(ctx).Execute()

Get cluster status.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetClusterStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetClusterStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterStatus`: ClusterStatus
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetClusterStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterStatusRequest struct via the builder pattern


### Return type

[**ClusterStatus**](ClusterStatus.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterSubnetsInfo

> []Subnet GetClusterSubnetsInfo(ctx).Execute()

Get cluster subnets info.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetClusterSubnetsInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetClusterSubnetsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusterSubnetsInfo`: []Subnet
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetClusterSubnetsInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterSubnetsInfoRequest struct via the builder pattern


### Return type

[**[]Subnet**](Subnet.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHardwareInfo

> HardwareInfo GetHardwareInfo(ctx).Execute()

Fetch Node Hardware Information



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetHardwareInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetHardwareInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHardwareInfo`: HardwareInfo
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetHardwareInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHardwareInfoRequest struct via the builder pattern


### Return type

[**HardwareInfo**](HardwareInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpmiFruInfo

> IpmiFruInfo GetIpmiFruInfo(ctx).NodeId(nodeId).NodeIp(nodeIp).Execute()

To get IPMI FRU info



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
	nodeId := "nodeId_example" // string | Specifies the node id of the node for which fru info is requested. This parameter is incompatible with 'nodeIp'. (optional)
	nodeIp := "nodeIp_example" // string | Specifies the IP Address of the node for which fru info is requested. This parameter is incompatible with 'nodeId'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetIpmiFruInfo(context.Background()).NodeId(nodeId).NodeIp(nodeIp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetIpmiFruInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpmiFruInfo`: IpmiFruInfo
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetIpmiFruInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIpmiFruInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeId** | **string** | Specifies the node id of the node for which fru info is requested. This parameter is incompatible with &#39;nodeIp&#39;. | 
 **nodeIp** | **string** | Specifies the IP Address of the node for which fru info is requested. This parameter is incompatible with &#39;nodeId&#39;. | 

### Return type

[**IpmiFruInfo**](IpmiFruInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpmiLanInfo

> IpmiLanInfo GetIpmiLanInfo(ctx).NodeId(nodeId).NodeIp(nodeIp).Execute()

To get IPMI LAN info



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
	nodeId := "nodeId_example" // string | Specifies the node id of the node for which lan info is requested. This parameter is incompatible with 'nodeIp'. (optional)
	nodeIp := "nodeIp_example" // string | Specifies the IP Address of the node for which lan info is requested. This parameter is incompatible with 'nodeId'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetIpmiLanInfo(context.Background()).NodeId(nodeId).NodeIp(nodeIp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetIpmiLanInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpmiLanInfo`: IpmiLanInfo
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetIpmiLanInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIpmiLanInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeId** | **string** | Specifies the node id of the node for which lan info is requested. This parameter is incompatible with &#39;nodeIp&#39;. | 
 **nodeIp** | **string** | Specifies the IP Address of the node for which lan info is requested. This parameter is incompatible with &#39;nodeId&#39;. | 

### Return type

[**IpmiLanInfo**](IpmiLanInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpmiSdrInfo

> IpmiSdrInfo GetIpmiSdrInfo(ctx).NodeId(nodeId).NodeIp(nodeIp).Execute()

To get IPMI SDR Info



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
	nodeId := "nodeId_example" // string | Specifies the node id of the node for which sdr is requested. This parameter is incompatible with 'nodeIp'. (optional)
	nodeIp := "nodeIp_example" // string | Specifies the IP Address of the node for which sdr is requested. This parameter is incompatible with 'nodeId'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetIpmiSdrInfo(context.Background()).NodeId(nodeId).NodeIp(nodeIp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetIpmiSdrInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpmiSdrInfo`: IpmiSdrInfo
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetIpmiSdrInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIpmiSdrInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeId** | **string** | Specifies the node id of the node for which sdr is requested. This parameter is incompatible with &#39;nodeIp&#39;. | 
 **nodeIp** | **string** | Specifies the IP Address of the node for which sdr is requested. This parameter is incompatible with &#39;nodeId&#39;. | 

### Return type

[**IpmiSdrInfo**](IpmiSdrInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpmiSel

> IpmiSel GetIpmiSel(ctx).NodeId(nodeId).NodeIp(nodeIp).Verbose(verbose).Execute()

To get IPMI SEL



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
	nodeId := "nodeId_example" // string | Specifies the node id of the node for which sel is requested. This parameter is incompatible with 'nodeIp'. (optional)
	nodeIp := "nodeIp_example" // string | Specifies the IP Address of the node for which sel is requested. This parameter is incompatible with 'nodeId'. (optional)
	verbose := true // bool | Specifies the Verbosity of log produced by sel request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetIpmiSel(context.Background()).NodeId(nodeId).NodeIp(nodeIp).Verbose(verbose).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetIpmiSel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpmiSel`: IpmiSel
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetIpmiSel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIpmiSelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeId** | **string** | Specifies the node id of the node for which sel is requested. This parameter is incompatible with &#39;nodeIp&#39;. | 
 **nodeIp** | **string** | Specifies the IP Address of the node for which sel is requested. This parameter is incompatible with &#39;nodeId&#39;. | 
 **verbose** | **bool** | Specifies the Verbosity of log produced by sel request. | 

### Return type

[**IpmiSel**](IpmiSel.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpmiSelInfo

> IpmiSelInfo GetIpmiSelInfo(ctx).NodeId(nodeId).NodeIp(nodeIp).Execute()

To get IPMI SEL Info



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
	nodeId := "nodeId_example" // string | Specifies the node id of the node for which sel is requested. This parameter is incompatible with 'nodeIp'. (optional)
	nodeIp := "nodeIp_example" // string | Specifies the IP Address of the node for which sel is requested. This parameter is incompatible with 'nodeId'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetIpmiSelInfo(context.Background()).NodeId(nodeId).NodeIp(nodeIp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetIpmiSelInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpmiSelInfo`: IpmiSelInfo
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetIpmiSelInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIpmiSelInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeId** | **string** | Specifies the node id of the node for which sel is requested. This parameter is incompatible with &#39;nodeIp&#39;. | 
 **nodeIp** | **string** | Specifies the IP Address of the node for which sel is requested. This parameter is incompatible with &#39;nodeId&#39;. | 

### Return type

[**IpmiSelInfo**](IpmiSelInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpmiUsers

> IpmiUsers GetIpmiUsers(ctx).NodeId(nodeId).NodeIp(nodeIp).Execute()

To get IPMI User Info for node



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
	nodeId := "nodeId_example" // string | Specifies the node id of the node for which ipmi users info is requested. This parameter is incompatible with 'nodeIp'. (optional)
	nodeIp := "nodeIp_example" // string | Specifies the IP address of the node for which ipmi users info is requested. This parameter is incompatible with 'nodeId'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetIpmiUsers(context.Background()).NodeId(nodeId).NodeIp(nodeIp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetIpmiUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpmiUsers`: IpmiUsers
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetIpmiUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIpmiUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeId** | **string** | Specifies the node id of the node for which ipmi users info is requested. This parameter is incompatible with &#39;nodeIp&#39;. | 
 **nodeIp** | **string** | Specifies the IP address of the node for which ipmi users info is requested. This parameter is incompatible with &#39;nodeId&#39;. | 

### Return type

[**IpmiUsers**](IpmiUsers.md)

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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetIsDMaaSCluster(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetIsDMaaSCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIsDMaaSCluster`: DMaaSInfo
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetIsDMaaSCluster`: %v\n", resp)
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


## GetKubernetesInfraHealthStatus

> GetKubernetesStatusResponse GetKubernetesInfraHealthStatus(ctx).Execute()

Get Kubernetes Infra Health Status



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetKubernetesInfraHealthStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetKubernetesInfraHealthStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesInfraHealthStatus`: GetKubernetesStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetKubernetesInfraHealthStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesInfraHealthStatusRequest struct via the builder pattern


### Return type

[**GetKubernetesStatusResponse**](GetKubernetesStatusResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkInterfaces

> ClusterInterfaces GetNetworkInterfaces(ctx).NodeId(nodeId).Cache(cache).BondInterfaceOnly(bondInterfaceOnly).IfaceGroupAssignedOnly(ifaceGroupAssignedOnly).IncludeUplinkSwitchInfo(includeUplinkSwitchInfo).IncludeBondSlaveDetails(includeBondSlaveDetails).IncludeStats(includeStats).Execute()

Get list of interfaces



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
	nodeId := int64(789) // int64 | Node id, used to get interfaces on a particular node. (optional)
	cache := true // bool | Get interfaces information from cache. (optional) (default to false)
	bondInterfaceOnly := true // bool | Specifies if only show bond interface info. (optional) (default to false)
	ifaceGroupAssignedOnly := true // bool | Specifies if only show interface group assigned interface info. (optional) (default to false)
	includeUplinkSwitchInfo := true // bool | Specifies if include uplink switch info. (optional) (default to false)
	includeBondSlaveDetails := true // bool | Specifies if include bond secondary detailed info. (optional) (default to false)
	includeStats := true // bool | Specifies if include stats. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetNetworkInterfaces(context.Background()).NodeId(nodeId).Cache(cache).BondInterfaceOnly(bondInterfaceOnly).IfaceGroupAssignedOnly(ifaceGroupAssignedOnly).IncludeUplinkSwitchInfo(includeUplinkSwitchInfo).IncludeBondSlaveDetails(includeBondSlaveDetails).IncludeStats(includeStats).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetNetworkInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkInterfaces`: ClusterInterfaces
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetNetworkInterfaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeId** | **int64** | Node id, used to get interfaces on a particular node. | 
 **cache** | **bool** | Get interfaces information from cache. | [default to false]
 **bondInterfaceOnly** | **bool** | Specifies if only show bond interface info. | [default to false]
 **ifaceGroupAssignedOnly** | **bool** | Specifies if only show interface group assigned interface info. | [default to false]
 **includeUplinkSwitchInfo** | **bool** | Specifies if include uplink switch info. | [default to false]
 **includeBondSlaveDetails** | **bool** | Specifies if include bond secondary detailed info. | [default to false]
 **includeStats** | **bool** | Specifies if include stats. | [default to false]

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


## GetNodes

> []Node GetNodes(ctx).Ids(ids).IncludeMarkedForRemoval(includeMarkedForRemoval).IncludeOnlyUnassignedNodes(includeOnlyUnassignedNodes).ClusterPartitionIds(clusterPartitionIds).FetchStats(fetchStats).ShowSystemDisks(showSystemDisks).Execute()

List Nodes of the cluster.



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
	ids := []int64{int64(123)} // []int64 | \"List of IDs to be returned. If empty, all nodes are returned.\" (optional)
	includeMarkedForRemoval := true // bool | IncludeMarkedForRemoval is used to specify whether to include nodes marked for removal. (optional)
	includeOnlyUnassignedNodes := true // bool | IncludeOnlyUnassignedNodes will return nodes that are not yet assigned to any cluster partition. If this parameter is specified as true and ClusterPartitionIdList is also non-empty, then no nodes will be returned. (optional)
	clusterPartitionIds := []int64{int64(123)} // []int64 | ClusterPartitionIdList specifies the list of Ids used to filter the nodes by specified cluster partition. (optional)
	fetchStats := true // bool | FetchStats is used to specify whether to call Stats service to fetch the stats for the nodes. Stats not displayed by default (optional)
	showSystemDisks := true // bool | ShowSystemdisks is used to specify whether to display system disks for the nodes. Not populated by default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetNodes(context.Background()).Ids(ids).IncludeMarkedForRemoval(includeMarkedForRemoval).IncludeOnlyUnassignedNodes(includeOnlyUnassignedNodes).ClusterPartitionIds(clusterPartitionIds).FetchStats(fetchStats).ShowSystemDisks(showSystemDisks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodes`: []Node
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | \&quot;List of IDs to be returned. If empty, all nodes are returned.\&quot; | 
 **includeMarkedForRemoval** | **bool** | IncludeMarkedForRemoval is used to specify whether to include nodes marked for removal. | 
 **includeOnlyUnassignedNodes** | **bool** | IncludeOnlyUnassignedNodes will return nodes that are not yet assigned to any cluster partition. If this parameter is specified as true and ClusterPartitionIdList is also non-empty, then no nodes will be returned. | 
 **clusterPartitionIds** | **[]int64** | ClusterPartitionIdList specifies the list of Ids used to filter the nodes by specified cluster partition. | 
 **fetchStats** | **bool** | FetchStats is used to specify whether to call Stats service to fetch the stats for the nodes. Stats not displayed by default | 
 **showSystemDisks** | **bool** | ShowSystemdisks is used to specify whether to display system disks for the nodes. Not populated by default. | 

### Return type

[**[]Node**](Node.md)

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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := int64(789) // int64 | Specifies the id of rack.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetRackById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetRackById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRackById`: Rack
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetRackById`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetRacks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetRacks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRacks`: Racks
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetRacks`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	diskIds := []int64{int64(123)} // []int64 | Specifies a list of disk ids, only disks having these ids will be returned. (optional)
	nodeIds := []int64{int64(123)} // []int64 | Specifies a list of node ids, only disks in these nodes will be returned. (optional)
	tiers := []string{"Tiers_example"} // []string | Specifies a list of disk tiers, only disks with given tiers will be returned. (optional)
	mountPath := "mountPath_example" // string | This field is deprecated. Providing this queryparam will not have any impact. Please use fileSystem query param to filter instead. (optional)
	fileSystem := "fileSystem_example" // string | Specified file system name to search. only disks with file system name that partially matches the specified name will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetRemoteDisks(context.Background()).DiskIds(diskIds).NodeIds(nodeIds).Tiers(tiers).MountPath(mountPath).FileSystem(fileSystem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetRemoteDisks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteDisks`: RemoteDisks
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetRemoteDisks`: %v\n", resp)
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


## GetSMTPConfiguration

> SMTPConfiguration GetSMTPConfiguration(ctx).Execute()

Get SMTP configuration.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetSMTPConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetSMTPConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSMTPConfiguration`: SMTPConfiguration
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetSMTPConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSMTPConfigurationRequest struct via the builder pattern


### Return type

[**SMTPConfiguration**](SMTPConfiguration.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSWUpdateHistory

> []ClusterSWUpdateHistoryEvent GetSWUpdateHistory(ctx).IncludeNodeHistory(includeNodeHistory).Execute()

Get cluster software history



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
	includeNodeHistory := true // bool | Flag to specify whether to fetch data from current node or all the nodes. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetSWUpdateHistory(context.Background()).IncludeNodeHistory(includeNodeHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetSWUpdateHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSWUpdateHistory`: []ClusterSWUpdateHistoryEvent
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetSWUpdateHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSWUpdateHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeNodeHistory** | **bool** | Flag to specify whether to fetch data from current node or all the nodes.  | 

### Return type

[**[]ClusterSWUpdateHistoryEvent**](ClusterSWUpdateHistoryEvent.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceGflags

> []ServiceGflags GetServiceGflags(ctx).Gflags(gflags).ServiceName(serviceName).Execute()

Gets cluster gflags for a service.



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
	gflags := []string{"Inner_example"} // []string | \"Specifies a list of gflag names. If specified, only gflags matching the gflag name list will be returned.\" (optional)
	serviceName := "serviceName_example" // string | Specifies the service name. If specified, only gflags matching the service name will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetServiceGflags(context.Background()).Gflags(gflags).ServiceName(serviceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetServiceGflags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceGflags`: []ServiceGflags
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetServiceGflags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceGflagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gflags** | **[]string** | \&quot;Specifies a list of gflag names. If specified, only gflags matching the gflag name list will be returned.\&quot; | 
 **serviceName** | **string** | Specifies the service name. If specified, only gflags matching the service name will be returned. | 

### Return type

[**[]ServiceGflags**](ServiceGflags.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSoftwareComponents

> SoftwareComponents GetSoftwareComponents(ctx).Execute()

Get Software Components



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetSoftwareComponents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetSoftwareComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSoftwareComponents`: SoftwareComponents
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetSoftwareComponents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSoftwareComponentsRequest struct via the builder pattern


### Return type

[**SoftwareComponents**](SoftwareComponents.md)

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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetSupportChannelConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetSupportChannelConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportChannelConfig`: SupportChannelConfig
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetSupportChannelConfig`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := int64(789) // int64 | Specifies id of node to identify.
	body := *openapiclient.NewNodeIdentifyParams(false) // NodeIdentifyParams | Specifies the parameter to identify node.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.IdentifyNode(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.IdentifyNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IdentifyNode`: NodeIdentifyParams
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.IdentifyNode`: %v\n", resp)
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


## ImportCrlFile

> ImportCrlFile(ctx).FileName(fileName).Crlfile(crlfile).Execute()

Import Crl File



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
	fileName := "fileName_example" // string | 
	crlfile := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.ImportCrlFile(context.Background()).FileName(fileName).Crlfile(crlfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.ImportCrlFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportCrlFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileName** | **string** |  | 
 **crlfile** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: multipart/form-data
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	nodeId := int64(789) // int64 | Specifies node id of the node to get list of disks (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.ListDisks(context.Background()).NodeId(nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.ListDisks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDisks`: DisksList
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.ListDisks`: %v\n", resp)
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


## ListFreeNodes

> FreeNodes ListFreeNodes(ctx).Ips(ips).Execute()

List the free Cohesity Nodes present on a network.



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
	ips := []string{"Inner_example"} // []string | \"Specifies a list of ips of nodes among which free and compatible nodes to be returned\" (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.ListFreeNodes(context.Background()).Ips(ips).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.ListFreeNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFreeNodes`: FreeNodes
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.ListFreeNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFreeNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ips** | **[]string** | \&quot;Specifies a list of ips of nodes among which free and compatible nodes to be returned\&quot; | 

### Return type

[**FreeNodes**](FreeNodes.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHosts

> HostMappings ListHosts(ctx).Execute()

List Host Mappings



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.ListHosts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.ListHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHosts`: HostMappings
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.ListHosts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListHostsRequest struct via the builder pattern


### Return type

[**HostMappings**](HostMappings.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkBaseosUpgrade

> MarkBaseosUpgradeInfo MarkBaseosUpgrade(ctx).Body(body).Execute()

Sets/clears the BaseOS upgrade cluster operation.



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
	body := *openapiclient.NewMarkBaseosUpgradeInfo(false) // MarkBaseosUpgradeInfo | Param to whether set/clear BaseOS uprgade  operation.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.MarkBaseosUpgrade(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.MarkBaseosUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkBaseosUpgrade`: MarkBaseosUpgradeInfo
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.MarkBaseosUpgrade`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkBaseosUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MarkBaseosUpgradeInfo**](MarkBaseosUpgradeInfo.md) | Param to whether set/clear BaseOS uprgade  operation. | 

### Return type

[**MarkBaseosUpgradeInfo**](MarkBaseosUpgradeInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := int64(789) // int64 | Specifies unique id of the disk to mark for removal.
	body := *openapiclient.NewDiskRemovalParams(false) // DiskRemovalParams | Specifies parameters to mark/cancel disk removal.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.MarkDiskRemoval(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.MarkDiskRemoval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkDiskRemoval`: RemoveDisk
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.MarkDiskRemoval`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := int64(789) // int64 | Specifies id of node to cancel removal.
	body := *openapiclient.NewNodeRemovalParams(false) // NodeRemovalParams | Specifies parameters to initiate/cancel node removal .

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.MarkNodeRemoval(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.MarkNodeRemoval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkNodeRemoval`: RemoveNode
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.MarkNodeRemoval`: %v\n", resp)
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


## NodeInformation

> NodeInfo NodeInformation(ctx).ShowServicesVersionInfo(showServicesVersionInfo).OnlyCheckNodeReachability(onlyCheckNodeReachability).Execute()

Fetch Node General Information



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
	showServicesVersionInfo := true // bool | Specifies whether to show version info of the services running on the node. (optional)
	onlyCheckNodeReachability := true // bool | Specifies to show only node reachability details (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.NodeInformation(context.Background()).ShowServicesVersionInfo(showServicesVersionInfo).OnlyCheckNodeReachability(onlyCheckNodeReachability).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.NodeInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NodeInformation`: NodeInfo
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.NodeInformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNodeInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showServicesVersionInfo** | **bool** | Specifies whether to show version info of the services running on the node. | 
 **onlyCheckNodeReachability** | **bool** | Specifies to show only node reachability details | [default to false]

### Return type

[**NodeInfo**](NodeInfo.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NodeStatus

> NodeStatusResult NodeStatus(ctx).Execute()

Fetch Node status Information



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.NodeStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.NodeStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NodeStatus`: NodeStatusResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.NodeStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNodeStatusRequest struct via the builder pattern


### Return type

[**NodeStatusResult**](NodeStatusResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicKeyRequest

> PublicKeyResponse PublicKeyRequest(ctx).Body(body).Execute()

Get the SSH public key.



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
	body := *openapiclient.NewPublicKeyRequest("WorkflowType_example") // PublicKeyRequest | Specifies the parameters required to retrieve SSH public key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PublicKeyRequest(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PublicKeyRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicKeyRequest`: PublicKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PublicKeyRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublicKeyRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PublicKeyRequest**](PublicKeyRequest.md) | Specifies the parameters required to retrieve SSH public key | 

### Return type

[**PublicKeyResponse**](PublicKeyResponse.md)

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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := int64(789) // int64 | Specifies the id of the remote disk to remove.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.RemoveRemoteDisk(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.RemoveRemoteDisk``: %v\n", err)
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


## ResetIpmiBmc

> IpmiTextResponse ResetIpmiBmc(ctx).Body(body).Execute()

To reset IPMI BMC for given node



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
	body := *openapiclient.NewResetIpmiBmcParams() // ResetIpmiBmcParams | Specifies the parameters to reset ipmi bmc for given node.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.ResetIpmiBmc(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.ResetIpmiBmc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetIpmiBmc`: IpmiTextResponse
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.ResetIpmiBmc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetIpmiBmcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ResetIpmiBmcParams**](ResetIpmiBmcParams.md) | Specifies the parameters to reset ipmi bmc for given node. | 

### Return type

[**IpmiTextResponse**](IpmiTextResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewNodePowerOperation("Operation_example") // NodePowerOperation | Specifies the reboot or shutdown operation.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.SetNodePower(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.SetNodePower``: %v\n", err)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewClusterAMQPTargetConfig() // ClusterAMQPTargetConfig | Specifies the parameters to update cluster AMQP target config.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.UpdateAMQPTargetConfig(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.UpdateAMQPTargetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAMQPTargetConfig`: ClusterAMQPTargetConfig
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.UpdateAMQPTargetConfig`: %v\n", resp)
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


## UpdateAirgapConfig

> AirgapConfig UpdateAirgapConfig(ctx).Body(body).Execute()

Update Airgap config



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
	body := *openapiclient.NewAirgapConfig() // AirgapConfig | Specifies the parameters to update airgap config.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.UpdateAirgapConfig(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.UpdateAirgapConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAirgapConfig`: AirgapConfig
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.UpdateAirgapConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAirgapConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AirgapConfig**](AirgapConfig.md) | Specifies the parameters to update airgap config. | 

### Return type

[**AirgapConfig**](AirgapConfig.md)

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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := int64(789) // int64 | Specifies the id of chassis.
	body := *openapiclient.NewChassis() // Chassis | Specifies the parameters to update chassis. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.UpdateChassisById(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.UpdateChassisById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateChassisById`: Chassis
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.UpdateChassisById`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewCluster() // Cluster | Specifies the parameters to update cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.UpdateCluster(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.UpdateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCluster`: Cluster
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.UpdateCluster`: %v\n", resp)
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


## UpdateClusterSnapshotPolicy

> ClusterSnapshotPolicy UpdateClusterSnapshotPolicy(ctx).Body(body).Execute()

Update cluster snapshot policy.



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
	body := *openapiclient.NewUpdateSnapshotPolicyParams() // UpdateSnapshotPolicyParams | Specifies the parameters to update cluster snapshot policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.UpdateClusterSnapshotPolicy(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.UpdateClusterSnapshotPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClusterSnapshotPolicy`: ClusterSnapshotPolicy
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.UpdateClusterSnapshotPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterSnapshotPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateSnapshotPolicyParams**](UpdateSnapshotPolicyParams.md) | Specifies the parameters to update cluster snapshot policy. | 

### Return type

[**ClusterSnapshotPolicy**](ClusterSnapshotPolicy.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClusterSoftware

> []ClusterOperationTypeAndId UpdateClusterSoftware(ctx).Body(body).Execute()

Update cluster software



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
	body := *openapiclient.NewClusterSWUpdateParams("OperationType_example") // ClusterSWUpdateParams | The parameters to update the software on the cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.UpdateClusterSoftware(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.UpdateClusterSoftware``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClusterSoftware`: []ClusterOperationTypeAndId
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.UpdateClusterSoftware`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterSoftwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClusterSWUpdateParams**](ClusterSWUpdateParams.md) | The parameters to update the software on the cluster. | 

### Return type

[**[]ClusterOperationTypeAndId**](ClusterOperationTypeAndId.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeatureFlag

> []FeatureFlag UpdateFeatureFlag(ctx).Body(body).Execute()

Update feature flag override status.



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
	body := *openapiclient.NewUpdateFeatureFlagParams() // UpdateFeatureFlagParams | Param for feature flag override request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.UpdateFeatureFlag(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.UpdateFeatureFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFeatureFlag`: []FeatureFlag
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.UpdateFeatureFlag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeatureFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateFeatureFlagParams**](UpdateFeatureFlagParams.md) | Param for feature flag override request. | 

### Return type

[**[]FeatureFlag**](FeatureFlag.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHosts

> HostMappings UpdateHosts(ctx).Body(body).Execute()

Update Host Mappings



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
	body := []openapiclient.HostEntry{*openapiclient.NewHostEntry([]string{"DomainNames_example"}, "Ip_example")} // []HostEntry | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.UpdateHosts(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.UpdateHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHosts`: HostMappings
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.UpdateHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]HostEntry**](HostEntry.md) |  | 

### Return type

[**HostMappings**](HostMappings.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIpmiUser

> IpmiTextResponse UpdateIpmiUser(ctx).Body(body).Execute()

To update IPMI User Info for node



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
	body := *openapiclient.NewUpdateIpmiUser() // UpdateIpmiUser | Specifies the parameters to add an ipmi user to node.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.UpdateIpmiUser(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.UpdateIpmiUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIpmiUser`: IpmiTextResponse
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.UpdateIpmiUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIpmiUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateIpmiUser**](UpdateIpmiUser.md) | Specifies the parameters to add an ipmi user to node. | 

### Return type

[**IpmiTextResponse**](IpmiTextResponse.md)

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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewDMaaSInfo(false) // DMaaSInfo | Param to update whether the cluster is a DMaaS cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.UpdateIsDMaaSCluster(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.UpdateIsDMaaSCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIsDMaaSCluster`: DMaaSInfo
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.UpdateIsDMaaSCluster`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := int64(789) // int64 | Specifies the id of rack.
	body := *openapiclient.NewRack() // Rack | Specifies the parameters to update rack. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.UpdateRackById(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.UpdateRackById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRackById`: Rack
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.UpdateRackById`: %v\n", resp)
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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewRacks() // Racks | Specifies the parameters to update racks.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.UpdateRacks(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.UpdateRacks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRacks`: Racks
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.UpdateRacks`: %v\n", resp)
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


## UpdateSMTPConfiguration

> SMTPConfiguration UpdateSMTPConfiguration(ctx).Body(body).Execute()

Update SMTP configuration.



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
	body := *openapiclient.NewUpdateSMTPParams("Hostname_example", int32(123)) // UpdateSMTPParams | Specifies the parameters to update cluster SMTP configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.UpdateSMTPConfiguration(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.UpdateSMTPConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSMTPConfiguration`: SMTPConfiguration
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.UpdateSMTPConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSMTPConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateSMTPParams**](UpdateSMTPParams.md) | Specifies the parameters to update cluster SMTP configuration. | 

### Return type

[**SMTPConfiguration**](SMTPConfiguration.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceGflags

> []ServiceGflags UpdateServiceGflags(ctx).Body(body).Execute()

Update the gflags



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
	body := *openapiclient.NewUpdateGflagParameters() // UpdateGflagParameters | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.UpdateServiceGflags(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.UpdateServiceGflags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServiceGflags`: []ServiceGflags
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.UpdateServiceGflags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceGflagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateGflagParameters**](UpdateGflagParameters.md) |  | 

### Return type

[**[]ServiceGflags**](ServiceGflags.md)

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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewSupportChannelConfig(NullableInt64(123), false) // SupportChannelConfig | Specifies the support channel configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.UpdateSupportChannelConfig(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.UpdateSupportChannelConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSupportChannelConfig`: SupportChannelConfig
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.UpdateSupportChannelConfig`: %v\n", resp)
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


## UpgradeCheckGetResults

> UpgradeChecksResults UpgradeCheckGetResults(ctx, testRunInstanceId).Execute()

Get upgrade checks results.



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
	testRunInstanceId := int64(789) // int64 | Specifies test run instance for which to fetch results

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.UpgradeCheckGetResults(context.Background(), testRunInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.UpgradeCheckGetResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeCheckGetResults`: UpgradeChecksResults
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.UpgradeCheckGetResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testRunInstanceId** | **int64** | Specifies test run instance for which to fetch results | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeCheckGetResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UpgradeChecksResults**](UpgradeChecksResults.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeCheckRunTests

> UpgradeCheckRunTestsResult UpgradeCheckRunTests(ctx).Body(body).Execute()

Run upgrade checks on cluster.



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
	body := *openapiclient.NewUpgradeCheckRunTestsRequest() // UpgradeCheckRunTestsRequest | Run upgrade checks on cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.UpgradeCheckRunTests(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.UpgradeCheckRunTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeCheckRunTests`: UpgradeCheckRunTestsResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.UpgradeCheckRunTests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeCheckRunTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpgradeCheckRunTestsRequest**](UpgradeCheckRunTestsRequest.md) | Run upgrade checks on cluster. | 

### Return type

[**UpgradeCheckRunTestsResult**](UpgradeCheckRunTestsResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFilePackage

> UploadFilePackage(ctx).PackageFile(packageFile).PackageType(packageType).Execute()

Upload package by files



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
	packageFile := os.NewFile(1234, "some_file") // *os.File | Binary content of the file.
	packageType := "packageType_example" // string | Package Type. (optional) (default to "Upgrade")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.UploadFilePackage(context.Background()).PackageFile(packageFile).PackageType(packageType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.UploadFilePackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadFilePackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packageFile** | ***os.File** | Binary content of the file. | 
 **packageType** | **string** | Package Type. | [default to &quot;Upgrade&quot;]

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateSMTPConfiguration

> ValidateSMTPConfiguration(ctx).Body(body).Execute()

Validate SMTP configuration.



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
	body := *openapiclient.NewTestSMTPConfig("Email_example") // TestSMTPConfig | Specifies the request parameters to validate SMTP configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.ValidateSMTPConfiguration(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.ValidateSMTPConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateSMTPConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TestSMTPConfig**](TestSMTPConfig.md) | Specifies the request parameters to validate SMTP configuration. | 

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

