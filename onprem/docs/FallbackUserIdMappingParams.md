# FallbackUserIdMappingParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Specifies the type of the mapping. | 
**FixedTypeParams** | Pointer to [**NullableAdFixedTypeParams**](AdFixedTypeParams.md) | Specifies the params for Fixed mapping type mapping. | [optional] 

## Methods

### NewFallbackUserIdMappingParams

`func NewFallbackUserIdMappingParams(type_ string, ) *FallbackUserIdMappingParams`

NewFallbackUserIdMappingParams instantiates a new FallbackUserIdMappingParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFallbackUserIdMappingParamsWithDefaults

`func NewFallbackUserIdMappingParamsWithDefaults() *FallbackUserIdMappingParams`

NewFallbackUserIdMappingParamsWithDefaults instantiates a new FallbackUserIdMappingParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FallbackUserIdMappingParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FallbackUserIdMappingParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FallbackUserIdMappingParams) SetType(v string)`

SetType sets Type field to given value.


### GetFixedTypeParams

`func (o *FallbackUserIdMappingParams) GetFixedTypeParams() AdFixedTypeParams`

GetFixedTypeParams returns the FixedTypeParams field if non-nil, zero value otherwise.

### GetFixedTypeParamsOk

`func (o *FallbackUserIdMappingParams) GetFixedTypeParamsOk() (*AdFixedTypeParams, bool)`

GetFixedTypeParamsOk returns a tuple with the FixedTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedTypeParams

`func (o *FallbackUserIdMappingParams) SetFixedTypeParams(v AdFixedTypeParams)`

SetFixedTypeParams sets FixedTypeParams field to given value.

### HasFixedTypeParams

`func (o *FallbackUserIdMappingParams) HasFixedTypeParams() bool`

HasFixedTypeParams returns a boolean if a field has been set.

### SetFixedTypeParamsNil

`func (o *FallbackUserIdMappingParams) SetFixedTypeParamsNil(b bool)`

 SetFixedTypeParamsNil sets the value for FixedTypeParams to be an explicit nil

### UnsetFixedTypeParams
`func (o *FallbackUserIdMappingParams) UnsetFixedTypeParams()`

UnsetFixedTypeParams ensures that no value is present for FixedTypeParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


