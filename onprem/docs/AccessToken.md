# \AccessToken

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessToken**](AccessToken.md#CreateAccessToken) | **Post** /accessTokens | Create a new API access token



## CreateAccessToken

> AccessTokenResponse CreateAccessToken(ctx).Body(body).Execute()

Create a new API access token



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
    body := *openapiclient.NewCreateAccessTokenRequestParams("Username_example", "Password_example") // CreateAccessTokenRequestParams | Specifies the parameters to generate an access token

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

    request := onprem.ApiCreateAccessTokenRequest{
        Body: &body
    }

    resp, r, err := api_client.AccessToken.CreateAccessToken(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessToken.CreateAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccessToken`: AccessTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessToken.CreateAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateAccessTokenRequestParams**](CreateAccessTokenRequestParams.md) | Specifies the parameters to generate an access token | 

### Return type

[**AccessTokenResponse**](AccessTokenResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

