# ViewsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataWrittenBytes** | Pointer to **NullableInt64** | Specifies the data written to all the views in bytes. | [optional] 
**DataWrittenBytesPrev** | Pointer to **NullableInt64** | Specifies the data written to all the views in bytes at a specific time. | [optional] 
**DataWrittenBytesPrevTimestampUsec** | Pointer to **NullableInt64** | Specifies the timestamp in micro seconds when &#39;dataWrittenBytesPrev&#39; was calculated. | [optional] 
**DataWrittenBytesTimestampUsec** | Pointer to **NullableInt64** | Specifies the timestamp in micro seconds when &#39;dataWrittenBytes&#39; was calculated. | [optional] 
**LocalTierResiliencyImpactBytes** | Pointer to **NullableInt64** | Specifies the size of the data that has been replicated to other nodes as per RF or Erasure Coding policy. | [optional] 
**LocalTierResiliencyImpactBytesPrev** | Pointer to **NullableInt64** | Specifies the size of the data that has been replicated to other nodes as per RF or Erasure Coding policy at a specific time. | [optional] 
**LocalTierResiliencyImpactBytesPrevTimestampUsec** | Pointer to **NullableInt64** | Specifies the timestamp in micro seconds when &#39;localTierResiliencyImpactBytesPrev&#39; was calculated. | [optional] 
**LocalTierResiliencyImpactBytesTimestampUsec** | Pointer to **NullableInt64** | Specifies the timestamp in micro seconds when &#39;localTierResiliencyImpactBytes&#39; was calculated. | [optional] 
**LogicalUsageBytes** | Pointer to **NullableInt64** | Specifies the logical usage of all the views in bytes. | [optional] 
**LogicalUsageBytesPrev** | Pointer to **NullableInt64** | Specifies the logical usage of all the views in bytes at a specific time. | [optional] 
**LogicalUsageBytesPrevTimestampUsec** | Pointer to **NullableInt64** | Specifies the timestamp in micro seconds when &#39;logicalUsageBytesPrev&#39; was calculated. | [optional] 
**LogicalUsageBytesTimestampUsec** | Pointer to **NullableInt64** | Specifies the timestamp in micro seconds when &#39;logicalUsageBytes&#39; was calculated. | [optional] 
**NumDirectories** | Pointer to **NullableInt64** | Specifies the number of directories. | [optional] 
**NumDirectoriesPrev** | Pointer to **NullableInt64** | Specifies the number of directories at a specific time. | [optional] 
**NumFiles** | Pointer to **NullableInt64** | Specifies the number of files. | [optional] 
**NumFilesPrev** | Pointer to **NullableInt64** | Specifies the number of files at a specific time. | [optional] 
**ProtectedViews** | Pointer to **NullableInt64** | Specifies the number of all protected views. | [optional] 
**ReplicatedInViews** | Pointer to **NullableInt64** | Specifies the number of all the views that are replicated from remote clusters. | [optional] 
**ReplicatedOutViews** | Pointer to **NullableInt64** | Specifies the number of all the views that are replicated out to remote clusters. | [optional] 
**StorageConsumedBytes** | Pointer to **NullableInt64** | Specifies the storage consumed of all the views in bytes. | [optional] 
**StorageConsumedBytesPrev** | Pointer to **NullableInt64** | Specifies the storage consumed by all the views in bytes at a specific time. | [optional] 
**StorageConsumedBytesPrevTimestampUsec** | Pointer to **NullableInt64** | Specifies the timestamp in micro seconds when &#39;storageConsumedBytesPrev&#39; was calculated. | [optional] 
**StorageConsumedBytesTimestampUsec** | Pointer to **NullableInt64** | Specifies the timestamp in micro seconds when &#39;storageConsumedBytes&#39; was calculated. | [optional] 
**TotalViews** | Pointer to **NullableInt64** | Specifies the number of all the views. | [optional] 
**ViewEntityId** | Pointer to **NullableString** | Specifies the entity id of all the views. | [optional] 

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

### GetDataWrittenBytes

`func (o *ViewsSummary) GetDataWrittenBytes() int64`

GetDataWrittenBytes returns the DataWrittenBytes field if non-nil, zero value otherwise.

