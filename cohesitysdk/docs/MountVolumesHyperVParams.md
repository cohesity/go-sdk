# MountVolumesHyperVParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BringDisksOnline** | Pointer to **NullableBool** | Optional setting which will determine if the volumes need to be onlined within the target environment after attaching the disks. NOTE: If this is set to true, then target_entity_credentials must be specified and HyperV Integration Services must be installed on the VM. | [optional] 
**TargetEntityCredentials** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 

## Methods

### NewMountVolumesHyperVParams

`func NewMountVolumesHyperVParams() *MountVolumesHyperVParams`

NewMountVolumesHyperVParams instantiates a new MountVolumesHyperVParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountVolumesHyperVParamsWithDefaults

`func NewMountVolumesHyperVParamsWithDefaults() *MountVolumesHyperVParams`

NewMountVolumesHyperVParamsWithDefaults instantiates a new MountVolumesHyperVParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBringDisksOnline

`func (o *MountVolumesHyperVParams) GetBringDisksOnline() bool`

GetBringDisksOnline returns the BringDisksOnline field if non-nil, zero value otherwise.

### GetBringDisksOnlineOk

`func (o *MountVolumesHyperVParams) GetBringDisksOnlineOk() (*bool, bool)`

GetBringDisksOnlineOk returns a tuple with the BringDisksOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBringDisksOnline

`func (o *MountVolumesHyperVParams) SetBringDisksOnline(v bool)`

SetBringDisksOnline sets BringDisksOnline field to given value.

### HasBringDisksOnline

`func (o *MountVolumesHyperVParams) HasBringDisksOnline() bool`

HasBringDisksOnline returns a boolean if a field has been set.

### SetBringDisksOnlineNil

`func (o *MountVolumesHyperVParams) SetBringDisksOnlineNil(b bool)`

 SetBringDisksOnlineNil sets the value for BringDisksOnline to be an explicit nil

### UnsetBringDisksOnline
`func (o *MountVolumesHyperVParams) UnsetBringDisksOnline()`

UnsetBringDisksOnline ensures that no value is present for BringDisksOnline, not even an explicit nil
### GetTargetEntityCredentials

`func (o *MountVolumesHyperVParams) GetTargetEntityCredentials() Credentials`

GetTargetEntityCredentials returns the TargetEntityCredentials field if non-nil, zero value otherwise.

### GetTargetEntityCredentialsOk

`func (o *MountVolumesHyperVParams) GetTargetEntityCredentialsOk() (*Credentials, bool)`

GetTargetEntityCredentialsOk returns a tuple with the TargetEntityCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntityCredentials

`func (o *MountVolumesHyperVParams) SetTargetEntityCredentials(v Credentials)`

SetTargetEntityCredentials sets TargetEntityCredentials field to given value.

### HasTargetEntityCredentials

`func (o *MountVolumesHyperVParams) HasTargetEntityCredentials() bool`

HasTargetEntityCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


