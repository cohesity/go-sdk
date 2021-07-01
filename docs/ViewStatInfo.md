# ViewStatInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | Specifies the cluster Id. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the cluster Incarnation Id. | [optional] 
**DataReadBytes** | Pointer to **NullableInt64** | Specifies the data read in bytes. | [optional] 
**DataWrittenBytes** | Pointer to **NullableInt64** | Specifies the data written in bytes. | [optional] 
**LogicalUsedBytes** | Pointer to **NullableInt64** | Specifies the logical size used in bytes. | [optional] 
**PeakReadThroughput** | Pointer to **NullableInt64** | Specifies the peak data read in bytes per second in the last day. | [optional] 
**PeakWriteThroughput** | Pointer to **NullableInt64** | Specifies the peak data written in bytes per second in the last day. | [optional] 
**PhysicalUsedBytes** | Pointer to **NullableInt64** | Specifies the physical size used in bytes. | [optional] 
**Protocols** | Pointer to **[]string** | Specifies the protocols of this view. | [optional] 
**StorageReductionRatio** | Pointer to **NullableFloat32** | Specifies the storage reduction ratio. | [optional] 
**ViewId** | Pointer to **NullableInt64** | Specifies the view Id. | [optional] 
**ViewName** | Pointer to **NullableString** | Specifies the view name. | [optional] 

## Methods

### NewViewStatInfo

`func NewViewStatInfo() *ViewStatInfo`

NewViewStatInfo instantiates a new ViewStatInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewStatInfoWithDefaults

`func NewViewStatInfoWithDefaults() *ViewStatInfo`

NewViewStatInfoWithDefaults instantiates a new ViewStatInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ViewStatInfo) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ViewStatInfo) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ViewStatInfo) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ViewStatInfo) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *ViewStatInfo) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *ViewStatInfo) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *ViewStatInfo) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *ViewStatInfo) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *ViewStatInfo) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *ViewStatInfo) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *ViewStatInfo) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *ViewStatInfo) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetDataReadBytes

`func (o *ViewStatInfo) GetDataReadBytes() int64`

GetDataReadBytes returns the DataReadBytes field if non-nil, zero value otherwise.

### GetDataReadBytesOk

`func (o *ViewStatInfo) GetDataReadBytesOk() (*int64, bool)`

GetDataReadBytesOk returns a tuple with the DataReadBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReadBytes

`func (o *ViewStatInfo) SetDataReadBytes(v int64)`

SetDataReadBytes sets DataReadBytes field to given value.

### HasDataReadBytes

`func (o *ViewStatInfo) HasDataReadBytes() bool`

HasDataReadBytes returns a boolean if a field has been set.

### SetDataReadBytesNil

`func (o *ViewStatInfo) SetDataReadBytesNil(b bool)`

 SetDataReadBytesNil sets the value for DataReadBytes to be an explicit nil

### UnsetDataReadBytes
`func (o *ViewStatInfo) UnsetDataReadBytes()`

UnsetDataReadBytes ensures that no value is present for DataReadBytes, not even an explicit nil
### GetDataWrittenBytes

`func (o *ViewStatInfo) GetDataWrittenBytes() int64`

GetDataWrittenBytes returns the DataWrittenBytes field if non-nil, zero value otherwise.

### GetDataWrittenBytesOk

`func (o *ViewStatInfo) GetDataWrittenBytesOk() (*int64, bool)`

GetDataWrittenBytesOk returns a tuple with the DataWrittenBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWrittenBytes

`func (o *ViewStatInfo) SetDataWrittenBytes(v int64)`

SetDataWrittenBytes sets DataWrittenBytes field to given value.

### HasDataWrittenBytes

`func (o *ViewStatInfo) HasDataWrittenBytes() bool`

HasDataWrittenBytes returns a boolean if a field has been set.

### SetDataWrittenBytesNil

`func (o *ViewStatInfo) SetDataWrittenBytesNil(b bool)`

 SetDataWrittenBytesNil sets the value for DataWrittenBytes to be an explicit nil

### UnsetDataWrittenBytes
`func (o *ViewStatInfo) UnsetDataWrittenBytes()`

UnsetDataWrittenBytes ensures that no value is present for DataWrittenBytes, not even an explicit nil
### GetLogicalUsedBytes

`func (o *ViewStatInfo) GetLogicalUsedBytes() int64`

GetLogicalUsedBytes returns the LogicalUsedBytes field if non-nil, zero value otherwise.

### GetLogicalUsedBytesOk

`func (o *ViewStatInfo) GetLogicalUsedBytesOk() (*int64, bool)`

GetLogicalUsedBytesOk returns a tuple with the LogicalUsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalUsedBytes

`func (o *ViewStatInfo) SetLogicalUsedBytes(v int64)`

SetLogicalUsedBytes sets LogicalUsedBytes field to given value.

### HasLogicalUsedBytes

`func (o *ViewStatInfo) HasLogicalUsedBytes() bool`

HasLogicalUsedBytes returns a boolean if a field has been set.

### SetLogicalUsedBytesNil

`func (o *ViewStatInfo) SetLogicalUsedBytesNil(b bool)`

 SetLogicalUsedBytesNil sets the value for LogicalUsedBytes to be an explicit nil

### UnsetLogicalUsedBytes
`func (o *ViewStatInfo) UnsetLogicalUsedBytes()`

UnsetLogicalUsedBytes ensures that no value is present for LogicalUsedBytes, not even an explicit nil
### GetPeakReadThroughput

`func (o *ViewStatInfo) GetPeakReadThroughput() int64`

GetPeakReadThroughput returns the PeakReadThroughput field if non-nil, zero value otherwise.

