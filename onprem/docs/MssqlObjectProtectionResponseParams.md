# MssqlObjectProtectionResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectProtectionType** | **string** | Specifies the MSSQL Object Protection type. | 
**CommonFileObjectProtectionTypeParams** | Pointer to [**CommonMssqlFileObjectProtectionParams**](CommonMssqlFileObjectProtectionParams.md) |  | [optional] 
**CommonNativeObjectProtectionTypeParams** | Pointer to [**CommonMssqlNativeObjectProtectionParams**](CommonMssqlNativeObjectProtectionParams.md) |  | [optional] 

## Methods

### NewMssqlObjectProtectionResponseParams

`func NewMssqlObjectProtectionResponseParams(objectProtectionType string, ) *MssqlObjectProtectionResponseParams`

NewMssqlObjectProtectionResponseParams instantiates a new MssqlObjectProtectionResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMssqlObjectProtectionResponseParamsWithDefaults

`func NewMssqlObjectProtectionResponseParamsWithDefaults() *MssqlObjectProtectionResponseParams`

NewMssqlObjectProtectionResponseParamsWithDefaults instantiates a new MssqlObjectProtectionResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectProtectionType

`func (o *MssqlObjectProtectionResponseParams) GetObjectProtectionType() string`

GetObjectProtectionType returns the ObjectProtectionType field if non-nil, zero value otherwise.

### GetObjectProtectionTypeOk

`func (o *MssqlObjectProtectionResponseParams) GetObjectProtectionTypeOk() (*string, bool)`

GetObjectProtectionTypeOk returns a tuple with the ObjectProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectProtectionType

`func (o *MssqlObjectProtectionResponseParams) SetObjectProtectionType(v string)`

SetObjectProtectionType sets ObjectProtectionType field to given value.


### GetCommonFileObjectProtectionTypeParams

`func (o *MssqlObjectProtectionResponseParams) GetCommonFileObjectProtectionTypeParams() CommonMssqlFileObjectProtectionParams`

GetCommonFileObjectProtectionTypeParams returns the CommonFileObjectProtectionTypeParams field if non-nil, zero value otherwise.

### GetCommonFileObjectProtectionTypeParamsOk

`func (o *MssqlObjectProtectionResponseParams) GetCommonFileObjectProtectionTypeParamsOk() (*CommonMssqlFileObjectProtectionParams, bool)`

GetCommonFileObjectProtectionTypeParamsOk returns a tuple with the CommonFileObjectProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonFileObjectProtectionTypeParams

`func (o *MssqlObjectProtectionResponseParams) SetCommonFileObjectProtectionTypeParams(v CommonMssqlFileObjectProtectionParams)`

SetCommonFileObjectProtectionTypeParams sets CommonFileObjectProtectionTypeParams field to given value.

### HasCommonFileObjectProtectionTypeParams

`func (o *MssqlObjectProtectionResponseParams) HasCommonFileObjectProtectionTypeParams() bool`

HasCommonFileObjectProtectionTypeParams returns a boolean if a field has been set.

### GetCommonNativeObjectProtectionTypeParams

`func (o *MssqlObjectProtectionResponseParams) GetCommonNativeObjectProtectionTypeParams() CommonMssqlNativeObjectProtectionParams`

GetCommonNativeObjectProtectionTypeParams returns the CommonNativeObjectProtectionTypeParams field if non-nil, zero value otherwise.

### GetCommonNativeObjectProtectionTypeParamsOk

`func (o *MssqlObjectProtectionResponseParams) GetCommonNativeObjectProtectionTypeParamsOk() (*CommonMssqlNativeObjectProtectionParams, bool)`

GetCommonNativeObjectProtectionTypeParamsOk returns a tuple with the CommonNativeObjectProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonNativeObjectProtectionTypeParams

`func (o *MssqlObjectProtectionResponseParams) SetCommonNativeObjectProtectionTypeParams(v CommonMssqlNativeObjectProtectionParams)`

SetCommonNativeObjectProtectionTypeParams sets CommonNativeObjectProtectionTypeParams field to given value.

### HasCommonNativeObjectProtectionTypeParams

`func (o *MssqlObjectProtectionResponseParams) HasCommonNativeObjectProtectionTypeParams() bool`

HasCommonNativeObjectProtectionTypeParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


