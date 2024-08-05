# FormFieldParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BooleanConfig** | Pointer to [**NullableBooleanFormFieldParams**](BooleanFormFieldParams.md) |  | [optional] 
**Id** | Pointer to **NullableString** | Id to identify the form field. This is also be used for assigning component html ids which can be leveraged for writing automation against the form field. | [optional] 
**Key** | Pointer to **NullableString** | Key against which the form field value will be returned | [optional] 
**Label** | Pointer to **NullableString** | Label to be shown on the UI screen | [optional] 
**NumberConfig** | Pointer to [**NullableNumberFormFieldParams**](NumberFormFieldParams.md) |  | [optional] 
**PasswordConfig** | Pointer to [**NullablePasswordFormFieldParams**](PasswordFormFieldParams.md) |  | [optional] 
**RadioGroupConfig** | Pointer to [**NullableRadioGroupFormFieldParams**](RadioGroupFormFieldParams.md) |  | [optional] 
**StringConfig** | Pointer to [**NullableStringFormFieldParams**](StringFormFieldParams.md) |  | [optional] 
**Type** | Pointer to **NullableString** | Type of the form field. Available types are &#39;string&#39;, &#39;password&#39;, &#39;number&#39;, &#39;boolean&#39;, &#39;radioGroup&#39; | [optional] 

## Methods

### NewFormFieldParams

`func NewFormFieldParams() *FormFieldParams`

NewFormFieldParams instantiates a new FormFieldParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldParamsWithDefaults

`func NewFormFieldParamsWithDefaults() *FormFieldParams`

NewFormFieldParamsWithDefaults instantiates a new FormFieldParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBooleanConfig

`func (o *FormFieldParams) GetBooleanConfig() BooleanFormFieldParams`

GetBooleanConfig returns the BooleanConfig field if non-nil, zero value otherwise.

### GetBooleanConfigOk

`func (o *FormFieldParams) GetBooleanConfigOk() (*BooleanFormFieldParams, bool)`

GetBooleanConfigOk returns a tuple with the BooleanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooleanConfig

`func (o *FormFieldParams) SetBooleanConfig(v BooleanFormFieldParams)`

SetBooleanConfig sets BooleanConfig field to given value.

### HasBooleanConfig

`func (o *FormFieldParams) HasBooleanConfig() bool`

HasBooleanConfig returns a boolean if a field has been set.

### SetBooleanConfigNil

`func (o *FormFieldParams) SetBooleanConfigNil(b bool)`

 SetBooleanConfigNil sets the value for BooleanConfig to be an explicit nil

### UnsetBooleanConfig
`func (o *FormFieldParams) UnsetBooleanConfig()`

UnsetBooleanConfig ensures that no value is present for BooleanConfig, not even an explicit nil
### GetId

`func (o *FormFieldParams) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormFieldParams) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormFieldParams) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormFieldParams) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *FormFieldParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *FormFieldParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetKey

`func (o *FormFieldParams) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormFieldParams) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormFieldParams) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FormFieldParams) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *FormFieldParams) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *FormFieldParams) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetLabel

