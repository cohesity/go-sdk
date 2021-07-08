# ClusterAuditLogsSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterAuditLogs** | Pointer to [**[]ClusterAuditLog**](ClusterAuditLog.md) | Array of Cluster Audit Logs.  Specifies a list of Cluster audit logs that match the specified filter criteria up to the limit specified in pageCount. | [optional] 
**TotalCount** | Pointer to **NullableInt64** | Specifies the total number of logs that match the specified filter criteria. (This number might be larger than the size of the Cluster Audit Logs array.) This count is provided to indicate if additional requests must be made to get the full result. | [optional] 

## Methods

### NewClusterAuditLogsSearchResult

`func NewClusterAuditLogsSearchResult() *ClusterAuditLogsSearchResult`

NewClusterAuditLogsSearchResult instantiates a new ClusterAuditLogsSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAuditLogsSearchResultWithDefaults

`func NewClusterAuditLogsSearchResultWithDefaults() *ClusterAuditLogsSearchResult`

NewClusterAuditLogsSearchResultWithDefaults instantiates a new ClusterAuditLogsSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterAuditLogs

`func (o *ClusterAuditLogsSearchResult) GetClusterAuditLogs() []ClusterAuditLog`

GetClusterAuditLogs returns the ClusterAuditLogs field if non-nil, zero value otherwise.

### GetClusterAuditLogsOk

`func (o *ClusterAuditLogsSearchResult) GetClusterAuditLogsOk() (*[]ClusterAuditLog, bool)`

GetClusterAuditLogsOk returns a tuple with the ClusterAuditLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterAuditLogs

`func (o *ClusterAuditLogsSearchResult) SetClusterAuditLogs(v []ClusterAuditLog)`

SetClusterAuditLogs sets ClusterAuditLogs field to given value.

### HasClusterAuditLogs

`func (o *ClusterAuditLogsSearchResult) HasClusterAuditLogs() bool`

HasClusterAuditLogs returns a boolean if a field has been set.

### SetClusterAuditLogsNil

`func (o *ClusterAuditLogsSearchResult) SetClusterAuditLogsNil(b bool)`

 SetClusterAuditLogsNil sets the value for ClusterAuditLogs to be an explicit nil

### UnsetClusterAuditLogs
`func (o *ClusterAuditLogsSearchResult) UnsetClusterAuditLogs()`

UnsetClusterAuditLogs ensures that no value is present for ClusterAuditLogs, not even an explicit nil
### GetTotalCount

`func (o *ClusterAuditLogsSearchResult) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ClusterAuditLogsSearchResult) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ClusterAuditLogsSearchResult) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ClusterAuditLogsSearchResult) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### SetTotalCountNil

`func (o *ClusterAuditLogsSearchResult) SetTotalCountNil(b bool)`

 SetTotalCountNil sets the value for TotalCount to be an explicit nil

### UnsetTotalCount
`func (o *ClusterAuditLogsSearchResult) UnsetTotalCount()`

UnsetTotalCount ensures that no value is present for TotalCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


