# SecurityConfigDataClassification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDataClassified** | Pointer to **NullableBool** | Specifies whether to mark the web page data classified/unclassified. | [optional] 
**ClassifiedDataMessage** | Pointer to **NullableString** | Specifies the classified data message. | [optional] 
**UnclassifiedDataMessage** | Pointer to **NullableString** | Specifies the unclassified data message. | [optional] 

## Methods

### NewSecurityConfigDataClassification

`func NewSecurityConfigDataClassification() *SecurityConfigDataClassification`

NewSecurityConfigDataClassification instantiates a new SecurityConfigDataClassification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityConfigDataClassificationWithDefaults

`func NewSecurityConfigDataClassificationWithDefaults() *SecurityConfigDataClassification`

NewSecurityConfigDataClassificationWithDefaults instantiates a new SecurityConfigDataClassification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDataClassified

`func (o *SecurityConfigDataClassification) GetIsDataClassified() bool`

GetIsDataClassified returns the IsDataClassified field if non-nil, zero value otherwise.

### GetIsDataClassifiedOk

`func (o *SecurityConfigDataClassification) GetIsDataClassifiedOk() (*bool, bool)`

GetIsDataClassifiedOk returns a tuple with the IsDataClassified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDataClassified

`func (o *SecurityConfigDataClassification) SetIsDataClassified(v bool)`

SetIsDataClassified sets IsDataClassified field to given value.

### HasIsDataClassified

`func (o *SecurityConfigDataClassification) HasIsDataClassified() bool`

HasIsDataClassified returns a boolean if a field has been set.

### SetIsDataClassifiedNil

`func (o *SecurityConfigDataClassification) SetIsDataClassifiedNil(b bool)`

 SetIsDataClassifiedNil sets the value for IsDataClassified to be an explicit nil

### UnsetIsDataClassified
`func (o *SecurityConfigDataClassification) UnsetIsDataClassified()`

UnsetIsDataClassified ensures that no value is present for IsDataClassified, not even an explicit nil
### GetClassifiedDataMessage

`func (o *SecurityConfigDataClassification) GetClassifiedDataMessage() string`

GetClassifiedDataMessage returns the ClassifiedDataMessage field if non-nil, zero value otherwise.

### GetClassifiedDataMessageOk

`func (o *SecurityConfigDataClassification) GetClassifiedDataMessageOk() (*string, bool)`

GetClassifiedDataMessageOk returns a tuple with the ClassifiedDataMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifiedDataMessage

`func (o *SecurityConfigDataClassification) SetClassifiedDataMessage(v string)`

SetClassifiedDataMessage sets ClassifiedDataMessage field to given value.

### HasClassifiedDataMessage

`func (o *SecurityConfigDataClassification) HasClassifiedDataMessage() bool`

HasClassifiedDataMessage returns a boolean if a field has been set.

### SetClassifiedDataMessageNil

`func (o *SecurityConfigDataClassification) SetClassifiedDataMessageNil(b bool)`

 SetClassifiedDataMessageNil sets the value for ClassifiedDataMessage to be an explicit nil

### UnsetClassifiedDataMessage
`func (o *SecurityConfigDataClassification) UnsetClassifiedDataMessage()`

UnsetClassifiedDataMessage ensures that no value is present for ClassifiedDataMessage, not even an explicit nil
### GetUnclassifiedDataMessage

`func (o *SecurityConfigDataClassification) GetUnclassifiedDataMessage() string`

GetUnclassifiedDataMessage returns the UnclassifiedDataMessage field if non-nil, zero value otherwise.

### GetUnclassifiedDataMessageOk

`func (o *SecurityConfigDataClassification) GetUnclassifiedDataMessageOk() (*string, bool)`

GetUnclassifiedDataMessageOk returns a tuple with the UnclassifiedDataMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnclassifiedDataMessage

`func (o *SecurityConfigDataClassification) SetUnclassifiedDataMessage(v string)`

SetUnclassifiedDataMessage sets UnclassifiedDataMessage field to given value.

### HasUnclassifiedDataMessage

`func (o *SecurityConfigDataClassification) HasUnclassifiedDataMessage() bool`

HasUnclassifiedDataMessage returns a boolean if a field has been set.

### SetUnclassifiedDataMessageNil

`func (o *SecurityConfigDataClassification) SetUnclassifiedDataMessageNil(b bool)`

 SetUnclassifiedDataMessageNil sets the value for UnclassifiedDataMessage to be an explicit nil

### UnsetUnclassifiedDataMessage
`func (o *SecurityConfigDataClassification) UnsetUnclassifiedDataMessage()`

UnsetUnclassifiedDataMessage ensures that no value is present for UnclassifiedDataMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


