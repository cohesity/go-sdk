# UserIdMappingParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CentrifyTypeParams** | Pointer to [**NullableUserIdMappingParamsCentrifyTypeParams**](UserIdMappingParamsCentrifyTypeParams.md) |  | [optional] 
**CustomAttributesTypeParams** | Pointer to [**NullableUserIdMappingParamsCustomAttributesTypeParams**](UserIdMappingParamsCustomAttributesTypeParams.md) |  | [optional] 
**FixedTypeParams** | Pointer to [**NullableUserIdMappingParamsFixedTypeParams**](UserIdMappingParamsFixedTypeParams.md) |  | [optional] 
**LdapProviderTypeParams** | Pointer to [**NullableUserIdMappingParamsLdapProviderTypeParams**](UserIdMappingParamsLdapProviderTypeParams.md) |  | [optional] 
**NisProviderTypeParams** | Pointer to [**NullableUserIdMappingParamsNisProviderTypeParams**](UserIdMappingParamsNisProviderTypeParams.md) |  | [optional] 
**Rfc2307TypeParams** | Pointer to [**NullableUserIdMappingParamsRfc2307TypeParams**](UserIdMappingParamsRfc2307TypeParams.md) |  | [optional] 
**Sfu30TypeParams** | Pointer to [**NullableUserIdMappingParamsSfu30TypeParams**](UserIdMappingParamsSfu30TypeParams.md) |  | [optional] 
**Type** | **string** | Specifies the type of the mapping. | 

## Methods

### NewUserIdMappingParams

`func NewUserIdMappingParams(type_ string, ) *UserIdMappingParams`

NewUserIdMappingParams instantiates a new UserIdMappingParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdMappingParamsWithDefaults

`func NewUserIdMappingParamsWithDefaults() *UserIdMappingParams`

NewUserIdMappingParamsWithDefaults instantiates a new UserIdMappingParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCentrifyTypeParams

`func (o *UserIdMappingParams) GetCentrifyTypeParams() UserIdMappingParamsCentrifyTypeParams`

GetCentrifyTypeParams returns the CentrifyTypeParams field if non-nil, zero value otherwise.

### GetCentrifyTypeParamsOk

`func (o *UserIdMappingParams) GetCentrifyTypeParamsOk() (*UserIdMappingParamsCentrifyTypeParams, bool)`

GetCentrifyTypeParamsOk returns a tuple with the CentrifyTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentrifyTypeParams

`func (o *UserIdMappingParams) SetCentrifyTypeParams(v UserIdMappingParamsCentrifyTypeParams)`

SetCentrifyTypeParams sets CentrifyTypeParams field to given value.

### HasCentrifyTypeParams

`func (o *UserIdMappingParams) HasCentrifyTypeParams() bool`

HasCentrifyTypeParams returns a boolean if a field has been set.

### SetCentrifyTypeParamsNil

`func (o *UserIdMappingParams) SetCentrifyTypeParamsNil(b bool)`

 SetCentrifyTypeParamsNil sets the value for CentrifyTypeParams to be an explicit nil

### UnsetCentrifyTypeParams
`func (o *UserIdMappingParams) UnsetCentrifyTypeParams()`

UnsetCentrifyTypeParams ensures that no value is present for CentrifyTypeParams, not even an explicit nil
### GetCustomAttributesTypeParams

`func (o *UserIdMappingParams) GetCustomAttributesTypeParams() UserIdMappingParamsCustomAttributesTypeParams`

GetCustomAttributesTypeParams returns the CustomAttributesTypeParams field if non-nil, zero value otherwise.

### GetCustomAttributesTypeParamsOk

`func (o *UserIdMappingParams) GetCustomAttributesTypeParamsOk() (*UserIdMappingParamsCustomAttributesTypeParams, bool)`

GetCustomAttributesTypeParamsOk returns a tuple with the CustomAttributesTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributesTypeParams

