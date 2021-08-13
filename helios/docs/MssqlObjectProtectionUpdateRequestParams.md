# MssqlObjectProtectionUpdateRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectProtectionType** | **string** | Specifies the MSSQL Object Protection type. | 
**CommonFileObjectProtectionTypeParams** | Pointer to [**CommonMssqlFileObjectProtectionParams**](CommonMssqlFileObjectProtectionParams.md) |  | [optional] 
**CommonNativeObjectProtectionTypeParams** | Pointer to [**CommonMssqlNativeObjectProtectionParams**](CommonMssqlNativeObjectProtectionParams.md) |  | [optional] 

## Methods

### NewMssqlObjectProtectionUpdateRequestParams

`func NewMssqlObjectProtectionUpdateRequestParams(objectProtectionType string, ) *MssqlObjectProtectionUpdateRequestParams`

NewMssqlObjectProtectionUpdateRequestParams instantiates a new MssqlObjectProtectionUpdateRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMssqlObjectProtectionUpdateRequestParamsWithDefaults

`func NewMssqlObjectProtectionUpdateRequestParamsWithDefaults() *MssqlObjectProtectionUpdateRequestParams`

NewMssqlObjectProtectionUpdateRequestParamsWithDefaults instantiates a new MssqlObjectProtectionUpdateRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectProtectionType

`func (o *MssqlObjectProtectionUpdateRequestParams) GetObjectProtectionType() string`

GetObjectProtectionType returns the ObjectProtectionType field if non-nil, zero value otherwise.

### GetObjectProtectionTypeOk

`func (o *MssqlObjectProtectionUpdateRequestParams) GetObjectProtectionTypeOk() (*string, bool)`

GetObjectProtectionTypeOk returns a tuple with the ObjectProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectProtectionType

`func (o *MssqlObjectProtectionUpdateRequestParams) SetObjectProtectionType(v string)`

SetObjectProtectionType sets ObjectProtectionType field to given value.


### GetCommonFileObjectProtectionTypeParams

`func (o *MssqlObjectProtectionUpdateRequestParams) GetCommonFileObjectProtectionTypeParams() CommonMssqlFileObjectProtectionParams`

GetCommonFileObjectProtectionTypeParams returns the CommonFileObjectProtectionTypeParams field if non-nil, zero value otherwise.

### GetCommonFileObjectProtectionTypeParamsOk

`func (o *MssqlObjectProtectionUpdateRequestParams) GetCommonFileObjectProtectionTypeParamsOk() (*CommonMssqlFileObjectProtectionParams, bool)`

GetCommonFileObjectProtectionTypeParamsOk returns a tuple with the CommonFileObjectProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonFileObjectProtectionTypeParams

`func (o *MssqlObjectProtectionUpdateRequestParams) SetCommonFileObjectProtectionTypeParams(v CommonMssqlFileObjectProtectionParams)`

SetCommonFileObjectProtectionTypeParams sets CommonFileObjectProtectionTypeParams field to given value.

### HasCommonFileObjectProtectionTypeParams

`func (o *MssqlObjectProtectionUpdateRequestParams) HasCommonFileObjectProtectionTypeParams() bool`

HasCommonFileObjectProtectionTypeParams returns a boolean if a field has been set.

### GetCommonNativeObjectProtectionTypeParams

`func (o *MssqlObjectProtectionUpdateRequestParams) GetCommonNativeObjectProtectionTypeParams() CommonMssqlNativeObjectProtectionParams`

GetCommonNativeObjectProtectionTypeParams returns the CommonNativeObjectProtectionTypeParams field if non-nil, zero value otherwise.

### GetCommonNativeObjectProtectionTypeParamsOk

`func (o *MssqlObjectProtectionUpdateRequestParams) GetCommonNativeObjectProtectionTypeParamsOk() (*CommonMssqlNativeObjectProtectionParams, bool)`

GetCommonNativeObjectProtectionTypeParamsOk returns a tuple with the CommonNativeObjectProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonNativeObjectProtectionTypeParams

`func (o *MssqlObjectProtectionUpdateRequestParams) SetCommonNativeObjectProtectionTypeParams(v CommonMssqlNativeObjectProtectionParams)`

SetCommonNativeObjectProtectionTypeParams sets CommonNativeObjectProtectionTypeParams field to given value.

### HasCommonNativeObjectProtectionTypeParams

`func (o *MssqlObjectProtectionUpdateRequestParams) HasCommonNativeObjectProtectionTypeParams() bool`

HasCommonNativeObjectProtectionTypeParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


