# KeystoneScopeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainScopeParams** | Pointer to [**KeystoneScopeParamsDomainScopeParams**](KeystoneScopeParamsDomainScopeParams.md) |  | [optional] 
**ProjectScopeParams** | Pointer to [**KeystoneScopeParamsProjectScopeParams**](KeystoneScopeParamsProjectScopeParams.md) |  | [optional] 
**Type** | **NullableString** | Specifies the scope type. | 

## Methods

### NewKeystoneScopeParams

`func NewKeystoneScopeParams(type_ NullableString, ) *KeystoneScopeParams`

NewKeystoneScopeParams instantiates a new KeystoneScopeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeystoneScopeParamsWithDefaults

`func NewKeystoneScopeParamsWithDefaults() *KeystoneScopeParams`

NewKeystoneScopeParamsWithDefaults instantiates a new KeystoneScopeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainScopeParams

`func (o *KeystoneScopeParams) GetDomainScopeParams() KeystoneScopeParamsDomainScopeParams`

GetDomainScopeParams returns the DomainScopeParams field if non-nil, zero value otherwise.

### GetDomainScopeParamsOk

`func (o *KeystoneScopeParams) GetDomainScopeParamsOk() (*KeystoneScopeParamsDomainScopeParams, bool)`

GetDomainScopeParamsOk returns a tuple with the DomainScopeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainScopeParams

`func (o *KeystoneScopeParams) SetDomainScopeParams(v KeystoneScopeParamsDomainScopeParams)`

SetDomainScopeParams sets DomainScopeParams field to given value.

### HasDomainScopeParams

`func (o *KeystoneScopeParams) HasDomainScopeParams() bool`

HasDomainScopeParams returns a boolean if a field has been set.

### GetProjectScopeParams

`func (o *KeystoneScopeParams) GetProjectScopeParams() KeystoneScopeParamsProjectScopeParams`

GetProjectScopeParams returns the ProjectScopeParams field if non-nil, zero value otherwise.

### GetProjectScopeParamsOk

`func (o *KeystoneScopeParams) GetProjectScopeParamsOk() (*KeystoneScopeParamsProjectScopeParams, bool)`

GetProjectScopeParamsOk returns a tuple with the ProjectScopeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectScopeParams

`func (o *KeystoneScopeParams) SetProjectScopeParams(v KeystoneScopeParamsProjectScopeParams)`

SetProjectScopeParams sets ProjectScopeParams field to given value.

### HasProjectScopeParams

`func (o *KeystoneScopeParams) HasProjectScopeParams() bool`

HasProjectScopeParams returns a boolean if a field has been set.

### GetType

`func (o *KeystoneScopeParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeystoneScopeParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeystoneScopeParams) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *KeystoneScopeParams) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *KeystoneScopeParams) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


