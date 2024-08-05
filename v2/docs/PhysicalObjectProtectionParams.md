# PhysicalObjectProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileObjectProtectionTypeParams** | Pointer to [**NullablePhysicalFileProtectionGroupParams**](PhysicalFileProtectionGroupParams.md) |  | [optional] 
**ObjectProtectionType** | **NullableString** | Specifies the Physical Object Protection type. | 
**VolumeObjectProtectionTypeParams** | Pointer to [**NullablePhysicalVolumeProtectionGroupParams**](PhysicalVolumeProtectionGroupParams.md) |  | [optional] 

## Methods

### NewPhysicalObjectProtectionParams

`func NewPhysicalObjectProtectionParams(objectProtectionType NullableString, ) *PhysicalObjectProtectionParams`

NewPhysicalObjectProtectionParams instantiates a new PhysicalObjectProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalObjectProtectionParamsWithDefaults

`func NewPhysicalObjectProtectionParamsWithDefaults() *PhysicalObjectProtectionParams`

NewPhysicalObjectProtectionParamsWithDefaults instantiates a new PhysicalObjectProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileObjectProtectionTypeParams

`func (o *PhysicalObjectProtectionParams) GetFileObjectProtectionTypeParams() PhysicalFileProtectionGroupParams`

GetFileObjectProtectionTypeParams returns the FileObjectProtectionTypeParams field if non-nil, zero value otherwise.

### GetFileObjectProtectionTypeParamsOk

`func (o *PhysicalObjectProtectionParams) GetFileObjectProtectionTypeParamsOk() (*PhysicalFileProtectionGroupParams, bool)`

GetFileObjectProtectionTypeParamsOk returns a tuple with the FileObjectProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileObjectProtectionTypeParams

`func (o *PhysicalObjectProtectionParams) SetFileObjectProtectionTypeParams(v PhysicalFileProtectionGroupParams)`

SetFileObjectProtectionTypeParams sets FileObjectProtectionTypeParams field to given value.

### HasFileObjectProtectionTypeParams

`func (o *PhysicalObjectProtectionParams) HasFileObjectProtectionTypeParams() bool`

HasFileObjectProtectionTypeParams returns a boolean if a field has been set.

### SetFileObjectProtectionTypeParamsNil

`func (o *PhysicalObjectProtectionParams) SetFileObjectProtectionTypeParamsNil(b bool)`

 SetFileObjectProtectionTypeParamsNil sets the value for FileObjectProtectionTypeParams to be an explicit nil

### UnsetFileObjectProtectionTypeParams
`func (o *PhysicalObjectProtectionParams) UnsetFileObjectProtectionTypeParams()`

UnsetFileObjectProtectionTypeParams ensures that no value is present for FileObjectProtectionTypeParams, not even an explicit nil
### GetObjectProtectionType

`func (o *PhysicalObjectProtectionParams) GetObjectProtectionType() string`

GetObjectProtectionType returns the ObjectProtectionType field if non-nil, zero value otherwise.

### GetObjectProtectionTypeOk

`func (o *PhysicalObjectProtectionParams) GetObjectProtectionTypeOk() (*string, bool)`

GetObjectProtectionTypeOk returns a tuple with the ObjectProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectProtectionType

`func (o *PhysicalObjectProtectionParams) SetObjectProtectionType(v string)`

SetObjectProtectionType sets ObjectProtectionType field to given value.


### SetObjectProtectionTypeNil

`func (o *PhysicalObjectProtectionParams) SetObjectProtectionTypeNil(b bool)`

 SetObjectProtectionTypeNil sets the value for ObjectProtectionType to be an explicit nil

### UnsetObjectProtectionType
`func (o *PhysicalObjectProtectionParams) UnsetObjectProtectionType()`

UnsetObjectProtectionType ensures that no value is present for ObjectProtectionType, not even an explicit nil
### GetVolumeObjectProtectionTypeParams

`func (o *PhysicalObjectProtectionParams) GetVolumeObjectProtectionTypeParams() PhysicalVolumeProtectionGroupParams`

GetVolumeObjectProtectionTypeParams returns the VolumeObjectProtectionTypeParams field if non-nil, zero value otherwise.

### GetVolumeObjectProtectionTypeParamsOk

`func (o *PhysicalObjectProtectionParams) GetVolumeObjectProtectionTypeParamsOk() (*PhysicalVolumeProtectionGroupParams, bool)`

GetVolumeObjectProtectionTypeParamsOk returns a tuple with the VolumeObjectProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeObjectProtectionTypeParams

`func (o *PhysicalObjectProtectionParams) SetVolumeObjectProtectionTypeParams(v PhysicalVolumeProtectionGroupParams)`

SetVolumeObjectProtectionTypeParams sets VolumeObjectProtectionTypeParams field to given value.

### HasVolumeObjectProtectionTypeParams

`func (o *PhysicalObjectProtectionParams) HasVolumeObjectProtectionTypeParams() bool`

HasVolumeObjectProtectionTypeParams returns a boolean if a field has been set.

### SetVolumeObjectProtectionTypeParamsNil

`func (o *PhysicalObjectProtectionParams) SetVolumeObjectProtectionTypeParamsNil(b bool)`

 SetVolumeObjectProtectionTypeParamsNil sets the value for VolumeObjectProtectionTypeParams to be an explicit nil

### UnsetVolumeObjectProtectionTypeParams
`func (o *PhysicalObjectProtectionParams) UnsetVolumeObjectProtectionTypeParams()`

UnsetVolumeObjectProtectionTypeParams ensures that no value is present for VolumeObjectProtectionTypeParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