`func (o *FormFieldParams) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormFieldParams) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormFieldParams) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FormFieldParams) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *FormFieldParams) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *FormFieldParams) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetNumberConfig

`func (o *FormFieldParams) GetNumberConfig() NumberFormFieldParams`

GetNumberConfig returns the NumberConfig field if non-nil, zero value otherwise.

### GetNumberConfigOk

`func (o *FormFieldParams) GetNumberConfigOk() (*NumberFormFieldParams, bool)`

GetNumberConfigOk returns a tuple with the NumberConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberConfig

`func (o *FormFieldParams) SetNumberConfig(v NumberFormFieldParams)`

SetNumberConfig sets NumberConfig field to given value.

### HasNumberConfig

`func (o *FormFieldParams) HasNumberConfig() bool`

HasNumberConfig returns a boolean if a field has been set.

### SetNumberConfigNil

`func (o *FormFieldParams) SetNumberConfigNil(b bool)`

 SetNumberConfigNil sets the value for NumberConfig to be an explicit nil

### UnsetNumberConfig
`func (o *FormFieldParams) UnsetNumberConfig()`

UnsetNumberConfig ensures that no value is present for NumberConfig, not even an explicit nil
### GetPasswordConfig

`func (o *FormFieldParams) GetPasswordConfig() PasswordFormFieldParams`

GetPasswordConfig returns the PasswordConfig field if non-nil, zero value otherwise.

### GetPasswordConfigOk

`func (o *FormFieldParams) GetPasswordConfigOk() (*PasswordFormFieldParams, bool)`

GetPasswordConfigOk returns a tuple with the PasswordConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordConfig

`func (o *FormFieldParams) SetPasswordConfig(v PasswordFormFieldParams)`

SetPasswordConfig sets PasswordConfig field to given value.

### HasPasswordConfig

`func (o *FormFieldParams) HasPasswordConfig() bool`

HasPasswordConfig returns a boolean if a field has been set.

### SetPasswordConfigNil

`func (o *FormFieldParams) SetPasswordConfigNil(b bool)`

 SetPasswordConfigNil sets the value for PasswordConfig to be an explicit nil

### UnsetPasswordConfig
`func (o *FormFieldParams) UnsetPasswordConfig()`

UnsetPasswordConfig ensures that no value is present for PasswordConfig, not even an explicit nil
### GetRadioGroupConfig

`func (o *FormFieldParams) GetRadioGroupConfig() RadioGroupFormFieldParams`

GetRadioGroupConfig returns the RadioGroupConfig field if non-nil, zero value otherwise.

### GetRadioGroupConfigOk

`func (o *FormFieldParams) GetRadioGroupConfigOk() (*RadioGroupFormFieldParams, bool)`

GetRadioGroupConfigOk returns a tuple with the RadioGroupConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioGroupConfig

`func (o *FormFieldParams) SetRadioGroupConfig(v RadioGroupFormFieldParams)`

SetRadioGroupConfig sets RadioGroupConfig field to given value.

### HasRadioGroupConfig

`func (o *FormFieldParams) HasRadioGroupConfig() bool`

HasRadioGroupConfig returns a boolean if a field has been set.

### SetRadioGroupConfigNil

`func (o *FormFieldParams) SetRadioGroupConfigNil(b bool)`

 SetRadioGroupConfigNil sets the value for RadioGroupConfig to be an explicit nil

### UnsetRadioGroupConfig
`func (o *FormFieldParams) UnsetRadioGroupConfig()`

UnsetRadioGroupConfig ensures that no value is present for RadioGroupConfig, not even an explicit nil
### GetStringConfig

`func (o *FormFieldParams) GetStringConfig() StringFormFieldParams`

GetStringConfig returns the StringConfig field if non-nil, zero value otherwise.

### GetStringConfigOk

`func (o *FormFieldParams) GetStringConfigOk() (*StringFormFieldParams, bool)`

GetStringConfigOk returns a tuple with the StringConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringConfig

`func (o *FormFieldParams) SetStringConfig(v StringFormFieldParams)`

SetStringConfig sets StringConfig field to given value.

### HasStringConfig

`func (o *FormFieldParams) HasStringConfig() bool`

HasStringConfig returns a boolean if a field has been set.

### SetStringConfigNil

`func (o *FormFieldParams) SetStringConfigNil(b bool)`

 SetStringConfigNil sets the value for StringConfig to be an explicit nil

### UnsetStringConfig
`func (o *FormFieldParams) UnsetStringConfig()`

UnsetStringConfig ensures that no value is present for StringConfig, not even an explicit nil
### GetType

`func (o *FormFieldParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldParams) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FormFieldParams) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *FormFieldParams) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FormFieldParams) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