### GetDataWrittenBytesOk

`func (o *ViewsSummary) GetDataWrittenBytesOk() (*int64, bool)`

GetDataWrittenBytesOk returns a tuple with the DataWrittenBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWrittenBytes

`func (o *ViewsSummary) SetDataWrittenBytes(v int64)`

SetDataWrittenBytes sets DataWrittenBytes field to given value.

### HasDataWrittenBytes

`func (o *ViewsSummary) HasDataWrittenBytes() bool`

HasDataWrittenBytes returns a boolean if a field has been set.

### SetDataWrittenBytesNil

`func (o *ViewsSummary) SetDataWrittenBytesNil(b bool)`

 SetDataWrittenBytesNil sets the value for DataWrittenBytes to be an explicit nil

### UnsetDataWrittenBytes
`func (o *ViewsSummary) UnsetDataWrittenBytes()`

UnsetDataWrittenBytes ensures that no value is present for DataWrittenBytes, not even an explicit nil
### GetDataWrittenBytesPrev

`func (o *ViewsSummary) GetDataWrittenBytesPrev() int64`

GetDataWrittenBytesPrev returns the DataWrittenBytesPrev field if non-nil, zero value otherwise.

### GetDataWrittenBytesPrevOk

`func (o *ViewsSummary) GetDataWrittenBytesPrevOk() (*int64, bool)`

GetDataWrittenBytesPrevOk returns a tuple with the DataWrittenBytesPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWrittenBytesPrev

`func (o *ViewsSummary) SetDataWrittenBytesPrev(v int64)`

SetDataWrittenBytesPrev sets DataWrittenBytesPrev field to given value.

### HasDataWrittenBytesPrev

`func (o *ViewsSummary) HasDataWrittenBytesPrev() bool`

HasDataWrittenBytesPrev returns a boolean if a field has been set.

### SetDataWrittenBytesPrevNil

`func (o *ViewsSummary) SetDataWrittenBytesPrevNil(b bool)`

 SetDataWrittenBytesPrevNil sets the value for DataWrittenBytesPrev to be an explicit nil

### UnsetDataWrittenBytesPrev
`func (o *ViewsSummary) UnsetDataWrittenBytesPrev()`

UnsetDataWrittenBytesPrev ensures that no value is present for DataWrittenBytesPrev, not even an explicit nil
### GetDataWrittenBytesPrevTimestampUsec

`func (o *ViewsSummary) GetDataWrittenBytesPrevTimestampUsec() int64`

GetDataWrittenBytesPrevTimestampUsec returns the DataWrittenBytesPrevTimestampUsec field if non-nil, zero value otherwise.

### GetDataWrittenBytesPrevTimestampUsecOk

`func (o *ViewsSummary) GetDataWrittenBytesPrevTimestampUsecOk() (*int64, bool)`

GetDataWrittenBytesPrevTimestampUsecOk returns a tuple with the DataWrittenBytesPrevTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWrittenBytesPrevTimestampUsec

`func (o *ViewsSummary) SetDataWrittenBytesPrevTimestampUsec(v int64)`

SetDataWrittenBytesPrevTimestampUsec sets DataWrittenBytesPrevTimestampUsec field to given value.

### HasDataWrittenBytesPrevTimestampUsec

`func (o *ViewsSummary) HasDataWrittenBytesPrevTimestampUsec() bool`

HasDataWrittenBytesPrevTimestampUsec returns a boolean if a field has been set.

### SetDataWrittenBytesPrevTimestampUsecNil

`func (o *ViewsSummary) SetDataWrittenBytesPrevTimestampUsecNil(b bool)`

 SetDataWrittenBytesPrevTimestampUsecNil sets the value for DataWrittenBytesPrevTimestampUsec to be an explicit nil

### UnsetDataWrittenBytesPrevTimestampUsec
`func (o *ViewsSummary) UnsetDataWrittenBytesPrevTimestampUsec()`

UnsetDataWrittenBytesPrevTimestampUsec ensures that no value is present for DataWrittenBytesPrevTimestampUsec, not even an explicit nil
### GetDataWrittenBytesTimestampUsec

