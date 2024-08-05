# \RoutesAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddStaticRoute**](RoutesAPI.md#AddStaticRoute) | **Post** /network/routes | Configure a static route



## AddStaticRoute

> StaticRouteParams AddStaticRoute(ctx).Body(body).Execute()

Configure a static route



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
	body := *openapiclient.NewStaticRouteParams("DestinationNetwork_example", "InterfaceGroup_example", "NextHop_example") // StaticRouteParams | Specifies the parameters to configure a static route on an interface.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.AddStaticRoute(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.AddStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddStaticRoute`: StaticRouteParams
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.AddStaticRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**StaticRouteParams**](StaticRouteParams.md) | Specifies the parameters to configure a static route on an interface. | 

### Return type

[**StaticRouteParams**](StaticRouteParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

