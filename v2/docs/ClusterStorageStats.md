# ClusterStorageStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataProtectionLogicalUsageBytes** | Pointer to **NullableInt64** | Specifies the logical size of protected objects in bytes. | [optional] 
**DataProtectionPhysicalUsageBytes** | Pointer to **NullableInt64** | Specifies the physical size of protected objects in bytes. | [optional] 
**FileServicesLogicalUsageBytes** | Pointer to **NullableInt64** | Specifies the logical size consumed by file services in bytes. | [optional] 
**FileServicesPhysicalUsageBytes** | Pointer to **NullableInt64** | Specifies the physical size consumed by file services in bytes. | [optional] 
**LocalAvailableBytes** | Pointer to **NullableInt64** | Specifies the local storage currently available on the cluster in bytes. | [optional] 
**LocalUsageBytes** | Pointer to **NullableInt64** | Specifies the local storage currently in use on the cluster in bytes. | [optional] 
**TotalCapacityBytes** | Pointer to **NullableInt64** | Specifies the total capacity of the cluster in bytes. | [optional] 

## Methods

### NewClusterStorageStats

`func NewClusterStorageStats() *ClusterStorageStats`

NewClusterStorageStats instantiates a new ClusterStorageStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStorageStatsWithDefaults

`func NewClusterStorageStatsWithDefaults() *ClusterStorageStats`

NewClusterStorageStatsWithDefaults instantiates a new ClusterStorageStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataProtectionLogicalUsageBytes

`func (o *ClusterStorageStats) GetDataProtectionLogicalUsageBytes() int64`

GetDataProtectionLogicalUsageBytes returns the DataProtectionLogicalUsageBytes field if non-nil, zero value otherwise.

### GetDataProtectionLogicalUsageBytesOk

`func (o *ClusterStorageStats) GetDataProtectionLogicalUsageBytesOk() (*int64, bool)`

GetDataProtectionLogicalUsageBytesOk returns a tuple with the DataProtectionLogicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtectionLogicalUsageBytes

`func (o *ClusterStorageStats) SetDataProtectionLogicalUsageBytes(v int64)`

SetDataProtectionLogicalUsageBytes sets DataProtectionLogicalUsageBytes field to given value.

### HasDataProtectionLogicalUsageBytes

`func (o *ClusterStorageStats) HasDataProtectionLogicalUsageBytes() bool`

HasDataProtectionLogicalUsageBytes returns a boolean if a field has been set.

### SetDataProtectionLogicalUsageBytesNil

`func (o *ClusterStorageStats) SetDataProtectionLogicalUsageBytesNil(b bool)`

 SetDataProtectionLogicalUsageBytesNil sets the value for DataProtectionLogicalUsageBytes to be an explicit nil

### UnsetDataProtectionLogicalUsageBytes
`func (o *ClusterStorageStats) UnsetDataProtectionLogicalUsageBytes()`

UnsetDataProtectionLogicalUsageBytes ensures that no value is present for DataProtectionLogicalUsageBytes, not even an explicit nil
### GetDataProtectionPhysicalUsageBytes

`func (o *ClusterStorageStats) GetDataProtectionPhysicalUsageBytes() int64`

GetDataProtectionPhysicalUsageBytes returns the DataProtectionPhysicalUsageBytes field if non-nil, zero value otherwise.

### GetDataProtectionPhysicalUsageBytesOk

`func (o *ClusterStorageStats) GetDataProtectionPhysicalUsageBytesOk() (*int64, bool)`

GetDataProtectionPhysicalUsageBytesOk returns a tuple with the DataProtectionPhysicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtectionPhysicalUsageBytes

`func (o *ClusterStorageStats) SetDataProtectionPhysicalUsageBytes(v int64)`

SetDataProtectionPhysicalUsageBytes sets DataProtectionPhysicalUsageBytes field to given value.

### HasDataProtectionPhysicalUsageBytes

`func (o *ClusterStorageStats) HasDataProtectionPhysicalUsageBytes() bool`

HasDataProtectionPhysicalUsageBytes returns a boolean if a field has been set.

### SetDataProtectionPhysicalUsageBytesNil

`func (o *ClusterStorageStats) SetDataProtectionPhysicalUsageBytesNil(b bool)`

 SetDataProtectionPhysicalUsageBytesNil sets the value for DataProtectionPhysicalUsageBytes to be an explicit nil

### UnsetDataProtectionPhysicalUsageBytes
`func (o *ClusterStorageStats) UnsetDataProtectionPhysicalUsageBytes()`

UnsetDataProtectionPhysicalUsageBytes ensures that no value is present for DataProtectionPhysicalUsageBytes, not even an explicit nil
### GetFileServicesLogicalUsageBytes

`func (o *ClusterStorageStats) GetFileServicesLogicalUsageBytes() int64`

GetFileServicesLogicalUsageBytes returns the FileServicesLogicalUsageBytes field if non-nil, zero value otherwise.

### GetFileServicesLogicalUsageBytesOk

`func (o *ClusterStorageStats) GetFileServicesLogicalUsageBytesOk() (*int64, bool)`

GetFileServicesLogicalUsageBytesOk returns a tuple with the FileServicesLogicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileServicesLogicalUsageBytes

`func (o *ClusterStorageStats) SetFileServicesLogicalUsageBytes(v int64)`

SetFileServicesLogicalUsageBytes sets FileServicesLogicalUsageBytes field to given value.

### HasFileServicesLogicalUsageBytes

`func (o *ClusterStorageStats) HasFileServicesLogicalUsageBytes() bool`

HasFileServicesLogicalUsageBytes returns a boolean if a field has been set.

### SetFileServicesLogicalUsageBytesNil