`func (o *ViewsSummary) GetDataWrittenBytesTimestampUsec() int64`

GetDataWrittenBytesTimestampUsec returns the DataWrittenBytesTimestampUsec field if non-nil, zero value otherwise.

### GetDataWrittenBytesTimestampUsecOk

`func (o *ViewsSummary) GetDataWrittenBytesTimestampUsecOk() (*int64, bool)`

GetDataWrittenBytesTimestampUsecOk returns a tuple with the DataWrittenBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWrittenBytesTimestampUsec

`func (o *ViewsSummary) SetDataWrittenBytesTimestampUsec(v int64)`

SetDataWrittenBytesTimestampUsec sets DataWrittenBytesTimestampUsec field to given value.

### HasDataWrittenBytesTimestampUsec

`func (o *ViewsSummary) HasDataWrittenBytesTimestampUsec() bool`

HasDataWrittenBytesTimestampUsec returns a boolean if a field has been set.

### SetDataWrittenBytesTimestampUsecNil

`func (o *ViewsSummary) SetDataWrittenBytesTimestampUsecNil(b bool)`

 SetDataWrittenBytesTimestampUsecNil sets the value for DataWrittenBytesTimestampUsec to be an explicit nil

### UnsetDataWrittenBytesTimestampUsec
`func (o *ViewsSummary) UnsetDataWrittenBytesTimestampUsec()`

UnsetDataWrittenBytesTimestampUsec ensures that no value is present for DataWrittenBytesTimestampUsec, not even an explicit nil
### GetLocalTierResiliencyImpactBytes

`func (o *ViewsSummary) GetLocalTierResiliencyImpactBytes() int64`

GetLocalTierResiliencyImpactBytes returns the LocalTierResiliencyImpactBytes field if non-nil, zero value otherwise.

### GetLocalTierResiliencyImpactBytesOk

`func (o *ViewsSummary) GetLocalTierResiliencyImpactBytesOk() (*int64, bool)`

GetLocalTierResiliencyImpactBytesOk returns a tuple with the LocalTierResiliencyImpactBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTierResiliencyImpactBytes

`func (o *ViewsSummary) SetLocalTierResiliencyImpactBytes(v int64)`

SetLocalTierResiliencyImpactBytes sets LocalTierResiliencyImpactBytes field to given value.

### HasLocalTierResiliencyImpactBytes

`func (o *ViewsSummary) HasLocalTierResiliencyImpactBytes() bool`

HasLocalTierResiliencyImpactBytes returns a boolean if a field has been set.

### SetLocalTierResiliencyImpactBytesNil

`func (o *ViewsSummary) SetLocalTierResiliencyImpactBytesNil(b bool)`

 SetLocalTierResiliencyImpactBytesNil sets the value for LocalTierResiliencyImpactBytes to be an explicit nil

### UnsetLocalTierResiliencyImpactBytes
`func (o *ViewsSummary) UnsetLocalTierResiliencyImpactBytes()`

UnsetLocalTierResiliencyImpactBytes ensures that no value is present for LocalTierResiliencyImpactBytes, not even an explicit nil
### GetLocalTierResiliencyImpactBytesPrev

`func (o *ViewsSummary) GetLocalTierResiliencyImpactBytesPrev() int64`

GetLocalTierResiliencyImpactBytesPrev returns the LocalTierResiliencyImpactBytesPrev field if non-nil, zero value otherwise.

### GetLocalTierResiliencyImpactBytesPrevOk

`func (o *ViewsSummary) GetLocalTierResiliencyImpactBytesPrevOk() (*int64, bool)`

GetLocalTierResiliencyImpactBytesPrevOk returns a tuple with the LocalTierResiliencyImpactBytesPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTierResiliencyImpactBytesPrev

`func (o *ViewsSummary) SetLocalTierResiliencyImpactBytesPrev(v int64)`

SetLocalTierResiliencyImpactBytesPrev sets LocalTierResiliencyImpactBytesPrev field to given value.

### HasLocalTierResiliencyImpactBytesPrev

`func (o *ViewsSummary) HasLocalTierResiliencyImpactBytesPrev() bool`

