# AlertProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **NullableString** | Specifies name of the property. | [optional] 
**Value** | Pointer to **NullableString** | Specifies value of the property. | [optional] 

## Methods

### NewAlertProperty

`func NewAlertProperty() *AlertProperty`

NewAlertProperty instantiates a new AlertProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertPropertyWithDefaults

`func NewAlertPropertyWithDefaults() *AlertProperty`

NewAlertPropertyWithDefaults instantiates a new AlertProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AlertProperty) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AlertProperty) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AlertProperty) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AlertProperty) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *AlertProperty) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *AlertProperty) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *AlertProperty) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AlertProperty) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AlertProperty) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AlertProperty) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *AlertProperty) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *AlertProperty) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


