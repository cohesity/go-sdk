# UserIdMappingParamsCustomAttributesTypeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FallbackOption** | [**NullableAdCustomAttributesTypeParamsFallbackOption**](AdCustomAttributesTypeParamsFallbackOption.md) |  | 
**GidAttrName** | **NullableString** | Specifies the custom field name in Active Directory user properties to get the GID. | 
**UidAttrName** | **NullableString** | Specifies the custom field name in Active Directory user properties to get the UID. | 

## Methods

### NewUserIdMappingParamsCustomAttributesTypeParams

`func NewUserIdMappingParamsCustomAttributesTypeParams(fallbackOption NullableAdCustomAttributesTypeParamsFallbackOption, gidAttrName NullableString, uidAttrName NullableString, ) *UserIdMappingParamsCustomAttributesTypeParams`

NewUserIdMappingParamsCustomAttributesTypeParams instantiates a new UserIdMappingParamsCustomAttributesTypeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdMappingParamsCustomAttributesTypeParamsWithDefaults

`func NewUserIdMappingParamsCustomAttributesTypeParamsWithDefaults() *UserIdMappingParamsCustomAttributesTypeParams`

NewUserIdMappingParamsCustomAttributesTypeParamsWithDefaults instantiates a new UserIdMappingParamsCustomAttributesTypeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFallbackOption

`func (o *UserIdMappingParamsCustomAttributesTypeParams) GetFallbackOption() AdCustomAttributesTypeParamsFallbackOption`

GetFallbackOption returns the FallbackOption field if non-nil, zero value otherwise.

### GetFallbackOptionOk

`func (o *UserIdMappingParamsCustomAttributesTypeParams) GetFallbackOptionOk() (*AdCustomAttributesTypeParamsFallbackOption, bool)`

GetFallbackOptionOk returns a tuple with the FallbackOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackOption

`func (o *UserIdMappingParamsCustomAttributesTypeParams) SetFallbackOption(v AdCustomAttributesTypeParamsFallbackOption)`

SetFallbackOption sets FallbackOption field to given value.


### SetFallbackOptionNil

`func (o *UserIdMappingParamsCustomAttributesTypeParams) SetFallbackOptionNil(b bool)`

 SetFallbackOptionNil sets the value for FallbackOption to be an explicit nil

### UnsetFallbackOption
`func (o *UserIdMappingParamsCustomAttributesTypeParams) UnsetFallbackOption()`

UnsetFallbackOption ensures that no value is present for FallbackOption, not even an explicit nil
### GetGidAttrName

`func (o *UserIdMappingParamsCustomAttributesTypeParams) GetGidAttrName() string`

GetGidAttrName returns the GidAttrName field if non-nil, zero value otherwise.

### GetGidAttrNameOk

`func (o *UserIdMappingParamsCustomAttributesTypeParams) GetGidAttrNameOk() (*string, bool)`

GetGidAttrNameOk returns a tuple with the GidAttrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGidAttrName

`func (o *UserIdMappingParamsCustomAttributesTypeParams) SetGidAttrName(v string)`

SetGidAttrName sets GidAttrName field to given value.


### SetGidAttrNameNil

`func (o *UserIdMappingParamsCustomAttributesTypeParams) SetGidAttrNameNil(b bool)`

 SetGidAttrNameNil sets the value for GidAttrName to be an explicit nil

### UnsetGidAttrName
`func (o *UserIdMappingParamsCustomAttributesTypeParams) UnsetGidAttrName()`

UnsetGidAttrName ensures that no value is present for GidAttrName, not even an explicit nil
### GetUidAttrName

`func (o *UserIdMappingParamsCustomAttributesTypeParams) GetUidAttrName() string`

GetUidAttrName returns the UidAttrName field if non-nil, zero value otherwise.

### GetUidAttrNameOk

`func (o *UserIdMappingParamsCustomAttributesTypeParams) GetUidAttrNameOk() (*string, bool)`

GetUidAttrNameOk returns a tuple with the UidAttrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidAttrName

`func (o *UserIdMappingParamsCustomAttributesTypeParams) SetUidAttrName(v string)`

SetUidAttrName sets UidAttrName field to given value.


### SetUidAttrNameNil

`func (o *UserIdMappingParamsCustomAttributesTypeParams) SetUidAttrNameNil(b bool)`

 SetUidAttrNameNil sets the value for UidAttrName to be an explicit nil

### UnsetUidAttrName
`func (o *UserIdMappingParamsCustomAttributesTypeParams) UnsetUidAttrName()`

UnsetUidAttrName ensures that no value is present for UidAttrName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


