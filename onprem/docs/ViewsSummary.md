# ViewsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalViews** | Pointer to **NullableInt64** | Specifies the number of all the views. | [optional] 
**ProtectedViews** | Pointer to **NullableInt64** | Specifies the number of all protected views. | [optional] 
**ReplicatedOutViews** | Pointer to **NullableInt64** | Specifies the number of all the views that are replicated out to remote clusters. | [optional] 
**ReplicatedInViews** | Pointer to **NullableInt64** | Specifies the number of all the views that are relicated from remote clusters. | [optional] 
**LogicalUsageBytes** | Pointer to **NullableInt64** | Specifies the logical usage of all the views in bytes. | [optional] 
**LogicalUsageBytesTimestampUsec** | Pointer to **NullableInt64** | Specifies the timestamp in micro seconds when &#39;logicalUsageBytes&#39; was calculated. | [optional] 
**StorageConsumedBytes** | Pointer to **NullableInt64** | Specifies the storage consumed of all the views in bytes. | [optional] 
**StorageConsumedBytesTimestampUsec** | Pointer to **NullableInt64** | Specifies the timestamp in micro seconds when &#39;storageConsumedBytes&#39; was calculated. | [optional] 

## Methods

### NewViewsSummary

`func NewViewsSummary() *ViewsSummary`

NewViewsSummary instantiates a new ViewsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewsSummaryWithDefaults

`func NewViewsSummaryWithDefaults() *ViewsSummary`

NewViewsSummaryWithDefaults instantiates a new ViewsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalViews

`func (o *ViewsSummary) GetTotalViews() int64`

GetTotalViews returns the TotalViews field if non-nil, zero value otherwise.

### GetTotalViewsOk

`func (o *ViewsSummary) GetTotalViewsOk() (*int64, bool)`

GetTotalViewsOk returns a tuple with the TotalViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalViews

`func (o *ViewsSummary) SetTotalViews(v int64)`

SetTotalViews sets TotalViews field to given value.

### HasTotalViews

`func (o *ViewsSummary) HasTotalViews() bool`

HasTotalViews returns a boolean if a field has been set.

### SetTotalViewsNil

`func (o *ViewsSummary) SetTotalViewsNil(b bool)`

 SetTotalViewsNil sets the value for TotalViews to be an explicit nil

### UnsetTotalViews
`func (o *ViewsSummary) UnsetTotalViews()`

UnsetTotalViews ensures that no value is present for TotalViews, not even an explicit nil
### GetProtectedViews

`func (o *ViewsSummary) GetProtectedViews() int64`

GetProtectedViews returns the ProtectedViews field if non-nil, zero value otherwise.

### GetProtectedViewsOk

`func (o *ViewsSummary) GetProtectedViewsOk() (*int64, bool)`

GetProtectedViewsOk returns a tuple with the ProtectedViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedViews

`func (o *ViewsSummary) SetProtectedViews(v int64)`

SetProtectedViews sets ProtectedViews field to given value.

### HasProtectedViews

`func (o *ViewsSummary) HasProtectedViews() bool`

HasProtectedViews returns a boolean if a field has been set.

### SetProtectedViewsNil

`func (o *ViewsSummary) SetProtectedViewsNil(b bool)`

 SetProtectedViewsNil sets the value for ProtectedViews to be an explicit nil

### UnsetProtectedViews
`func (o *ViewsSummary) UnsetProtectedViews()`

UnsetProtectedViews ensures that no value is present for ProtectedViews, not even an explicit nil
### GetReplicatedOutViews

`func (o *ViewsSummary) GetReplicatedOutViews() int64`

GetReplicatedOutViews returns the ReplicatedOutViews field if non-nil, zero value otherwise.

### GetReplicatedOutViewsOk

`func (o *ViewsSummary) GetReplicatedOutViewsOk() (*int64, bool)`

GetReplicatedOutViewsOk returns a tuple with the ReplicatedOutViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicatedOutViews

`func (o *ViewsSummary) SetReplicatedOutViews(v int64)`

SetReplicatedOutViews sets ReplicatedOutViews field to given value.

### HasReplicatedOutViews

`func (o *ViewsSummary) HasReplicatedOutViews() bool`

HasReplicatedOutViews returns a boolean if a field has been set.

### SetReplicatedOutViewsNil

`func (o *ViewsSummary) SetReplicatedOutViewsNil(b bool)`

 SetReplicatedOutViewsNil sets the value for ReplicatedOutViews to be an explicit nil

### UnsetReplicatedOutViews
`func (o *ViewsSummary) UnsetReplicatedOutViews()`

UnsetReplicatedOutViews ensures that no value is present for ReplicatedOutViews, not even an explicit nil
### GetReplicatedInViews

`func (o *ViewsSummary) GetReplicatedInViews() int64`

GetReplicatedInViews returns the ReplicatedInViews field if non-nil, zero value otherwise.

### GetReplicatedInViewsOk

`func (o *ViewsSummary) GetReplicatedInViewsOk() (*int64, bool)`

GetReplicatedInViewsOk returns a tuple with the ReplicatedInViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicatedInViews

`func (o *ViewsSummary) SetReplicatedInViews(v int64)`

SetReplicatedInViews sets ReplicatedInViews field to given value.

### HasReplicatedInViews

`func (o *ViewsSummary) HasReplicatedInViews() bool`

HasReplicatedInViews returns a boolean if a field has been set.

### SetReplicatedInViewsNil

`func (o *ViewsSummary) SetReplicatedInViewsNil(b bool)`

 SetReplicatedInViewsNil sets the value for ReplicatedInViews to be an explicit nil

### UnsetReplicatedInViews
`func (o *ViewsSummary) UnsetReplicatedInViews()`

