# \PolicyAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProtectionPolicy**](PolicyAPI.md#CreateProtectionPolicy) | **Post** /data-protect/policies | Create a Protection Policy.
[**DeleteProtectionPolicy**](PolicyAPI.md#DeleteProtectionPolicy) | **Delete** /data-protect/policies/{id} | Delete a Protection Policy.
[**GetPolicySummary**](PolicyAPI.md#GetPolicySummary) | **Get** /data-protect/policies/{id}/summary | Get the protection policy summary
[**GetPolicyTemplateById**](PolicyAPI.md#GetPolicyTemplateById) | **Get** /data-protect/policy-templates/{id} | List details about a single Policy Template.
[**GetPolicyTemplates**](PolicyAPI.md#GetPolicyTemplates) | **Get** /data-protect/policy-templates | List Policy Templates filtered by query parameters.
[**GetProtectionPolicies**](PolicyAPI.md#GetProtectionPolicies) | **Get** /data-protect/policies | List Protection Policies based on provided filtering parameters.
[**GetProtectionPolicyById**](PolicyAPI.md#GetProtectionPolicyById) | **Get** /data-protect/policies/{id} | List details about a single Protection Policy.
[**UpdateProtectionPolicy**](PolicyAPI.md#UpdateProtectionPolicy) | **Put** /data-protect/policies/{id} | Update a Protection Policy.



## CreateProtectionPolicy

> ProtectionPolicyResponse CreateProtectionPolicy(ctx).Body(body).Execute()

Create a Protection Policy.



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
	body := *openapiclient.NewProtectionPolicyRequest(*openapiclient.NewBackupPolicy(*openapiclient.NewRegularBackupPolicy()), "Name_example") // ProtectionPolicyRequest | Request to create a Protection Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.CreateProtectionPolicy(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.CreateProtectionPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProtectionPolicy`: ProtectionPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.CreateProtectionPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProtectionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProtectionPolicyRequest**](ProtectionPolicyRequest.md) | Request to create a Protection Policy. | 

### Return type

[**ProtectionPolicyResponse**](ProtectionPolicyResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	id := "id_example" // string | Specifies a unique id of the Protection Policy to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PolicyAPI.DeleteProtectionPolicy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.DeleteProtectionPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the Protection Policy to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProtectionPolicyRequest struct via the builder pattern


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


## GetPolicySummary

> PolicySummaryResponse GetPolicySummary(ctx, id).RequestInitiatorType(requestInitiatorType).IncludeAggregatedLastRunSummary(includeAggregatedLastRunSummary).IncludeAggregatedRunsSummary(includeAggregatedRunsSummary).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).PageCount(pageCount).PaginationCookie(paginationCookie).Execute()

Get the protection policy summary



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
	id := "id_example" // string | Specifies the id of the policy whose summary should be retrieved. If this is not set, the API will return error.
	requestInitiatorType := "requestInitiatorType_example" // string | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. (optional)
	includeAggregatedLastRunSummary := true // bool | Specifies whether to include summary of the last Protection Run of each Protection Source (optional)
	includeAggregatedRunsSummary := true // bool | Specifies whether to include summary of all Protection Runs of the Protection Source or Protection Jobs. If this is set to true, then only the Protection Runs from the provided 'startTimeUsecs' and 'endTimeUsecs' are processed. (optional)
	startTimeUsecs := int64(789) // int64 | Filter by a start time specified as a Unix epoch Timestamp (in microseconds). Only Job Runs that started after the specified time are included in the aggregated runs summary result. (optional)
	endTimeUsecs := int64(789) // int64 | Filter by a end time specified as a Unix epoch Timestamp (in microseconds). Only Job Runs that completed before the specified time are included in the aggregated runs summary result. (optional)
	pageCount := int64(789) // int64 | Specifies the limit of the number of Protection Sources or Protection Jobs to be returned as a part of the Protection Policy Summary. (optional)
	paginationCookie := "paginationCookie_example" // string | If set, i.e. there are more results to display, use this value to get the next set of results, by using this value in paginationCookie param for the next request to GetProtectionPolicySummary. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.GetPolicySummary(context.Background(), id).RequestInitiatorType(requestInitiatorType).IncludeAggregatedLastRunSummary(includeAggregatedLastRunSummary).IncludeAggregatedRunsSummary(includeAggregatedRunsSummary).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).PageCount(pageCount).PaginationCookie(paginationCookie).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.GetPolicySummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicySummary`: PolicySummaryResponse
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.GetPolicySummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies the id of the policy whose summary should be retrieved. If this is not set, the API will return error. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicySummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestInitiatorType** | **string** | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | 
 **includeAggregatedLastRunSummary** | **bool** | Specifies whether to include summary of the last Protection Run of each Protection Source | 
 **includeAggregatedRunsSummary** | **bool** | Specifies whether to include summary of all Protection Runs of the Protection Source or Protection Jobs. If this is set to true, then only the Protection Runs from the provided &#39;startTimeUsecs&#39; and &#39;endTimeUsecs&#39; are processed. | 
 **startTimeUsecs** | **int64** | Filter by a start time specified as a Unix epoch Timestamp (in microseconds). Only Job Runs that started after the specified time are included in the aggregated runs summary result. | 
 **endTimeUsecs** | **int64** | Filter by a end time specified as a Unix epoch Timestamp (in microseconds). Only Job Runs that completed before the specified time are included in the aggregated runs summary result. | 
 **pageCount** | **int64** | Specifies the limit of the number of Protection Sources or Protection Jobs to be returned as a part of the Protection Policy Summary. | 
 **paginationCookie** | **string** | If set, i.e. there are more results to display, use this value to get the next set of results, by using this value in paginationCookie param for the next request to GetProtectionPolicySummary. | 

### Return type

[**PolicySummaryResponse**](PolicySummaryResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyTemplateById

> PolicyTemplateResponse GetPolicyTemplateById(ctx, id).Execute()

List details about a single Policy Template.



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
	id := "id_example" // string | Specifies a unique id of the Policy Template to return.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.GetPolicyTemplateById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.GetPolicyTemplateById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicyTemplateById`: PolicyTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.GetPolicyTemplateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the Policy Template to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyTemplateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyTemplateResponse**](PolicyTemplateResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyTemplates

> PolicyTemplatesResponseWithPagination GetPolicyTemplates(ctx).RequestInitiatorType(requestInitiatorType).Ids(ids).PolicyNames(policyNames).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()

List Policy Templates filtered by query parameters.



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
	requestInitiatorType := "requestInitiatorType_example" // string | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. (optional)
	ids := []string{"Inner_example"} // []string | Filter policies by a list of policy template ids. (optional)
	policyNames := []string{"Inner_example"} // []string | Filter policies by a list of policy names. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the organizations for which objects are to be returned. (optional)
	includeTenants := true // bool | IncludeTenantPolicies specifies if objects of all the organizations under the hierarchy of the logged in user's organization should be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.GetPolicyTemplates(context.Background()).RequestInitiatorType(requestInitiatorType).Ids(ids).PolicyNames(policyNames).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.GetPolicyTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicyTemplates`: PolicyTemplatesResponseWithPagination
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.GetPolicyTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestInitiatorType** | **string** | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | 
 **ids** | **[]string** | Filter policies by a list of policy template ids. | 
 **policyNames** | **[]string** | Filter policies by a list of policy names. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the organizations for which objects are to be returned. | 
 **includeTenants** | **bool** | IncludeTenantPolicies specifies if objects of all the organizations under the hierarchy of the logged in user&#39;s organization should be returned. | 

### Return type

[**PolicyTemplatesResponseWithPagination**](PolicyTemplatesResponseWithPagination.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionPolicies

> ProtectionPolicyResponseWithPagination GetProtectionPolicies(ctx).RequestInitiatorType(requestInitiatorType).Ids(ids).PolicyNames(policyNames).TenantIds(tenantIds).IncludeTenants(includeTenants).Types(types).ExcludeLinkedPolicies(excludeLinkedPolicies).IncludeReplicatedPolicies(includeReplicatedPolicies).IncludeStats(includeStats).VaultIds(vaultIds).Execute()

List Protection Policies based on provided filtering parameters.



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
	requestInitiatorType := "requestInitiatorType_example" // string | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. (optional)
	ids := []string{"Inner_example"} // []string | Filter policies by a list of policy ids. (optional)
	policyNames := []string{"Inner_example"} // []string | Filter policies by a list of policy names. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the organizations for which objects are to be returned. (optional)
	includeTenants := true // bool | IncludeTenantPolicies specifies if objects of all the organizations under the hierarchy of the logged in user's organization should be returned. (optional)
	types := []string{"Types_example"} // []string | Types specifies the policy type of policies to be returned (optional) (default to ["Regular"])
	excludeLinkedPolicies := true // bool | If excludeLinkedPolicies is set to true then only local policies created on cluster will be returned. The result will exclude all linked policies created from policy templates. (optional)
	includeReplicatedPolicies := true // bool | If includeReplicatedPolicies is set to true, then response will also contain replicated policies. By default, replication policies are not included in the response. (optional)
	includeStats := true // bool | If includeStats is set to true, then response will return number of protection groups and objects. By default, the protection stats are not included in the response. (optional)
	vaultIds := []int64{int64(123)} // []int64 | Filter by a list of Vault ids. Policies archiving to any of the specified vaults will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.GetProtectionPolicies(context.Background()).RequestInitiatorType(requestInitiatorType).Ids(ids).PolicyNames(policyNames).TenantIds(tenantIds).IncludeTenants(includeTenants).Types(types).ExcludeLinkedPolicies(excludeLinkedPolicies).IncludeReplicatedPolicies(includeReplicatedPolicies).IncludeStats(includeStats).VaultIds(vaultIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.GetProtectionPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtectionPolicies`: ProtectionPolicyResponseWithPagination
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.GetProtectionPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestInitiatorType** | **string** | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | 
 **ids** | **[]string** | Filter policies by a list of policy ids. | 
 **policyNames** | **[]string** | Filter policies by a list of policy names. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the organizations for which objects are to be returned. | 
 **includeTenants** | **bool** | IncludeTenantPolicies specifies if objects of all the organizations under the hierarchy of the logged in user&#39;s organization should be returned. | 
 **types** | **[]string** | Types specifies the policy type of policies to be returned | [default to [&quot;Regular&quot;]]
 **excludeLinkedPolicies** | **bool** | If excludeLinkedPolicies is set to true then only local policies created on cluster will be returned. The result will exclude all linked policies created from policy templates. | 
 **includeReplicatedPolicies** | **bool** | If includeReplicatedPolicies is set to true, then response will also contain replicated policies. By default, replication policies are not included in the response. | 
 **includeStats** | **bool** | If includeStats is set to true, then response will return number of protection groups and objects. By default, the protection stats are not included in the response. | 
 **vaultIds** | **[]int64** | Filter by a list of Vault ids. Policies archiving to any of the specified vaults will be returned. | 

### Return type

[**ProtectionPolicyResponseWithPagination**](ProtectionPolicyResponseWithPagination.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectionPolicyById

> ProtectionPolicyResponse GetProtectionPolicyById(ctx, id).RequestInitiatorType(requestInitiatorType).Execute()

List details about a single Protection Policy.



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
	id := "id_example" // string | Specifies a unique id of the Protection Policy to return.
	requestInitiatorType := "requestInitiatorType_example" // string | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.GetProtectionPolicyById(context.Background(), id).RequestInitiatorType(requestInitiatorType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.GetProtectionPolicyById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtectionPolicyById`: ProtectionPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.GetProtectionPolicyById`: %v\n", resp)
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

 **requestInitiatorType** | **string** | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | 

### Return type

[**ProtectionPolicyResponse**](ProtectionPolicyResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProtectionPolicy

> ProtectionPolicyResponse UpdateProtectionPolicy(ctx, id).Body(body).Execute()

Update a Protection Policy.



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
	id := "id_example" // string | Specifies a unique id of the Protection Policy to update.
	body := *openapiclient.NewProtectionPolicyRequest(*openapiclient.NewBackupPolicy(*openapiclient.NewRegularBackupPolicy()), "Name_example") // ProtectionPolicyRequest | Request to update a Protection Policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.UpdateProtectionPolicy(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.UpdateProtectionPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProtectionPolicy`: ProtectionPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.UpdateProtectionPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the Protection Policy to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProtectionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProtectionPolicyRequest**](ProtectionPolicyRequest.md) | Request to update a Protection Policy. | 

### Return type

[**ProtectionPolicyResponse**](ProtectionPolicyResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