### GetPeakReadThroughputOk

`func (o *ViewStatInfo) GetPeakReadThroughputOk() (*int64, bool)`

GetPeakReadThroughputOk returns a tuple with the PeakReadThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakReadThroughput

`func (o *ViewStatInfo) SetPeakReadThroughput(v int64)`

SetPeakReadThroughput sets PeakReadThroughput field to given value.

### HasPeakReadThroughput

`func (o *ViewStatInfo) HasPeakReadThroughput() bool`

HasPeakReadThroughput returns a boolean if a field has been set.

### SetPeakReadThroughputNil

`func (o *ViewStatInfo) SetPeakReadThroughputNil(b bool)`

 SetPeakReadThroughputNil sets the value for PeakReadThroughput to be an explicit nil

### UnsetPeakReadThroughput
`func (o *ViewStatInfo) UnsetPeakReadThroughput()`

UnsetPeakReadThroughput ensures that no value is present for PeakReadThroughput, not even an explicit nil
### GetPeakWriteThroughput

`func (o *ViewStatInfo) GetPeakWriteThroughput() int64`

GetPeakWriteThroughput returns the PeakWriteThroughput field if non-nil, zero value otherwise.

### GetPeakWriteThroughputOk

`func (o *ViewStatInfo) GetPeakWriteThroughputOk() (*int64, bool)`

GetPeakWriteThroughputOk returns a tuple with the PeakWriteThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakWriteThroughput

`func (o *ViewStatInfo) SetPeakWriteThroughput(v int64)`

SetPeakWriteThroughput sets PeakWriteThroughput field to given value.

### HasPeakWriteThroughput

`func (o *ViewStatInfo) HasPeakWriteThroughput() bool`

HasPeakWriteThroughput returns a boolean if a field has been set.

### SetPeakWriteThroughputNil

`func (o *ViewStatInfo) SetPeakWriteThroughputNil(b bool)`

 SetPeakWriteThroughputNil sets the value for PeakWriteThroughput to be an explicit nil

### UnsetPeakWriteThroughput
`func (o *ViewStatInfo) UnsetPeakWriteThroughput()`

UnsetPeakWriteThroughput ensures that no value is present for PeakWriteThroughput, not even an explicit nil
### GetPhysicalUsedBytes

`func (o *ViewStatInfo) GetPhysicalUsedBytes() int64`

GetPhysicalUsedBytes returns the PhysicalUsedBytes field if non-nil, zero value otherwise.

### GetPhysicalUsedBytesOk

`func (o *ViewStatInfo) GetPhysicalUsedBytesOk() (*int64, bool)`

GetPhysicalUsedBytesOk returns a tuple with the PhysicalUsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalUsedBytes

`func (o *ViewStatInfo) SetPhysicalUsedBytes(v int64)`

SetPhysicalUsedBytes sets PhysicalUsedBytes field to given value.

### HasPhysicalUsedBytes

`func (o *ViewStatInfo) HasPhysicalUsedBytes() bool`

HasPhysicalUsedBytes returns a boolean if a field has been set.

### SetPhysicalUsedBytesNil

`func (o *ViewStatInfo) SetPhysicalUsedBytesNil(b bool)`

 SetPhysicalUsedBytesNil sets the value for PhysicalUsedBytes to be an explicit nil

### UnsetPhysicalUsedBytes
`func (o *ViewStatInfo) UnsetPhysicalUsedBytes()`

UnsetPhysicalUsedBytes ensures that no value is present for PhysicalUsedBytes, not even an explicit nil
### GetProtocols

`func (o *ViewStatInfo) GetProtocols() []string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *ViewStatInfo) GetProtocolsOk() (*[]string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *ViewStatInfo) SetProtocols(v []string)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *ViewStatInfo) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### GetStorageReductionRatio

`func (o *ViewStatInfo) GetStorageReductionRatio() float32`

GetStorageReductionRatio returns the StorageReductionRatio field if non-nil, zero value otherwise.

### GetStorageReductionRatioOk

`func (o *ViewStatInfo) GetStorageReductionRatioOk() (*float32, bool)`

GetStorageReductionRatioOk returns a tuple with the StorageReductionRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageReductionRatio

`func (o *ViewStatInfo) SetStorageReductionRatio(v float32)`

SetStorageReductionRatio sets StorageReductionRatio field to given value.

### HasStorageReductionRatio

`func (o *ViewStatInfo) HasStorageReductionRatio() bool`

HasStorageReductionRatio returns a boolean if a field has been set.

### SetStorageReductionRatioNil

`func (o *ViewStatInfo) SetStorageReductionRatioNil(b bool)`

 SetStorageReductionRatioNil sets the value for StorageReductionRatio to be an explicit nil

### UnsetStorageReductionRatio
`func (o *ViewStatInfo) UnsetStorageReductionRatio()`

UnsetStorageReductionRatio ensures that no value is present for StorageReductionRatio, not even an explicit nil
### GetViewId

`func (o *ViewStatInfo) GetViewId() int64`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *ViewStatInfo) GetViewIdOk() (*int64, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *ViewStatInfo) SetViewId(v int64)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *ViewStatInfo) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### SetViewIdNil

`func (o *ViewStatInfo) SetViewIdNil(b bool)`

 SetViewIdNil sets the value for ViewId to be an explicit nil

### UnsetViewId
`func (o *ViewStatInfo) UnsetViewId()`

UnsetViewId ensures that no value is present for ViewId, not even an explicit nil
### GetViewName

`func (o *ViewStatInfo) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *ViewStatInfo) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *ViewStatInfo) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *ViewStatInfo) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *ViewStatInfo) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *ViewStatInfo) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