HasLocalTierResiliencyImpactBytesPrev returns a boolean if a field has been set.

### SetLocalTierResiliencyImpactBytesPrevNil

`func (o *ViewsSummary) SetLocalTierResiliencyImpactBytesPrevNil(b bool)`

 SetLocalTierResiliencyImpactBytesPrevNil sets the value for LocalTierResiliencyImpactBytesPrev to be an explicit nil

### UnsetLocalTierResiliencyImpactBytesPrev
`func (o *ViewsSummary) UnsetLocalTierResiliencyImpactBytesPrev()`

UnsetLocalTierResiliencyImpactBytesPrev ensures that no value is present for LocalTierResiliencyImpactBytesPrev, not even an explicit nil
### GetLocalTierResiliencyImpactBytesPrevTimestampUsec

`func (o *ViewsSummary) GetLocalTierResiliencyImpactBytesPrevTimestampUsec() int64`

GetLocalTierResiliencyImpactBytesPrevTimestampUsec returns the LocalTierResiliencyImpactBytesPrevTimestampUsec field if non-nil, zero value otherwise.

### GetLocalTierResiliencyImpactBytesPrevTimestampUsecOk

`func (o *ViewsSummary) GetLocalTierResiliencyImpactBytesPrevTimestampUsecOk() (*int64, bool)`

GetLocalTierResiliencyImpactBytesPrevTimestampUsecOk returns a tuple with the LocalTierResiliencyImpactBytesPrevTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTierResiliencyImpactBytesPrevTimestampUsec

`func (o *ViewsSummary) SetLocalTierResiliencyImpactBytesPrevTimestampUsec(v int64)`

SetLocalTierResiliencyImpactBytesPrevTimestampUsec sets LocalTierResiliencyImpactBytesPrevTimestampUsec field to given value.

### HasLocalTierResiliencyImpactBytesPrevTimestampUsec

`func (o *ViewsSummary) HasLocalTierResiliencyImpactBytesPrevTimestampUsec() bool`

HasLocalTierResiliencyImpactBytesPrevTimestampUsec returns a boolean if a field has been set.

### SetLocalTierResiliencyImpactBytesPrevTimestampUsecNil

`func (o *ViewsSummary) SetLocalTierResiliencyImpactBytesPrevTimestampUsecNil(b bool)`

 SetLocalTierResiliencyImpactBytesPrevTimestampUsecNil sets the value for LocalTierResiliencyImpactBytesPrevTimestampUsec to be an explicit nil

### UnsetLocalTierResiliencyImpactBytesPrevTimestampUsec
`func (o *ViewsSummary) UnsetLocalTierResiliencyImpactBytesPrevTimestampUsec()`

UnsetLocalTierResiliencyImpactBytesPrevTimestampUsec ensures that no value is present for LocalTierResiliencyImpactBytesPrevTimestampUsec, not even an explicit nil
### GetLocalTierResiliencyImpactBytesTimestampUsec

`func (o *ViewsSummary) GetLocalTierResiliencyImpactBytesTimestampUsec() int64`

GetLocalTierResiliencyImpactBytesTimestampUsec returns the LocalTierResiliencyImpactBytesTimestampUsec field if non-nil, zero value otherwise.

### GetLocalTierResiliencyImpactBytesTimestampUsecOk

`func (o *ViewsSummary) GetLocalTierResiliencyImpactBytesTimestampUsecOk() (*int64, bool)`

GetLocalTierResiliencyImpactBytesTimestampUsecOk returns a tuple with the LocalTierResiliencyImpactBytesTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTierResiliencyImpactBytesTimestampUsec

`func (o *ViewsSummary) SetLocalTierResiliencyImpactBytesTimestampUsec(v int64)`

SetLocalTierResiliencyImpactBytesTimestampUsec sets LocalTierResiliencyImpactBytesTimestampUsec field to given value.

### HasLocalTierResiliencyImpactBytesTimestampUsec

`func (o *ViewsSummary) HasLocalTierResiliencyImpactBytesTimestampUsec() bool`

HasLocalTierResiliencyImpactBytesTimestampUsec returns a boolean if a field has been set.

### SetLocalTierResiliencyImpactBytesTimestampUsecNil

