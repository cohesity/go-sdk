# AdCustomAttributesTypeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UidAttrName** | **NullableString** | Specifies the custom field name in Active Directory user properties to get the UID. | 
**GidAttrName** | **NullableString** | Specifies the custom field name in Active Directory user properties to get the GID. | 
**FallbackOption** | [**NullableFallbackUserIdMappingParams**](FallbackUserIdMappingParams.md) | Specifies a fallback user id mapping param in case the primary config does not work. | 

## Methods

### NewAdCustomAttributesTypeParams

`func NewAdCustomAttributesTypeParams(uidAttrName NullableString, gidAttrName NullableString, fallbackOption NullableFallbackUserIdMappingParams, ) *AdCustomAttributesTypeParams`

NewAdCustomAttributesTypeParams instantiates a new AdCustomAttributesTypeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdCustomAttributesTypeParamsWithDefaults

`func NewAdCustomAttributesTypeParamsWithDefaults() *AdCustomAttributesTypeParams`

NewAdCustomAttributesTypeParamsWithDefaults instantiates a new AdCustomAttributesTypeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUidAttrName

`func (o *AdCustomAttributesTypeParams) GetUidAttrName() string`

GetUidAttrName returns the UidAttrName field if non-nil, zero value otherwise.

### GetUidAttrNameOk

`func (o *AdCustomAttributesTypeParams) GetUidAttrNameOk() (*string, bool)`

GetUidAttrNameOk returns a tuple with the UidAttrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidAttrName

`func (o *AdCustomAttributesTypeParams) SetUidAttrName(v string)`

SetUidAttrName sets UidAttrName field to given value.


### SetUidAttrNameNil

`func (o *AdCustomAttributesTypeParams) SetUidAttrNameNil(b bool)`

 SetUidAttrNameNil sets the value for UidAttrName to be an explicit nil

### UnsetUidAttrName
`func (o *AdCustomAttributesTypeParams) UnsetUidAttrName()`

UnsetUidAttrName ensures that no value is present for UidAttrName, not even an explicit nil
### GetGidAttrName

`func (o *AdCustomAttributesTypeParams) GetGidAttrName() string`

GetGidAttrName returns the GidAttrName field if non-nil, zero value otherwise.

### GetGidAttrNameOk

`func (o *AdCustomAttributesTypeParams) GetGidAttrNameOk() (*string, bool)`

GetGidAttrNameOk returns a tuple with the GidAttrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGidAttrName

`func (o *AdCustomAttributesTypeParams) SetGidAttrName(v string)`

SetGidAttrName sets GidAttrName field to given value.


### SetGidAttrNameNil

`func (o *AdCustomAttributesTypeParams) SetGidAttrNameNil(b bool)`

 SetGidAttrNameNil sets the value for GidAttrName to be an explicit nil

### UnsetGidAttrName
`func (o *AdCustomAttributesTypeParams) UnsetGidAttrName()`

UnsetGidAttrName ensures that no value is present for GidAttrName, not even an explicit nil
### GetFallbackOption

`func (o *AdCustomAttributesTypeParams) GetFallbackOption() FallbackUserIdMappingParams`

GetFallbackOption returns the FallbackOption field if non-nil, zero value otherwise.

### GetFallbackOptionOk

`func (o *AdCustomAttributesTypeParams) GetFallbackOptionOk() (*FallbackUserIdMappingParams, bool)`

GetFallbackOptionOk returns a tuple with the FallbackOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackOption

`func (o *AdCustomAttributesTypeParams) SetFallbackOption(v FallbackUserIdMappingParams)`

SetFallbackOption sets FallbackOption field to given value.


### SetFallbackOptionNil

`func (o *AdCustomAttributesTypeParams) SetFallbackOptionNil(b bool)`

 SetFallbackOptionNil sets the value for FallbackOption to be an explicit nil

### UnsetFallbackOption
`func (o *AdCustomAttributesTypeParams) UnsetFallbackOption()`

UnsetFallbackOption ensures that no value is present for FallbackOption, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