`func (o *UserIdMappingParams) SetCustomAttributesTypeParams(v UserIdMappingParamsCustomAttributesTypeParams)`

SetCustomAttributesTypeParams sets CustomAttributesTypeParams field to given value.

### HasCustomAttributesTypeParams

`func (o *UserIdMappingParams) HasCustomAttributesTypeParams() bool`

HasCustomAttributesTypeParams returns a boolean if a field has been set.

### SetCustomAttributesTypeParamsNil

`func (o *UserIdMappingParams) SetCustomAttributesTypeParamsNil(b bool)`

 SetCustomAttributesTypeParamsNil sets the value for CustomAttributesTypeParams to be an explicit nil

### UnsetCustomAttributesTypeParams
`func (o *UserIdMappingParams) UnsetCustomAttributesTypeParams()`

UnsetCustomAttributesTypeParams ensures that no value is present for CustomAttributesTypeParams, not even an explicit nil
### GetFixedTypeParams

`func (o *UserIdMappingParams) GetFixedTypeParams() UserIdMappingParamsFixedTypeParams`

GetFixedTypeParams returns the FixedTypeParams field if non-nil, zero value otherwise.

### GetFixedTypeParamsOk

`func (o *UserIdMappingParams) GetFixedTypeParamsOk() (*UserIdMappingParamsFixedTypeParams, bool)`

GetFixedTypeParamsOk returns a tuple with the FixedTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedTypeParams

`func (o *UserIdMappingParams) SetFixedTypeParams(v UserIdMappingParamsFixedTypeParams)`

SetFixedTypeParams sets FixedTypeParams field to given value.

### HasFixedTypeParams

`func (o *UserIdMappingParams) HasFixedTypeParams() bool`

HasFixedTypeParams returns a boolean if a field has been set.

### SetFixedTypeParamsNil

`func (o *UserIdMappingParams) SetFixedTypeParamsNil(b bool)`

 SetFixedTypeParamsNil sets the value for FixedTypeParams to be an explicit nil

### UnsetFixedTypeParams
`func (o *UserIdMappingParams) UnsetFixedTypeParams()`

UnsetFixedTypeParams ensures that no value is present for FixedTypeParams, not even an explicit nil
### GetLdapProviderTypeParams

`func (o *UserIdMappingParams) GetLdapProviderTypeParams() UserIdMappingParamsLdapProviderTypeParams`

GetLdapProviderTypeParams returns the LdapProviderTypeParams field if non-nil, zero value otherwise.

### GetLdapProviderTypeParamsOk

`func (o *UserIdMappingParams) GetLdapProviderTypeParamsOk() (*UserIdMappingParamsLdapProviderTypeParams, bool)`

GetLdapProviderTypeParamsOk returns a tuple with the LdapProviderTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProviderTypeParams

`func (o *UserIdMappingParams) SetLdapProviderTypeParams(v UserIdMappingParamsLdapProviderTypeParams)`

SetLdapProviderTypeParams sets LdapProviderTypeParams field to given value.

### HasLdapProviderTypeParams

`func (o *UserIdMappingParams) HasLdapProviderTypeParams() bool`

HasLdapProviderTypeParams returns a boolean if a field has been set.

### SetLdapProviderTypeParamsNil

`func (o *UserIdMappingParams) SetLdapProviderTypeParamsNil(b bool)`

 SetLdapProviderTypeParamsNil sets the value for LdapProviderTypeParams to be an explicit nil

### UnsetLdapProviderTypeParams
`func (o *UserIdMappingParams) UnsetLdapProviderTypeParams()`

UnsetLdapProviderTypeParams ensures that no value is present for LdapProviderTypeParams, not even an explicit nil
### GetNisProviderTypeParams

`func (o *UserIdMappingParams) GetNisProviderTypeParams() UserIdMappingParamsNisProviderTypeParams`

GetNisProviderTypeParams returns the NisProviderTypeParams field if non-nil, zero value otherwise.

### GetNisProviderTypeParamsOk

