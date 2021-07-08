# RestoreAcropolisVMParamNetworkConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkStateChange** | Pointer to **NullableInt32** | Network state to be applied to the restored VM. | [optional] 
**NicVec** | Pointer to [**[]RestoreAcropolisVMParamNetworkConfigInfoNICSpec**](RestoreAcropolisVMParamNetworkConfigInfoNICSpec.md) | This field is applicable only if the network_state_change is set to &#39;kAttachNewNetwork&#39;. | [optional] 

## Methods

### NewRestoreAcropolisVMParamNetworkConfigInfo

`func NewRestoreAcropolisVMParamNetworkConfigInfo() *RestoreAcropolisVMParamNetworkConfigInfo`

NewRestoreAcropolisVMParamNetworkConfigInfo instantiates a new RestoreAcropolisVMParamNetworkConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreAcropolisVMParamNetworkConfigInfoWithDefaults

`func NewRestoreAcropolisVMParamNetworkConfigInfoWithDefaults() *RestoreAcropolisVMParamNetworkConfigInfo`

NewRestoreAcropolisVMParamNetworkConfigInfoWithDefaults instantiates a new RestoreAcropolisVMParamNetworkConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkStateChange

`func (o *RestoreAcropolisVMParamNetworkConfigInfo) GetNetworkStateChange() int32`

GetNetworkStateChange returns the NetworkStateChange field if non-nil, zero value otherwise.

### GetNetworkStateChangeOk

`func (o *RestoreAcropolisVMParamNetworkConfigInfo) GetNetworkStateChangeOk() (*int32, bool)`

GetNetworkStateChangeOk returns a tuple with the NetworkStateChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStateChange

`func (o *RestoreAcropolisVMParamNetworkConfigInfo) SetNetworkStateChange(v int32)`

SetNetworkStateChange sets NetworkStateChange field to given value.

### HasNetworkStateChange

`func (o *RestoreAcropolisVMParamNetworkConfigInfo) HasNetworkStateChange() bool`

HasNetworkStateChange returns a boolean if a field has been set.

### SetNetworkStateChangeNil

`func (o *RestoreAcropolisVMParamNetworkConfigInfo) SetNetworkStateChangeNil(b bool)`

 SetNetworkStateChangeNil sets the value for NetworkStateChange to be an explicit nil

### UnsetNetworkStateChange
`func (o *RestoreAcropolisVMParamNetworkConfigInfo) UnsetNetworkStateChange()`

UnsetNetworkStateChange ensures that no value is present for NetworkStateChange, not even an explicit nil
### GetNicVec

`func (o *RestoreAcropolisVMParamNetworkConfigInfo) GetNicVec() []RestoreAcropolisVMParamNetworkConfigInfoNICSpec`

GetNicVec returns the NicVec field if non-nil, zero value otherwise.

### GetNicVecOk

`func (o *RestoreAcropolisVMParamNetworkConfigInfo) GetNicVecOk() (*[]RestoreAcropolisVMParamNetworkConfigInfoNICSpec, bool)`

GetNicVecOk returns a tuple with the NicVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicVec

`func (o *RestoreAcropolisVMParamNetworkConfigInfo) SetNicVec(v []RestoreAcropolisVMParamNetworkConfigInfoNICSpec)`

SetNicVec sets NicVec field to given value.

### HasNicVec

`func (o *RestoreAcropolisVMParamNetworkConfigInfo) HasNicVec() bool`

HasNicVec returns a boolean if a field has been set.

### SetNicVecNil

`func (o *RestoreAcropolisVMParamNetworkConfigInfo) SetNicVecNil(b bool)`

 SetNicVecNil sets the value for NicVec to be an explicit nil

### UnsetNicVec
`func (o *RestoreAcropolisVMParamNetworkConfigInfo) UnsetNicVec()`

UnsetNicVec ensures that no value is present for NicVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


