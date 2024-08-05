# TagSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **NullableString** | Specifies the Key of the tag. | [optional] 
**Value** | Pointer to **NullableString** | Specifies the Value of the tag. | [optional] 

## Methods

### NewTagSet

`func NewTagSet() *TagSet`

NewTagSet instantiates a new TagSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagSetWithDefaults

`func NewTagSetWithDefaults() *TagSet`

NewTagSetWithDefaults instantiates a new TagSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *TagSet) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TagSet) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TagSet) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TagSet) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *TagSet) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *TagSet) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *TagSet) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TagSet) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TagSet) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TagSet) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *TagSet) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *TagSet) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


