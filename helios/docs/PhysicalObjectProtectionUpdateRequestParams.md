# PhysicalObjectProtectionUpdateRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectProtectionType** | **NullableString** | Specifies the Physical Object Protection type. | 
**FileObjectProtectionTypeParams** | Pointer to [**NullablePhysicalFileProtectionGroupParams**](PhysicalFileProtectionGroupParams.md) |  | [optional] 
**VolumeObjectProtectionTypeParams** | Pointer to [**NullablePhysicalVolumeProtectionGroupParams**](PhysicalVolumeProtectionGroupParams.md) |  | [optional] 

## Methods

### NewPhysicalObjectProtectionUpdateRequestParams

`func NewPhysicalObjectProtectionUpdateRequestParams(objectProtectionType NullableString, ) *PhysicalObjectProtectionUpdateRequestParams`

NewPhysicalObjectProtectionUpdateRequestParams instantiates a new PhysicalObjectProtectionUpdateRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalObjectProtectionUpdateRequestParamsWithDefaults

`func NewPhysicalObjectProtectionUpdateRequestParamsWithDefaults() *PhysicalObjectProtectionUpdateRequestParams`

NewPhysicalObjectProtectionUpdateRequestParamsWithDefaults instantiates a new PhysicalObjectProtectionUpdateRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectProtectionType

`func (o *PhysicalObjectProtectionUpdateRequestParams) GetObjectProtectionType() string`

GetObjectProtectionType returns the ObjectProtectionType field if non-nil, zero value otherwise.

### GetObjectProtectionTypeOk

`func (o *PhysicalObjectProtectionUpdateRequestParams) GetObjectProtectionTypeOk() (*string, bool)`

GetObjectProtectionTypeOk returns a tuple with the ObjectProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectProtectionType

`func (o *PhysicalObjectProtectionUpdateRequestParams) SetObjectProtectionType(v string)`

SetObjectProtectionType sets ObjectProtectionType field to given value.


### SetObjectProtectionTypeNil

`func (o *PhysicalObjectProtectionUpdateRequestParams) SetObjectProtectionTypeNil(b bool)`

 SetObjectProtectionTypeNil sets the value for ObjectProtectionType to be an explicit nil

### UnsetObjectProtectionType
`func (o *PhysicalObjectProtectionUpdateRequestParams) UnsetObjectProtectionType()`

UnsetObjectProtectionType ensures that no value is present for ObjectProtectionType, not even an explicit nil
### GetFileObjectProtectionTypeParams

`func (o *PhysicalObjectProtectionUpdateRequestParams) GetFileObjectProtectionTypeParams() PhysicalFileProtectionGroupParams`

GetFileObjectProtectionTypeParams returns the FileObjectProtectionTypeParams field if non-nil, zero value otherwise.

### GetFileObjectProtectionTypeParamsOk

`func (o *PhysicalObjectProtectionUpdateRequestParams) GetFileObjectProtectionTypeParamsOk() (*PhysicalFileProtectionGroupParams, bool)`

GetFileObjectProtectionTypeParamsOk returns a tuple with the FileObjectProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileObjectProtectionTypeParams

`func (o *PhysicalObjectProtectionUpdateRequestParams) SetFileObjectProtectionTypeParams(v PhysicalFileProtectionGroupParams)`

SetFileObjectProtectionTypeParams sets FileObjectProtectionTypeParams field to given value.

### HasFileObjectProtectionTypeParams

`func (o *PhysicalObjectProtectionUpdateRequestParams) HasFileObjectProtectionTypeParams() bool`

HasFileObjectProtectionTypeParams returns a boolean if a field has been set.

### SetFileObjectProtectionTypeParamsNil

`func (o *PhysicalObjectProtectionUpdateRequestParams) SetFileObjectProtectionTypeParamsNil(b bool)`

 SetFileObjectProtectionTypeParamsNil sets the value for FileObjectProtectionTypeParams to be an explicit nil

### UnsetFileObjectProtectionTypeParams
`func (o *PhysicalObjectProtectionUpdateRequestParams) UnsetFileObjectProtectionTypeParams()`

UnsetFileObjectProtectionTypeParams ensures that no value is present for FileObjectProtectionTypeParams, not even an explicit nil
### GetVolumeObjectProtectionTypeParams

`func (o *PhysicalObjectProtectionUpdateRequestParams) GetVolumeObjectProtectionTypeParams() PhysicalVolumeProtectionGroupParams`

GetVolumeObjectProtectionTypeParams returns the VolumeObjectProtectionTypeParams field if non-nil, zero value otherwise.

### GetVolumeObjectProtectionTypeParamsOk

`func (o *PhysicalObjectProtectionUpdateRequestParams) GetVolumeObjectProtectionTypeParamsOk() (*PhysicalVolumeProtectionGroupParams, bool)`

GetVolumeObjectProtectionTypeParamsOk returns a tuple with the VolumeObjectProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeObjectProtectionTypeParams

`func (o *PhysicalObjectProtectionUpdateRequestParams) SetVolumeObjectProtectionTypeParams(v PhysicalVolumeProtectionGroupParams)`

SetVolumeObjectProtectionTypeParams sets VolumeObjectProtectionTypeParams field to given value.

### HasVolumeObjectProtectionTypeParams

`func (o *PhysicalObjectProtectionUpdateRequestParams) HasVolumeObjectProtectionTypeParams() bool`

HasVolumeObjectProtectionTypeParams returns a boolean if a field has been set.

### SetVolumeObjectProtectionTypeParamsNil

`func (o *PhysicalObjectProtectionUpdateRequestParams) SetVolumeObjectProtectionTypeParamsNil(b bool)`

 SetVolumeObjectProtectionTypeParamsNil sets the value for VolumeObjectProtectionTypeParams to be an explicit nil

### UnsetVolumeObjectProtectionTypeParams
`func (o *PhysicalObjectProtectionUpdateRequestParams) UnsetVolumeObjectProtectionTypeParams()`

UnsetVolumeObjectProtectionTypeParams ensures that no value is present for VolumeObjectProtectionTypeParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


