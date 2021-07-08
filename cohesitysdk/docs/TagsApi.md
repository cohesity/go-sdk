# \TagsApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyTags**](TagsApi.md#ApplyTags) | **Put** /public/tags/apply | Apply Tags on existing objects in the Cohesity Cluster.
[**RemoveTags**](TagsApi.md#RemoveTags) | **Put** /public/tags/remove | Remove Tags on existing objects in the Cohesity Cluster.



## ApplyTags

> TagsOperationResult ApplyTags(ctx).Body(body).Execute()

Apply Tags on existing objects in the Cohesity Cluster.



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
    body := *openapiclient.NewTagsOperationParameters() // TagsOperationParameters | Request to add or update document tags. (optional)

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

    request := cohesitysdk.ApiApplyTagsRequest{
        Body: &body
    }

    resp, r, err := api_client.TagsApi.ApplyTags(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.ApplyTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplyTags`: TagsOperationResult
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.ApplyTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TagsOperationParameters**](TagsOperationParameters.md) | Request to add or update document tags. | 

### Return type

[**TagsOperationResult**](TagsOperationResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTags

> TagsOperationResult RemoveTags(ctx).Body(body).Execute()

Remove Tags on existing objects in the Cohesity Cluster.



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
    body := *openapiclient.NewTagsOperationParameters() // TagsOperationParameters | Request to add or update document tags. (optional)

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

    request := cohesitysdk.ApiRemoveTagsRequest{
        Body: &body
    }

    resp, r, err := api_client.TagsApi.RemoveTags(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.RemoveTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveTags`: TagsOperationResult
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.RemoveTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TagsOperationParameters**](TagsOperationParameters.md) | Request to add or update document tags. | 

### Return type

[**TagsOperationResult**](TagsOperationResult.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