`func (o *ViewsSummary) SetLocalTierResiliencyImpactBytesTimestampUsecNil(b bool)`

 SetLocalTierResiliencyImpactBytesTimestampUsecNil sets the value for LocalTierResiliencyImpactBytesTimestampUsec to be an explicit nil

### UnsetLocalTierResiliencyImpactBytesTimestampUsec
`func (o *ViewsSummary) UnsetLocalTierResiliencyImpactBytesTimestampUsec()`

UnsetLocalTierResiliencyImpactBytesTimestampUsec ensures that no value is present for LocalTierResiliencyImpactBytesTimestampUsec, not even an explicit nil
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
### GetLogicalUsageBytesPrev

`func (o *ViewsSummary) GetLogicalUsageBytesPrev() int64`

GetLogicalUsageBytesPrev returns the LogicalUsageBytesPrev field if non-nil, zero value otherwise.

### GetLogicalUsageBytesPrevOk

`func (o *ViewsSummary) GetLogicalUsageBytesPrevOk() (*int64, bool)`

GetLogicalUsageBytesPrevOk returns a tuple with the LogicalUsageBytesPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalUsageBytesPrev

`func (o *ViewsSummary) SetLogicalUsageBytesPrev(v int64)`

SetLogicalUsageBytesPrev sets LogicalUsageBytesPrev field to given value.

### HasLogicalUsageBytesPrev

`func (o *ViewsSummary) HasLogicalUsageBytesPrev() bool`

HasLogicalUsageBytesPrev returns a boolean if a field has been set.

### SetLogicalUsageBytesPrevNil

`func (o *ViewsSummary) SetLogicalUsageBytesPrevNil(b bool)`

 SetLogicalUsageBytesPrevNil sets the value for LogicalUsageBytesPrev to be an explicit nil

### UnsetLogicalUsageBytesPrev
`func (o *ViewsSummary) UnsetLogicalUsageBytesPrev()`

UnsetLogicalUsageBytesPrev ensures that no value is present for LogicalUsageBytesPrev, not even an explicit nil
### GetLogicalUsageBytesPrevTimestampUsec

`func (o *ViewsSummary) GetLogicalUsageBytesPrevTimestampUsec() int64`

GetLogicalUsageBytesPrevTimestampUsec returns the LogicalUsageBytesPrevTimestampUsec field if non-nil, zero value otherwise.

### GetLogicalUsageBytesPrevTimestampUsecOk

`func (o *ViewsSummary) GetLogicalUsageBytesPrevTimestampUsecOk() (*int64, bool)`

GetLogicalUsageBytesPrevTimestampUsecOk returns a tuple with the LogicalUsageBytesPrevTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalUsageBytesPrevTimestampUsec

`func (o *ViewsSummary) SetLogicalUsageBytesPrevTimestampUsec(v int64)`

SetLogicalUsageBytesPrevTimestampUsec sets LogicalUsageBytesPrevTimestampUsec field to given value.

### HasLogicalUsageBytesPrevTimestampUsec

`func (o *ViewsSummary) HasLogicalUsageBytesPrevTimestampUsec() bool`

HasLogicalUsageBytesPrevTimestampUsec returns a boolean if a field has been set.

### SetLogicalUsageBytesPrevTimestampUsecNil

`func (o *ViewsSummary) SetLogicalUsageBytesPrevTimestampUsecNil(b bool)`

 SetLogicalUsageBytesPrevTimestampUsecNil sets the value for LogicalUsageBytesPrevTimestampUsec to be an explicit nil

### UnsetLogicalUsageBytesPrevTimestampUsec
`func (o *ViewsSummary) UnsetLogicalUsageBytesPrevTimestampUsec()`

UnsetLogicalUsageBytesPrevTimestampUsec ensures that no value is present for LogicalUsageBytesPrevTimestampUsec, not even an explicit nil
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
### GetNumDirectories

`func (o *ViewsSummary) GetNumDirectories() int64`

GetNumDirectories returns the NumDirectories field if non-nil, zero value otherwise.

### GetNumDirectoriesOk

`func (o *ViewsSummary) GetNumDirectoriesOk() (*int64, bool)`

