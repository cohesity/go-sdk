# PhysicalMountVolumesNewTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountTarget** | [**RecoverTarget**](RecoverTarget.md) |  | 
**ServerCredentials** | Pointer to [**NullablePhysicalMountVolumesNewTargetConfigServerCredentials**](PhysicalMountVolumesNewTargetConfigServerCredentials.md) |  | [optional] 

## Methods

### NewPhysicalMountVolumesNewTargetConfig

`func NewPhysicalMountVolumesNewTargetConfig(mountTarget RecoverTarget, ) *PhysicalMountVolumesNewTargetConfig`

NewPhysicalMountVolumesNewTargetConfig instantiates a new PhysicalMountVolumesNewTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalMountVolumesNewTargetConfigWithDefaults

`func NewPhysicalMountVolumesNewTargetConfigWithDefaults() *PhysicalMountVolumesNewTargetConfig`

NewPhysicalMountVolumesNewTargetConfigWithDefaults instantiates a new PhysicalMountVolumesNewTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountTarget

`func (o *PhysicalMountVolumesNewTargetConfig) GetMountTarget() RecoverTarget`

GetMountTarget returns the MountTarget field if non-nil, zero value otherwise.

### GetMountTargetOk

`func (o *PhysicalMountVolumesNewTargetConfig) GetMountTargetOk() (*RecoverTarget, bool)`

GetMountTargetOk returns a tuple with the MountTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountTarget

`func (o *PhysicalMountVolumesNewTargetConfig) SetMountTarget(v RecoverTarget)`

SetMountTarget sets MountTarget field to given value.


### GetServerCredentials

`func (o *PhysicalMountVolumesNewTargetConfig) GetServerCredentials() PhysicalMountVolumesNewTargetConfigServerCredentials`

GetServerCredentials returns the ServerCredentials field if non-nil, zero value otherwise.

### GetServerCredentialsOk

`func (o *PhysicalMountVolumesNewTargetConfig) GetServerCredentialsOk() (*PhysicalMountVolumesNewTargetConfigServerCredentials, bool)`

GetServerCredentialsOk returns a tuple with the ServerCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCredentials

`func (o *PhysicalMountVolumesNewTargetConfig) SetServerCredentials(v PhysicalMountVolumesNewTargetConfigServerCredentials)`

SetServerCredentials sets ServerCredentials field to given value.

### HasServerCredentials

`func (o *PhysicalMountVolumesNewTargetConfig) HasServerCredentials() bool`

HasServerCredentials returns a boolean if a field has been set.

### SetServerCredentialsNil

`func (o *PhysicalMountVolumesNewTargetConfig) SetServerCredentialsNil(b bool)`

 SetServerCredentialsNil sets the value for ServerCredentials to be an explicit nil

### UnsetServerCredentials
`func (o *PhysicalMountVolumesNewTargetConfig) UnsetServerCredentials()`

UnsetServerCredentials ensures that no value is present for ServerCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


