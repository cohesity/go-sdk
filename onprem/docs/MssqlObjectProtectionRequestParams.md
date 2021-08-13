# MssqlObjectProtectionRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectProtectionType** | **string** | Specifies the MSSQL Object Protection type. | 
**FileObjectProtectionTypeParams** | Pointer to [**MssqlFileObjectProtectionParams**](MssqlFileObjectProtectionParams.md) |  | [optional] 
**NativeObjectProtectionTypeParams** | Pointer to [**MssqlNativeObjectProtectionParams**](MssqlNativeObjectProtectionParams.md) |  | [optional] 

## Methods

### NewMssqlObjectProtectionRequestParams

`func NewMssqlObjectProtectionRequestParams(objectProtectionType string, ) *MssqlObjectProtectionRequestParams`

NewMssqlObjectProtectionRequestParams instantiates a new MssqlObjectProtectionRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMssqlObjectProtectionRequestParamsWithDefaults

`func NewMssqlObjectProtectionRequestParamsWithDefaults() *MssqlObjectProtectionRequestParams`

NewMssqlObjectProtectionRequestParamsWithDefaults instantiates a new MssqlObjectProtectionRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectProtectionType

`func (o *MssqlObjectProtectionRequestParams) GetObjectProtectionType() string`

GetObjectProtectionType returns the ObjectProtectionType field if non-nil, zero value otherwise.

### GetObjectProtectionTypeOk

`func (o *MssqlObjectProtectionRequestParams) GetObjectProtectionTypeOk() (*string, bool)`

GetObjectProtectionTypeOk returns a tuple with the ObjectProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectProtectionType

`func (o *MssqlObjectProtectionRequestParams) SetObjectProtectionType(v string)`

SetObjectProtectionType sets ObjectProtectionType field to given value.


### GetFileObjectProtectionTypeParams

`func (o *MssqlObjectProtectionRequestParams) GetFileObjectProtectionTypeParams() MssqlFileObjectProtectionParams`

GetFileObjectProtectionTypeParams returns the FileObjectProtectionTypeParams field if non-nil, zero value otherwise.

### GetFileObjectProtectionTypeParamsOk

`func (o *MssqlObjectProtectionRequestParams) GetFileObjectProtectionTypeParamsOk() (*MssqlFileObjectProtectionParams, bool)`

GetFileObjectProtectionTypeParamsOk returns a tuple with the FileObjectProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileObjectProtectionTypeParams

`func (o *MssqlObjectProtectionRequestParams) SetFileObjectProtectionTypeParams(v MssqlFileObjectProtectionParams)`

SetFileObjectProtectionTypeParams sets FileObjectProtectionTypeParams field to given value.

### HasFileObjectProtectionTypeParams

`func (o *MssqlObjectProtectionRequestParams) HasFileObjectProtectionTypeParams() bool`

HasFileObjectProtectionTypeParams returns a boolean if a field has been set.

### GetNativeObjectProtectionTypeParams

`func (o *MssqlObjectProtectionRequestParams) GetNativeObjectProtectionTypeParams() MssqlNativeObjectProtectionParams`

GetNativeObjectProtectionTypeParams returns the NativeObjectProtectionTypeParams field if non-nil, zero value otherwise.

### GetNativeObjectProtectionTypeParamsOk

`func (o *MssqlObjectProtectionRequestParams) GetNativeObjectProtectionTypeParamsOk() (*MssqlNativeObjectProtectionParams, bool)`

GetNativeObjectProtectionTypeParamsOk returns a tuple with the NativeObjectProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeObjectProtectionTypeParams

`func (o *MssqlObjectProtectionRequestParams) SetNativeObjectProtectionTypeParams(v MssqlNativeObjectProtectionParams)`

SetNativeObjectProtectionTypeParams sets NativeObjectProtectionTypeParams field to given value.

### HasNativeObjectProtectionTypeParams

`func (o *MssqlObjectProtectionRequestParams) HasNativeObjectProtectionTypeParams() bool`

HasNativeObjectProtectionTypeParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


