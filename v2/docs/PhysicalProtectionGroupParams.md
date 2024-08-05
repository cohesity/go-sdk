# PhysicalProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileProtectionTypeParams** | Pointer to [**NullablePhysicalFileProtectionGroupParams**](PhysicalFileProtectionGroupParams.md) |  | [optional] 
**ProtectionType** | **string** | Specifies the Physical Protection Group type. | 
**VolumeProtectionTypeParams** | Pointer to [**NullablePhysicalVolumeProtectionGroupParams**](PhysicalVolumeProtectionGroupParams.md) |  | [optional] 

## Methods

### NewPhysicalProtectionGroupParams

`func NewPhysicalProtectionGroupParams(protectionType string, ) *PhysicalProtectionGroupParams`

NewPhysicalProtectionGroupParams instantiates a new PhysicalProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalProtectionGroupParamsWithDefaults

`func NewPhysicalProtectionGroupParamsWithDefaults() *PhysicalProtectionGroupParams`

NewPhysicalProtectionGroupParamsWithDefaults instantiates a new PhysicalProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileProtectionTypeParams

`func (o *PhysicalProtectionGroupParams) GetFileProtectionTypeParams() PhysicalFileProtectionGroupParams`

GetFileProtectionTypeParams returns the FileProtectionTypeParams field if non-nil, zero value otherwise.

### GetFileProtectionTypeParamsOk

`func (o *PhysicalProtectionGroupParams) GetFileProtectionTypeParamsOk() (*PhysicalFileProtectionGroupParams, bool)`

GetFileProtectionTypeParamsOk returns a tuple with the FileProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileProtectionTypeParams

`func (o *PhysicalProtectionGroupParams) SetFileProtectionTypeParams(v PhysicalFileProtectionGroupParams)`

SetFileProtectionTypeParams sets FileProtectionTypeParams field to given value.

### HasFileProtectionTypeParams

`func (o *PhysicalProtectionGroupParams) HasFileProtectionTypeParams() bool`

HasFileProtectionTypeParams returns a boolean if a field has been set.

### SetFileProtectionTypeParamsNil

`func (o *PhysicalProtectionGroupParams) SetFileProtectionTypeParamsNil(b bool)`

 SetFileProtectionTypeParamsNil sets the value for FileProtectionTypeParams to be an explicit nil

### UnsetFileProtectionTypeParams
`func (o *PhysicalProtectionGroupParams) UnsetFileProtectionTypeParams()`

UnsetFileProtectionTypeParams ensures that no value is present for FileProtectionTypeParams, not even an explicit nil
### GetProtectionType

`func (o *PhysicalProtectionGroupParams) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *PhysicalProtectionGroupParams) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *PhysicalProtectionGroupParams) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.


### GetVolumeProtectionTypeParams

`func (o *PhysicalProtectionGroupParams) GetVolumeProtectionTypeParams() PhysicalVolumeProtectionGroupParams`

GetVolumeProtectionTypeParams returns the VolumeProtectionTypeParams field if non-nil, zero value otherwise.

### GetVolumeProtectionTypeParamsOk

`func (o *PhysicalProtectionGroupParams) GetVolumeProtectionTypeParamsOk() (*PhysicalVolumeProtectionGroupParams, bool)`

GetVolumeProtectionTypeParamsOk returns a tuple with the VolumeProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeProtectionTypeParams

`func (o *PhysicalProtectionGroupParams) SetVolumeProtectionTypeParams(v PhysicalVolumeProtectionGroupParams)`

SetVolumeProtectionTypeParams sets VolumeProtectionTypeParams field to given value.

### HasVolumeProtectionTypeParams

`func (o *PhysicalProtectionGroupParams) HasVolumeProtectionTypeParams() bool`

HasVolumeProtectionTypeParams returns a boolean if a field has been set.

### SetVolumeProtectionTypeParamsNil

`func (o *PhysicalProtectionGroupParams) SetVolumeProtectionTypeParamsNil(b bool)`

 SetVolumeProtectionTypeParamsNil sets the value for VolumeProtectionTypeParams to be an explicit nil

### UnsetVolumeProtectionTypeParams
`func (o *PhysicalProtectionGroupParams) UnsetVolumeProtectionTypeParams()`

UnsetVolumeProtectionTypeParams ensures that no value is present for VolumeProtectionTypeParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