UnsetReplicatedInViews ensures that no value is present for ReplicatedInViews, not even an explicit nil
### GetLogicalUsageBytes

`func (o *ViewsSummary) GetLogicalUsageBytes() int64`

GetLogicalUsageBytes returns the LogicalUsageBytes field if non-nil, zero value otherwise.

### GetLogicalUsageBytesOk

`func (o *ViewsSummary) GetLogicalUsageBytesOk() (*int64, bool)`

GetLogicalUsageBytesOk returns a tuple with the LogicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalUsageBytes

`func (o *ViewsSummary) SetLogicalUsageBytes(v int64)`

SetLogicalUsageBytes sets LogicalUsageBytes field to given value.

### HasLogicalUsageBytes

`func (o *ViewsSummary) HasLogicalUsageBytes() bool`

HasLogicalUsageBytes returns a boolean if a field has been set.

### SetLogicalUsageBytesNil

`func (o *ViewsSummary) SetLogicalUsageBytesNil(b bool)`

 SetLogicalUsageBytesNil sets the value for LogicalUsageBytes to be an explicit nil

### UnsetLogicalUsageBytes
`func (o *ViewsSummary) UnsetLogicalUsageBytes()`

UnsetLogicalUsageBytes ensures that no value is present for LogicalUsageBytes, not even an explicit nil
### GetLogicalUsageBytesTimestampUsec

`func (o *ViewsSummary) GetLogicalUsageBytesTimestampUsec() int64`

GetLogicalUsageBytesTimestampUsec returns the LogicalUsageBytesTimestampUsec field if non-nil, zero value otherwise.

### GetLogicalUsageBytesTimestampUsecOk

`func (o *ViewsSummary) GetLogicalUsageBytesTimestampUsecOk() (*int64, bool)`

GetLogicalUsageBytesTimestampUsecOk returns a tuple with the LogicalUsageBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalUsageBytesTimestampUsec

`func (o *ViewsSummary) SetLogicalUsageBytesTimestampUsec(v int64)`

SetLogicalUsageBytesTimestampUsec sets LogicalUsageBytesTimestampUsec field to given value.

### HasLogicalUsageBytesTimestampUsec

`func (o *ViewsSummary) HasLogicalUsageBytesTimestampUsec() bool`

HasLogicalUsageBytesTimestampUsec returns a boolean if a field has been set.

### SetLogicalUsageBytesTimestampUsecNil

`func (o *ViewsSummary) SetLogicalUsageBytesTimestampUsecNil(b bool)`

 SetLogicalUsageBytesTimestampUsecNil sets the value for LogicalUsageBytesTimestampUsec to be an explicit nil

### UnsetLogicalUsageBytesTimestampUsec
`func (o *ViewsSummary) UnsetLogicalUsageBytesTimestampUsec()`

UnsetLogicalUsageBytesTimestampUsec ensures that no value is present for LogicalUsageBytesTimestampUsec, not even an explicit nil
### GetStorageConsumedBytes

`func (o *ViewsSummary) GetStorageConsumedBytes() int64`

GetStorageConsumedBytes returns the StorageConsumedBytes field if non-nil, zero value otherwise.

### GetStorageConsumedBytesOk

`func (o *ViewsSummary) GetStorageConsumedBytesOk() (*int64, bool)`

GetStorageConsumedBytesOk returns a tuple with the StorageConsumedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConsumedBytes

`func (o *ViewsSummary) SetStorageConsumedBytes(v int64)`

SetStorageConsumedBytes sets StorageConsumedBytes field to given value.

### HasStorageConsumedBytes

`func (o *ViewsSummary) HasStorageConsumedBytes() bool`

HasStorageConsumedBytes returns a boolean if a field has been set.

### SetStorageConsumedBytesNil

`func (o *ViewsSummary) SetStorageConsumedBytesNil(b bool)`

 SetStorageConsumedBytesNil sets the value for StorageConsumedBytes to be an explicit nil

### UnsetStorageConsumedBytes
`func (o *ViewsSummary) UnsetStorageConsumedBytes()`

UnsetStorageConsumedBytes ensures that no value is present for StorageConsumedBytes, not even an explicit nil
### GetStorageConsumedBytesTimestampUsec

`func (o *ViewsSummary) GetStorageConsumedBytesTimestampUsec() int64`

GetStorageConsumedBytesTimestampUsec returns the StorageConsumedBytesTimestampUsec field if non-nil, zero value otherwise.

### GetStorageConsumedBytesTimestampUsecOk

`func (o *ViewsSummary) GetStorageConsumedBytesTimestampUsecOk() (*int64, bool)`

GetStorageConsumedBytesTimestampUsecOk returns a tuple with the StorageConsumedBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConsumedBytesTimestampUsec

`func (o *ViewsSummary) SetStorageConsumedBytesTimestampUsec(v int64)`

SetStorageConsumedBytesTimestampUsec sets StorageConsumedBytesTimestampUsec field to given value.

### HasStorageConsumedBytesTimestampUsec

`func (o *ViewsSummary) HasStorageConsumedBytesTimestampUsec() bool`

HasStorageConsumedBytesTimestampUsec returns a boolean if a field has been set.

### SetStorageConsumedBytesTimestampUsecNil

`func (o *ViewsSummary) SetStorageConsumedBytesTimestampUsecNil(b bool)`

 SetStorageConsumedBytesTimestampUsecNil sets the value for StorageConsumedBytesTimestampUsec to be an explicit nil

### UnsetStorageConsumedBytesTimestampUsec
`func (o *ViewsSummary) UnsetStorageConsumedBytesTimestampUsec()`

UnsetStorageConsumedBytesTimestampUsec ensures that no value is present for StorageConsumedBytesTimestampUsec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


