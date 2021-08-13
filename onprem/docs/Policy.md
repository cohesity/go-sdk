# \Policy

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProtectionPolicy**](Policy.md#CreateProtectionPolicy) | **Post** /data-protect/policies | Create a Protection Policy.
[**DeleteProtectionPolicy**](Policy.md#DeleteProtectionPolicy) | **Delete** /data-protect/policies/{id} | Delete a Protection Policy.
[**GetPolicyTemplateById**](Policy.md#GetPolicyTemplateById) | **Get** /data-protect/policy-templates/{id} | List details about a single Policy Template.
[**GetPolicyTemplates**](Policy.md#GetPolicyTemplates) | **Get** /data-protect/policy-templates | List Policy Templates filtered by query parameters.
[**GetProtectionPolicies**](Policy.md#GetProtectionPolicies) | **Get** /data-protect/policies | List Protection Policies based on provided filtering parameters.
[**GetProtectionPolicyById**](Policy.md#GetProtectionPolicyById) | **Get** /data-protect/policies/{id} | List details about a single Protection Policy.
[**UpdateProtectionPolicy**](Policy.md#UpdateProtectionPolicy) | **Put** /data-protect/policies/{id} | Update a Protection Policy.



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
    onprem "onprem"
)

func main() {
    body := *openapiclient.NewProtectionPolicyRequest("Name_example", *openapiclient.NewBackupPolicy(*openapiclient.NewRegularBackupPolicy(*openapiclient.NewRetention("Unit_example", NullableInt64(123))))) // ProtectionPolicyRequest | Request to create a Protection Policy.

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

    request := onprem.ApiCreateProtectionPolicyRequest{
        Body: &body
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
    onprem "onprem"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the Protection Policy to delete.

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

    request := onprem.ApiDeleteProtectionPolicyRequest{
        Id: &id
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
    onprem "onprem"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the Policy Template to return.

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

    request := onprem.ApiGetPolicyTemplateByIdRequest{
        Id: &id
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

> PolicyTemplatesResponseWithPagination GetPolicyTemplates(ctx).Ids(ids).PolicyNames(policyNames).TenantIds(tenantIds).IncludeTenants(includeTenants).Execute()

List Policy Templates filtered by query parameters.



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

    request := onprem.ApiGetPolicyTemplatesRequest{
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

> ProtectionPolicyResponseWithPagination GetProtectionPolicies(ctx).Ids(ids).PolicyNames(policyNames).TenantIds(tenantIds).IncludeTenants(includeTenants).Types(types).ExcludeLinkedPolicies(excludeLinkedPolicies).IncludeReplicatedPolicies(includeReplicatedPolicies).Execute()

List Protection Policies based on provided filtering parameters.



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

    request := onprem.ApiGetProtectionPoliciesRequest{
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

> ProtectionPolicyResponse GetProtectionPolicyById(ctx, id).Execute()

List details about a single Protection Policy.



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

    request := onprem.ApiGetProtectionPolicyByIdRequest{
        Id: &id
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
    onprem "onprem"
)

func main() {
    id := "id_example" // string | Specifies a unique id of the Protection Policy to update.
    body := *openapiclient.NewProtectionPolicyRequest("Name_example", *openapiclient.NewBackupPolicy(*openapiclient.NewRegularBackupPolicy(*openapiclient.NewRetention("Unit_example", NullableInt64(123))))) // ProtectionPolicyRequest | Request to update a Protection Policy.

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

    request := onprem.ApiUpdateProtectionPolicyRequest{
        Id: &id
        Body: &body
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

