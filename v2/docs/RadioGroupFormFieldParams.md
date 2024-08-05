# RadioGroupFormFieldParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **NullableString** | Default selection of the radio button. | [optional] 
**RadioButtons** | Pointer to [**[]RadioButtonFormFieldParams**](RadioButtonFormFieldParams.md) | List of radio buttons part of the radio group. | [optional] 
**Required** | Pointer to **NullableBool** | Specifies whether the field is mandatory. | [optional] 

## Methods

### NewRadioGroupFormFieldParams

`func NewRadioGroupFormFieldParams() *RadioGroupFormFieldParams`

NewRadioGroupFormFieldParams instantiates a new RadioGroupFormFieldParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadioGroupFormFieldParamsWithDefaults

`func NewRadioGroupFormFieldParamsWithDefaults() *RadioGroupFormFieldParams`

NewRadioGroupFormFieldParamsWithDefaults instantiates a new RadioGroupFormFieldParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *RadioGroupFormFieldParams) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *RadioGroupFormFieldParams) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *RadioGroupFormFieldParams) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *RadioGroupFormFieldParams) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *RadioGroupFormFieldParams) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *RadioGroupFormFieldParams) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetRadioButtons

`func (o *RadioGroupFormFieldParams) GetRadioButtons() []RadioButtonFormFieldParams`

GetRadioButtons returns the RadioButtons field if non-nil, zero value otherwise.

### GetRadioButtonsOk

`func (o *RadioGroupFormFieldParams) GetRadioButtonsOk() (*[]RadioButtonFormFieldParams, bool)`

GetRadioButtonsOk returns a tuple with the RadioButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioButtons

`func (o *RadioGroupFormFieldParams) SetRadioButtons(v []RadioButtonFormFieldParams)`

SetRadioButtons sets RadioButtons field to given value.

### HasRadioButtons

`func (o *RadioGroupFormFieldParams) HasRadioButtons() bool`

HasRadioButtons returns a boolean if a field has been set.

### SetRadioButtonsNil

`func (o *RadioGroupFormFieldParams) SetRadioButtonsNil(b bool)`

 SetRadioButtonsNil sets the value for RadioButtons to be an explicit nil

### UnsetRadioButtons
`func (o *RadioGroupFormFieldParams) UnsetRadioButtons()`

UnsetRadioButtons ensures that no value is present for RadioButtons, not even an explicit nil
### GetRequired

`func (o *RadioGroupFormFieldParams) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *RadioGroupFormFieldParams) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *RadioGroupFormFieldParams) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *RadioGroupFormFieldParams) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *RadioGroupFormFieldParams) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *RadioGroupFormFieldParams) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


