# \Policy

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHeliosPolicy**](Policy.md#CreateHeliosPolicy) | **Post** /mcm/data-protect/policies | Create a Policy.
[**CreateProtectionPolicy**](Policy.md#CreateProtectionPolicy) | **Post** /data-protect/policies | Create a Protection Policy.
[**DeleteHeliosPolicy**](Policy.md#DeleteHeliosPolicy) | **Delete** /mcm/data-protect/policies/{id} | Delete a Policy.
[**DeleteProtectionPolicy**](Policy.md#DeleteProtectionPolicy) | **Delete** /data-protect/policies/{id} | Delete a Protection Policy.
[**GetHeliosPolicies**](Policy.md#GetHeliosPolicies) | **Get** /mcm/data-protect/policies | List Policies based on provided filtering parameters.
[**GetHeliosPolicyById**](Policy.md#GetHeliosPolicyById) | **Get** /mcm/data-protect/policies/{id} | List details about a single Protection Policy.
[**GetPolicyTemplateById**](Policy.md#GetPolicyTemplateById) | **Get** /data-protect/policy-templates/{id} | List details about a single Policy Template.
[**GetPolicyTemplates**](Policy.md#GetPolicyTemplates) | **Get** /data-protect/policy-templates | List Policy Templates filtered by query parameters.
[**GetProtectionPolicies**](Policy.md#GetProtectionPolicies) | **Get** /data-protect/policies | List Protection Policies based on provided filtering parameters.
[**GetProtectionPolicyById**](Policy.md#GetProtectionPolicyById) | **Get** /data-protect/policies/{id} | List details about a single Protection Policy.
[**UpdateHeliosPolicy**](Policy.md#UpdateHeliosPolicy) | **Put** /mcm/data-protect/policies/{id} | Update a Protection Policy.
[**UpdateProtectionPolicy**](Policy.md#UpdateProtectionPolicy) | **Put** /data-protect/policies/{id} | Update a Protection Policy.



## CreateHeliosPolicy

> HeliosPolicyResponse CreateHeliosPolicy(ctx).Body(body).RegionId(regionId).Execute()

Create a Policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    body := *openapiclient.NewHeliosPolicyRequest("Name_example", "Type_example") // HeliosPolicyRequest | Request to create a Helios Policy.
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiCreateHeliosPolicyRequest{
        Body: &body
        RegionId: &regionId
    }

    resp, r, err := api_client.Policy.CreateHeliosPolicy(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Policy.CreateHeliosPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHeliosPolicy`: HeliosPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `Policy.CreateHeliosPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHeliosPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HeliosPolicyRequest**](HeliosPolicyRequest.md) | Request to create a Helios Policy. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**HeliosPolicyResponse**](HeliosPolicyResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProtectionPolicy

> ProtectionPolicyResponse CreateProtectionPolicy(ctx).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Create a Protection Policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    body := *openapiclient.NewProtectionPolicyRequest("Name_example", *openapiclient.NewBackupPolicy(*openapiclient.NewRegularBackupPolicy(*openapiclient.NewRetention("Unit_example", NullableInt64(123))))) // ProtectionPolicyRequest | Request to create a Protection Policy.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiCreateProtectionPolicyRequest{
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Policy.CreateProtectionPolicy(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Policy.CreateProtectionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProtectionPolicy`: ProtectionPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `Policy.CreateProtectionPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProtectionPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProtectionPolicyRequest**](ProtectionPolicyRequest.md) | Request to create a Protection Policy. | 
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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


## DeleteHeliosPolicy

> DeleteHeliosPolicy(ctx, id).RegionId(regionId).Execute()

Delete a Policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the Policy to delete.
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiDeleteHeliosPolicyRequest{
        Id: &id
        RegionId: &regionId
    }

    resp, r, err := api_client.Policy.DeleteHeliosPolicy(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Policy.DeleteHeliosPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the Policy to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHeliosPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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


## DeleteProtectionPolicy

> DeleteProtectionPolicy(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Delete a Protection Policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the Protection Policy to delete.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiDeleteProtectionPolicyRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Policy.DeleteProtectionPolicy(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Policy.DeleteProtectionPolicy``: %v\n", err)
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

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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


## GetHeliosPolicies

> HeliosPoliciesResponseWithPagination GetHeliosPolicies(ctx).RegionId(regionId).Ids(ids).PolicyNames(policyNames).Types(types).ClusterIdentifiers(clusterIdentifiers).TenantIds(tenantIds).Execute()

List Policies based on provided filtering parameters.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    ids := []string{"Inner_example"} // []string | Filter policies by a list of policy ids. (optional)
    policyNames := []string{"Inner_example"} // []string | Filter policies by a list of policy names. (optional)
    types := []string{"Types_example"} // []string | Type specifies the policy type of policies to be returned. If not specified, all types of policies are fetched. (optional)
    clusterIdentifiers := []string{"Inner_example"} // []string | Specifies the list of cluster identifiers. This is applicable only for type OnPremPolicy. The format is clusterId:clusterIncarnationId. (optional)
    tenantIds := []string{"Inner_example"} // []string | List of Tenant Ids to filter from. This is applicable only for type OnPremPolicy. (optional)

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

    request := helios.ApiGetHeliosPoliciesRequest{
        RegionId: &regionId
        Ids: &ids
        PolicyNames: &policyNames
        Types: &types
        ClusterIdentifiers: &clusterIdentifiers
        TenantIds: &tenantIds
    }

    resp, r, err := api_client.Policy.GetHeliosPolicies(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Policy.GetHeliosPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHeliosPolicies`: HeliosPoliciesResponseWithPagination
    fmt.Fprintf(os.Stdout, "Response from `Policy.GetHeliosPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHeliosPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **ids** | **[]string** | Filter policies by a list of policy ids. | 
 **policyNames** | **[]string** | Filter policies by a list of policy names. | 
 **types** | **[]string** | Type specifies the policy type of policies to be returned. If not specified, all types of policies are fetched. | 
 **clusterIdentifiers** | **[]string** | Specifies the list of cluster identifiers. This is applicable only for type OnPremPolicy. The format is clusterId:clusterIncarnationId. | 
 **tenantIds** | **[]string** | List of Tenant Ids to filter from. This is applicable only for type OnPremPolicy. | 

### Return type

[**HeliosPoliciesResponseWithPagination**](HeliosPoliciesResponseWithPagination.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHeliosPolicyById

> HeliosPolicyResponse GetHeliosPolicyById(ctx, id).RegionId(regionId).Execute()

List details about a single Protection Policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the Protection Policy to return.
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiGetHeliosPolicyByIdRequest{
        Id: &id
        RegionId: &regionId
    }

    resp, r, err := api_client.Policy.GetHeliosPolicyById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Policy.GetHeliosPolicyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHeliosPolicyById`: HeliosPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `Policy.GetHeliosPolicyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the Protection Policy to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHeliosPolicyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**HeliosPolicyResponse**](HeliosPolicyResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyTemplateById

> PolicyTemplateResponse GetPolicyTemplateById(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

List details about a single Policy Template.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the Policy Template to return.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiGetPolicyTemplateByIdRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Policy.GetPolicyTemplateById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Policy.GetPolicyTemplateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyTemplateById`: PolicyTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `Policy.GetPolicyTemplateById`: %v\n", resp)
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

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

> PolicyTemplatesResponseWithPagination GetPolicyTemplates(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Ids(ids).PolicyNames(policyNames).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()

List Policy Templates filtered by query parameters.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    ids := []string{"Inner_example"} // []string | Filter policies by a list of policy template ids. (optional)
    policyNames := []string{"Inner_example"} // []string | Filter policies by a list of policy names. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the organizations for which objects are to be returned. (optional)
    includeTenants := true // bool | IncludeTenantPolicies specifies if objects of all the organizations under the hierarchy of the logged in user's organization should be returned. (optional)

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

    request := helios.ApiGetPolicyTemplatesRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        Ids: &ids
        PolicyNames: &policyNames
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
    }

    resp, r, err := api_client.Policy.GetPolicyTemplates(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Policy.GetPolicyTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyTemplates`: PolicyTemplatesResponseWithPagination
    fmt.Fprintf(os.Stdout, "Response from `Policy.GetPolicyTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
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

> ProtectionPolicyResponseWithPagination GetProtectionPolicies(ctx).AccessClusterId(accessClusterId).RegionId(regionId).Ids(ids).PolicyNames(policyNames).TenantIds(tenantIds).IncludeTenants(includeTenants).Types(types).ExcludeLinkedPolicies(excludeLinkedPolicies).IncludeReplicatedPolicies(includeReplicatedPolicies).Execute()

List Protection Policies based on provided filtering parameters.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)
    ids := []string{"Inner_example"} // []string | Filter policies by a list of policy ids. (optional)
    policyNames := []string{"Inner_example"} // []string | Filter policies by a list of policy names. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the organizations for which objects are to be returned. (optional)
    includeTenants := true // bool | IncludeTenantPolicies specifies if objects of all the organizations under the hierarchy of the logged in user's organization should be returned. (optional)
    types := []string{"Types_example"} // []string | Types specifies the policy type of policies to be returned (optional) (default to ["Regular"])
    excludeLinkedPolicies := true // bool | If excludeLinkedPolicies is set to true then only local policies created on cluster will be returned. The result will exclude all linked policies created from policy templates. (optional)
    includeReplicatedPolicies := true // bool | If includeReplicatedPolicies is set to true, then response will also contain replicated policies. By default, replication policies are not included in the response. (optional)

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

    request := helios.ApiGetProtectionPoliciesRequest{
        AccessClusterId: &accessClusterId
        RegionId: &regionId
        Ids: &ids
        PolicyNames: &policyNames
        TenantIds: &tenantIds
        IncludeTenants: &includeTenants
        Types: &types
        ExcludeLinkedPolicies: &excludeLinkedPolicies
        IncludeReplicatedPolicies: &includeReplicatedPolicies
    }

    resp, r, err := api_client.Policy.GetProtectionPolicies(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Policy.GetProtectionPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionPolicies`: ProtectionPolicyResponseWithPagination
    fmt.Fprintf(os.Stdout, "Response from `Policy.GetProtectionPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectionPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 
 **ids** | **[]string** | Filter policies by a list of policy ids. | 
 **policyNames** | **[]string** | Filter policies by a list of policy names. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the organizations for which objects are to be returned. | 
 **includeTenants** | **bool** | IncludeTenantPolicies specifies if objects of all the organizations under the hierarchy of the logged in user&#39;s organization should be returned. | 
 **types** | **[]string** | Types specifies the policy type of policies to be returned | [default to [&quot;Regular&quot;]]
 **excludeLinkedPolicies** | **bool** | If excludeLinkedPolicies is set to true then only local policies created on cluster will be returned. The result will exclude all linked policies created from policy templates. | 
 **includeReplicatedPolicies** | **bool** | If includeReplicatedPolicies is set to true, then response will also contain replicated policies. By default, replication policies are not included in the response. | 

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

> ProtectionPolicyResponse GetProtectionPolicyById(ctx, id).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

List details about a single Protection Policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the Protection Policy to return.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiGetProtectionPolicyByIdRequest{
        Id: &id
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Policy.GetProtectionPolicyById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Policy.GetProtectionPolicyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProtectionPolicyById`: ProtectionPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `Policy.GetProtectionPolicyById`: %v\n", resp)
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

 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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


## UpdateHeliosPolicy

> HeliosPolicyResponse UpdateHeliosPolicy(ctx, id).Body(body).RegionId(regionId).Execute()

Update a Protection Policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the Protection Policy to update.
    body := *openapiclient.NewHeliosPolicyRequest("Name_example", "Type_example") // HeliosPolicyRequest | Request to update a Protection Policy.
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiUpdateHeliosPolicyRequest{
        Id: &id
        Body: &body
        RegionId: &regionId
    }

    resp, r, err := api_client.Policy.UpdateHeliosPolicy(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Policy.UpdateHeliosPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHeliosPolicy`: HeliosPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `Policy.UpdateHeliosPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Specifies a unique id of the Protection Policy to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHeliosPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HeliosPolicyRequest**](HeliosPolicyRequest.md) | Request to update a Protection Policy. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

### Return type

[**HeliosPolicyResponse**](HeliosPolicyResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProtectionPolicy

> ProtectionPolicyResponse UpdateProtectionPolicy(ctx, id).Body(body).AccessClusterId(accessClusterId).RegionId(regionId).Execute()

Update a Protection Policy.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    helios "helios"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the Protection Policy to update.
    body := *openapiclient.NewProtectionPolicyRequest("Name_example", *openapiclient.NewBackupPolicy(*openapiclient.NewRegularBackupPolicy(*openapiclient.NewRetention("Unit_example", NullableInt64(123))))) // ProtectionPolicyRequest | Request to update a Protection Policy.
    accessClusterId := int64(789) // int64 | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. (optional)
    regionId := "regionId_example" // string | This field uniquely represents a region and is used for making Helios calls to a specific region. (optional)

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

    request := helios.ApiUpdateProtectionPolicyRequest{
        Id: &id
        Body: &body
        AccessClusterId: &accessClusterId
        RegionId: &regionId
    }

    resp, r, err := api_client.Policy.UpdateProtectionPolicy(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Policy.UpdateProtectionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProtectionPolicy`: ProtectionPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `Policy.UpdateProtectionPolicy`: %v\n", resp)
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
 **accessClusterId** | **int64** | This field uniquely represents a Cohesity Cluster and is used for making on-prem calls from Helios. | 
 **regionId** | **string** | This field uniquely represents a region and is used for making Helios calls to a specific region. | 

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

