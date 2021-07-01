# \AccessTokensApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateAccessToken**](AccessTokensApi.md#GenerateAccessToken) | **Post** /public/accessTokens | Generate an Access Token.



## GenerateAccessToken

> AccessToken GenerateAccessToken(ctx).Body(body).Execute()

Generate an Access Token.



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
    body := *openapiclient.NewAccessTokenCredential("Password_example", "Username_example") // AccessTokenCredential | Request to generate access token.

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

    request := cohesitysdk.ApiGenerateAccessTokenRequest{
        body: &Body
    }

    resp, r, err := api_client.AccessTokensApi.GenerateAccessToken(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensApi.GenerateAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateAccessToken`: AccessToken
    fmt.Fprintf(os.Stdout, "Response from `AccessTokensApi.GenerateAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AccessTokenCredential**](AccessTokenCredential.md) | Request to generate access token. | 

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

