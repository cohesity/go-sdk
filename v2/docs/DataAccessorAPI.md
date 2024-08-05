# \DataAccessorAPI

All URIs are relative to */v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDataAccessSession**](DataAccessorAPI.md#CreateDataAccessSession) | **Post** /data-access/snapshots/sessions | Create Data Access Session
[**DifferenceOfGraphNodes**](DataAccessorAPI.md#DifferenceOfGraphNodes) | **Post** /data-access/snapshots/sessions/{sessionId}/graph-nodes/differences | Difference of Graph nodes
[**GetDataAccessSessions**](DataAccessorAPI.md#GetDataAccessSessions) | **Get** /data-access/snapshots/sessions | Lists all the Data Access Sessions
[**GetGraphNodeDetails**](DataAccessorAPI.md#GetGraphNodeDetails) | **Post** /data-access/snapshots/sessions/{sessionId}/graph-nodes/{nodeId} | Get Graph Node details
[**GetGraphNodeRelationsDifferences**](DataAccessorAPI.md#GetGraphNodeRelationsDifferences) | **Post** /data-access/snapshots/sessions/{sessionId}/graph-nodes/query-relations/{nodeId}/differences | Query for difference of graph node relations
[**SearchGraphNodes**](DataAccessorAPI.md#SearchGraphNodes) | **Post** /data-access/snapshots/sessions/graph-nodes/query | Search Graph nodes
[**TearDownDataAccessSession**](DataAccessorAPI.md#TearDownDataAccessSession) | **Delete** /data-access/snapshots/sessions/{sessionId} | Tear down data access session for a given id



## CreateDataAccessSession

> CreateDataAccessSessionResponseParams CreateDataAccessSession(ctx).Body(body).Execute()

Create Data Access Session



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
	body := *openapiclient.NewCreateDataAccessSessionRequestParams(*openapiclient.NewDataAccessSnapshotInfo(int64(123), "SnapshotId_example"), int64(123)) // CreateDataAccessSessionRequestParams | Specifies the parameters to create data access session

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAccessorAPI.CreateDataAccessSession(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAccessorAPI.CreateDataAccessSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDataAccessSession`: CreateDataAccessSessionResponseParams
	fmt.Fprintf(os.Stdout, "Response from `DataAccessorAPI.CreateDataAccessSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataAccessSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateDataAccessSessionRequestParams**](CreateDataAccessSessionRequestParams.md) | Specifies the parameters to create data access session | 

### Return type

[**CreateDataAccessSessionResponseParams**](CreateDataAccessSessionResponseParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DifferenceOfGraphNodes

> QueryGraphNodesDiffResult DifferenceOfGraphNodes(ctx, sessionId).Body(body).Execute()

Difference of Graph nodes



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
	sessionId := "sessionId_example" // string | Specifies the id of a session.
	body := *openapiclient.NewQueryGraphNodesDiffParams() // QueryGraphNodesDiffParams | Specifies the parameters to determine graph nodes diff for a given session id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAccessorAPI.DifferenceOfGraphNodes(context.Background(), sessionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAccessorAPI.DifferenceOfGraphNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DifferenceOfGraphNodes`: QueryGraphNodesDiffResult
	fmt.Fprintf(os.Stdout, "Response from `DataAccessorAPI.DifferenceOfGraphNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Specifies the id of a session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDifferenceOfGraphNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**QueryGraphNodesDiffParams**](QueryGraphNodesDiffParams.md) | Specifies the parameters to determine graph nodes diff for a given session id. | 

### Return type

[**QueryGraphNodesDiffResult**](QueryGraphNodesDiffResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataAccessSessions

> GetDataAccessSessionsResponseParams GetDataAccessSessions(ctx).SessionIds(sessionIds).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).SnapshotEnvironments(snapshotEnvironments).Statuses(statuses).PaginationCookie(paginationCookie).Count(count).Execute()

Lists all the Data Access Sessions



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
	sessionIds := []string{"Inner_example"} // []string | Filter Data Access Sessions for given session ids. (optional)
	startTimeUsecs := int64(789) // int64 | Returns the data access sessions which are started after the specific time. This value should be in Unix timestamp epoch in microseconds. (optional)
	endTimeUsecs := int64(789) // int64 | Returns the data access sessions which are started before the specific time. This value should be in Unix timestamp epoch in microseconds. (optional)
	snapshotEnvironments := []string{"SnapshotEnvironments_example"} // []string | Specifies the snapshot environment types to filter data access sessions. (optional)
	statuses := []string{"Statuses_example"} // []string | Specifies the list of session states to filter data access sessions (optional)
	paginationCookie := "paginationCookie_example" // string | Specifies a cookie which can be passed in by the user in order to retrieve the next page of results. (optional)
	count := int32(56) // int32 | Specifies the number of objects to be fetched for the specified pagination cookie. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAccessorAPI.GetDataAccessSessions(context.Background()).SessionIds(sessionIds).StartTimeUsecs(startTimeUsecs).EndTimeUsecs(endTimeUsecs).SnapshotEnvironments(snapshotEnvironments).Statuses(statuses).PaginationCookie(paginationCookie).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAccessorAPI.GetDataAccessSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataAccessSessions`: GetDataAccessSessionsResponseParams
	fmt.Fprintf(os.Stdout, "Response from `DataAccessorAPI.GetDataAccessSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataAccessSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionIds** | **[]string** | Filter Data Access Sessions for given session ids. | 
 **startTimeUsecs** | **int64** | Returns the data access sessions which are started after the specific time. This value should be in Unix timestamp epoch in microseconds. | 
 **endTimeUsecs** | **int64** | Returns the data access sessions which are started before the specific time. This value should be in Unix timestamp epoch in microseconds. | 
 **snapshotEnvironments** | **[]string** | Specifies the snapshot environment types to filter data access sessions. | 
 **statuses** | **[]string** | Specifies the list of session states to filter data access sessions | 
 **paginationCookie** | **string** | Specifies a cookie which can be passed in by the user in order to retrieve the next page of results. | 
 **count** | **int32** | Specifies the number of objects to be fetched for the specified pagination cookie. | 

### Return type

[**GetDataAccessSessionsResponseParams**](GetDataAccessSessionsResponseParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGraphNodeDetails

> GetGraphNodeDetailsResult GetGraphNodeDetails(ctx, sessionId, nodeId).Body(body).Execute()

Get Graph Node details



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
	sessionId := "sessionId_example" // string | Specifies the id of a session.
	nodeId := "nodeId_example" // string | Specifies the id of a node.
	body := *openapiclient.NewGetGraphNodeDetailsRequestParams() // GetGraphNodeDetailsRequestParams | Specifies the parameters to get node details in the graph for a given node id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAccessorAPI.GetGraphNodeDetails(context.Background(), sessionId, nodeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAccessorAPI.GetGraphNodeDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGraphNodeDetails`: GetGraphNodeDetailsResult
	fmt.Fprintf(os.Stdout, "Response from `DataAccessorAPI.GetGraphNodeDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Specifies the id of a session. | 
**nodeId** | **string** | Specifies the id of a node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGraphNodeDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**GetGraphNodeDetailsRequestParams**](GetGraphNodeDetailsRequestParams.md) | Specifies the parameters to get node details in the graph for a given node id. | 

### Return type

[**GetGraphNodeDetailsResult**](GetGraphNodeDetailsResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGraphNodeRelationsDifferences

> DiffGraphNodeRelation GetGraphNodeRelationsDifferences(ctx, sessionId, nodeId).Body(body).Execute()

Query for difference of graph node relations



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
	sessionId := "sessionId_example" // string | Specifies the id of a session.
	nodeId := "nodeId_example" // string | Specifies the id of a graph node.
	body := *openapiclient.NewGetGraphNodeRelationsDiffParams() // GetGraphNodeRelationsDiffParams | Specifies the parameters to search graph node relations for a given node id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAccessorAPI.GetGraphNodeRelationsDifferences(context.Background(), sessionId, nodeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAccessorAPI.GetGraphNodeRelationsDifferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGraphNodeRelationsDifferences`: DiffGraphNodeRelation
	fmt.Fprintf(os.Stdout, "Response from `DataAccessorAPI.GetGraphNodeRelationsDifferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Specifies the id of a session. | 
**nodeId** | **string** | Specifies the id of a graph node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGraphNodeRelationsDifferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**GetGraphNodeRelationsDiffParams**](GetGraphNodeRelationsDiffParams.md) | Specifies the parameters to search graph node relations for a given node id. | 

### Return type

[**DiffGraphNodeRelation**](DiffGraphNodeRelation.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchGraphNodes

> SearchGraphNodesResponseParams SearchGraphNodes(ctx).Body(body).Execute()

Search Graph nodes



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
	body := *openapiclient.NewSearchGraphNodesRequestParams("SessionId_example") // SearchGraphNodesRequestParams | Specifies the parameters to query nodes in the graph for a given session id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataAccessorAPI.SearchGraphNodes(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAccessorAPI.SearchGraphNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGraphNodes`: SearchGraphNodesResponseParams
	fmt.Fprintf(os.Stdout, "Response from `DataAccessorAPI.SearchGraphNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchGraphNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SearchGraphNodesRequestParams**](SearchGraphNodesRequestParams.md) | Specifies the parameters to query nodes in the graph for a given session id. | 

### Return type

[**SearchGraphNodesResponseParams**](SearchGraphNodesResponseParams.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TearDownDataAccessSession

> TearDownDataAccessSession(ctx, sessionId).Execute()

Tear down data access session for a given id



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
	sessionId := "sessionId_example" // string | Specifies the id of the data access session.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataAccessorAPI.TearDownDataAccessSession(context.Background(), sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataAccessorAPI.TearDownDataAccessSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Specifies the id of the data access session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTearDownDataAccessSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

