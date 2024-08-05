# RadioButtonFormFieldParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Id to identify the radio button. This is also be used for assigning component html ids which can be leveraged for writing automation against the button. | [optional] 
**Label** | Pointer to **NullableString** | Label to be shown for the radio button on the UI. | [optional] 
**Panel** | Pointer to [**NullableFormPanelParams**](FormPanelParams.md) |  | [optional] 
**Value** | Pointer to **NullableString** | Value associated with the radio button. | [optional] 

## Methods

### NewRadioButtonFormFieldParams

`func NewRadioButtonFormFieldParams() *RadioButtonFormFieldParams`

NewRadioButtonFormFieldParams instantiates a new RadioButtonFormFieldParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadioButtonFormFieldParamsWithDefaults

`func NewRadioButtonFormFieldParamsWithDefaults() *RadioButtonFormFieldParams`

NewRadioButtonFormFieldParamsWithDefaults instantiates a new RadioButtonFormFieldParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RadioButtonFormFieldParams) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RadioButtonFormFieldParams) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RadioButtonFormFieldParams) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RadioButtonFormFieldParams) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RadioButtonFormFieldParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RadioButtonFormFieldParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLabel

`func (o *RadioButtonFormFieldParams) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RadioButtonFormFieldParams) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RadioButtonFormFieldParams) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RadioButtonFormFieldParams) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *RadioButtonFormFieldParams) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *RadioButtonFormFieldParams) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetPanel

`func (o *RadioButtonFormFieldParams) GetPanel() FormPanelParams`

GetPanel returns the Panel field if non-nil, zero value otherwise.

### GetPanelOk

`func (o *RadioButtonFormFieldParams) GetPanelOk() (*FormPanelParams, bool)`

GetPanelOk returns a tuple with the Panel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanel

`func (o *RadioButtonFormFieldParams) SetPanel(v FormPanelParams)`

SetPanel sets Panel field to given value.

### HasPanel

`func (o *RadioButtonFormFieldParams) HasPanel() bool`

HasPanel returns a boolean if a field has been set.

### SetPanelNil

`func (o *RadioButtonFormFieldParams) SetPanelNil(b bool)`

 SetPanelNil sets the value for Panel to be an explicit nil

### UnsetPanel
`func (o *RadioButtonFormFieldParams) UnsetPanel()`

UnsetPanel ensures that no value is present for Panel, not even an explicit nil
### GetValue

`func (o *RadioButtonFormFieldParams) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RadioButtonFormFieldParams) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RadioButtonFormFieldParams) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RadioButtonFormFieldParams) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *RadioButtonFormFieldParams) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *RadioButtonFormFieldParams) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


