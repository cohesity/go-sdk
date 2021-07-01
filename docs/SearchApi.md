# \SearchApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchProtectionRuns**](SearchApi.md#SearchProtectionRuns) | **Get** /public/search/protectionRuns | 
[**SearchProtectionSources**](SearchApi.md#SearchProtectionSources) | **Get** /public/search/protectionSources | Performs search on all the objects which are registered to a cluster.



## SearchProtectionRuns

> ProtectionRunResponse SearchProtectionRuns(ctx).Uuid(uuid).Execute()





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
    uuid := "uuid_example" // string | Specifies the unique id of the Protection Source.

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

    request := cohesitysdk.ApiSearchProtectionRunsRequest{
        uuid: &Uuid
    }

    resp, r, err := api_client.SearchApi.SearchProtectionRuns(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchProtectionRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchProtectionRuns`: ProtectionRunResponse
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchProtectionRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProtectionRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **string** | Specifies the unique id of the Protection Source. | 

### Return type

[**ProtectionRunResponse**](ProtectionRunResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchProtectionSources

> []ProtectionSourceResponse SearchProtectionSources(ctx).Office365ProtectionSourceTypes(office365ProtectionSourceTypes).DepartmentList(departmentList).TitleList(titleList).CountryList(countryList).SearchString(searchString).ProtectionStatus(protectionStatus).Environments(environments).LastProtectionJobRunStatus(lastProtectionJobRunStatus).PhysicalServerHostTypes(physicalServerHostTypes).RegisteredSourceUuids(registeredSourceUuids).StartIndex(startIndex).PageCount(pageCount).Execute()

Performs search on all the objects which are registered to a cluster.



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
    office365ProtectionSourceTypes := []string{"Office365ProtectionSourceTypes_example"} // []string | Specifies the Array of Office365 source types. Specifies the type of Office 365 entity 'kDomain' indicates the O365 domain through which authentication occurs. 'kOutlook' indicates the Exchange online entities. 'kMailbox' indicates the user's mailbox account. 'kUsers' indicates the container for User entities. 'kGroups' indicates the container for Group entities. 'kSites' indicates the container for Site entities. 'kUser' indicates an Office365 User entity. 'kGroup' indicates an Office365 Group entity. 'kSite' indicates an Office365 SharePoint Site entity. (optional)
    departmentList := []string{"Inner_example"} // []string | Specifies the list of departments to which an Office365 user may belong. (optional)
    titleList := []string{"Inner_example"} // []string | Specifies the list of titles/desgination applicable to Office365 users. (optional)
    countryList := []string{"Inner_example"} // []string | Specifies the list of countries to which Office365 user belongs. (optional)
    searchString := "searchString_example" // string | Specifies the search string to query the name of the Protection Source or the name of the job protecting a Protection Source. (optional)
    protectionStatus := []string{"Inner_example"} // []string | Specifies the protection status of the object. If specified, the objects are filtered based on current protection status of that object on the cluster. Possible values that can be passed are \"protected\", \"unprotected\" or both. If not specified, all the objects are returned irrespective of protection status of the object. (optional)
    environments := []string{"Environments_example"} // []string | Specifies the environment type by which Protection Sources will be filtered. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. 'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment. 'kSQL' indicates the SQL Protection Source environment. 'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment. 'Nimble' indicates the Nimble Storage Protection Source environment. 'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment. 'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhsicalFiles' indicates the Physical Files Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment. (optional)
    lastProtectionJobRunStatus := []int32{int32(123)} // []int32 | Specifies the last Protection Job run status of the object. If specified, objects will be filtered based on last job run status. (optional)
    physicalServerHostTypes := []string{"PhysicalServerHostTypes_example"} // []string | Specifies physical server host OS type. If specified, the physical server objects will be filtered based on OS type of the server. 'kLinux' indicates the Linux operating system. 'kWindows' indicates the Microsoft Windows operating system. 'kAix' indicates the IBM AIX operating system. 'kSolaris' indicates the Oracle Solaris operating system. 'kSapHana' indicates the Sap Hana database system developed by SAP SE. 'kOther' indicates the other types of operating system. (optional)
    registeredSourceUuids := []string{"Inner_example"} // []string | Specifies the list of Registered Sources Uuids. Only items from the listed Registered Sources are returned. (optional)
    startIndex := int32(56) // int32 | Specifies an index number that can be used to return subsets of items in multiple requests. Break up the items to return into multiple requests by setting pageCount and using startIndex to return a subsets of items.  For example, set startIndex to 0 to get the first set of items for the first request. Increment startIndex by pageCount to get the next set of items for a next request. (optional)
    pageCount := int32(56) // int32 | Specifies the number of items to return in the response for pagination purposes. Default the pageCount to MaxSearchResponseSize if this is unspecified. (optional)

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

    request := cohesitysdk.ApiSearchProtectionSourcesRequest{
        office365ProtectionSourceTypes: &Office365ProtectionSourceTypes
        departmentList: &DepartmentList
        titleList: &TitleList
        countryList: &CountryList
        searchString: &SearchString
        protectionStatus: &ProtectionStatus
        environments: &Environments
        lastProtectionJobRunStatus: &LastProtectionJobRunStatus
        physicalServerHostTypes: &PhysicalServerHostTypes
        registeredSourceUuids: &RegisteredSourceUuids
        startIndex: &StartIndex
        pageCount: &PageCount
    }

    resp, r, err := api_client.SearchApi.SearchProtectionSources(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchProtectionSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchProtectionSources`: []ProtectionSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchProtectionSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProtectionSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **office365ProtectionSourceTypes** | **[]string** | Specifies the Array of Office365 source types. Specifies the type of Office 365 entity &#39;kDomain&#39; indicates the O365 domain through which authentication occurs. &#39;kOutlook&#39; indicates the Exchange online entities. &#39;kMailbox&#39; indicates the user&#39;s mailbox account. &#39;kUsers&#39; indicates the container for User entities. &#39;kGroups&#39; indicates the container for Group entities. &#39;kSites&#39; indicates the container for Site entities. &#39;kUser&#39; indicates an Office365 User entity. &#39;kGroup&#39; indicates an Office365 Group entity. &#39;kSite&#39; indicates an Office365 SharePoint Site entity. | 
 **departmentList** | **[]string** | Specifies the list of departments to which an Office365 user may belong. | 
 **titleList** | **[]string** | Specifies the list of titles/desgination applicable to Office365 users. | 
 **countryList** | **[]string** | Specifies the list of countries to which Office365 user belongs. | 
 **searchString** | **string** | Specifies the search string to query the name of the Protection Source or the name of the job protecting a Protection Source. | 
 **protectionStatus** | **[]string** | Specifies the protection status of the object. If specified, the objects are filtered based on current protection status of that object on the cluster. Possible values that can be passed are \&quot;protected\&quot;, \&quot;unprotected\&quot; or both. If not specified, all the objects are returned irrespective of protection status of the object. | 
 **environments** | **[]string** | Specifies the environment type by which Protection Sources will be filtered. Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | 
 **lastProtectionJobRunStatus** | **[]int32** | Specifies the last Protection Job run status of the object. If specified, objects will be filtered based on last job run status. | 
 **physicalServerHostTypes** | **[]string** | Specifies physical server host OS type. If specified, the physical server objects will be filtered based on OS type of the server. &#39;kLinux&#39; indicates the Linux operating system. &#39;kWindows&#39; indicates the Microsoft Windows operating system. &#39;kAix&#39; indicates the IBM AIX operating system. &#39;kSolaris&#39; indicates the Oracle Solaris operating system. &#39;kSapHana&#39; indicates the Sap Hana database system developed by SAP SE. &#39;kOther&#39; indicates the other types of operating system. | 
 **registeredSourceUuids** | **[]string** | Specifies the list of Registered Sources Uuids. Only items from the listed Registered Sources are returned. | 
 **startIndex** | **int32** | Specifies an index number that can be used to return subsets of items in multiple requests. Break up the items to return into multiple requests by setting pageCount and using startIndex to return a subsets of items.  For example, set startIndex to 0 to get the first set of items for the first request. Increment startIndex by pageCount to get the next set of items for a next request. | 
 **pageCount** | **int32** | Specifies the number of items to return in the response for pagination purposes. Default the pageCount to MaxSearchResponseSize if this is unspecified. | 

### Return type

[**[]ProtectionSourceResponse**](ProtectionSourceResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

