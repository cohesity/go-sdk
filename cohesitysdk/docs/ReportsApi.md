# \ReportsApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgentDeploymentReport**](ReportsApi.md#GetAgentDeploymentReport) | **Get** /public/reports/agents | 
[**GetDataTransferFromVaultsReportRequest**](ReportsApi.md#GetDataTransferFromVaultsReportRequest) | **Get** /public/reports/dataTransferFromVaults | Get summary statistics about transferring data from Vaults (External Targets) to this Cohesity Cluster.
[**GetDataTransferToVaultsReportRequest**](ReportsApi.md#GetDataTransferToVaultsReportRequest) | **Get** /public/reports/dataTransferToVaults | Get summary statistics about transferring data from the Cohesity Cluster to Vaults (External Targets).
[**GetGdprReport**](ReportsApi.md#GetGdprReport) | **Get** /public/reports/gdpr | Get gdpr report related information.
[**GetProtectedObjectsTrendsReportRequest**](ReportsApi.md#GetProtectedObjectsTrendsReportRequest) | **Get** /public/reports/protectedObjectsTrends | Get protected objects trends grouped by days/weeks/months.
[**GetProtectionSourcesJobRunsReportRequest**](ReportsApi.md#GetProtectionSourcesJobRunsReportRequest) | **Get** /public/reports/protectionSourcesJobRuns | Get protection details about the specified list of leaf Protection Source Objects (such as a VMs).
[**GetProtectionSourcesJobsSummaryReportRequest**](ReportsApi.md#GetProtectionSourcesJobsSummaryReportRequest) | **Get** /public/reports/protectionSourcesJobsSummary | Get Job Run (Snapshot) summary statistics about the leaf Protection Sources Objects that match the specified filter criteria.



## GetAgentDeploymentReport

> []AgentDeploymentStatusResponse GetAgentDeploymentReport(ctx).HostOsType(hostOsType).CompactVersion(compactVersion).OutputFormat(outputFormat).HealthStatus(healthStatus).Execute()





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
    hostOsType := []string{"HostOsType_example"} // []string | Specifies the host type on which the Cohesity agent is installed. Setting this parameter will filter the response based on host OS type on which agent is running. 'kLinux' indicates the Linux operating system. 'kWindows' indicates the Microsoft Windows operating system. 'kAix' indicates the IBM AIX operating system. 'kSolaris' indicates the Oracle Solaris operating system. 'kSapHana' indicates the Sap Hana database system developed by SAP SE. 'kOther' indicates the other types of operating system. (optional)
    compactVersion := "compactVersion_example" // string | Specifies the compact version of Cohesity agent. For example, 6.0.1. Setting this parameter will filter the response based on installed agent version. (optional)
    outputFormat := "outputFormat_example" // string | Specifies the format for the output such as 'csv' or 'json'. If not specified, the json format is returned. If 'csv' is specified, a comma-separated list with a heading row is returned. (optional)
    healthStatus := []string{"HealthStatus_example"} // []string | Specifies the health status of the Cohesity agent. Setting this parameter will filter the response based on agent health status. Specifies the status of the agent running on a physical source. 'kUnknown' indicates the Agent is not known. No attempt to connect to the Agent has occurred. 'kUnreachable' indicates the Agent is not reachable. 'kHealthy' indicates the Agent is healthy. 'kDegraded' indicates the Agent is running but in a degraded state. (optional)

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

    request := cohesitysdk.ApiGetAgentDeploymentReportRequest{
        HostOsType: &hostOsType
        CompactVersion: &compactVersion
        OutputFormat: &outputFormat
        HealthStatus: &healthStatus
    }

    resp, r, err := api_client.ReportsApi.GetAgentDeploymentReport(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetAgentDeploymentReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentDeploymentReport`: []AgentDeploymentStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetAgentDeploymentReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentDeploymentReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostOsType** | **[]string** | Specifies the host type on which the Cohesity agent is installed. Setting this parameter will filter the response based on host OS type on which agent is running. &#39;kLinux&#39; indicates the Linux operating system. &#39;kWindows&#39; indicates the Microsoft Windows operating system. &#39;kAix&#39; indicates the IBM AIX operating system. &#39;kSolaris&#39; indicates the Oracle Solaris operating system. &#39;kSapHana&#39; indicates the Sap Hana database system developed by SAP SE. &#39;kOther&#39; indicates the other types of operating system. | 
 **compactVersion** | **string** | Specifies the compact version of Cohesity agent. For example, 6.0.1. Setting this parameter will filter the response based on installed agent version. | 
 **outputFormat** | **string** | Specifies the format for the output such as &#39;csv&#39; or &#39;json&#39;. If not specified, the json format is returned. If &#39;csv&#39; is specified, a comma-separated list with a heading row is returned. | 
 **healthStatus** | **[]string** | Specifies the health status of the Cohesity agent. Setting this parameter will filter the response based on agent health status. Specifies the status of the agent running on a physical source. &#39;kUnknown&#39; indicates the Agent is not known. No attempt to connect to the Agent has occurred. &#39;kUnreachable&#39; indicates the Agent is not reachable. &#39;kHealthy&#39; indicates the Agent is healthy. &#39;kDegraded&#39; indicates the Agent is running but in a degraded state. | 

### Return type

[**[]AgentDeploymentStatusResponse**](AgentDeploymentStatusResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTransferFromVaultsReportRequest

> DataTransferFromVaultsSummaryResponse GetDataTransferFromVaultsReportRequest(ctx).VaultIds(vaultIds).StartTimeMsecs(startTimeMsecs).EndTimeMsecs(endTimeMsecs).OutputFormat(outputFormat).GroupBy(groupBy).Execute()

Get summary statistics about transferring data from Vaults (External Targets) to this Cohesity Cluster.



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
    vaultIds := []int64{int64(123)} // []int64 | Filter by a list of Vault ids.
    startTimeMsecs := int64(789) // int64 | Filter by a start time. Specify the start time as a Unix epoch Timestamp (in milliseconds). If startTimeMsecs and endTimeMsecs are not specified, the time period is the last 7 days. (optional)
    endTimeMsecs := int64(789) // int64 | Filter by end time. Specify the end time as a Unix epoch Timestamp (in milliseconds). If startTimeMsecs and endTimeMsecs are not specified, the time period is the last 7 days. (optional)
    outputFormat := "outputFormat_example" // string | Specifies the format for the output such as 'csv' or 'json'. If not specified, the json format is returned. If 'csv' is specified, a comma-separated list with a heading row is returned. (optional)
    groupBy := int32(56) // int32 | Specifies wheather the report should be grouped by target when scheduled or downloaded. If not set or set to false, report is grouped by protection jobs. It is ignored if outformat is not \"csv\" and response contains whole report. (optional)

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

    request := cohesitysdk.ApiGetDataTransferFromVaultsReportRequestRequest{
        VaultIds: &vaultIds
        StartTimeMsecs: &startTimeMsecs
        EndTimeMsecs: &endTimeMsecs
        OutputFormat: &outputFormat
        GroupBy: &groupBy
    }

    resp, r, err := api_client.ReportsApi.GetDataTransferFromVaultsReportRequest(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetDataTransferFromVaultsReportRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataTransferFromVaultsReportRequest`: DataTransferFromVaultsSummaryResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetDataTransferFromVaultsReportRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTransferFromVaultsReportRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vaultIds** | **[]int64** | Filter by a list of Vault ids. | 
 **startTimeMsecs** | **int64** | Filter by a start time. Specify the start time as a Unix epoch Timestamp (in milliseconds). If startTimeMsecs and endTimeMsecs are not specified, the time period is the last 7 days. | 
 **endTimeMsecs** | **int64** | Filter by end time. Specify the end time as a Unix epoch Timestamp (in milliseconds). If startTimeMsecs and endTimeMsecs are not specified, the time period is the last 7 days. | 
 **outputFormat** | **string** | Specifies the format for the output such as &#39;csv&#39; or &#39;json&#39;. If not specified, the json format is returned. If &#39;csv&#39; is specified, a comma-separated list with a heading row is returned. | 
 **groupBy** | **int32** | Specifies wheather the report should be grouped by target when scheduled or downloaded. If not set or set to false, report is grouped by protection jobs. It is ignored if outformat is not \&quot;csv\&quot; and response contains whole report. | 

### Return type

[**DataTransferFromVaultsSummaryResponse**](DataTransferFromVaultsSummaryResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataTransferToVaultsReportRequest

> DataTransferToVaultsSummaryResponse GetDataTransferToVaultsReportRequest(ctx).VaultIds(vaultIds).StartTimeMsecs(startTimeMsecs).EndTimeMsecs(endTimeMsecs).OutputFormat(outputFormat).GroupBy(groupBy).Execute()

Get summary statistics about transferring data from the Cohesity Cluster to Vaults (External Targets).



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
    vaultIds := []int64{int64(123)} // []int64 | Filter by a list of Vault ids.
    startTimeMsecs := int64(789) // int64 | Filter by a start time. Specify the start time as a Unix epoch Timestamp (in milliseconds). If startTimeMsecs and endTimeMsecs are not specified, the time period is the last 7 days. (optional)
    endTimeMsecs := int64(789) // int64 | Filter by end time. Specify the end time as a Unix epoch Timestamp (in milliseconds). If startTimeMsecs and endTimeMsecs are not specified, the time period is the last 7 days. (optional)
    outputFormat := "outputFormat_example" // string | Specifies the format for the output such as 'csv' or 'json'. If not specified, the json format is returned. If 'csv' is specified, a comma-separated list with a heading row is returned. (optional)
    groupBy := int32(56) // int32 | Specifies wheather the report should be grouped by target when scheduled or downloaded. If not set or set to false, report is grouped by protection jobs. It is ignored if outformat is not \"csv\" and response contains whole report. (optional)

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

    request := cohesitysdk.ApiGetDataTransferToVaultsReportRequestRequest{
        VaultIds: &vaultIds
        StartTimeMsecs: &startTimeMsecs
        EndTimeMsecs: &endTimeMsecs
        OutputFormat: &outputFormat
        GroupBy: &groupBy
    }

    resp, r, err := api_client.ReportsApi.GetDataTransferToVaultsReportRequest(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetDataTransferToVaultsReportRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataTransferToVaultsReportRequest`: DataTransferToVaultsSummaryResponse
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetDataTransferToVaultsReportRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataTransferToVaultsReportRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vaultIds** | **[]int64** | Filter by a list of Vault ids. | 
 **startTimeMsecs** | **int64** | Filter by a start time. Specify the start time as a Unix epoch Timestamp (in milliseconds). If startTimeMsecs and endTimeMsecs are not specified, the time period is the last 7 days. | 
 **endTimeMsecs** | **int64** | Filter by end time. Specify the end time as a Unix epoch Timestamp (in milliseconds). If startTimeMsecs and endTimeMsecs are not specified, the time period is the last 7 days. | 
 **outputFormat** | **string** | Specifies the format for the output such as &#39;csv&#39; or &#39;json&#39;. If not specified, the json format is returned. If &#39;csv&#39; is specified, a comma-separated list with a heading row is returned. | 
 **groupBy** | **int32** | Specifies wheather the report should be grouped by target when scheduled or downloaded. If not set or set to false, report is grouped by protection jobs. It is ignored if outformat is not \&quot;csv\&quot; and response contains whole report. | 

### Return type

[**DataTransferToVaultsSummaryResponse**](DataTransferToVaultsSummaryResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGdprReport

> []ObjectInformation GetGdprReport(ctx).Id(id).AccessibleUsers(accessibleUsers).ParentSourceId(parentSourceId).OutputFormat(outputFormat).Actions(actions).Search(search).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).Execute()

Get gdpr report related information.



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
    id := []int64{int64(123)} // []int64 | Specifies the objects for which to get the gdpr information. (optional)
    accessibleUsers := []string{"Inner_example"} // []string | Specifies the users for which to get the accessible objects. (optional)
    parentSourceId := []int64{int64(123)} // []int64 | Specifies the parent sources of objects for which to get info for. (optional)
    outputFormat := "outputFormat_example" // string | Specifies the format for the output such as 'csv' or 'json'. If not specified, the json format is returned. If 'csv' is specified, a comma-separated list with a heading row is returned. (optional)
    actions := []string{"Inner_example"} // []string | Specifies the action for the audit logs. (optional)
    search := "search_example" // string | Specifies the search string for the audit logs. (optional)
    startTimeUsecs := int64(789) // int64 | Specifies the start time for the audit logs as a Unix epoch Timestamp (in microseconds). (optional)
    endTimeUsecs := int64(789) // int64 | Specifies the end time for the audit logsas a Unix epoch Timestamp (in microseconds). (optional)

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

    request := cohesitysdk.ApiGetGdprReportRequest{
        Id: &id
        AccessibleUsers: &accessibleUsers
        ParentSourceId: &parentSourceId
        OutputFormat: &outputFormat
        Actions: &actions
        Search: &search
        StartTimeUsecs: &startTimeUsecs
        EndTimeUsecs: &endTimeUsecs
    }

    resp, r, err := api_client.ReportsApi.GetGdprReport(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetGdprReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGdprReport`: []ObjectInformation
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetGdprReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGdprReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]int64** | Specifies the objects for which to get the gdpr information. | 
 **accessibleUsers** | **[]string** | Specifies the users for which to get the accessible objects. | 
 **parentSourceId** | **[]int64** | Specifies the parent sources of objects for which to get info for. | 
 **outputFormat** | **string** | Specifies the format for the output such as &#39;csv&#39; or &#39;json&#39;. If not specified, the json format is returned. If &#39;csv&#39; is specified, a comma-separated list with a heading row is returned. | 
 **actions** | **[]string** | Specifies the action for the audit logs. | 
 **search** | **string** | Specifies the search string for the audit logs. | 
 **startTimeUsecs** | **int64** | Specifies the start time for the audit logs as a Unix epoch Timestamp (in microseconds). | 
 **endTimeUsecs** | **int64** | Specifies the end time for the audit logsas a Unix epoch Timestamp (in microseconds). | 

### Return type

[**[]ObjectInformation**](ObjectInformation.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectedObjectsTrendsReportRequest

> []ProtectionTrend GetProtectedObjectsTrendsReportRequest(ctx).Timezone(timezone).JobIds(jobIds).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).Environments(environments).ProtectedObjectIds(protectedObjectIds).RegisteredSourceId(registeredSourceId).Rollup(rollup).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).Execute()

Get protected objects trends grouped by days/weeks/months.



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
    timezone := "timezone_example" // string | Specifies the timezone to use when calculating day/week/month Specify the timezone in the following format: \"Area/Location\", for example: \"America/New_York\".
    jobIds := []int64{int64(123)} // []int64 | Filter by a list of Job ids. Snapshots summary statistics for the specified Protection Jobs are reported. (optional)
    startTimeUsecs := int64(789) // int64 | Filter by a start time. Snapshot summary statistics for Job Runs that started after the specified time are reported. Specify the start time as a Unix epoch Timestamp (in microseconds). (optional)
    endTimeUsecs := int64(789) // int64 | Filter by end time. Snapshot summary statistics for Job Runs that ended before the specified time are returned. Specify the end time as a Unix epoch Timestamp (in microseconds). (optional)
    environments := []string{"Environments_example"} // []string | Filter by a list of environment types such as 'kVMware', 'kView', etc. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. 'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment. 'kSQL' indicates the SQL Protection Source environment. 'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment. 'Nimble' indicates the Nimble Storage Protection Source environment. 'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment. 'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhsicalFiles' indicates the Physical Files Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment. (optional)
    protectedObjectIds := []int64{int64(123)} // []int64 | Filter by a list of leaf Protection Sources Objects (such as VMs). Snapshot summary statistics for the listed Protection Source Objects are reported. (optional)
    registeredSourceId := int64(789) // int64 | Specifies an id of a top level Registered Source such as a vCenter Server. If specified, Snapshot summary statistics for all the leaf Protection Sources (such as VMs) that are children of this Registered Source are reported. NOTE: If specified, filtering by other fields is not supported. (optional)
    rollup := "rollup_example" // string | Roll up type for grouping. Valid values are day, week, month (optional)
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

    request := cohesitysdk.ApiGetProtectedObjectsTrendsReportRequestRequest{
        Timezone: &timezone
        JobIds: &jobIds
        StartTimeUsecs: &startTimeUsecs
        EndTimeUsecs: &endTimeUsecs
        Environments: &environments
        ProtectedObjectIds: &protectedObjectIds
        RegisteredSourceId: &registeredSourceId
        Rollup: &rollup
        TenantIds: &tenantIds
        AllUnderHierarchy: &allUnderHierarchy
    }

    resp, r, err := api_client.ReportsApi.GetProtectedObjectsTrendsReportRequest(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetProtectedObjectsTrendsReportRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectedObjectsTrendsReportRequest`: []ProtectionTrend
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetProtectedObjectsTrendsReportRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectedObjectsTrendsReportRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timezone** | **string** | Specifies the timezone to use when calculating day/week/month Specify the timezone in the following format: \&quot;Area/Location\&quot;, for example: \&quot;America/New_York\&quot;. | 
 **jobIds** | **[]int64** | Filter by a list of Job ids. Snapshots summary statistics for the specified Protection Jobs are reported. | 
 **startTimeUsecs** | **int64** | Filter by a start time. Snapshot summary statistics for Job Runs that started after the specified time are reported. Specify the start time as a Unix epoch Timestamp (in microseconds). | 
 **endTimeUsecs** | **int64** | Filter by end time. Snapshot summary statistics for Job Runs that ended before the specified time are returned. Specify the end time as a Unix epoch Timestamp (in microseconds). | 
 **environments** | **[]string** | Filter by a list of environment types such as &#39;kVMware&#39;, &#39;kView&#39;, etc. Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | 
 **protectedObjectIds** | **[]int64** | Filter by a list of leaf Protection Sources Objects (such as VMs). Snapshot summary statistics for the listed Protection Source Objects are reported. | 
 **registeredSourceId** | **int64** | Specifies an id of a top level Registered Source such as a vCenter Server. If specified, Snapshot summary statistics for all the leaf Protection Sources (such as VMs) that are children of this Registered Source are reported. NOTE: If specified, filtering by other fields is not supported. | 
 **rollup** | **string** | Roll up type for grouping. Valid values are day, week, month | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 

### Return type

[**[]ProtectionTrend**](ProtectionTrend.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionSourcesJobRunsReportRequest

> []ProtectionSourcesJobRunsReportElement GetProtectionSourcesJobRunsReportRequest(ctx).ProtectionSourceIds(protectionSourceIds).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).JobIds(jobIds).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).Environments(environments).OutputFormat(outputFormat).PageCount(pageCount).RunStatus(runStatus).Execute()

Get protection details about the specified list of leaf Protection Source Objects (such as a VMs).



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
    protectionSourceIds := []int64{int64(123)} // []int64 | Filter by a list of leaf Protection Sources Objects (such as VMs). Snapshots of the specified Protection Source Objects are returned.
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    allUnderHierarchy := true // bool | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)
    jobIds := []int64{int64(123)} // []int64 | Filter by a list of Job ids. Snapshots for the specified Protection Jobs are listed. (optional)
    startTimeUsecs := int64(789) // int64 | Filter by a start time. Snapshots that started after the specified time are returned. Specify the start time as a Unix epoch Timestamp (in microseconds). (optional)
    endTimeUsecs := int64(789) // int64 | Filter by a end time. Snapshots that ended before the specified time are returned. Specify the end time as a Unix epoch Timestamp (in microseconds). (optional)
    environments := []string{"Environments_example"} // []string | Filter by a list of environment types such as 'kVMware', 'kView', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. (optional)
    outputFormat := "outputFormat_example" // string | Specifies the format for the output such as 'cvs' or 'json'. If not specified, the json format is returned. If 'csv' is specified, a comma-separated list with a heading row is returned. (optional)
    pageCount := int32(56) // int32 | Specifies the number of Snapshots to return in the response for pagination purposes. Used in combination with the paginationCookie in the response to return multiple sets of Snapshots. (optional)
    runStatus := []string{"RunStatus_example"} // []string | Filter by a list of run statuses such as 'kRunning', 'kSuccess', 'kFailure' etc. Snapshots of Job Runs with the specified run statuses are reported. 'kSuccess' indicates that the Job Run was successful. 'kRunning' indicates that the Job Run is currently running. 'kWarning' indicates that the Job Run was successful but warnings were issued. 'kCancelled' indicates that the Job Run was canceled. 'kError' indicates the Job Run encountered an error and did not run to completion. (optional)

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

    request := cohesitysdk.ApiGetProtectionSourcesJobRunsReportRequestRequest{
        ProtectionSourceIds: &protectionSourceIds
        TenantIds: &tenantIds
        AllUnderHierarchy: &allUnderHierarchy
        JobIds: &jobIds
        StartTimeUsecs: &startTimeUsecs
        EndTimeUsecs: &endTimeUsecs
        Environments: &environments
        OutputFormat: &outputFormat
        PageCount: &pageCount
        RunStatus: &runStatus
    }

    resp, r, err := api_client.ReportsApi.GetProtectionSourcesJobRunsReportRequest(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetProtectionSourcesJobRunsReportRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionSourcesJobRunsReportRequest`: []ProtectionSourcesJobRunsReportElement
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetProtectionSourcesJobRunsReportRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionSourcesJobRunsReportRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **protectionSourceIds** | **[]int64** | Filter by a list of leaf Protection Sources Objects (such as VMs). Snapshots of the specified Protection Source Objects are returned. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 
 **jobIds** | **[]int64** | Filter by a list of Job ids. Snapshots for the specified Protection Jobs are listed. | 
 **startTimeUsecs** | **int64** | Filter by a start time. Snapshots that started after the specified time are returned. Specify the start time as a Unix epoch Timestamp (in microseconds). | 
 **endTimeUsecs** | **int64** | Filter by a end time. Snapshots that ended before the specified time are returned. Specify the end time as a Unix epoch Timestamp (in microseconds). | 
 **environments** | **[]string** | Filter by a list of environment types such as &#39;kVMware&#39;, &#39;kView&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. | 
 **outputFormat** | **string** | Specifies the format for the output such as &#39;cvs&#39; or &#39;json&#39;. If not specified, the json format is returned. If &#39;csv&#39; is specified, a comma-separated list with a heading row is returned. | 
 **pageCount** | **int32** | Specifies the number of Snapshots to return in the response for pagination purposes. Used in combination with the paginationCookie in the response to return multiple sets of Snapshots. | 
 **runStatus** | **[]string** | Filter by a list of run statuses such as &#39;kRunning&#39;, &#39;kSuccess&#39;, &#39;kFailure&#39; etc. Snapshots of Job Runs with the specified run statuses are reported. &#39;kSuccess&#39; indicates that the Job Run was successful. &#39;kRunning&#39; indicates that the Job Run is currently running. &#39;kWarning&#39; indicates that the Job Run was successful but warnings were issued. &#39;kCancelled&#39; indicates that the Job Run was canceled. &#39;kError&#39; indicates the Job Run encountered an error and did not run to completion. | 

### Return type

[**[]ProtectionSourcesJobRunsReportElement**](ProtectionSourcesJobRunsReportElement.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionSourcesJobsSummaryReportRequest

> ProtectionSourcesJobsSummaryReportResponseWrapper GetProtectionSourcesJobsSummaryReportRequest(ctx).JobIds(jobIds).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).Environments(environments).ProtectionSourceIds(protectionSourceIds).Statuses(statuses).OutputFormat(outputFormat).RegisteredSourceId(registeredSourceId).ConsecutiveFailures(consecutiveFailures).ReportName(reportName).ReportType(reportType).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).Execute()

Get Job Run (Snapshot) summary statistics about the leaf Protection Sources Objects that match the specified filter criteria.



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
    jobIds := []int64{int64(123)} // []int64 | Filter by a list of Job ids. Snapshots summary statistics for the specified Protection Jobs are reported. (optional)
    startTimeUsecs := int64(789) // int64 | Filter by a start time. Snapshot summary statistics for Job Runs that started after the specified time are reported. Specify the start time as a Unix epoch Timestamp (in microseconds). (optional)
    endTimeUsecs := int64(789) // int64 | Filter by end time. Snapshot summary statistics for Job Runs that ended before the specified time are returned. Specify the end time as a Unix epoch Timestamp (in microseconds). (optional)
    environments := []string{"Environments_example"} // []string | Filter by a list of environment types such as 'kVMware', 'kView', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. (optional)
    protectionSourceIds := []int64{int64(123)} // []int64 | Filter by a list of leaf Protection Sources Objects (such as VMs). Snapshot summary statistics for the listed Protection Source Objects are reported. (optional)
    statuses := []string{"Statuses_example"} // []string | Filter by a list of run statuses. 'kSuccess' indicates that the Job Run was successful. 'kRunning' indicates that the Job Run is currently running. 'kWarning' indicates that the Job Run was successful but warnings were issued. 'kCancelled' indicates that the Job Run was canceled. 'kError' indicates the Job Run encountered an error and did not run to completion. (optional)
    outputFormat := "outputFormat_example" // string | Specifies the format for the output such as 'csv' or 'json'. If not specified, the json format is returned. If 'csv' is specified, a comma-separated list with a heading row is returned. (optional)
    registeredSourceId := int64(789) // int64 | Specifies an id of a top level Registered Source such as a vCenter Server. If specified, Snapshot summary statistics for all the leaf Protection Sources (such as VMs) that are children of this Registered Source are reported. NOTE: If specified, filtering by other fields is not supported. (optional)
    consecutiveFailures := int32(56) // int32 | Filters out those jobs which have number of consecutive run failures less than consecutiveFailures. (optional)
    reportName := "reportName_example" // string | Specifies the custom report name the user wants to set for this report. (optional)
    reportType := int32(56) // int32 | Specifies the report type that will be used to set the right label & subject line for the report when downloaded or emailed because same API is used for 3 reports currently 1. kAvailableLocalSnapshotsReport 2. kFailedObjectsReport 3. kProtectionSummaryByObjectTypeReport (optional)
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

    request := cohesitysdk.ApiGetProtectionSourcesJobsSummaryReportRequestRequest{
        JobIds: &jobIds
        StartTimeUsecs: &startTimeUsecs
        EndTimeUsecs: &endTimeUsecs
        Environments: &environments
        ProtectionSourceIds: &protectionSourceIds
        Statuses: &statuses
        OutputFormat: &outputFormat
        RegisteredSourceId: &registeredSourceId
        ConsecutiveFailures: &consecutiveFailures
        ReportName: &reportName
        ReportType: &reportType
        TenantIds: &tenantIds
        AllUnderHierarchy: &allUnderHierarchy
    }

    resp, r, err := api_client.ReportsApi.GetProtectionSourcesJobsSummaryReportRequest(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.GetProtectionSourcesJobsSummaryReportRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionSourcesJobsSummaryReportRequest`: ProtectionSourcesJobsSummaryReportResponseWrapper
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.GetProtectionSourcesJobsSummaryReportRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionSourcesJobsSummaryReportRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobIds** | **[]int64** | Filter by a list of Job ids. Snapshots summary statistics for the specified Protection Jobs are reported. | 
 **startTimeUsecs** | **int64** | Filter by a start time. Snapshot summary statistics for Job Runs that started after the specified time are reported. Specify the start time as a Unix epoch Timestamp (in microseconds). | 
 **endTimeUsecs** | **int64** | Filter by end time. Snapshot summary statistics for Job Runs that ended before the specified time are returned. Specify the end time as a Unix epoch Timestamp (in microseconds). | 
 **environments** | **[]string** | Filter by a list of environment types such as &#39;kVMware&#39;, &#39;kView&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. | 
 **protectionSourceIds** | **[]int64** | Filter by a list of leaf Protection Sources Objects (such as VMs). Snapshot summary statistics for the listed Protection Source Objects are reported. | 
 **statuses** | **[]string** | Filter by a list of run statuses. &#39;kSuccess&#39; indicates that the Job Run was successful. &#39;kRunning&#39; indicates that the Job Run is currently running. &#39;kWarning&#39; indicates that the Job Run was successful but warnings were issued. &#39;kCancelled&#39; indicates that the Job Run was canceled. &#39;kError&#39; indicates the Job Run encountered an error and did not run to completion. | 
 **outputFormat** | **string** | Specifies the format for the output such as &#39;csv&#39; or &#39;json&#39;. If not specified, the json format is returned. If &#39;csv&#39; is specified, a comma-separated list with a heading row is returned. | 
 **registeredSourceId** | **int64** | Specifies an id of a top level Registered Source such as a vCenter Server. If specified, Snapshot summary statistics for all the leaf Protection Sources (such as VMs) that are children of this Registered Source are reported. NOTE: If specified, filtering by other fields is not supported. | 
 **consecutiveFailures** | **int32** | Filters out those jobs which have number of consecutive run failures less than consecutiveFailures. | 
 **reportName** | **string** | Specifies the custom report name the user wants to set for this report. | 
 **reportType** | **int32** | Specifies the report type that will be used to set the right label &amp; subject line for the report when downloaded or emailed because same API is used for 3 reports currently 1. kAvailableLocalSnapshotsReport 2. kFailedObjectsReport 3. kProtectionSummaryByObjectTypeReport | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 

### Return type

[**ProtectionSourcesJobsSummaryReportResponseWrapper**](ProtectionSourcesJobsSummaryReportResponseWrapper.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