GetNumDirectoriesOk returns a tuple with the NumDirectories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDirectories

`func (o *ViewsSummary) SetNumDirectories(v int64)`

SetNumDirectories sets NumDirectories field to given value.

### HasNumDirectories

`func (o *ViewsSummary) HasNumDirectories() bool`

HasNumDirectories returns a boolean if a field has been set.

### SetNumDirectoriesNil

`func (o *ViewsSummary) SetNumDirectoriesNil(b bool)`

 SetNumDirectoriesNil sets the value for NumDirectories to be an explicit nil

### UnsetNumDirectories
`func (o *ViewsSummary) UnsetNumDirectories()`

UnsetNumDirectories ensures that no value is present for NumDirectories, not even an explicit nil
### GetNumDirectoriesPrev

`func (o *ViewsSummary) GetNumDirectoriesPrev() int64`

GetNumDirectoriesPrev returns the NumDirectoriesPrev field if non-nil, zero value otherwise.

### GetNumDirectoriesPrevOk

`func (o *ViewsSummary) GetNumDirectoriesPrevOk() (*int64, bool)`

GetNumDirectoriesPrevOk returns a tuple with the NumDirectoriesPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDirectoriesPrev

`func (o *ViewsSummary) SetNumDirectoriesPrev(v int64)`

SetNumDirectoriesPrev sets NumDirectoriesPrev field to given value.

### HasNumDirectoriesPrev

`func (o *ViewsSummary) HasNumDirectoriesPrev() bool`

HasNumDirectoriesPrev returns a boolean if a field has been set.

### SetNumDirectoriesPrevNil

`func (o *ViewsSummary) SetNumDirectoriesPrevNil(b bool)`

 SetNumDirectoriesPrevNil sets the value for NumDirectoriesPrev to be an explicit nil

### UnsetNumDirectoriesPrev
`func (o *ViewsSummary) UnsetNumDirectoriesPrev()`

UnsetNumDirectoriesPrev ensures that no value is present for NumDirectoriesPrev, not even an explicit nil
### GetNumFiles

`func (o *ViewsSummary) GetNumFiles() int64`

GetNumFiles returns the NumFiles field if non-nil, zero value otherwise.

### GetNumFilesOk

`func (o *ViewsSummary) GetNumFilesOk() (*int64, bool)`

GetNumFilesOk returns a tuple with the NumFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFiles

`func (o *ViewsSummary) SetNumFiles(v int64)`

SetNumFiles sets NumFiles field to given value.

### HasNumFiles

`func (o *ViewsSummary) HasNumFiles() bool`

HasNumFiles returns a boolean if a field has been set.

### SetNumFilesNil

`func (o *ViewsSummary) SetNumFilesNil(b bool)`

 SetNumFilesNil sets the value for NumFiles to be an explicit nil

### UnsetNumFiles
`func (o *ViewsSummary) UnsetNumFiles()`

UnsetNumFiles ensures that no value is present for NumFiles, not even an explicit nil
### GetNumFilesPrev

`func (o *ViewsSummary) GetNumFilesPrev() int64`

GetNumFilesPrev returns the NumFilesPrev field if non-nil, zero value otherwise.

### GetNumFilesPrevOk

`func (o *ViewsSummary) GetNumFilesPrevOk() (*int64, bool)`

GetNumFilesPrevOk returns a tuple with the NumFilesPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFilesPrev

`func (o *ViewsSummary) SetNumFilesPrev(v int64)`

SetNumFilesPrev sets NumFilesPrev field to given value.

### HasNumFilesPrev

`func (o *ViewsSummary) HasNumFilesPrev() bool`

HasNumFilesPrev returns a boolean if a field has been set.

### SetNumFilesPrevNil

`func (o *ViewsSummary) SetNumFilesPrevNil(b bool)`

 SetNumFilesPrevNil sets the value for NumFilesPrev to be an explicit nil

### UnsetNumFilesPrev
`func (o *ViewsSummary) UnsetNumFilesPrev()`

UnsetNumFilesPrev ensures that no value is present for NumFilesPrev, not even an explicit nil
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
### GetStorageConsumedBytesPrev