`func (o *UserIdMappingParams) GetNisProviderTypeParamsOk() (*UserIdMappingParamsNisProviderTypeParams, bool)`

GetNisProviderTypeParamsOk returns a tuple with the NisProviderTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNisProviderTypeParams

`func (o *UserIdMappingParams) SetNisProviderTypeParams(v UserIdMappingParamsNisProviderTypeParams)`

SetNisProviderTypeParams sets NisProviderTypeParams field to given value.

### HasNisProviderTypeParams

`func (o *UserIdMappingParams) HasNisProviderTypeParams() bool`

HasNisProviderTypeParams returns a boolean if a field has been set.

### SetNisProviderTypeParamsNil

`func (o *UserIdMappingParams) SetNisProviderTypeParamsNil(b bool)`

 SetNisProviderTypeParamsNil sets the value for NisProviderTypeParams to be an explicit nil

### UnsetNisProviderTypeParams
`func (o *UserIdMappingParams) UnsetNisProviderTypeParams()`

UnsetNisProviderTypeParams ensures that no value is present for NisProviderTypeParams, not even an explicit nil
### GetRfc2307TypeParams

`func (o *UserIdMappingParams) GetRfc2307TypeParams() UserIdMappingParamsRfc2307TypeParams`

GetRfc2307TypeParams returns the Rfc2307TypeParams field if non-nil, zero value otherwise.

### GetRfc2307TypeParamsOk

`func (o *UserIdMappingParams) GetRfc2307TypeParamsOk() (*UserIdMappingParamsRfc2307TypeParams, bool)`

GetRfc2307TypeParamsOk returns a tuple with the Rfc2307TypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfc2307TypeParams

`func (o *UserIdMappingParams) SetRfc2307TypeParams(v UserIdMappingParamsRfc2307TypeParams)`

SetRfc2307TypeParams sets Rfc2307TypeParams field to given value.

### HasRfc2307TypeParams

`func (o *UserIdMappingParams) HasRfc2307TypeParams() bool`

HasRfc2307TypeParams returns a boolean if a field has been set.

### SetRfc2307TypeParamsNil

`func (o *UserIdMappingParams) SetRfc2307TypeParamsNil(b bool)`

 SetRfc2307TypeParamsNil sets the value for Rfc2307TypeParams to be an explicit nil

### UnsetRfc2307TypeParams
`func (o *UserIdMappingParams) UnsetRfc2307TypeParams()`

UnsetRfc2307TypeParams ensures that no value is present for Rfc2307TypeParams, not even an explicit nil
### GetSfu30TypeParams

`func (o *UserIdMappingParams) GetSfu30TypeParams() UserIdMappingParamsSfu30TypeParams`

GetSfu30TypeParams returns the Sfu30TypeParams field if non-nil, zero value otherwise.

### GetSfu30TypeParamsOk

`func (o *UserIdMappingParams) GetSfu30TypeParamsOk() (*UserIdMappingParamsSfu30TypeParams, bool)`

GetSfu30TypeParamsOk returns a tuple with the Sfu30TypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfu30TypeParams

`func (o *UserIdMappingParams) SetSfu30TypeParams(v UserIdMappingParamsSfu30TypeParams)`

SetSfu30TypeParams sets Sfu30TypeParams field to given value.

### HasSfu30TypeParams

`func (o *UserIdMappingParams) HasSfu30TypeParams() bool`

HasSfu30TypeParams returns a boolean if a field has been set.

### SetSfu30TypeParamsNil

`func (o *UserIdMappingParams) SetSfu30TypeParamsNil(b bool)`

 SetSfu30TypeParamsNil sets the value for Sfu30TypeParams to be an explicit nil

### UnsetSfu30TypeParams
`func (o *UserIdMappingParams) UnsetSfu30TypeParams()`

UnsetSfu30TypeParams ensures that no value is present for Sfu30TypeParams, not even an explicit nil
### GetType

`func (o *UserIdMappingParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserIdMappingParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserIdMappingParams) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


