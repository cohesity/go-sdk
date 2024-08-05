# \FirewallAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListFirewallSettings**](FirewallAPI.md#ListFirewallSettings) | **Get** /network/firewall | List all firewall settings.
[**RemoveFirewallProfiles**](FirewallAPI.md#RemoveFirewallProfiles) | **Put** /network/firewall/profile/remove | Remove firewall profiles.
[**UpdateFirewallProfile**](FirewallAPI.md#UpdateFirewallProfile) | **Put** /network/firewall/profile | Update firewall profiles &amp; their attachments.



## ListFirewallSettings

> FirewallEntry ListFirewallSettings(ctx).Execute()

List all firewall settings.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallAPI.ListFirewallSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.ListFirewallSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFirewallSettings`: FirewallEntry
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.ListFirewallSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFirewallSettingsRequest struct via the builder pattern


### Return type

[**FirewallEntry**](FirewallEntry.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveFirewallProfiles

> SuccessResp RemoveFirewallProfiles(ctx).Body(body).Execute()

Remove firewall profiles.



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
	body := *openapiclient.NewFirewallProfileNamesParams([]string{"Names_example"}) // FirewallProfileNamesParams | Specifies the parameters to remove firewall profiles and their attachments.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallAPI.RemoveFirewallProfiles(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.RemoveFirewallProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveFirewallProfiles`: SuccessResp
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.RemoveFirewallProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFirewallProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FirewallProfileNamesParams**](FirewallProfileNamesParams.md) | Specifies the parameters to remove firewall profiles and their attachments. | 

### Return type

[**SuccessResp**](SuccessResp.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFirewallProfile

> FirewallProfileParams UpdateFirewallProfile(ctx).Body(body).Execute()

Update firewall profiles & their attachments.



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
	body := *openapiclient.NewFirewallProfileParams("Action_example", "Name_example") // FirewallProfileParams | Specifies the parameters to configure firewall profiles and/or their attachments.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallAPI.UpdateFirewallProfile(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.UpdateFirewallProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFirewallProfile`: FirewallProfileParams
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.UpdateFirewallProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFirewallProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FirewallProfileParams**](FirewallProfileParams.md) | Specifies the parameters to configure firewall profiles and/or their attachments. | 

### Return type

[**FirewallProfileParams**](FirewallProfileParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

