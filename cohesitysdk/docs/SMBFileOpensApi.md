# \SMBFileOpensApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseSmbFileOpen**](SMBFileOpensApi.md#CloseSmbFileOpen) | **Post** /public/smbFileOpens | Closes an active SMB file open.
[**GetSmbFileOpens**](SMBFileOpensApi.md#GetSmbFileOpens) | **Get** /public/smbFileOpens | List the active SMB file opens that match the filter criteria specified using parameters.



## CloseSmbFileOpen

> CloseSmbFileOpen(ctx).Body(body).Execute()

Closes an active SMB file open.



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
    body := *openapiclient.NewCloseSmbFileOpenParameters() // CloseSmbFileOpenParameters | Request to close an active SMB file open.

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

    request := cohesitysdk.ApiCloseSmbFileOpenRequest{
        Body: &body
    }

    resp, r, err := api_client.SMBFileOpensApi.CloseSmbFileOpen(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMBFileOpensApi.CloseSmbFileOpen``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloseSmbFileOpenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CloseSmbFileOpenParameters**](CloseSmbFileOpenParameters.md) | Request to close an active SMB file open. | 

### Return type

 (empty response body)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmbFileOpens

> SmbActiveFileOpensResponse GetSmbFileOpens(ctx).FilePath(filePath).ViewName(viewName).PageCount(pageCount).Cookie(cookie).Execute()

List the active SMB file opens that match the filter criteria specified using parameters.



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
    filePath := "filePath_example" // string | Specifies the filepath in the view relative to the root filesystem. If this field is specified, viewName field must also be specified. (optional)
    viewName := "viewName_example" // string | Specifies the name of the View in which to search. If a view name is not specified, all the views in the Cluster is searched. This field is mandatory if filePath field is specified. (optional)
    pageCount := int32(56) // int32 | Specifies the maximum number of active opens to return in the response. This field cannot be set above 1000. If this is not set, maximum of 1000 entries are returned. (optional)
    cookie := "cookie_example" // string | Specifies the opaque string returned in the previous response. If this is set, next set of active opens just after the previous response are returned. If this is not set, first set of active opens are returned. (optional)

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

    request := cohesitysdk.ApiGetSmbFileOpensRequest{
        FilePath: &filePath
        ViewName: &viewName
        PageCount: &pageCount
        Cookie: &cookie
    }

    resp, r, err := api_client.SMBFileOpensApi.GetSmbFileOpens(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMBFileOpensApi.GetSmbFileOpens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmbFileOpens`: SmbActiveFileOpensResponse
    fmt.Fprintf(os.Stdout, "Response from `SMBFileOpensApi.GetSmbFileOpens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSmbFileOpensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filePath** | **string** | Specifies the filepath in the view relative to the root filesystem. If this field is specified, viewName field must also be specified. | 
 **viewName** | **string** | Specifies the name of the View in which to search. If a view name is not specified, all the views in the Cluster is searched. This field is mandatory if filePath field is specified. | 
 **pageCount** | **int32** | Specifies the maximum number of active opens to return in the response. This field cannot be set above 1000. If this is not set, maximum of 1000 entries are returned. | 
 **cookie** | **string** | Specifies the opaque string returned in the previous response. If this is set, next set of active opens just after the previous response are returned. If this is not set, first set of active opens are returned. | 

### Return type

[**SmbActiveFileOpensResponse**](SmbActiveFileOpensResponse.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

