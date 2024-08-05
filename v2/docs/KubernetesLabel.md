# KubernetesLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **NullableString** | The key of the label, used to identify the label. | [optional] [readonly] 
**Value** | Pointer to **NullableString** | The value associated with the label key. | [optional] [readonly] 

## Methods

### NewKubernetesLabel

`func NewKubernetesLabel() *KubernetesLabel`

NewKubernetesLabel instantiates a new KubernetesLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesLabelWithDefaults

`func NewKubernetesLabelWithDefaults() *KubernetesLabel`

NewKubernetesLabelWithDefaults instantiates a new KubernetesLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *KubernetesLabel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KubernetesLabel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KubernetesLabel) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *KubernetesLabel) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *KubernetesLabel) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *KubernetesLabel) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *KubernetesLabel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KubernetesLabel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KubernetesLabel) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *KubernetesLabel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *KubernetesLabel) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *KubernetesLabel) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


