# \ProtectionPoliciesApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProtectionPolicy**](ProtectionPoliciesApi.md#CreateProtectionPolicy) | **Post** /public/protectionPolicies | Create a Protection Policy.
[**DeleteProtectionPolicy**](ProtectionPoliciesApi.md#DeleteProtectionPolicy) | **Delete** /public/protectionPolicies/{id} | Delete a Protection Policy.
[**GetProtectionPolicies**](ProtectionPoliciesApi.md#GetProtectionPolicies) | **Get** /public/protectionPolicies | List Protection Policies filtered by some parameters.
[**GetProtectionPolicyById**](ProtectionPoliciesApi.md#GetProtectionPolicyById) | **Get** /public/protectionPolicies/{id} | List details about a single Protection Policy.
[**GetProtectionPolicySummary**](ProtectionPoliciesApi.md#GetProtectionPolicySummary) | **Get** /public/protectionPolicySummary | List Protection Policy Summary.
[**UpdateProtectionPolicy**](ProtectionPoliciesApi.md#UpdateProtectionPolicy) | **Put** /public/protectionPolicies/{id} | Update a Protection Policy.



## CreateProtectionPolicy

> ProtectionPolicy CreateProtectionPolicy(ctx).Body(body).Execute()

Create a Protection Policy.



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
    body := *openapiclient.NewProtectionPolicyRequest() // ProtectionPolicyRequest | Request to create a Protection Policy.

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

    request := cohesitysdk.ApiCreateProtectionPolicyRequest{
        Body: &body
    }

    resp, r, err := api_client.ProtectionPoliciesApi.CreateProtectionPolicy(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionPoliciesApi.CreateProtectionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProtectionPolicy`: ProtectionPolicy
    fmt.Fprintf(os.Stdout, "Response from `ProtectionPoliciesApi.CreateProtectionPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProtectionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProtectionPolicyRequest**](ProtectionPolicyRequest.md) | Request to create a Protection Policy. | 

### Return type

[**ProtectionPolicy**](ProtectionPolicy.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProtectionPolicy

> DeleteProtectionPolicy(ctx, id).Execute()

Delete a Protection Policy.



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
    id := "id_example" // string | Specifies a unique id of the Protection Policy to return.

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

    request := cohesitysdk.ApiDeleteProtectionPolicyRequest{
        Id: &id
    }

    resp, r, err := api_client.ProtectionPoliciesApi.DeleteProtectionPolicy(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionPoliciesApi.DeleteProtectionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the Protection Policy to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProtectionPolicyRequest struct via the builder pattern


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


## GetProtectionPolicies

> []ProtectionPolicy GetProtectionPolicies(ctx).Ids(ids).Names(names).Environments(environments).VaultIds(vaultIds).Origin(origin).Types(types).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).Execute()

List Protection Policies filtered by some parameters.



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
    ids := []string{"Inner_example"} // []string | Filter by a list of Protection Policy ids. (optional)
    names := []string{"Inner_example"} // []string | Filter by a list of Protection Policy names. (optional)
    environments := []string{"Environments_example"} // []string | Filter by Environment type such as 'kVMware', 'kView', etc. Only Policies protecting the specified environment type are returned. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. (optional)
    vaultIds := []int64{int64(123)} // []int64 | Filter by a list of Vault ids. Policies archiving to any of the specified vaults will be returned. (optional)
    origin := "origin_example" // string | Specifies the origin of the protection policy. 'kHelios' means a global policy which was created on Helios. 'kLocal' means a local policy which was created on the cluster. (optional)
    types := []string{"Types_example"} // []string | Specifies the type of the protection policy. 'kRegular' means a regular Protection Policy. 'kRPO' means an RPO Protection Policy. (optional)
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

    request := cohesitysdk.ApiGetProtectionPoliciesRequest{
        Ids: &ids
        Names: &names
        Environments: &environments
        VaultIds: &vaultIds
        Origin: &origin
        Types: &types
        TenantIds: &tenantIds
        AllUnderHierarchy: &allUnderHierarchy
    }

    resp, r, err := api_client.ProtectionPoliciesApi.GetProtectionPolicies(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionPoliciesApi.GetProtectionPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionPolicies`: []ProtectionPolicy
    fmt.Fprintf(os.Stdout, "Response from `ProtectionPoliciesApi.GetProtectionPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | Filter by a list of Protection Policy ids. | 
 **names** | **[]string** | Filter by a list of Protection Policy names. | 
 **environments** | **[]string** | Filter by Environment type such as &#39;kVMware&#39;, &#39;kView&#39;, etc. Only Policies protecting the specified environment type are returned. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. | 
 **vaultIds** | **[]int64** | Filter by a list of Vault ids. Policies archiving to any of the specified vaults will be returned. | 
 **origin** | **string** | Specifies the origin of the protection policy. &#39;kHelios&#39; means a global policy which was created on Helios. &#39;kLocal&#39; means a local policy which was created on the cluster. | 
 **types** | **[]string** | Specifies the type of the protection policy. &#39;kRegular&#39; means a regular Protection Policy. &#39;kRPO&#39; means an RPO Protection Policy. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 

### Return type

[**[]ProtectionPolicy**](ProtectionPolicy.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionPolicyById

> ProtectionPolicy GetProtectionPolicyById(ctx, id).Execute()

List details about a single Protection Policy.



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
    id := "id_example" // string | Specifies a unique id of the Protection Policy to return.

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

    request := cohesitysdk.ApiGetProtectionPolicyByIdRequest{
        Id: &id
    }

    resp, r, err := api_client.ProtectionPoliciesApi.GetProtectionPolicyById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionPoliciesApi.GetProtectionPolicyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionPolicyById`: ProtectionPolicy
    fmt.Fprintf(os.Stdout, "Response from `ProtectionPoliciesApi.GetProtectionPolicyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the Protection Policy to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionPolicyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProtectionPolicy**](ProtectionPolicy.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionPolicySummary

> ProtectionPolicySummary GetProtectionPolicySummary(ctx).Id(id).IncludeAggregatedLastRunSummary(includeAggregatedLastRunSummary).IncludeAggregatedRunsSummary(includeAggregatedRunsSummary).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).PageCount(pageCount).PaginationCookie(paginationCookie).Execute()

List Protection Policy Summary.

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
    id := "id_example" // string | Specifies the id of the policy whose summary should be retrieved. If this is not set, the API will return error.
    includeAggregatedLastRunSummary := true // bool | Specifies whether to include summary of the last Protection Run of each Protection Source. (optional)
    includeAggregatedRunsSummary := true // bool | Specifies whether to include summary of all Protection Runs of the Protection Source or Protection Jobs. If this is set to true, then only the Protection Runs from the provided 'startTimeUsecs' and 'endTimeUsecs' are processed. (optional)
    startTimeUsecs := int64(789) // int64 | Filter by a start time specified as a Unix epoch Timestamp (in microseconds). Only Job Runs that started after the specified time are included in the aggregated runs summary result. (optional)
    endTimeUsecs := int64(789) // int64 | Filter by a end time specified as a Unix epoch Timestamp (in microseconds). Only Job Runs that completed before the specified end time are included int he aggregated runs summary result. (optional)
    pageCount := int64(789) // int64 | Specifies the limit of the number of Protection Sources or Protection Jobs to be returned as a part of the Protection Policy Summary. (optional)
    paginationCookie := "paginationCookie_example" // string | If set, i.e. there are more results to display, use this value to get the next set of results, by using this value in paginationCookie param for the next request to GetProtectionPolicySummary. (optional)

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

    request := cohesitysdk.ApiGetProtectionPolicySummaryRequest{
        Id: &id
        IncludeAggregatedLastRunSummary: &includeAggregatedLastRunSummary
        IncludeAggregatedRunsSummary: &includeAggregatedRunsSummary
        StartTimeUsecs: &startTimeUsecs
        EndTimeUsecs: &endTimeUsecs
        PageCount: &pageCount
        PaginationCookie: &paginationCookie
    }

    resp, r, err := api_client.ProtectionPoliciesApi.GetProtectionPolicySummary(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionPoliciesApi.GetProtectionPolicySummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionPolicySummary`: ProtectionPolicySummary
    fmt.Fprintf(os.Stdout, "Response from `ProtectionPoliciesApi.GetProtectionPolicySummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionPolicySummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Specifies the id of the policy whose summary should be retrieved. If this is not set, the API will return error. | 
 **includeAggregatedLastRunSummary** | **bool** | Specifies whether to include summary of the last Protection Run of each Protection Source. | 
 **includeAggregatedRunsSummary** | **bool** | Specifies whether to include summary of all Protection Runs of the Protection Source or Protection Jobs. If this is set to true, then only the Protection Runs from the provided &#39;startTimeUsecs&#39; and &#39;endTimeUsecs&#39; are processed. | 
 **startTimeUsecs** | **int64** | Filter by a start time specified as a Unix epoch Timestamp (in microseconds). Only Job Runs that started after the specified time are included in the aggregated runs summary result. | 
 **endTimeUsecs** | **int64** | Filter by a end time specified as a Unix epoch Timestamp (in microseconds). Only Job Runs that completed before the specified end time are included int he aggregated runs summary result. | 
 **pageCount** | **int64** | Specifies the limit of the number of Protection Sources or Protection Jobs to be returned as a part of the Protection Policy Summary. | 
 **paginationCookie** | **string** | If set, i.e. there are more results to display, use this value to get the next set of results, by using this value in paginationCookie param for the next request to GetProtectionPolicySummary. | 

### Return type

[**ProtectionPolicySummary**](ProtectionPolicySummary.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProtectionPolicy

> ProtectionPolicy UpdateProtectionPolicy(ctx, id).Body(body).Execute()

Update a Protection Policy.



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
    id := "id_example" // string | Specifies a unique id of the Protection Policy to return.
    body := *openapiclient.NewProtectionPolicyRequest() // ProtectionPolicyRequest | Request to update a Protection Policy.

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

    request := cohesitysdk.ApiUpdateProtectionPolicyRequest{
        Id: &id
        Body: &body
    }

    resp, r, err := api_client.ProtectionPoliciesApi.UpdateProtectionPolicy(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProtectionPoliciesApi.UpdateProtectionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProtectionPolicy`: ProtectionPolicy
    fmt.Fprintf(os.Stdout, "Response from `ProtectionPoliciesApi.UpdateProtectionPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the Protection Policy to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProtectionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProtectionPolicyRequest**](ProtectionPolicyRequest.md) | Request to update a Protection Policy. | 

### Return type

[**ProtectionPolicy**](ProtectionPolicy.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