`func (o *ClusterStorageStats) SetFileServicesLogicalUsageBytesNil(b bool)`

 SetFileServicesLogicalUsageBytesNil sets the value for FileServicesLogicalUsageBytes to be an explicit nil

### UnsetFileServicesLogicalUsageBytes
`func (o *ClusterStorageStats) UnsetFileServicesLogicalUsageBytes()`

UnsetFileServicesLogicalUsageBytes ensures that no value is present for FileServicesLogicalUsageBytes, not even an explicit nil
### GetFileServicesPhysicalUsageBytes

`func (o *ClusterStorageStats) GetFileServicesPhysicalUsageBytes() int64`

GetFileServicesPhysicalUsageBytes returns the FileServicesPhysicalUsageBytes field if non-nil, zero value otherwise.

### GetFileServicesPhysicalUsageBytesOk

`func (o *ClusterStorageStats) GetFileServicesPhysicalUsageBytesOk() (*int64, bool)`

GetFileServicesPhysicalUsageBytesOk returns a tuple with the FileServicesPhysicalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileServicesPhysicalUsageBytes

`func (o *ClusterStorageStats) SetFileServicesPhysicalUsageBytes(v int64)`

SetFileServicesPhysicalUsageBytes sets FileServicesPhysicalUsageBytes field to given value.

### HasFileServicesPhysicalUsageBytes

`func (o *ClusterStorageStats) HasFileServicesPhysicalUsageBytes() bool`

HasFileServicesPhysicalUsageBytes returns a boolean if a field has been set.

### SetFileServicesPhysicalUsageBytesNil

`func (o *ClusterStorageStats) SetFileServicesPhysicalUsageBytesNil(b bool)`

 SetFileServicesPhysicalUsageBytesNil sets the value for FileServicesPhysicalUsageBytes to be an explicit nil

### UnsetFileServicesPhysicalUsageBytes
`func (o *ClusterStorageStats) UnsetFileServicesPhysicalUsageBytes()`

UnsetFileServicesPhysicalUsageBytes ensures that no value is present for FileServicesPhysicalUsageBytes, not even an explicit nil
### GetLocalAvailableBytes

`func (o *ClusterStorageStats) GetLocalAvailableBytes() int64`

GetLocalAvailableBytes returns the LocalAvailableBytes field if non-nil, zero value otherwise.

### GetLocalAvailableBytesOk

`func (o *ClusterStorageStats) GetLocalAvailableBytesOk() (*int64, bool)`

GetLocalAvailableBytesOk returns a tuple with the LocalAvailableBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAvailableBytes

`func (o *ClusterStorageStats) SetLocalAvailableBytes(v int64)`

SetLocalAvailableBytes sets LocalAvailableBytes field to given value.

### HasLocalAvailableBytes

`func (o *ClusterStorageStats) HasLocalAvailableBytes() bool`

HasLocalAvailableBytes returns a boolean if a field has been set.

### SetLocalAvailableBytesNil

`func (o *ClusterStorageStats) SetLocalAvailableBytesNil(b bool)`

 SetLocalAvailableBytesNil sets the value for LocalAvailableBytes to be an explicit nil

### UnsetLocalAvailableBytes
`func (o *ClusterStorageStats) UnsetLocalAvailableBytes()`

UnsetLocalAvailableBytes ensures that no value is present for LocalAvailableBytes, not even an explicit nil
### GetLocalUsageBytes

`func (o *ClusterStorageStats) GetLocalUsageBytes() int64`

GetLocalUsageBytes returns the LocalUsageBytes field if non-nil, zero value otherwise.

### GetLocalUsageBytesOk

`func (o *ClusterStorageStats) GetLocalUsageBytesOk() (*int64, bool)`

GetLocalUsageBytesOk returns a tuple with the LocalUsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUsageBytes

`func (o *ClusterStorageStats) SetLocalUsageBytes(v int64)`

SetLocalUsageBytes sets LocalUsageBytes field to given value.

### HasLocalUsageBytes

`func (o *ClusterStorageStats) HasLocalUsageBytes() bool`

HasLocalUsageBytes returns a boolean if a field has been set.

### SetLocalUsageBytesNil

`func (o *ClusterStorageStats) SetLocalUsageBytesNil(b bool)`

 SetLocalUsageBytesNil sets the value for LocalUsageBytes to be an explicit nil

### UnsetLocalUsageBytes
`func (o *ClusterStorageStats) UnsetLocalUsageBytes()`

UnsetLocalUsageBytes ensures that no value is present for LocalUsageBytes, not even an explicit nil
### GetTotalCapacityBytes

`func (o *ClusterStorageStats) GetTotalCapacityBytes() int64`

GetTotalCapacityBytes returns the TotalCapacityBytes field if non-nil, zero value otherwise.

### GetTotalCapacityBytesOk

`func (o *ClusterStorageStats) GetTotalCapacityBytesOk() (*int64, bool)`

GetTotalCapacityBytesOk returns a tuple with the TotalCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacityBytes

`func (o *ClusterStorageStats) SetTotalCapacityBytes(v int64)`

SetTotalCapacityBytes sets TotalCapacityBytes field to given value.

### HasTotalCapacityBytes

`func (o *ClusterStorageStats) HasTotalCapacityBytes() bool`

HasTotalCapacityBytes returns a boolean if a field has been set.

### SetTotalCapacityBytesNil

`func (o *ClusterStorageStats) SetTotalCapacityBytesNil(b bool)`

 SetTotalCapacityBytesNil sets the value for TotalCapacityBytes to be an explicit nil

### UnsetTotalCapacityBytes
`func (o *ClusterStorageStats) UnsetTotalCapacityBytes()`

UnsetTotalCapacityBytes ensures that no value is present for TotalCapacityBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


