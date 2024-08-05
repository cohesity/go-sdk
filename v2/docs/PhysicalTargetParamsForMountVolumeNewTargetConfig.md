# PhysicalTargetParamsForMountVolumeNewTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountTarget** | [**RecoverTarget**](RecoverTarget.md) |  | 
**ServerCredentials** | Pointer to [**NullablePhysicalMountVolumesNewTargetConfigServerCredentials**](PhysicalMountVolumesNewTargetConfigServerCredentials.md) |  | [optional] 

## Methods

### NewPhysicalTargetParamsForMountVolumeNewTargetConfig

`func NewPhysicalTargetParamsForMountVolumeNewTargetConfig(mountTarget RecoverTarget, ) *PhysicalTargetParamsForMountVolumeNewTargetConfig`

NewPhysicalTargetParamsForMountVolumeNewTargetConfig instantiates a new PhysicalTargetParamsForMountVolumeNewTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalTargetParamsForMountVolumeNewTargetConfigWithDefaults

`func NewPhysicalTargetParamsForMountVolumeNewTargetConfigWithDefaults() *PhysicalTargetParamsForMountVolumeNewTargetConfig`

NewPhysicalTargetParamsForMountVolumeNewTargetConfigWithDefaults instantiates a new PhysicalTargetParamsForMountVolumeNewTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountTarget

`func (o *PhysicalTargetParamsForMountVolumeNewTargetConfig) GetMountTarget() RecoverTarget`

GetMountTarget returns the MountTarget field if non-nil, zero value otherwise.

### GetMountTargetOk

`func (o *PhysicalTargetParamsForMountVolumeNewTargetConfig) GetMountTargetOk() (*RecoverTarget, bool)`

GetMountTargetOk returns a tuple with the MountTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountTarget

`func (o *PhysicalTargetParamsForMountVolumeNewTargetConfig) SetMountTarget(v RecoverTarget)`

SetMountTarget sets MountTarget field to given value.


### GetServerCredentials

`func (o *PhysicalTargetParamsForMountVolumeNewTargetConfig) GetServerCredentials() PhysicalMountVolumesNewTargetConfigServerCredentials`

GetServerCredentials returns the ServerCredentials field if non-nil, zero value otherwise.

### GetServerCredentialsOk

`func (o *PhysicalTargetParamsForMountVolumeNewTargetConfig) GetServerCredentialsOk() (*PhysicalMountVolumesNewTargetConfigServerCredentials, bool)`

GetServerCredentialsOk returns a tuple with the ServerCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCredentials

`func (o *PhysicalTargetParamsForMountVolumeNewTargetConfig) SetServerCredentials(v PhysicalMountVolumesNewTargetConfigServerCredentials)`

SetServerCredentials sets ServerCredentials field to given value.

### HasServerCredentials

`func (o *PhysicalTargetParamsForMountVolumeNewTargetConfig) HasServerCredentials() bool`

HasServerCredentials returns a boolean if a field has been set.

### SetServerCredentialsNil

`func (o *PhysicalTargetParamsForMountVolumeNewTargetConfig) SetServerCredentialsNil(b bool)`

 SetServerCredentialsNil sets the value for ServerCredentials to be an explicit nil

### UnsetServerCredentials
`func (o *PhysicalTargetParamsForMountVolumeNewTargetConfig) UnsetServerCredentials()`

UnsetServerCredentials ensures that no value is present for ServerCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


