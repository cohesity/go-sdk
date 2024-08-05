# \PrivilegeAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPrivileges**](PrivilegeAPI.md#GetPrivileges) | **Get** /privileges | Get Privileges.



## GetPrivileges

> Privileges GetPrivileges(ctx).Names(names).Execute()

Get Privileges.



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
	names := []string{"Inner_example"} // []string | Filter by a list of Privilege names. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivilegeAPI.GetPrivileges(context.Background()).Names(names).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivilegeAPI.GetPrivileges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivileges`: Privileges
	fmt.Fprintf(os.Stdout, "Response from `PrivilegeAPI.GetPrivileges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivilegesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | **[]string** | Filter by a list of Privilege names. | 

### Return type

[**Privileges**](Privileges.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

