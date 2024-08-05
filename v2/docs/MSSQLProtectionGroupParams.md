# MSSQLProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileProtectionTypeParams** | Pointer to [**MSSQLFileProtectionGroupParams**](MSSQLFileProtectionGroupParams.md) |  | [optional] 
**NativeProtectionTypeParams** | Pointer to [**MSSQLNativeProtectionGroupParams**](MSSQLNativeProtectionGroupParams.md) |  | [optional] 
**ProtectionType** | **string** | Specifies the MSSQL Protection Group type. | 
**VolumeProtectionTypeParams** | Pointer to [**MSSQLVolumeProtectionGroupParams**](MSSQLVolumeProtectionGroupParams.md) |  | [optional] 

## Methods

### NewMSSQLProtectionGroupParams

`func NewMSSQLProtectionGroupParams(protectionType string, ) *MSSQLProtectionGroupParams`

NewMSSQLProtectionGroupParams instantiates a new MSSQLProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMSSQLProtectionGroupParamsWithDefaults

`func NewMSSQLProtectionGroupParamsWithDefaults() *MSSQLProtectionGroupParams`

NewMSSQLProtectionGroupParamsWithDefaults instantiates a new MSSQLProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileProtectionTypeParams

`func (o *MSSQLProtectionGroupParams) GetFileProtectionTypeParams() MSSQLFileProtectionGroupParams`

GetFileProtectionTypeParams returns the FileProtectionTypeParams field if non-nil, zero value otherwise.

### GetFileProtectionTypeParamsOk

`func (o *MSSQLProtectionGroupParams) GetFileProtectionTypeParamsOk() (*MSSQLFileProtectionGroupParams, bool)`

GetFileProtectionTypeParamsOk returns a tuple with the FileProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileProtectionTypeParams

`func (o *MSSQLProtectionGroupParams) SetFileProtectionTypeParams(v MSSQLFileProtectionGroupParams)`

SetFileProtectionTypeParams sets FileProtectionTypeParams field to given value.

### HasFileProtectionTypeParams

`func (o *MSSQLProtectionGroupParams) HasFileProtectionTypeParams() bool`

HasFileProtectionTypeParams returns a boolean if a field has been set.

### GetNativeProtectionTypeParams

`func (o *MSSQLProtectionGroupParams) GetNativeProtectionTypeParams() MSSQLNativeProtectionGroupParams`

GetNativeProtectionTypeParams returns the NativeProtectionTypeParams field if non-nil, zero value otherwise.

### GetNativeProtectionTypeParamsOk

`func (o *MSSQLProtectionGroupParams) GetNativeProtectionTypeParamsOk() (*MSSQLNativeProtectionGroupParams, bool)`

GetNativeProtectionTypeParamsOk returns a tuple with the NativeProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeProtectionTypeParams

`func (o *MSSQLProtectionGroupParams) SetNativeProtectionTypeParams(v MSSQLNativeProtectionGroupParams)`

SetNativeProtectionTypeParams sets NativeProtectionTypeParams field to given value.

### HasNativeProtectionTypeParams

`func (o *MSSQLProtectionGroupParams) HasNativeProtectionTypeParams() bool`

HasNativeProtectionTypeParams returns a boolean if a field has been set.

### GetProtectionType

`func (o *MSSQLProtectionGroupParams) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *MSSQLProtectionGroupParams) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *MSSQLProtectionGroupParams) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.


### GetVolumeProtectionTypeParams

`func (o *MSSQLProtectionGroupParams) GetVolumeProtectionTypeParams() MSSQLVolumeProtectionGroupParams`

GetVolumeProtectionTypeParams returns the VolumeProtectionTypeParams field if non-nil, zero value otherwise.

### GetVolumeProtectionTypeParamsOk

`func (o *MSSQLProtectionGroupParams) GetVolumeProtectionTypeParamsOk() (*MSSQLVolumeProtectionGroupParams, bool)`

GetVolumeProtectionTypeParamsOk returns a tuple with the VolumeProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeProtectionTypeParams

`func (o *MSSQLProtectionGroupParams) SetVolumeProtectionTypeParams(v MSSQLVolumeProtectionGroupParams)`

SetVolumeProtectionTypeParams sets VolumeProtectionTypeParams field to given value.

### HasVolumeProtectionTypeParams

`func (o *MSSQLProtectionGroupParams) HasVolumeProtectionTypeParams() bool`

HasVolumeProtectionTypeParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


