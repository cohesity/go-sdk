# DataTransferInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPrivateNetwork** | Pointer to **NullableBool** | Specifies whether to use private network or public network. | [optional] 
**PrivateNetworkInfoList** | Pointer to [**[]PrivateNetworkInfo**](PrivateNetworkInfo.md) | Specifies Information required to create endpoints in private networks for all regions whose VMs are getting protected. | [optional] 
**UseProtectionJobInfo** | Pointer to **NullableBool** | Specifies Whether to use private network info which was used in backup of VMs.This should be populated only for restore job. | [optional] 

## Methods

### NewDataTransferInfo

`func NewDataTransferInfo() *DataTransferInfo`

NewDataTransferInfo instantiates a new DataTransferInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTransferInfoWithDefaults

`func NewDataTransferInfoWithDefaults() *DataTransferInfo`

NewDataTransferInfoWithDefaults instantiates a new DataTransferInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPrivateNetwork

`func (o *DataTransferInfo) GetIsPrivateNetwork() bool`

GetIsPrivateNetwork returns the IsPrivateNetwork field if non-nil, zero value otherwise.

### GetIsPrivateNetworkOk

`func (o *DataTransferInfo) GetIsPrivateNetworkOk() (*bool, bool)`

GetIsPrivateNetworkOk returns a tuple with the IsPrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivateNetwork

`func (o *DataTransferInfo) SetIsPrivateNetwork(v bool)`

SetIsPrivateNetwork sets IsPrivateNetwork field to given value.

### HasIsPrivateNetwork

`func (o *DataTransferInfo) HasIsPrivateNetwork() bool`

HasIsPrivateNetwork returns a boolean if a field has been set.

### SetIsPrivateNetworkNil

`func (o *DataTransferInfo) SetIsPrivateNetworkNil(b bool)`

 SetIsPrivateNetworkNil sets the value for IsPrivateNetwork to be an explicit nil

### UnsetIsPrivateNetwork
`func (o *DataTransferInfo) UnsetIsPrivateNetwork()`

UnsetIsPrivateNetwork ensures that no value is present for IsPrivateNetwork, not even an explicit nil
### GetPrivateNetworkInfoList

`func (o *DataTransferInfo) GetPrivateNetworkInfoList() []PrivateNetworkInfo`

GetPrivateNetworkInfoList returns the PrivateNetworkInfoList field if non-nil, zero value otherwise.

### GetPrivateNetworkInfoListOk

`func (o *DataTransferInfo) GetPrivateNetworkInfoListOk() (*[]PrivateNetworkInfo, bool)`

GetPrivateNetworkInfoListOk returns a tuple with the PrivateNetworkInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkInfoList

`func (o *DataTransferInfo) SetPrivateNetworkInfoList(v []PrivateNetworkInfo)`

SetPrivateNetworkInfoList sets PrivateNetworkInfoList field to given value.

### HasPrivateNetworkInfoList

`func (o *DataTransferInfo) HasPrivateNetworkInfoList() bool`

HasPrivateNetworkInfoList returns a boolean if a field has been set.

### GetUseProtectionJobInfo

`func (o *DataTransferInfo) GetUseProtectionJobInfo() bool`

GetUseProtectionJobInfo returns the UseProtectionJobInfo field if non-nil, zero value otherwise.

### GetUseProtectionJobInfoOk

`func (o *DataTransferInfo) GetUseProtectionJobInfoOk() (*bool, bool)`

GetUseProtectionJobInfoOk returns a tuple with the UseProtectionJobInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProtectionJobInfo

`func (o *DataTransferInfo) SetUseProtectionJobInfo(v bool)`

SetUseProtectionJobInfo sets UseProtectionJobInfo field to given value.

### HasUseProtectionJobInfo

`func (o *DataTransferInfo) HasUseProtectionJobInfo() bool`

HasUseProtectionJobInfo returns a boolean if a field has been set.

### SetUseProtectionJobInfoNil

`func (o *DataTransferInfo) SetUseProtectionJobInfoNil(b bool)`

 SetUseProtectionJobInfoNil sets the value for UseProtectionJobInfo to be an explicit nil

### UnsetUseProtectionJobInfo
`func (o *DataTransferInfo) UnsetUseProtectionJobInfo()`

UnsetUseProtectionJobInfo ensures that no value is present for UseProtectionJobInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