`func (o *ViewsSummary) GetStorageConsumedBytesPrev() int64`

GetStorageConsumedBytesPrev returns the StorageConsumedBytesPrev field if non-nil, zero value otherwise.

### GetStorageConsumedBytesPrevOk

`func (o *ViewsSummary) GetStorageConsumedBytesPrevOk() (*int64, bool)`

GetStorageConsumedBytesPrevOk returns a tuple with the StorageConsumedBytesPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConsumedBytesPrev

`func (o *ViewsSummary) SetStorageConsumedBytesPrev(v int64)`

SetStorageConsumedBytesPrev sets StorageConsumedBytesPrev field to given value.

### HasStorageConsumedBytesPrev

`func (o *ViewsSummary) HasStorageConsumedBytesPrev() bool`

HasStorageConsumedBytesPrev returns a boolean if a field has been set.

### SetStorageConsumedBytesPrevNil

`func (o *ViewsSummary) SetStorageConsumedBytesPrevNil(b bool)`

 SetStorageConsumedBytesPrevNil sets the value for StorageConsumedBytesPrev to be an explicit nil

### UnsetStorageConsumedBytesPrev
`func (o *ViewsSummary) UnsetStorageConsumedBytesPrev()`

UnsetStorageConsumedBytesPrev ensures that no value is present for StorageConsumedBytesPrev, not even an explicit nil
### GetStorageConsumedBytesPrevTimestampUsec

`func (o *ViewsSummary) GetStorageConsumedBytesPrevTimestampUsec() int64`

GetStorageConsumedBytesPrevTimestampUsec returns the StorageConsumedBytesPrevTimestampUsec field if non-nil, zero value otherwise.

### GetStorageConsumedBytesPrevTimestampUsecOk

`func (o *ViewsSummary) GetStorageConsumedBytesPrevTimestampUsecOk() (*int64, bool)`

GetStorageConsumedBytesPrevTimestampUsecOk returns a tuple with the StorageConsumedBytesPrevTimestampUsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConsumedBytesPrevTimestampUsec

`func (o *ViewsSummary) SetStorageConsumedBytesPrevTimestampUsec(v int64)`

SetStorageConsumedBytesPrevTimestampUsec sets StorageConsumedBytesPrevTimestampUsec field to given value.

### HasStorageConsumedBytesPrevTimestampUsec

`func (o *ViewsSummary) HasStorageConsumedBytesPrevTimestampUsec() bool`

HasStorageConsumedBytesPrevTimestampUsec returns a boolean if a field has been set.

### SetStorageConsumedBytesPrevTimestampUsecNil

`func (o *ViewsSummary) SetStorageConsumedBytesPrevTimestampUsecNil(b bool)`

 SetStorageConsumedBytesPrevTimestampUsecNil sets the value for StorageConsumedBytesPrevTimestampUsec to be an explicit nil

### UnsetStorageConsumedBytesPrevTimestampUsec
`func (o *ViewsSummary) UnsetStorageConsumedBytesPrevTimestampUsec()`

UnsetStorageConsumedBytesPrevTimestampUsec ensures that no value is present for StorageConsumedBytesPrevTimestampUsec, not even an explicit nil
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
### GetViewEntityId

`func (o *ViewsSummary) GetViewEntityId() string`

GetViewEntityId returns the ViewEntityId field if non-nil, zero value otherwise.

### GetViewEntityIdOk

`func (o *ViewsSummary) GetViewEntityIdOk() (*string, bool)`

GetViewEntityIdOk returns a tuple with the ViewEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewEntityId

`func (o *ViewsSummary) SetViewEntityId(v string)`

SetViewEntityId sets ViewEntityId field to given value.

### HasViewEntityId

`func (o *ViewsSummary) HasViewEntityId() bool`

HasViewEntityId returns a boolean if a field has been set.

### SetViewEntityIdNil

`func (o *ViewsSummary) SetViewEntityIdNil(b bool)`

 SetViewEntityIdNil sets the value for ViewEntityId to be an explicit nil

### UnsetViewEntityId
`func (o *ViewsSummary) UnsetViewEntityId()`

UnsetViewEntityId ensures that no value is present for ViewEntityId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


