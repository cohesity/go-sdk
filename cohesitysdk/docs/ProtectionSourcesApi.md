# \ProtectionSourcesApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadCftFile**](ProtectionSourcesApi.md#DownloadCftFile) | **Get** /public/protectionSources/downloadCftFile | Download the CFT.
[**DownloadPhysicalAgent**](ProtectionSourcesApi.md#DownloadPhysicalAgent) | **Get** /public/physicalAgents/download | Download the physical agent for a host type.
[**GetProtectionSourcesObjectById**](ProtectionSourcesApi.md#GetProtectionSourcesObjectById) | **Get** /public/protectionSources/objects/{id} | Get details about a single Protection Source Object.
[**GetProtectionSourcesObjects**](ProtectionSourcesApi.md#GetProtectionSourcesObjects) | **Get** /public/protectionSources/objects | List details about Protection Source objects.
[**ListApplicationServers**](ProtectionSourcesApi.md#ListApplicationServers) | **Get** /public/protectionSources/applicationServers | Returns the registered Application Servers and their Object subtrees.
[**ListDataStoreInformation**](ProtectionSourcesApi.md#ListDataStoreInformation) | **Get** /public/protectionSources/datastores | Returns the datastore information in VMware environment.
[**ListExchangeDagHosts**](ProtectionSourcesApi.md#ListExchangeDagHosts) | **Get** /public/protectionSources/exchangeDagHosts | 
[**ListProtectedObjects**](ProtectionSourcesApi.md#ListProtectedObjects) | **Get** /public/protectionSources/protectedObjects | Returns the list of protected Objects in a registered Protection Source.
[**ListProtectionSources**](ProtectionSourcesApi.md#ListProtectionSources) | **Get** /public/protectionSources | Returns the registered Protection Sources and their Object subtrees.
[**ListProtectionSourcesRegistrationInfo**](ProtectionSourcesApi.md#ListProtectionSourcesRegistrationInfo) | **Get** /public/protectionSources/registrationInfo | 
[**ListProtectionSourcesRootNodes**](ProtectionSourcesApi.md#ListProtectionSourcesRootNodes) | **Get** /public/protectionSources/rootNodes | Returns the top level (root) Protection Sources with registration information.
[**ListSqlAagHostsAndDatabases**](ProtectionSourcesApi.md#ListSqlAagHostsAndDatabases) | **Get** /public/protectionSources/sqlAagHostsAndDatabases | Returns the registered Protection Sources and their Object subtrees.
[**ListVirtualMachines**](ProtectionSourcesApi.md#ListVirtualMachines) | **Get** /public/protectionSources/virtualMachines | Returns the Virtual Machines in a vCenter Server.
[**RefreshProtectionSourceById**](ProtectionSourcesApi.md#RefreshProtectionSourceById) | **Post** /public/protectionSources/refresh/{id} | Refresh the Object hierarchy of the Protection Sources tree.
[**RegisterApplicationServers**](ProtectionSourcesApi.md#RegisterApplicationServers) | **Post** /public/protectionSources/applicationServers | Register a Protection Source as running one or more Application Servers like Microsoft SQL servers or Microsoft Exchange servers.
[**RegisterProtectionSource**](ProtectionSourcesApi.md#RegisterProtectionSource) | **Post** /public/protectionSources/register | Register a Protection Source.
[**RunDiagnostics**](ProtectionSourcesApi.md#RunDiagnostics) | **Post** /public/protectionSources/diagnostics/{id} | Collect diagnostics of the protection source for a host type.
[**UnregisterApplicationServers**](ProtectionSourcesApi.md#UnregisterApplicationServers) | **Delete** /public/protectionSources/applicationServers/{id} | Unregister Application Servers like Microsoft SQL servers or Microsoft Exchange servers running on a Protection Source.
[**UnregisterProtectionSource**](ProtectionSourcesApi.md#UnregisterProtectionSource) | **Delete** /public/protectionSources/{id} | Unregister a previously registered Protection Source.
[**UpdateApplicationServers**](ProtectionSourcesApi.md#UpdateApplicationServers) | **Put** /public/protectionSources/applicationServers | Modifies the registration parameters of Application Servers in a Protection Source.
[**UpdateProtectionSource**](ProtectionSourcesApi.md#UpdateProtectionSource) | **Patch** /public/protectionSources/{id} | Update a previously registered Protection Source with new details.
[**UpgradePhysicalAgents**](ProtectionSourcesApi.md#UpgradePhysicalAgents) | **Post** /public/physicalAgents/upgrade | Initiate a request to upgrade Cohesity agents on one or more Physical Servers registered on the Cohesity Cluster.



## DownloadCftFile

> DownloadCftResponse DownloadCftFile(ctx).Body(body).Execute()

Download the CFT.

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
    body := *openapiclient.NewDownloadCftParams() // DownloadCftParams | Specifies the request to download CFT. (optional)

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

    request := cohesitysdk.ApiDownloadCftFileRequest{
        Body: &body
    }

    resp, r, err := api_client.ProtectionSourcesApi.DownloadCftFile(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.DownloadCftFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadCftFile`: DownloadCftResponse
    fmt.Fprintf(os.Stdout, "Response from `ProtectionSourcesApi.DownloadCftFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadCftFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DownloadCftParams**](DownloadCftParams.md) | Specifies the request to download CFT. | 

### Return type

[**DownloadCftResponse**](DownloadCftResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadPhysicalAgent

> []int32 DownloadPhysicalAgent(ctx).HostType(hostType).PkgType(pkgType).AgentType(agentType).Execute()

Download the physical agent for a host type.



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
    hostType := "hostType_example" // string | Specifies the host type for which user wants to download the physical agent. 'kLinux' indicates the Linux operating system. 'kWindows' indicates the Microsoft Windows operating system. 'kAix' indicates the IBM AIX operating system. 'kSolaris' indicates the Oracle Solaris operating system. 'kSapHana' indicates the Sap Hana database system developed by SAP SE. 'kOther' indicates the other types of operating system.
    pkgType := "pkgType_example" // string | Specifies the Linux installer package type applicable only to Linux OS and the value can be any of (\"kScript\",\"kRPM\", \"kSuseRPM\", \"kDEB\") 'kScript' indicates a script based agent installer. 'kRPM' indicates a RPM agent installer. 'kSuseRPM' indicates a Open Suse RPM installer. 'kDEB' indicates a Debian agent installer. (optional)
    agentType := "agentType_example" // string | Specifies agent type. Can be \"kGo\" for go agent and \"kJava\" for java agent and \"kCpp\" for c++ agent. 'kCpp' indicates a c++ agent. 'kJava' indicates a java agent. 'kGo' indicates a go agent. (optional)

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

    request := cohesitysdk.ApiDownloadPhysicalAgentRequest{
        HostType: &hostType
        PkgType: &pkgType
        AgentType: &agentType
    }

    resp, r, err := api_client.ProtectionSourcesApi.DownloadPhysicalAgent(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.DownloadPhysicalAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadPhysicalAgent`: []int32
    fmt.Fprintf(os.Stdout, "Response from `ProtectionSourcesApi.DownloadPhysicalAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadPhysicalAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostType** | **string** | Specifies the host type for which user wants to download the physical agent. &#39;kLinux&#39; indicates the Linux operating system. &#39;kWindows&#39; indicates the Microsoft Windows operating system. &#39;kAix&#39; indicates the IBM AIX operating system. &#39;kSolaris&#39; indicates the Oracle Solaris operating system. &#39;kSapHana&#39; indicates the Sap Hana database system developed by SAP SE. &#39;kOther&#39; indicates the other types of operating system. | 
 **pkgType** | **string** | Specifies the Linux installer package type applicable only to Linux OS and the value can be any of (\&quot;kScript\&quot;,\&quot;kRPM\&quot;, \&quot;kSuseRPM\&quot;, \&quot;kDEB\&quot;) &#39;kScript&#39; indicates a script based agent installer. &#39;kRPM&#39; indicates a RPM agent installer. &#39;kSuseRPM&#39; indicates a Open Suse RPM installer. &#39;kDEB&#39; indicates a Debian agent installer. | 
 **agentType** | **string** | Specifies agent type. Can be \&quot;kGo\&quot; for go agent and \&quot;kJava\&quot; for java agent and \&quot;kCpp\&quot; for c++ agent. &#39;kCpp&#39; indicates a c++ agent. &#39;kJava&#39; indicates a java agent. &#39;kGo&#39; indicates a go agent. | 

### Return type

**[]int32**

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionSourcesObjectById

> ProtectionSource GetProtectionSourcesObjectById(ctx, id).Execute()

Get details about a single Protection Source Object.



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
    id := int64(789) // int64 | Specifies a unique id of the Protection Source to return.

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

    request := cohesitysdk.ApiGetProtectionSourcesObjectByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.ProtectionSourcesApi.GetProtectionSourcesObjectById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.GetProtectionSourcesObjectById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionSourcesObjectById`: ProtectionSource
    fmt.Fprintf(os.Stdout, "Response from `ProtectionSourcesApi.GetProtectionSourcesObjectById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the Protection Source to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionSourcesObjectByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProtectionSource**](ProtectionSource.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionSourcesObjects

> []ProtectionSource GetProtectionSourcesObjects(ctx).ObjectIds(objectIds).Execute()

List details about Protection Source objects.



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
    objectIds := []int64{int64(123)} // []int64 | Specifies the ids of the Protection Source objects to return. (optional)

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

    request := cohesitysdk.ApiGetProtectionSourcesObjectsRequest{
        ObjectIds: &objectIds
    }

    resp, r, err := api_client.ProtectionSourcesApi.GetProtectionSourcesObjects(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.GetProtectionSourcesObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionSourcesObjects`: []ProtectionSource
    fmt.Fprintf(os.Stdout, "Response from `ProtectionSourcesApi.GetProtectionSourcesObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionSourcesObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectIds** | **[]int64** | Specifies the ids of the Protection Source objects to return. | 

### Return type

[**[]ProtectionSource**](ProtectionSource.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationServers

> []RegisteredApplicationServer ListApplicationServers(ctx).ProtectionSourcesRootNodeId(protectionSourcesRootNodeId).Environment(environment).ProtectionSourceId(protectionSourceId).Application(application).Execute()

Returns the registered Application Servers and their Object subtrees.



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
    protectionSourcesRootNodeId := int64(789) // int64 | Specifies the Protection Source Id of the root node of a Protection Sources tree. A root node represents a registered Source on the Cohesity Cluster, such as a vCenter Server. (optional)
    environment := "environment_example" // string | Specifies the environment such as 'kPhysical' or 'kVMware' of the Protection Source tree. overrideDescription: true Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. 'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment. 'kSQL' indicates the SQL Protection Source environment. 'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment. 'Nimble' indicates the Nimble Storage Protection Source environment. 'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment. 'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhsicalFiles' indicates the Physical Files Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment. (optional)
    protectionSourceId := int64(789) // int64 | Specifies the Protection Source Id of the 'kPhysical' or 'kVMware' entity in the Protection Source tree hosting the applications. (optional)
    application := "application_example" // string | Specifies the application such as 'kSQL', 'kExchange' running on the Protection Source. overrideDescription: true Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. 'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment. 'kSQL' indicates the SQL Protection Source environment. 'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment. 'Nimble' indicates the Nimble Storage Protection Source environment. 'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment. 'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhsicalFiles' indicates the Physical Files Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment. (optional)

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

    request := cohesitysdk.ApiListApplicationServersRequest{
        ProtectionSourcesRootNodeId: &protectionSourcesRootNodeId
        Environment: &environment
        ProtectionSourceId: &protectionSourceId
        Application: &application
    }

    resp, r, err := api_client.ProtectionSourcesApi.ListApplicationServers(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.ListApplicationServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationServers`: []RegisteredApplicationServer
    fmt.Fprintf(os.Stdout, "Response from `ProtectionSourcesApi.ListApplicationServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **protectionSourcesRootNodeId** | **int64** | Specifies the Protection Source Id of the root node of a Protection Sources tree. A root node represents a registered Source on the Cohesity Cluster, such as a vCenter Server. | 
 **environment** | **string** | Specifies the environment such as &#39;kPhysical&#39; or &#39;kVMware&#39; of the Protection Source tree. overrideDescription: true Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | 
 **protectionSourceId** | **int64** | Specifies the Protection Source Id of the &#39;kPhysical&#39; or &#39;kVMware&#39; entity in the Protection Source tree hosting the applications. | 
 **application** | **string** | Specifies the application such as &#39;kSQL&#39;, &#39;kExchange&#39; running on the Protection Source. overrideDescription: true Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | 

### Return type

[**[]RegisteredApplicationServer**](RegisteredApplicationServer.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataStoreInformation

> []ProtectionSource ListDataStoreInformation(ctx).SourceId(sourceId).Execute()

Returns the datastore information in VMware environment.

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
    sourceId := int64(789) // int64 | Specifies the id of the virtual machine in vmware environment.

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

    request := cohesitysdk.ApiListDataStoreInformationRequest{
        SourceId: &sourceId
    }

    resp, r, err := api_client.ProtectionSourcesApi.ListDataStoreInformation(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.ListDataStoreInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDataStoreInformation`: []ProtectionSource
    fmt.Fprintf(os.Stdout, "Response from `ProtectionSourcesApi.ListDataStoreInformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDataStoreInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceId** | **int64** | Specifies the id of the virtual machine in vmware environment. | 

### Return type

[**[]ProtectionSource**](ProtectionSource.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExchangeDagHosts

> ExchangeDagHostsResponse ListExchangeDagHosts(ctx).Endpoint(endpoint).ProtectionSourceId(protectionSourceId).Execute()





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
    endpoint := "endpoint_example" // string | Specifies the endpoint of Exchange DAG or a host which is member of Exchange DAG or a standalone exchange server. (optional)
    protectionSourceId := int64(789) // int64 | Specifies the Protection Source Id of the Exchange DAG source. (optional)

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

    request := cohesitysdk.ApiListExchangeDagHostsRequest{
        Endpoint: &endpoint
        ProtectionSourceId: &protectionSourceId
    }

    resp, r, err := api_client.ProtectionSourcesApi.ListExchangeDagHosts(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.ListExchangeDagHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExchangeDagHosts`: ExchangeDagHostsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProtectionSourcesApi.ListExchangeDagHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExchangeDagHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpoint** | **string** | Specifies the endpoint of Exchange DAG or a host which is member of Exchange DAG or a standalone exchange server. | 
 **protectionSourceId** | **int64** | Specifies the Protection Source Id of the Exchange DAG source. | 

### Return type

[**ExchangeDagHostsResponse**](ExchangeDagHostsResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProtectedObjects

> []ProtectedVmInfo ListProtectedObjects(ctx).Environment(environment).Id(id).AllUnderHierarchy(allUnderHierarchy).IncludeRpoSnapshots(includeRpoSnapshots).Execute()

Returns the list of protected Objects in a registered Protection Source.

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
    environment := "environment_example" // string | Specifies the environment type of the registered Protection Source such as 'kVMware', 'kSQL', 'kView' 'kPhysical', 'kPuppeteer', 'kPure', 'kNetapp', 'kGenericNas', 'kHyperV', 'kAcropolis', or 'kAzure'. For example, set this parameter to 'kVMware' if the registered Protection Source is of 'kVMware' environment type. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. 'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment. 'kSQL' indicates the SQL Protection Source environment. 'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment. 'Nimble' indicates the Nimble Storage Protection Source environment. 'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment. 'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhsicalFiles' indicates the Physical Files Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment.
    id := int64(789) // int64 | Specifies the Id of a registered Protection Source of the type given in environment.
    allUnderHierarchy := true // bool | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)
    includeRpoSnapshots := true // bool | If true, then the Protected Objects protected by RPO policies will also be returned. (optional)

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

    request := cohesitysdk.ApiListProtectedObjectsRequest{
        Environment: &environment
        Id: &id
        AllUnderHierarchy: &allUnderHierarchy
        IncludeRpoSnapshots: &includeRpoSnapshots
    }

    resp, r, err := api_client.ProtectionSourcesApi.ListProtectedObjects(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.ListProtectedObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProtectedObjects`: []ProtectedVmInfo
    fmt.Fprintf(os.Stdout, "Response from `ProtectionSourcesApi.ListProtectedObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProtectedObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Specifies the environment type of the registered Protection Source such as &#39;kVMware&#39;, &#39;kSQL&#39;, &#39;kView&#39; &#39;kPhysical&#39;, &#39;kPuppeteer&#39;, &#39;kPure&#39;, &#39;kNetapp&#39;, &#39;kGenericNas&#39;, &#39;kHyperV&#39;, &#39;kAcropolis&#39;, or &#39;kAzure&#39;. For example, set this parameter to &#39;kVMware&#39; if the registered Protection Source is of &#39;kVMware&#39; environment type. Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | 
 **id** | **int64** | Specifies the Id of a registered Protection Source of the type given in environment. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 
 **includeRpoSnapshots** | **bool** | If true, then the Protected Objects protected by RPO policies will also be returned. | 

### Return type

[**[]ProtectedVmInfo**](ProtectedVmInfo.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProtectionSources

> []ProtectionSourceNode ListProtectionSources(ctx).AfterCursorEntityId(afterCursorEntityId).BeforeCursorEntityId(beforeCursorEntityId).NodeId(nodeId).PageSize(pageSize).Id(id).NumLevels(numLevels).ExcludeTypes(excludeTypes).ExcludeOffice365Types(excludeOffice365Types).ExcludeAwsTypes(excludeAwsTypes).IncludeDatastores(includeDatastores).IncludeNetworks(includeNetworks).IncludeVMFolders(includeVMFolders).IncludeSystemVApps(includeSystemVApps).Environments(environments).Environment(environment).IncludeEntityPermissionInfo(includeEntityPermissionInfo).Sids(sids).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).Execute()

Returns the registered Protection Sources and their Object subtrees.



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
    afterCursorEntityId := int64(789) // int64 | Specifies the entity id starting from which the items are to be returned. (optional)
    beforeCursorEntityId := int64(789) // int64 | Specifies the entity id upto which the items are to be returned. (optional)
    nodeId := int64(789) // int64 | Specifies the entity id for the Node at any level within the Source entity hierarchy whose children are to be paginated. (optional)
    pageSize := int64(789) // int64 | Specifies the maximum number of entities to be returned within the page. (optional)
    id := int64(789) // int64 | Return the Object subtree for the passed in Protection Source id. (optional)
    numLevels := int32(56) // int32 | Specifies the expected number of levels from the root node to be returned in the entity hierarchy response. (optional)
    excludeTypes := []string{"ExcludeTypes_example"} // []string | Filter out the Object types (and their subtrees) that match the passed in types such as 'kVCenter', 'kFolder', 'kDatacenter', 'kComputeResource', 'kResourcePool', 'kDatastore', 'kHostSystem', 'kVirtualMachine', etc. For example, set this parameter to 'kResourcePool' to exclude Resource Pool Objects from being returned. (optional)
    excludeOffice365Types := []string{"ExcludeOffice365Types_example"} // []string | Specifies the Object types to be filtered out for Office 365 that match the passed in types such as 'kDomain', 'kOutlook', 'kMailbox', etc. For example, set this parameter to 'kMailbox' to exclude Mailbox Objects from being returned. (optional)
    excludeAwsTypes := []string{"ExcludeAwsTypes_example"} // []string | Specifies the Object types to be filtered out for AWS that match the passed in types such as 'kEC2Instance', 'kRDSInstance' etc. For example, set this parameter to 'kEC2Instance' to exclude ec2 instance from being returned. (optional)
    includeDatastores := true // bool | Set this parameter to true to also return kDatastore object types found in the Source in addition to their Object subtrees. By default, datastores are not returned. (optional)
    includeNetworks := true // bool | Set this parameter to true to also return kNetwork object types found in the Source in addition to their Object subtrees. By default, network objects are not returned. (optional)
    includeVMFolders := true // bool | Set this parameter to true to also return kVMFolder object types found in the Source in addition to their Object subtrees. By default, VM folder objects are not returned. (optional)
    includeSystemVApps := true // bool | Set this parameter to true to also return system VApp object types found in the Source in addition to their Object subtrees. By default, VM folder objects are not returned. (optional)
    environments := []string{"Environments_example"} // []string | Return only Protection Sources that match the passed in environment type such as 'kVMware', 'kSQL', 'kView' 'kPhysical', 'kPuppeteer', 'kPure', 'kNetapp', 'kGenericNas', 'kHyperV', 'kAcropolis', or 'kAzure'. For example, set this parameter to 'kVMware' to only return the Sources (and their Object subtrees) found in the 'kVMware' (VMware vCenter Server) environment.  NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. (optional)
    environment := "environment_example" // string | This field is deprecated. Use environments instead. deprecated: true (optional)
    includeEntityPermissionInfo := true // bool | If specified, then a list of entites with permissions assigned to them are returned. (optional)
    sids := []string{"Inner_example"} // []string | Filter the object subtree for the sids given in the list. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    allUnderHierarchy := true // bool | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)

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

    request := cohesitysdk.ApiListProtectionSourcesRequest{
        AfterCursorEntityId: &afterCursorEntityId
        BeforeCursorEntityId: &beforeCursorEntityId
        NodeId: &nodeId
        PageSize: &pageSize
        Id: &id
        NumLevels: &numLevels
        ExcludeTypes: &excludeTypes
        ExcludeOffice365Types: &excludeOffice365Types
        ExcludeAwsTypes: &excludeAwsTypes
        IncludeDatastores: &includeDatastores
        IncludeNetworks: &includeNetworks
        IncludeVMFolders: &includeVMFolders
        IncludeSystemVApps: &includeSystemVApps
        Environments: &environments
        Environment: &environment
        IncludeEntityPermissionInfo: &includeEntityPermissionInfo
        Sids: &sids
        TenantIds: &tenantIds
        AllUnderHierarchy: &allUnderHierarchy
    }

    resp, r, err := api_client.ProtectionSourcesApi.ListProtectionSources(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.ListProtectionSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProtectionSources`: []ProtectionSourceNode
    fmt.Fprintf(os.Stdout, "Response from `ProtectionSourcesApi.ListProtectionSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProtectionSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afterCursorEntityId** | **int64** | Specifies the entity id starting from which the items are to be returned. | 
 **beforeCursorEntityId** | **int64** | Specifies the entity id upto which the items are to be returned. | 
 **nodeId** | **int64** | Specifies the entity id for the Node at any level within the Source entity hierarchy whose children are to be paginated. | 
 **pageSize** | **int64** | Specifies the maximum number of entities to be returned within the page. | 
 **id** | **int64** | Return the Object subtree for the passed in Protection Source id. | 
 **numLevels** | **int32** | Specifies the expected number of levels from the root node to be returned in the entity hierarchy response. | 
 **excludeTypes** | **[]string** | Filter out the Object types (and their subtrees) that match the passed in types such as &#39;kVCenter&#39;, &#39;kFolder&#39;, &#39;kDatacenter&#39;, &#39;kComputeResource&#39;, &#39;kResourcePool&#39;, &#39;kDatastore&#39;, &#39;kHostSystem&#39;, &#39;kVirtualMachine&#39;, etc. For example, set this parameter to &#39;kResourcePool&#39; to exclude Resource Pool Objects from being returned. | 
 **excludeOffice365Types** | **[]string** | Specifies the Object types to be filtered out for Office 365 that match the passed in types such as &#39;kDomain&#39;, &#39;kOutlook&#39;, &#39;kMailbox&#39;, etc. For example, set this parameter to &#39;kMailbox&#39; to exclude Mailbox Objects from being returned. | 
 **excludeAwsTypes** | **[]string** | Specifies the Object types to be filtered out for AWS that match the passed in types such as &#39;kEC2Instance&#39;, &#39;kRDSInstance&#39; etc. For example, set this parameter to &#39;kEC2Instance&#39; to exclude ec2 instance from being returned. | 
 **includeDatastores** | **bool** | Set this parameter to true to also return kDatastore object types found in the Source in addition to their Object subtrees. By default, datastores are not returned. | 
 **includeNetworks** | **bool** | Set this parameter to true to also return kNetwork object types found in the Source in addition to their Object subtrees. By default, network objects are not returned. | 
 **includeVMFolders** | **bool** | Set this parameter to true to also return kVMFolder object types found in the Source in addition to their Object subtrees. By default, VM folder objects are not returned. | 
 **includeSystemVApps** | **bool** | Set this parameter to true to also return system VApp object types found in the Source in addition to their Object subtrees. By default, VM folder objects are not returned. | 
 **environments** | **[]string** | Return only Protection Sources that match the passed in environment type such as &#39;kVMware&#39;, &#39;kSQL&#39;, &#39;kView&#39; &#39;kPhysical&#39;, &#39;kPuppeteer&#39;, &#39;kPure&#39;, &#39;kNetapp&#39;, &#39;kGenericNas&#39;, &#39;kHyperV&#39;, &#39;kAcropolis&#39;, or &#39;kAzure&#39;. For example, set this parameter to &#39;kVMware&#39; to only return the Sources (and their Object subtrees) found in the &#39;kVMware&#39; (VMware vCenter Server) environment.  NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. | 
 **environment** | **string** | This field is deprecated. Use environments instead. deprecated: true | 
 **includeEntityPermissionInfo** | **bool** | If specified, then a list of entites with permissions assigned to them are returned. | 
 **sids** | **[]string** | Filter the object subtree for the sids given in the list. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 

### Return type

[**[]ProtectionSourceNode**](ProtectionSourceNode.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProtectionSourcesRegistrationInfo

> GetRegistrationInfoResponse ListProtectionSourcesRegistrationInfo(ctx).Environments(environments).Ids(ids).IncludeEntityPermissionInfo(includeEntityPermissionInfo).Sids(sids).IncludeApplicationsTreeInfo(includeApplicationsTreeInfo).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).Execute()





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
    environments := []string{"Environments_example"} // []string | Return only Protection Sources that match the passed in environment type such as 'kVMware', 'kSQL', 'kView' 'kPhysical', 'kPuppeteer', 'kPure', 'kNetapp', 'kGenericNas', 'kHyperV', 'kAcropolis', or 'kAzure'. For example, set this parameter to 'kVMware' to only return the Sources (and their Object subtrees) found in the 'kVMware' (VMware vCenter Server) environment.  NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. (optional)
    ids := []int64{int64(123)} // []int64 | Return only the registered root nodes whose Ids are given in the list. (optional)
    includeEntityPermissionInfo := true // bool | If specified, then a list of entities with permissions assigned to them are returned. (optional)
    sids := []string{"Inner_example"} // []string | Filter the registered root nodes for the sids given in the list. (optional)
    includeApplicationsTreeInfo := true // bool | Specifies whether to return applications tree info or not. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    allUnderHierarchy := true // bool | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)

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

    request := cohesitysdk.ApiListProtectionSourcesRegistrationInfoRequest{
        Environments: &environments
        Ids: &ids
        IncludeEntityPermissionInfo: &includeEntityPermissionInfo
        Sids: &sids
        IncludeApplicationsTreeInfo: &includeApplicationsTreeInfo
        TenantIds: &tenantIds
        AllUnderHierarchy: &allUnderHierarchy
    }

    resp, r, err := api_client.ProtectionSourcesApi.ListProtectionSourcesRegistrationInfo(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.ListProtectionSourcesRegistrationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProtectionSourcesRegistrationInfo`: GetRegistrationInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `ProtectionSourcesApi.ListProtectionSourcesRegistrationInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProtectionSourcesRegistrationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environments** | **[]string** | Return only Protection Sources that match the passed in environment type such as &#39;kVMware&#39;, &#39;kSQL&#39;, &#39;kView&#39; &#39;kPhysical&#39;, &#39;kPuppeteer&#39;, &#39;kPure&#39;, &#39;kNetapp&#39;, &#39;kGenericNas&#39;, &#39;kHyperV&#39;, &#39;kAcropolis&#39;, or &#39;kAzure&#39;. For example, set this parameter to &#39;kVMware&#39; to only return the Sources (and their Object subtrees) found in the &#39;kVMware&#39; (VMware vCenter Server) environment.  NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. | 
 **ids** | **[]int64** | Return only the registered root nodes whose Ids are given in the list. | 
 **includeEntityPermissionInfo** | **bool** | If specified, then a list of entities with permissions assigned to them are returned. | 
 **sids** | **[]string** | Filter the registered root nodes for the sids given in the list. | 
 **includeApplicationsTreeInfo** | **bool** | Specifies whether to return applications tree info or not. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 

### Return type

[**GetRegistrationInfoResponse**](GetRegistrationInfoResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProtectionSourcesRootNodes

> []ProtectionSourceNode ListProtectionSourcesRootNodes(ctx).Id(id).Environments(environments).Environment(environment).Execute()

Returns the top level (root) Protection Sources with registration information.



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
    id := int64(789) // int64 | Return the registration information for the Protection Source id. (optional)
    environments := []string{"Environments_example"} // []string | Return only the root Protection Sources that match the passed in environment type such as 'kVMware', 'kSQL', 'kView', 'kPuppeteer', 'kPhysical', 'kPure', 'kNetapp', 'kGenericNas', 'kHyperV', 'kAcropolis' 'kAzure'. For example, set this parameter to 'kVMware' to only return the root Protection Sources found in the 'kVMware' (VMware vCenter) environment. In addition, the registration information for each Source is returned.  NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. (optional)
    environment := "environment_example" // string | This field is deprecated. Use environments instead. deprecated: true (optional)

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

    request := cohesitysdk.ApiListProtectionSourcesRootNodesRequest{
        Id: &id
        Environments: &environments
        Environment: &environment
    }

    resp, r, err := api_client.ProtectionSourcesApi.ListProtectionSourcesRootNodes(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.ListProtectionSourcesRootNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProtectionSourcesRootNodes`: []ProtectionSourceNode
    fmt.Fprintf(os.Stdout, "Response from `ProtectionSourcesApi.ListProtectionSourcesRootNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProtectionSourcesRootNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64** | Return the registration information for the Protection Source id. | 
 **environments** | **[]string** | Return only the root Protection Sources that match the passed in environment type such as &#39;kVMware&#39;, &#39;kSQL&#39;, &#39;kView&#39;, &#39;kPuppeteer&#39;, &#39;kPhysical&#39;, &#39;kPure&#39;, &#39;kNetapp&#39;, &#39;kGenericNas&#39;, &#39;kHyperV&#39;, &#39;kAcropolis&#39; &#39;kAzure&#39;. For example, set this parameter to &#39;kVMware&#39; to only return the root Protection Sources found in the &#39;kVMware&#39; (VMware vCenter) environment. In addition, the registration information for each Source is returned.  NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. | 
 **environment** | **string** | This field is deprecated. Use environments instead. deprecated: true | 

### Return type

[**[]ProtectionSourceNode**](ProtectionSourceNode.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSqlAagHostsAndDatabases

> []SqlAagHostAndDatabases ListSqlAagHostsAndDatabases(ctx).SqlProtectionSourceIds(sqlProtectionSourceIds).Execute()

Returns the registered Protection Sources and their Object subtrees.



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
    sqlProtectionSourceIds := []int64{int64(123)} // []int64 | Specifies a list of Ids of Protection Sources registered as SQL servers. These sources may have one or more SQL databases in them. Some of them may be part of AAGs(Always on Availability Group).

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

    request := cohesitysdk.ApiListSqlAagHostsAndDatabasesRequest{
        SqlProtectionSourceIds: &sqlProtectionSourceIds
    }

    resp, r, err := api_client.ProtectionSourcesApi.ListSqlAagHostsAndDatabases(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.ListSqlAagHostsAndDatabases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSqlAagHostsAndDatabases`: []SqlAagHostAndDatabases
    fmt.Fprintf(os.Stdout, "Response from `ProtectionSourcesApi.ListSqlAagHostsAndDatabases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSqlAagHostsAndDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sqlProtectionSourceIds** | **[]int64** | Specifies a list of Ids of Protection Sources registered as SQL servers. These sources may have one or more SQL databases in them. Some of them may be part of AAGs(Always on Availability Group). | 

### Return type

[**[]SqlAagHostAndDatabases**](SqlAagHostAndDatabases.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualMachines

> []ProtectionSource ListVirtualMachines(ctx).VCenterId(vCenterId).Names(names).Uuids(uuids).Protected(protected).Execute()

Returns the Virtual Machines in a vCenter Server.



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
    vCenterId := int64(789) // int64 | Limit the VMs returned to the set of VMs found in a specific vCenter Server. Pass in the root Protection Source id for the vCenter Server to search for VMs. (optional)
    names := []string{"Inner_example"} // []string | Limit the returned VMs to those that exactly match the passed in VM name. To match multiple VM names, specify multiple \"names\" parameters that each specify a single VM name. The string must exactly match the passed in VM name and wild cards are not supported. (optional)
    uuids := []string{"Inner_example"} // []string | Limit the returned VMs to those that exactly match the passed in UUIDs. (optional)
    protected := true // bool | Limit the returned VMs to those that have been protected by a Protection Job. By default, both protected and unprotected VMs are returned. (optional)

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

    request := cohesitysdk.ApiListVirtualMachinesRequest{
        VCenterId: &vCenterId
        Names: &names
        Uuids: &uuids
        Protected: &protected
    }

    resp, r, err := api_client.ProtectionSourcesApi.ListVirtualMachines(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.ListVirtualMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVirtualMachines`: []ProtectionSource
    fmt.Fprintf(os.Stdout, "Response from `ProtectionSourcesApi.ListVirtualMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vCenterId** | **int64** | Limit the VMs returned to the set of VMs found in a specific vCenter Server. Pass in the root Protection Source id for the vCenter Server to search for VMs. | 
 **names** | **[]string** | Limit the returned VMs to those that exactly match the passed in VM name. To match multiple VM names, specify multiple \&quot;names\&quot; parameters that each specify a single VM name. The string must exactly match the passed in VM name and wild cards are not supported. | 
 **uuids** | **[]string** | Limit the returned VMs to those that exactly match the passed in UUIDs. | 
 **protected** | **bool** | Limit the returned VMs to those that have been protected by a Protection Job. By default, both protected and unprotected VMs are returned. | 

### Return type

[**[]ProtectionSource**](ProtectionSource.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshProtectionSourceById

> RefreshProtectionSourceById(ctx, id).Execute()

Refresh the Object hierarchy of the Protection Sources tree.



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
    id := int64(789) // int64 | Id of the root node of the Protection Sources tree to refresh.  Force a refresh of the Object hierarchy for the passed in Protection Sources Id.

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

    request := cohesitysdk.ApiRefreshProtectionSourceByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.ProtectionSourcesApi.RefreshProtectionSourceById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.RefreshProtectionSourceById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Id of the root node of the Protection Sources tree to refresh.  Force a refresh of the Object hierarchy for the passed in Protection Sources Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshProtectionSourceByIdRequest struct via the builder pattern


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


## RegisterApplicationServers

> ProtectionSource RegisterApplicationServers(ctx).Body(body).Execute()

Register a Protection Source as running one or more Application Servers like Microsoft SQL servers or Microsoft Exchange servers.



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
    body := *openapiclient.NewRegisterApplicationServersParameters() // RegisterApplicationServersParameters | Request to register Application Servers in a Protection Source.

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

    request := cohesitysdk.ApiRegisterApplicationServersRequest{
        Body: &body
    }

    resp, r, err := api_client.ProtectionSourcesApi.RegisterApplicationServers(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.RegisterApplicationServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterApplicationServers`: ProtectionSource
    fmt.Fprintf(os.Stdout, "Response from `ProtectionSourcesApi.RegisterApplicationServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterApplicationServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RegisterApplicationServersParameters**](RegisterApplicationServersParameters.md) | Request to register Application Servers in a Protection Source. | 

### Return type

[**ProtectionSource**](ProtectionSource.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterProtectionSource

> ProtectionSource RegisterProtectionSource(ctx).Body(body).Execute()

Register a Protection Source.



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
    body := *openapiclient.NewRegisterProtectionSourceParameters() // RegisterProtectionSourceParameters | Request to register a protection source.

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

    request := cohesitysdk.ApiRegisterProtectionSourceRequest{
        Body: &body
    }

    resp, r, err := api_client.ProtectionSourcesApi.RegisterProtectionSource(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.RegisterProtectionSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterProtectionSource`: ProtectionSource
    fmt.Fprintf(os.Stdout, "Response from `ProtectionSourcesApi.RegisterProtectionSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterProtectionSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RegisterProtectionSourceParameters**](RegisterProtectionSourceParameters.md) | Request to register a protection source. | 

### Return type

[**ProtectionSource**](ProtectionSource.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunDiagnostics

> RunDiagnosticsMessage RunDiagnostics(ctx, id).Execute()

Collect diagnostics of the protection source for a host type.



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
    id := int64(789) // int64 | Specifies the entity id.

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

    request := cohesitysdk.ApiRunDiagnosticsRequest{
        Id: &id
    }

    resp, r, err := api_client.ProtectionSourcesApi.RunDiagnostics(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.RunDiagnostics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunDiagnostics`: RunDiagnosticsMessage
    fmt.Fprintf(os.Stdout, "Response from `ProtectionSourcesApi.RunDiagnostics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies the entity id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunDiagnosticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RunDiagnosticsMessage**](RunDiagnosticsMessage.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterApplicationServers

> ProtectionSource UnregisterApplicationServers(ctx, id).Body(body).Execute()

Unregister Application Servers like Microsoft SQL servers or Microsoft Exchange servers running on a Protection Source.



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
    id := int64(789) // int64 | Specifies a unique id of the Protection Source to unregister the Application Servers. If the Protection Source is currently being backed up, unregister operation will fail.
    body := *openapiclient.NewUnRegisterApplicationServersParameters() // UnRegisterApplicationServersParameters | Request to register a protection source.

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

    request := cohesitysdk.ApiUnregisterApplicationServersRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.ProtectionSourcesApi.UnregisterApplicationServers(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.UnregisterApplicationServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnregisterApplicationServers`: ProtectionSource
    fmt.Fprintf(os.Stdout, "Response from `ProtectionSourcesApi.UnregisterApplicationServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the Protection Source to unregister the Application Servers. If the Protection Source is currently being backed up, unregister operation will fail. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterApplicationServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UnRegisterApplicationServersParameters**](UnRegisterApplicationServersParameters.md) | Request to register a protection source. | 

### Return type

[**ProtectionSource**](ProtectionSource.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterProtectionSource

> UnregisterProtectionSource(ctx, id).Execute()

Unregister a previously registered Protection Source.

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
    id := int64(789) // int64 | Specifies a unique id of the Protection Source to unregister. If the Protection Source is currently being backed up, unregister operation will fail.

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

    request := cohesitysdk.ApiUnregisterProtectionSourceRequest{
        Id: &id
    }

    resp, r, err := api_client.ProtectionSourcesApi.UnregisterProtectionSource(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.UnregisterProtectionSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the Protection Source to unregister. If the Protection Source is currently being backed up, unregister operation will fail. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterProtectionSourceRequest struct via the builder pattern


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


## UpdateApplicationServers

> ProtectionSource UpdateApplicationServers(ctx).Body(body).Execute()

Modifies the registration parameters of Application Servers in a Protection Source.



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
    body := *openapiclient.NewUpdateApplicationServerParameters() // UpdateApplicationServerParameters | Request to modify the Application Servers registration of a Protection Source.

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

    request := cohesitysdk.ApiUpdateApplicationServersRequest{
        Body: &body
    }

    resp, r, err := api_client.ProtectionSourcesApi.UpdateApplicationServers(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.UpdateApplicationServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationServers`: ProtectionSource
    fmt.Fprintf(os.Stdout, "Response from `ProtectionSourcesApi.UpdateApplicationServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateApplicationServerParameters**](UpdateApplicationServerParameters.md) | Request to modify the Application Servers registration of a Protection Source. | 

### Return type

[**ProtectionSource**](ProtectionSource.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProtectionSource

> ProtectionSourceNode UpdateProtectionSource(ctx, id).Body(body).Execute()

Update a previously registered Protection Source with new details.

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
    id := int64(789) // int64 | Specifies a unique id of the Protection Source to update.
    body := *openapiclient.NewUpdateProtectionSourceParameters() // UpdateProtectionSourceParameters | Request to update protection source. (optional)

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

    request := cohesitysdk.ApiUpdateProtectionSourceRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.ProtectionSourcesApi.UpdateProtectionSource(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.UpdateProtectionSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProtectionSource`: ProtectionSourceNode
    fmt.Fprintf(os.Stdout, "Response from `ProtectionSourcesApi.UpdateProtectionSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Specifies a unique id of the Protection Source to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProtectionSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateProtectionSourceParameters**](UpdateProtectionSourceParameters.md) | Request to update protection source. | 

### Return type

[**ProtectionSourceNode**](ProtectionSourceNode.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradePhysicalAgents

> UpgradePhysicalAgentsMessage UpgradePhysicalAgents(ctx).Body(body).Execute()

Initiate a request to upgrade Cohesity agents on one or more Physical Servers registered on the Cohesity Cluster.



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
    body := *openapiclient.NewUpgradePhysicalServerAgents([]int64{int64(123)}) // UpgradePhysicalServerAgents | Request to upgrade agents on Physical Servers. (optional)

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

    request := cohesitysdk.ApiUpgradePhysicalAgentsRequest{
        Body: &body
    }

    resp, r, err := api_client.ProtectionSourcesApi.UpgradePhysicalAgents(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionSourcesApi.UpgradePhysicalAgents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradePhysicalAgents`: UpgradePhysicalAgentsMessage
    fmt.Fprintf(os.Stdout, "Response from `ProtectionSourcesApi.UpgradePhysicalAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpgradePhysicalAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpgradePhysicalServerAgents**](UpgradePhysicalServerAgents.md) | Request to upgrade agents on Physical Servers. | 

### Return type

[**UpgradePhysicalAgentsMessage**](UpgradePhysicalAgentsMessage.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

