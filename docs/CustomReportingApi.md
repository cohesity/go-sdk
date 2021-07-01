# \CustomReportingApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPostgres**](CustomReportingApi.md#GetPostgres) | **Get** /public/postgres | List the postgres database running nodes on the cohesity cluster.



## GetPostgres

> []PostgresNodeInfo GetPostgres(ctx).Execute()

List the postgres database running nodes on the cohesity cluster.



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

    request := cohesitysdk.ApiGetPostgresRequest{
    }

    resp, r, err := api_client.CustomReportingApi.GetPostgres(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomReportingApi.GetPostgres``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostgres`: []PostgresNodeInfo
    fmt.Fprintf(os.Stdout, "Response from `CustomReportingApi.GetPostgres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostgresRequest struct via the builder pattern


### Return type

[**[]PostgresNodeInfo**](PostgresNodeInfo.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

