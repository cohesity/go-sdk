# \AlertAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlertSummary**](AlertAPI.md#GetAlertSummary) | **Get** /alerts-summary | Get alerts summary.
[**GetAlerts**](AlertAPI.md#GetAlerts) | **Get** /alerts | Get alerts.



## GetAlertSummary

> AlertsSummaryResponse GetAlertSummary(ctx).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).IncludeTenants(includeTenants).TenantIds(tenantIds).StatesList(statesList).Execute()

Get alerts summary.



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
	startTimeUsecs := int64(789) // int64 | Filter by start time. Specify the start time as a Unix epoch Timestamp (in microseconds). By default it is current time minus a day. (optional)
	endTimeUsecs := int64(789) // int64 | Filter by end time. Specify the end time as a Unix epoch Timestamp (in microseconds). By default it is current time. (optional)
	includeTenants := true // bool | IncludeTenants specifies if alerts of all the tenants under the hierarchy of the logged in user's organization should be used to compute summary. (optional)
	tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which alerts are to be used to compute summary. (optional)
	statesList := []string{"StatesList_example"} // []string | Specifies list of alert states to filter alerts by. If not specified, only open alerts will be used to get summary. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertAPI.GetAlertSummary(context.Background()).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).IncludeTenants(includeTenants).TenantIds(tenantIds).StatesList(statesList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.GetAlertSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertSummary`: AlertsSummaryResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertAPI.GetAlertSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTimeUsecs** | **int64** | Filter by start time. Specify the start time as a Unix epoch Timestamp (in microseconds). By default it is current time minus a day. | 
 **endTimeUsecs** | **int64** | Filter by end time. Specify the end time as a Unix epoch Timestamp (in microseconds). By default it is current time. | 
 **includeTenants** | **bool** | IncludeTenants specifies if alerts of all the tenants under the hierarchy of the logged in user&#39;s organization should be used to compute summary. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which alerts are to be used to compute summary. | 
 **statesList** | **[]string** | Specifies list of alert states to filter alerts by. If not specified, only open alerts will be used to get summary. | 

### Return type

[**AlertsSummaryResponse**](AlertsSummaryResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlerts

> AlertList GetAlerts(ctx).AlertIds(alertIds).AlertTypes(alertTypes).AlertCategories(alertCategories).AlertStates(alertStates).AlertSeverities(alertSeverities).AlertTypeBuckets(alertTypeBuckets).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).MaxAlerts(maxAlerts).PropertyKey(propertyKey).PropertyValue(propertyValue).AlertName(alertName).ResolutionIds(resolutionIds).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).Execute()

Get alerts.

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
	alertIds := []string{"Inner_example"} // []string | Filter by list of alert ids. (optional)
	alertTypes := []int32{int32(123)} // []int32 | Filter by list of alert types. (optional)
	alertCategories := []string{"AlertCategories_example"} // []string | Filter by list of alert categories. (optional)
	alertStates := []string{"AlertStates_example"} // []string | Filter by list of alert states. (optional)
	alertSeverities := []string{"AlertSeverities_example"} // []string | Filter by list of alert severity types. (optional)
	alertTypeBuckets := []string{"AlertTypeBuckets_example"} // []string | Filter by list of alert type buckets. (optional)
	startTimeUsecs := int64(789) // int64 | Specifies start time Unix epoch time in microseconds to filter alerts by. (optional)
	endTimeUsecs := int64(789) // int64 | Specifies end time Unix epoch time in microseconds to filter alerts by. (optional)
	maxAlerts := int32(56) // int32 | Specifies maximum number of alerts to return.The default value is 100 and maximum allowed value is 1000 (optional)
	propertyKey := "propertyKey_example" // string | Specifies name of the property to filter alerts by. (optional)
	propertyValue := "propertyValue_example" // string | Specifies value of the property to filter alerts by. (optional)
	alertName := "alertName_example" // string | Specifies name of alert to filter alerts by. (optional)
	resolutionIds := []int64{int64(123)} // []int64 | Specifies alert resolution ids to filter alerts by. (optional)
	tenantIds := []string{"Inner_example"} // []string | Filter by tenant ids. (optional)
	allUnderHierarchy := true // bool | Filter by objects of all the tenants under the hierarchy of the logged in user's organization. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertAPI.GetAlerts(context.Background()).AlertIds(alertIds).AlertTypes(alertTypes).AlertCategories(alertCategories).AlertStates(alertStates).AlertSeverities(alertSeverities).AlertTypeBuckets(alertTypeBuckets).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).MaxAlerts(maxAlerts).PropertyKey(propertyKey).PropertyValue(propertyValue).AlertName(alertName).ResolutionIds(resolutionIds).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.GetAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlerts`: AlertList
	fmt.Fprintf(os.Stdout, "Response from `AlertAPI.GetAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertIds** | **[]string** | Filter by list of alert ids. | 
 **alertTypes** | **[]int32** | Filter by list of alert types. | 
 **alertCategories** | **[]string** | Filter by list of alert categories. | 
 **alertStates** | **[]string** | Filter by list of alert states. | 
 **alertSeverities** | **[]string** | Filter by list of alert severity types. | 
 **alertTypeBuckets** | **[]string** | Filter by list of alert type buckets. | 
 **startTimeUsecs** | **int64** | Specifies start time Unix epoch time in microseconds to filter alerts by. | 
 **endTimeUsecs** | **int64** | Specifies end time Unix epoch time in microseconds to filter alerts by. | 
 **maxAlerts** | **int32** | Specifies maximum number of alerts to return.The default value is 100 and maximum allowed value is 1000 | 
 **propertyKey** | **string** | Specifies name of the property to filter alerts by. | 
 **propertyValue** | **string** | Specifies value of the property to filter alerts by. | 
 **alertName** | **string** | Specifies name of alert to filter alerts by. | 
 **resolutionIds** | **[]int64** | Specifies alert resolution ids to filter alerts by. | 
 **tenantIds** | **[]string** | Filter by tenant ids. | 
 **allUnderHierarchy** | **bool** | Filter by objects of all the tenants under the hierarchy of the logged in user&#39;s organization. | 

### Return type

[**AlertList**](AlertList.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

