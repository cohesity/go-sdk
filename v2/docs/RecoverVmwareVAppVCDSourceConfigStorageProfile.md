# RecoverVmwareVAppVCDSourceConfigStorageProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the name of the storage profile. | [optional] [readonly] 
**VcdUuid** | **NullableString** | Specifies the UUID assigned by the VCD to the storage profile. | 

## Methods

### NewRecoverVmwareVAppVCDSourceConfigStorageProfile

`func NewRecoverVmwareVAppVCDSourceConfigStorageProfile(vcdUuid NullableString, ) *RecoverVmwareVAppVCDSourceConfigStorageProfile`

NewRecoverVmwareVAppVCDSourceConfigStorageProfile instantiates a new RecoverVmwareVAppVCDSourceConfigStorageProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareVAppVCDSourceConfigStorageProfileWithDefaults

`func NewRecoverVmwareVAppVCDSourceConfigStorageProfileWithDefaults() *RecoverVmwareVAppVCDSourceConfigStorageProfile`

NewRecoverVmwareVAppVCDSourceConfigStorageProfileWithDefaults instantiates a new RecoverVmwareVAppVCDSourceConfigStorageProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RecoverVmwareVAppVCDSourceConfigStorageProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecoverVmwareVAppVCDSourceConfigStorageProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecoverVmwareVAppVCDSourceConfigStorageProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecoverVmwareVAppVCDSourceConfigStorageProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RecoverVmwareVAppVCDSourceConfigStorageProfile) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RecoverVmwareVAppVCDSourceConfigStorageProfile) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVcdUuid

`func (o *RecoverVmwareVAppVCDSourceConfigStorageProfile) GetVcdUuid() string`

GetVcdUuid returns the VcdUuid field if non-nil, zero value otherwise.

### GetVcdUuidOk

`func (o *RecoverVmwareVAppVCDSourceConfigStorageProfile) GetVcdUuidOk() (*string, bool)`

GetVcdUuidOk returns a tuple with the VcdUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdUuid

`func (o *RecoverVmwareVAppVCDSourceConfigStorageProfile) SetVcdUuid(v string)`

SetVcdUuid sets VcdUuid field to given value.


### SetVcdUuidNil

`func (o *RecoverVmwareVAppVCDSourceConfigStorageProfile) SetVcdUuidNil(b bool)`

 SetVcdUuidNil sets the value for VcdUuid to be an explicit nil

### UnsetVcdUuid
`func (o *RecoverVmwareVAppVCDSourceConfigStorageProfile) UnsetVcdUuid()`

UnsetVcdUuid ensures that no value is present for VcdUuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


