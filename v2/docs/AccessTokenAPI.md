# \AccessTokenAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessToken**](AccessTokenAPI.md#CreateAccessToken) | **Post** /access-tokens | Create a new API access token



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
	openapiclient "github.com/cohesity/go-sdk"
)

func main() {
	body := *openapiclient.NewCreateAccessTokenRequestParams() // CreateAccessTokenRequestParams | Specifies the parameters to generate an access token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessTokenAPI.CreateAccessToken(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessTokenAPI.CreateAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccessToken`: AccessTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessTokenAPI.CreateAccessToken`: %v\n", resp)
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

