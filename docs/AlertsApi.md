# \AlertsApi

All URIs are relative to *http://localhost/irisservices/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationRule**](AlertsApi.md#CreateNotificationRule) | **Post** /public/alertNotificationRules | Creates a new alert notification rule.
[**CreateResolution**](AlertsApi.md#CreateResolution) | **Post** /public/alertResolutions | Create an Alert Resolution.
[**DeleteNotificationRule**](AlertsApi.md#DeleteNotificationRule) | **Delete** /public/alertNotificationRules/{ruleId} | Deletes an alert notification rule.
[**GetAlertById**](AlertsApi.md#GetAlertById) | **Get** /public/alerts/{id} | List details about a single Alert.
[**GetAlertCategories**](AlertsApi.md#GetAlertCategories) | **Get** /public/alertCategories | Get alert categories in the Cohesity cluster.
[**GetAlertTypes**](AlertsApi.md#GetAlertTypes) | **Get** /public/alertTypes | Get registered alerts in the Cohesity cluster.
[**GetAlerts**](AlertsApi.md#GetAlerts) | **Get** /public/alerts | List the Alerts on the Cohesity Cluster.
[**GetNotificationRules**](AlertsApi.md#GetNotificationRules) | **Get** /public/alertNotificationRules | Gets all alert notification rules.
[**GetResolutionById**](AlertsApi.md#GetResolutionById) | **Get** /public/alertResolutions/{id} | List details about a single Alert Resolution.
[**GetResolutions**](AlertsApi.md#GetResolutions) | **Get** /public/alertResolutions | List the Alert Resolutions on the Cohesity Cluster.
[**UpdateNotificationRule**](AlertsApi.md#UpdateNotificationRule) | **Put** /public/alertNotificationRules | Updates an existing alert notification rule.
[**UpdateResolution**](AlertsApi.md#UpdateResolution) | **Put** /public/alertResolutions/{id} | Apply an existing Alert Resolution to additional Alerts.



## CreateNotificationRule

> NotificationRule CreateNotificationRule(ctx).Body(body).Execute()

Creates a new alert notification rule.



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
    body := *openapiclient.NewNotificationRule() // NotificationRule | Create Notification Rule argument. (optional)

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

    request := cohesitysdk.ApiCreateNotificationRuleRequest{
        body: &Body
    }

    resp, r, err := api_client.AlertsApi.CreateNotificationRule(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.CreateNotificationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationRule`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.CreateNotificationRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NotificationRule**](NotificationRule.md) | Create Notification Rule argument. | 

### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateResolution

> AlertResolution CreateResolution(ctx).Body(body).Execute()

Create an Alert Resolution.



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
    body := *openapiclient.NewAlertResolutionRequest() // AlertResolutionRequest | Request to create an Alert Resolution and apply it to the specified Alerts.

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

    request := cohesitysdk.ApiCreateResolutionRequest{
        body: &Body
    }

    resp, r, err := api_client.AlertsApi.CreateResolution(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.CreateResolution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateResolution`: AlertResolution
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.CreateResolution`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateResolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AlertResolutionRequest**](AlertResolutionRequest.md) | Request to create an Alert Resolution and apply it to the specified Alerts. | 

### Return type

[**AlertResolution**](AlertResolution.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationRule

> DeleteNotificationRule(ctx, ruleId).Execute()

Deletes an alert notification rule.



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
    ruleId := int64(789) // int64 | Specifies the rule id.

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

    request := cohesitysdk.ApiDeleteNotificationRuleRequest{
        ruleId: &RuleId
    }

    resp, r, err := api_client.AlertsApi.DeleteNotificationRule(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.DeleteNotificationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **int64** | Specifies the rule id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertById

> Alert GetAlertById(ctx, id).Execute()

List details about a single Alert.



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
    id := "id_example" // string | Unique id of the Alert to return.

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

    request := cohesitysdk.ApiGetAlertByIdRequest{
        id: &Id
    }

    resp, r, err := api_client.AlertsApi.GetAlertById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetAlertById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertById`: Alert
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetAlertById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique id of the Alert to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Alert**](Alert.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertCategories

> []AlertCategoryName GetAlertCategories(ctx).Execute()

Get alert categories in the Cohesity cluster.



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

    request := cohesitysdk.ApiGetAlertCategoriesRequest{
    }

    resp, r, err := api_client.AlertsApi.GetAlertCategories(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetAlertCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertCategories`: []AlertCategoryName
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetAlertCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertCategoriesRequest struct via the builder pattern


### Return type

[**[]AlertCategoryName**](AlertCategoryName.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertTypes

> []AlertMetadata GetAlertTypes(ctx).Body(body).Execute()

Get registered alerts in the Cohesity cluster.



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
    body := *openapiclient.NewGetAlertTypesParams() // GetAlertTypesParams | Get alert types params.

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

    request := cohesitysdk.ApiGetAlertTypesRequest{
        body: &Body
    }

    resp, r, err := api_client.AlertsApi.GetAlertTypes(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetAlertTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertTypes`: []AlertMetadata
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetAlertTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GetAlertTypesParams**](GetAlertTypesParams.md) | Get alert types params. | 

### Return type

[**[]AlertMetadata**](AlertMetadata.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlerts

> []Alert GetAlerts(ctx).MaxAlerts(maxAlerts).AlertIdList(alertIdList).AlertTypeList(alertTypeList).AlertCategoryList(alertCategoryList).PropertyKey(propertyKey).PropertyValue(propertyValue).StartDateUsecs(startDateUsecs).EndDateUsecs(endDateUsecs).AlertStateList(alertStateList).AlertSeverityList(alertSeverityList).AlertTypeBucketList(alertTypeBucketList).ResolutionIdList(resolutionIdList).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).Execute()

List the Alerts on the Cohesity Cluster.



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
    maxAlerts := int32(56) // int32 | Specifies the number of returned Alerts to be returned. The newest Alerts are returned.
    alertIdList := []string{"Inner_example"} // []string | Specifies list of Alert ids to filter alerts by. (optional)
    alertTypeList := []int32{int32(123)} // []int32 | Specifies list of Alert Types to filter alerts by. (optional)
    alertCategoryList := []string{"AlertCategoryList_example"} // []string | Specifies list of Alert Categories. (optional)
    propertyKey := "propertyKey_example" // string | Specifies name of the property to filter alerts by. (optional)
    propertyValue := "propertyValue_example" // string | Specifies value of the property to filter alerts by. (optional)
    startDateUsecs := int64(789) // int64 | Specifies start time Unix epoch time in microseconds to filter alerts by. (optional)
    endDateUsecs := int64(789) // int64 | Specifies end time Unix epoch time in microseconds to filter alerts by. (optional)
    alertStateList := []string{"AlertStateList_example"} // []string | Specifies list of Alert States to filter alerts by. (optional)
    alertSeverityList := []string{"AlertSeverityList_example"} // []string | Specifies list of Alert severity to filter alerts by. (optional)
    alertTypeBucketList := []string{"AlertTypeBucketList_example"} // []string | Specifies the list of Alert type bucket. Specifies the Alert type bucket. kSoftware - Alerts which are related to Cohesity services. kHardware - Alerts related to hardware on which Cohesity software is running. kService - Alerts related to other external services. kOther - Alerts not of one of above categories. (optional)
    resolutionIdList := []int64{int64(123)} // []int64 | Specifies alert resolution ids to filter alerts by. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    allUnderHierarchy := true // bool | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)

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

    request := cohesitysdk.ApiGetAlertsRequest{
        maxAlerts: &MaxAlerts
        alertIdList: &AlertIdList
        alertTypeList: &AlertTypeList
        alertCategoryList: &AlertCategoryList
        propertyKey: &PropertyKey
        propertyValue: &PropertyValue
        startDateUsecs: &StartDateUsecs
        endDateUsecs: &EndDateUsecs
        alertStateList: &AlertStateList
        alertSeverityList: &AlertSeverityList
        alertTypeBucketList: &AlertTypeBucketList
        resolutionIdList: &ResolutionIdList
        tenantIds: &TenantIds
        allUnderHierarchy: &AllUnderHierarchy
    }

    resp, r, err := api_client.AlertsApi.GetAlerts(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlerts`: []Alert
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxAlerts** | **int32** | Specifies the number of returned Alerts to be returned. The newest Alerts are returned. | 
 **alertIdList** | **[]string** | Specifies list of Alert ids to filter alerts by. | 
 **alertTypeList** | **[]int32** | Specifies list of Alert Types to filter alerts by. | 
 **alertCategoryList** | **[]string** | Specifies list of Alert Categories. | 
 **propertyKey** | **string** | Specifies name of the property to filter alerts by. | 
 **propertyValue** | **string** | Specifies value of the property to filter alerts by. | 
 **startDateUsecs** | **int64** | Specifies start time Unix epoch time in microseconds to filter alerts by. | 
 **endDateUsecs** | **int64** | Specifies end time Unix epoch time in microseconds to filter alerts by. | 
 **alertStateList** | **[]string** | Specifies list of Alert States to filter alerts by. | 
 **alertSeverityList** | **[]string** | Specifies list of Alert severity to filter alerts by. | 
 **alertTypeBucketList** | **[]string** | Specifies the list of Alert type bucket. Specifies the Alert type bucket. kSoftware - Alerts which are related to Cohesity services. kHardware - Alerts related to hardware on which Cohesity software is running. kService - Alerts related to other external services. kOther - Alerts not of one of above categories. | 
 **resolutionIdList** | **[]int64** | Specifies alert resolution ids to filter alerts by. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 

### Return type

[**[]Alert**](Alert.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationRules

> []NotificationRule GetNotificationRules(ctx).Execute()

Gets all alert notification rules.



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

    request := cohesitysdk.ApiGetNotificationRulesRequest{
    }

    resp, r, err := api_client.AlertsApi.GetNotificationRules(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetNotificationRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationRules`: []NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetNotificationRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationRulesRequest struct via the builder pattern


### Return type

[**[]NotificationRule**](NotificationRule.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResolutionById

> AlertResolution GetResolutionById(ctx, id).Execute()

List details about a single Alert Resolution.



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
    id := int64(789) // int64 | Unique id of the Alert Resolution to return.

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

    request := cohesitysdk.ApiGetResolutionByIdRequest{
        id: &Id
    }

    resp, r, err := api_client.AlertsApi.GetResolutionById(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetResolutionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResolutionById`: AlertResolution
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetResolutionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Unique id of the Alert Resolution to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResolutionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertResolution**](AlertResolution.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResolutions

> []AlertResolution GetResolutions(ctx).MaxResolutions(maxResolutions).ResolutionIdList(resolutionIdList).AlertIdList(alertIdList).StartDateUsecs(startDateUsecs).EndDateUsecs(endDateUsecs).TenantIds(tenantIds).AllUnderHierarchy(allUnderHierarchy).Execute()

List the Alert Resolutions on the Cohesity Cluster.



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
    maxResolutions := int32(56) // int32 | Specifies the number of returned Resolutions to be returned. The newest Resolutions are returned.
    resolutionIdList := []int64{int64(123)} // []int64 | Specifies list of Alert Resolution ids to filter resolutions by. (optional)
    alertIdList := []string{"Inner_example"} // []string | Specifies list of Alert Resolution ids to filter resolutions by. (optional)
    startDateUsecs := int64(789) // int64 | Specifies Start Time Unix epoch in microseconds to filter resolutions by. (optional)
    endDateUsecs := int64(789) // int64 | Specifies End Time Unix epoch in microseconds to filter resolutions by. (optional)
    tenantIds := []string{"Inner_example"} // []string | TenantIds contains ids of the tenants for which objects are to be returned. (optional)
    allUnderHierarchy := true // bool | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned. (optional)

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

    request := cohesitysdk.ApiGetResolutionsRequest{
        maxResolutions: &MaxResolutions
        resolutionIdList: &ResolutionIdList
        alertIdList: &AlertIdList
        startDateUsecs: &StartDateUsecs
        endDateUsecs: &EndDateUsecs
        tenantIds: &TenantIds
        allUnderHierarchy: &AllUnderHierarchy
    }

    resp, r, err := api_client.AlertsApi.GetResolutions(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetResolutions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResolutions`: []AlertResolution
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetResolutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResolutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxResolutions** | **int32** | Specifies the number of returned Resolutions to be returned. The newest Resolutions are returned. | 
 **resolutionIdList** | **[]int64** | Specifies list of Alert Resolution ids to filter resolutions by. | 
 **alertIdList** | **[]string** | Specifies list of Alert Resolution ids to filter resolutions by. | 
 **startDateUsecs** | **int64** | Specifies Start Time Unix epoch in microseconds to filter resolutions by. | 
 **endDateUsecs** | **int64** | Specifies End Time Unix epoch in microseconds to filter resolutions by. | 
 **tenantIds** | **[]string** | TenantIds contains ids of the tenants for which objects are to be returned. | 
 **allUnderHierarchy** | **bool** | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user&#39;s organization should be returned. | 

### Return type

[**[]AlertResolution**](AlertResolution.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationRule

> NotificationRule UpdateNotificationRule(ctx).Execute()

Updates an existing alert notification rule.



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

    request := cohesitysdk.ApiUpdateNotificationRuleRequest{
    }

    resp, r, err := api_client.AlertsApi.UpdateNotificationRule(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.UpdateNotificationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationRule`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.UpdateNotificationRule`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationRuleRequest struct via the builder pattern


### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResolution

> AlertResolution UpdateResolution(ctx, id).Body(body).Execute()

Apply an existing Alert Resolution to additional Alerts.



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
    id := int64(789) // int64 | Unique id of the Alert Resolution to return.
    body := *openapiclient.NewUpdateResolutionParams() // UpdateResolutionParams | Request to apply an existing resolution to the specified Alerts.

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

    request := cohesitysdk.ApiUpdateResolutionRequest{
        id: &Id
        body: &Body
    }

    resp, r, err := api_client.AlertsApi.UpdateResolution(request)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.UpdateResolution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateResolution`: AlertResolution
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.UpdateResolution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Unique id of the Alert Resolution to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateResolutionParams**](UpdateResolutionParams.md) | Request to apply an existing resolution to the specified Alerts. | 

### Return type

[**AlertResolution**](AlertResolution.md)

### Authorization

[TokenHeader](../README.md#TokenHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

