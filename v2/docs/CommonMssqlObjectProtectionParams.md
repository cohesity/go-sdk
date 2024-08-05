# CommonMssqlObjectProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonFileObjectProtectionTypeParams** | Pointer to [**CommonMssqlFileObjectProtectionParams**](CommonMssqlFileObjectProtectionParams.md) |  | [optional] 
**CommonNativeObjectProtectionTypeParams** | Pointer to [**CommonMssqlNativeObjectProtectionParams**](CommonMssqlNativeObjectProtectionParams.md) |  | [optional] 
**ObjectProtectionType** | **string** | Specifies the MSSQL Object Protection type. | 

## Methods

### NewCommonMssqlObjectProtectionParams

`func NewCommonMssqlObjectProtectionParams(objectProtectionType string, ) *CommonMssqlObjectProtectionParams`

NewCommonMssqlObjectProtectionParams instantiates a new CommonMssqlObjectProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonMssqlObjectProtectionParamsWithDefaults

`func NewCommonMssqlObjectProtectionParamsWithDefaults() *CommonMssqlObjectProtectionParams`

NewCommonMssqlObjectProtectionParamsWithDefaults instantiates a new CommonMssqlObjectProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonFileObjectProtectionTypeParams

`func (o *CommonMssqlObjectProtectionParams) GetCommonFileObjectProtectionTypeParams() CommonMssqlFileObjectProtectionParams`

GetCommonFileObjectProtectionTypeParams returns the CommonFileObjectProtectionTypeParams field if non-nil, zero value otherwise.

### GetCommonFileObjectProtectionTypeParamsOk

`func (o *CommonMssqlObjectProtectionParams) GetCommonFileObjectProtectionTypeParamsOk() (*CommonMssqlFileObjectProtectionParams, bool)`

GetCommonFileObjectProtectionTypeParamsOk returns a tuple with the CommonFileObjectProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonFileObjectProtectionTypeParams

`func (o *CommonMssqlObjectProtectionParams) SetCommonFileObjectProtectionTypeParams(v CommonMssqlFileObjectProtectionParams)`

SetCommonFileObjectProtectionTypeParams sets CommonFileObjectProtectionTypeParams field to given value.

### HasCommonFileObjectProtectionTypeParams

`func (o *CommonMssqlObjectProtectionParams) HasCommonFileObjectProtectionTypeParams() bool`

HasCommonFileObjectProtectionTypeParams returns a boolean if a field has been set.

### GetCommonNativeObjectProtectionTypeParams

`func (o *CommonMssqlObjectProtectionParams) GetCommonNativeObjectProtectionTypeParams() CommonMssqlNativeObjectProtectionParams`

GetCommonNativeObjectProtectionTypeParams returns the CommonNativeObjectProtectionTypeParams field if non-nil, zero value otherwise.

### GetCommonNativeObjectProtectionTypeParamsOk

`func (o *CommonMssqlObjectProtectionParams) GetCommonNativeObjectProtectionTypeParamsOk() (*CommonMssqlNativeObjectProtectionParams, bool)`

GetCommonNativeObjectProtectionTypeParamsOk returns a tuple with the CommonNativeObjectProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonNativeObjectProtectionTypeParams

`func (o *CommonMssqlObjectProtectionParams) SetCommonNativeObjectProtectionTypeParams(v CommonMssqlNativeObjectProtectionParams)`

SetCommonNativeObjectProtectionTypeParams sets CommonNativeObjectProtectionTypeParams field to given value.

### HasCommonNativeObjectProtectionTypeParams

`func (o *CommonMssqlObjectProtectionParams) HasCommonNativeObjectProtectionTypeParams() bool`

HasCommonNativeObjectProtectionTypeParams returns a boolean if a field has been set.

### GetObjectProtectionType

`func (o *CommonMssqlObjectProtectionParams) GetObjectProtectionType() string`

GetObjectProtectionType returns the ObjectProtectionType field if non-nil, zero value otherwise.

### GetObjectProtectionTypeOk

`func (o *CommonMssqlObjectProtectionParams) GetObjectProtectionTypeOk() (*string, bool)`

GetObjectProtectionTypeOk returns a tuple with the ObjectProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectProtectionType

`func (o *CommonMssqlObjectProtectionParams) SetObjectProtectionType(v string)`

SetObjectProtectionType sets ObjectProtectionType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


