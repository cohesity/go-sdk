# VMFilterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseSensitive** | Pointer to **NullableBool** | Specifies whether the provided filter string is case sensitive or not. This needs to be explicitly set to true if user is trying to filter by case sensitive expressions. The default value is assumed to be false. | [optional] [default to false]

## Methods

### NewVMFilterAllOf

`func NewVMFilterAllOf() *VMFilterAllOf`

NewVMFilterAllOf instantiates a new VMFilterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMFilterAllOfWithDefaults

`func NewVMFilterAllOfWithDefaults() *VMFilterAllOf`

NewVMFilterAllOfWithDefaults instantiates a new VMFilterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseSensitive

`func (o *VMFilterAllOf) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *VMFilterAllOf) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *VMFilterAllOf) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *VMFilterAllOf) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### SetCaseSensitiveNil

`func (o *VMFilterAllOf) SetCaseSensitiveNil(b bool)`

 SetCaseSensitiveNil sets the value for CaseSensitive to be an explicit nil

### UnsetCaseSensitive
`func (o *VMFilterAllOf) UnsetCaseSensitive()`

UnsetCaseSensitive ensures that no value is present for CaseSensitive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


