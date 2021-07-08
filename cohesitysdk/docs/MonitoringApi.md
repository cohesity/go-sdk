# \MonitoringApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllJobRuns**](MonitoringApi.md#GetAllJobRuns) | **Get** /public/monitoring/jobs | List runs of all the jobs.
[**GetJobRunInfo**](MonitoringApi.md#GetJobRunInfo) | **Get** /public/monitoring/jobRunInfo | List info related to a job run.
[**GetRunObjectsDetails**](MonitoringApi.md#GetRunObjectsDetails) | **Get** /public/monitoring/objectDetails | List details of objects in a job run.



## GetAllJobRuns

> []GetAllJobRunsResult GetAllJobRuns(ctx).StartTimeMsecs(startTimeMsecs).EndTimeMsecs(endTimeMsecs).JobTypes(jobTypes).EnvTypes(envTypes).Page(page).PageSize(pageSize).Execute()

List runs of all the jobs.



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
    startTimeMsecs := int64(789) // int64 | Specifies the start time of the time range of interest.
    endTimeMsecs := int64(789) // int64 | Specifies the end time of the time range of interest.
    jobTypes := []string{"Inner_example"} // []string | Specifies the job types for which runs are needed. (optional)
    envTypes := []string{"EnvTypes_example"} // []string | Specifies the environment types of the job. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. 'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment. 'kSQL' indicates the SQL Protection Source environment. 'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment. 'Nimble' indicates the Nimble Storage Protection Source environment. 'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment. 'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhsicalFiles' indicates the Physical Files Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment. (optional)
    page := int32(56) // int32 | Specifies the page number in case of pagination of response. (optional)
    pageSize := int32(56) // int32 | Specifies the size of the page in case of pagination of response. (optional)

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

    request := cohesitysdk.ApiGetAllJobRunsRequest{
        StartTimeMsecs: &startTimeMsecs
        EndTimeMsecs: &endTimeMsecs
        JobTypes: &jobTypes
        EnvTypes: &envTypes
        Page: &page
        PageSize: &pageSize
    }

    resp, r, err := api_client.MonitoringApi.GetAllJobRuns(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.GetAllJobRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllJobRuns`: []GetAllJobRunsResult
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.GetAllJobRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllJobRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimeMsecs** | **int64** | Specifies the start time of the time range of interest. | 
 **endTimeMsecs** | **int64** | Specifies the end time of the time range of interest. | 
 **jobTypes** | **[]string** | Specifies the job types for which runs are needed. | 
 **envTypes** | **[]string** | Specifies the environment types of the job. Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | 
 **page** | **int32** | Specifies the page number in case of pagination of response. | 
 **pageSize** | **int32** | Specifies the size of the page in case of pagination of response. | 

### Return type

[**[]GetAllJobRunsResult**](GetAllJobRunsResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobRunInfo

> GetJobRunInfoResult GetJobRunInfo(ctx).JobType(jobType).JobId(jobId).JobRunId(jobRunId).Execute()

List info related to a job run.



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
    jobType := "jobType_example" // string | Specifies the job type of the needed run.
    jobId := int64(789) // int64 | Specifies the job id of the needed run.
    jobRunId := int64(789) // int64 | Specifies the job run id of the needed run.

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

    request := cohesitysdk.ApiGetJobRunInfoRequest{
        JobType: &jobType
        JobId: &jobId
        JobRunId: &jobRunId
    }

    resp, r, err := api_client.MonitoringApi.GetJobRunInfo(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.GetJobRunInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobRunInfo`: GetJobRunInfoResult
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.GetJobRunInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobRunInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobType** | **string** | Specifies the job type of the needed run. | 
 **jobId** | **int64** | Specifies the job id of the needed run. | 
 **jobRunId** | **int64** | Specifies the job run id of the needed run. | 

### Return type

[**GetJobRunInfoResult**](GetJobRunInfoResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunObjectsDetails

> []GetObjectsDetailsResult GetRunObjectsDetails(ctx).JobType(jobType).JobId(jobId).JobRunId(jobRunId).Execute()

List details of objects in a job run.



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
    jobType := "jobType_example" // string | Specifies the job type of the needed run.
    jobId := int64(789) // int64 | Specifies the job id of the needed run.
    jobRunId := int64(789) // int64 | Specifies the job run id of the needed run.

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

    request := cohesitysdk.ApiGetRunObjectsDetailsRequest{
        JobType: &jobType
        JobId: &jobId
        JobRunId: &jobRunId
    }

    resp, r, err := api_client.MonitoringApi.GetRunObjectsDetails(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.GetRunObjectsDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRunObjectsDetails`: []GetObjectsDetailsResult
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.GetRunObjectsDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRunObjectsDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobType** | **string** | Specifies the job type of the needed run. | 
 **jobId** | **int64** | Specifies the job id of the needed run. | 
 **jobRunId** | **int64** | Specifies the job run id of the needed run. | 

### Return type

[**[]GetObjectsDetailsResult**](GetObjectsDetailsResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

