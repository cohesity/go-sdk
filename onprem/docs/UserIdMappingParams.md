# UserIdMappingParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Specifies the type of the mapping. | 
**Rfc2307TypeParams** | Pointer to [**NullableAdRfc2307TypeParams**](AdRfc2307TypeParams.md) | Specifies the params for Rfc2307 mapping type mapping. | [optional] 
**Sfu30TypeParams** | Pointer to [**NullableAdSfu30TypeParams**](AdSfu30TypeParams.md) | Specifies the params for Sfu30 mapping type mapping. | [optional] 
**LdapProviderTypeParams** | Pointer to [**NullableAdLdapProviderTypeParams**](AdLdapProviderTypeParams.md) | Specifies the params for LdapProvider mapping type mapping. | [optional] 
**NisProviderTypeParams** | Pointer to [**NullableAdNisProviderTypeParams**](AdNisProviderTypeParams.md) | Specifies the params for NisProvider mapping type mapping. | [optional] 
**CentrifyTypeParams** | Pointer to [**NullableAdCentrifyTypeParams**](AdCentrifyTypeParams.md) | Specifies the params for Centrify mapping type mapping. | [optional] 
**FixedTypeParams** | Pointer to [**NullableAdFixedTypeParams**](AdFixedTypeParams.md) | Specifies the params for Fixed mapping type mapping. | [optional] 
**CustomAttributesTypeParams** | Pointer to [**NullableAdCustomAttributesTypeParams**](AdCustomAttributesTypeParams.md) | Specifies the params for CustomAttributes mapping type mapping. | [optional] 

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


### GetRfc2307TypeParams

`func (o *UserIdMappingParams) GetRfc2307TypeParams() AdRfc2307TypeParams`

GetRfc2307TypeParams returns the Rfc2307TypeParams field if non-nil, zero value otherwise.

### GetRfc2307TypeParamsOk

`func (o *UserIdMappingParams) GetRfc2307TypeParamsOk() (*AdRfc2307TypeParams, bool)`

GetRfc2307TypeParamsOk returns a tuple with the Rfc2307TypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfc2307TypeParams

`func (o *UserIdMappingParams) SetRfc2307TypeParams(v AdRfc2307TypeParams)`

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

`func (o *UserIdMappingParams) GetSfu30TypeParams() AdSfu30TypeParams`

GetSfu30TypeParams returns the Sfu30TypeParams field if non-nil, zero value otherwise.

### GetSfu30TypeParamsOk

`func (o *UserIdMappingParams) GetSfu30TypeParamsOk() (*AdSfu30TypeParams, bool)`

GetSfu30TypeParamsOk returns a tuple with the Sfu30TypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfu30TypeParams

`func (o *UserIdMappingParams) SetSfu30TypeParams(v AdSfu30TypeParams)`

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
### GetLdapProviderTypeParams

`func (o *UserIdMappingParams) GetLdapProviderTypeParams() AdLdapProviderTypeParams`

GetLdapProviderTypeParams returns the LdapProviderTypeParams field if non-nil, zero value otherwise.

### GetLdapProviderTypeParamsOk

`func (o *UserIdMappingParams) GetLdapProviderTypeParamsOk() (*AdLdapProviderTypeParams, bool)`

GetLdapProviderTypeParamsOk returns a tuple with the LdapProviderTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapProviderTypeParams

`func (o *UserIdMappingParams) SetLdapProviderTypeParams(v AdLdapProviderTypeParams)`

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

`func (o *UserIdMappingParams) GetNisProviderTypeParams() AdNisProviderTypeParams`

GetNisProviderTypeParams returns the NisProviderTypeParams field if non-nil, zero value otherwise.

### GetNisProviderTypeParamsOk

`func (o *UserIdMappingParams) GetNisProviderTypeParamsOk() (*AdNisProviderTypeParams, bool)`

GetNisProviderTypeParamsOk returns a tuple with the NisProviderTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNisProviderTypeParams

`func (o *UserIdMappingParams) SetNisProviderTypeParams(v AdNisProviderTypeParams)`

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
### GetCentrifyTypeParams

`func (o *UserIdMappingParams) GetCentrifyTypeParams() AdCentrifyTypeParams`

GetCentrifyTypeParams returns the CentrifyTypeParams field if non-nil, zero value otherwise.

### GetCentrifyTypeParamsOk

`func (o *UserIdMappingParams) GetCentrifyTypeParamsOk() (*AdCentrifyTypeParams, bool)`

GetCentrifyTypeParamsOk returns a tuple with the CentrifyTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentrifyTypeParams

`func (o *UserIdMappingParams) SetCentrifyTypeParams(v AdCentrifyTypeParams)`

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
### GetFixedTypeParams

`func (o *UserIdMappingParams) GetFixedTypeParams() AdFixedTypeParams`

GetFixedTypeParams returns the FixedTypeParams field if non-nil, zero value otherwise.

### GetFixedTypeParamsOk

`func (o *UserIdMappingParams) GetFixedTypeParamsOk() (*AdFixedTypeParams, bool)`

GetFixedTypeParamsOk returns a tuple with the FixedTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedTypeParams

`func (o *UserIdMappingParams) SetFixedTypeParams(v AdFixedTypeParams)`

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
### GetCustomAttributesTypeParams

`func (o *UserIdMappingParams) GetCustomAttributesTypeParams() AdCustomAttributesTypeParams`

GetCustomAttributesTypeParams returns the CustomAttributesTypeParams field if non-nil, zero value otherwise.

### GetCustomAttributesTypeParamsOk

`func (o *UserIdMappingParams) GetCustomAttributesTypeParamsOk() (*AdCustomAttributesTypeParams, bool)`

GetCustomAttributesTypeParamsOk returns a tuple with the CustomAttributesTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributesTypeParams

`func (o *UserIdMappingParams) SetCustomAttributesTypeParams(v AdCustomAttributesTypeParams)`

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


