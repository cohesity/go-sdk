# NumberFormFieldParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **NullableString** | Default value for the field | [optional] 
**Description** | Pointer to **NullableString** | Description for the field to be shown on UI screen | [optional] 
**MaximumValue** | Pointer to **NullableString** | Maximum allowable value for the field | [optional] 
**MinimumValue** | Pointer to **NullableString** | Minimum allowable value for the field | [optional] 
**Placeholder** | Pointer to **NullableString** | Placeholder for the form field | [optional] 
**Required** | Pointer to **NullableBool** | Specifies whether the field is mandatory | [optional] 

## Methods

### NewNumberFormFieldParams

`func NewNumberFormFieldParams() *NumberFormFieldParams`

NewNumberFormFieldParams instantiates a new NumberFormFieldParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberFormFieldParamsWithDefaults

`func NewNumberFormFieldParamsWithDefaults() *NumberFormFieldParams`

NewNumberFormFieldParamsWithDefaults instantiates a new NumberFormFieldParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *NumberFormFieldParams) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *NumberFormFieldParams) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *NumberFormFieldParams) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *NumberFormFieldParams) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *NumberFormFieldParams) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *NumberFormFieldParams) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetDescription

`func (o *NumberFormFieldParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NumberFormFieldParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NumberFormFieldParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NumberFormFieldParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NumberFormFieldParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NumberFormFieldParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMaximumValue

`func (o *NumberFormFieldParams) GetMaximumValue() string`

GetMaximumValue returns the MaximumValue field if non-nil, zero value otherwise.

### GetMaximumValueOk

`func (o *NumberFormFieldParams) GetMaximumValueOk() (*string, bool)`

GetMaximumValueOk returns a tuple with the MaximumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumValue

`func (o *NumberFormFieldParams) SetMaximumValue(v string)`

SetMaximumValue sets MaximumValue field to given value.

### HasMaximumValue

`func (o *NumberFormFieldParams) HasMaximumValue() bool`

HasMaximumValue returns a boolean if a field has been set.

### SetMaximumValueNil

`func (o *NumberFormFieldParams) SetMaximumValueNil(b bool)`

 SetMaximumValueNil sets the value for MaximumValue to be an explicit nil

### UnsetMaximumValue
`func (o *NumberFormFieldParams) UnsetMaximumValue()`

UnsetMaximumValue ensures that no value is present for MaximumValue, not even an explicit nil
### GetMinimumValue

`func (o *NumberFormFieldParams) GetMinimumValue() string`

GetMinimumValue returns the MinimumValue field if non-nil, zero value otherwise.

### GetMinimumValueOk

`func (o *NumberFormFieldParams) GetMinimumValueOk() (*string, bool)`

GetMinimumValueOk returns a tuple with the MinimumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumValue

`func (o *NumberFormFieldParams) SetMinimumValue(v string)`

SetMinimumValue sets MinimumValue field to given value.

### HasMinimumValue

`func (o *NumberFormFieldParams) HasMinimumValue() bool`

HasMinimumValue returns a boolean if a field has been set.

### SetMinimumValueNil

`func (o *NumberFormFieldParams) SetMinimumValueNil(b bool)`

 SetMinimumValueNil sets the value for MinimumValue to be an explicit nil

### UnsetMinimumValue
`func (o *NumberFormFieldParams) UnsetMinimumValue()`

UnsetMinimumValue ensures that no value is present for MinimumValue, not even an explicit nil
### GetPlaceholder

`func (o *NumberFormFieldParams) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *NumberFormFieldParams) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *NumberFormFieldParams) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *NumberFormFieldParams) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### SetPlaceholderNil

`func (o *NumberFormFieldParams) SetPlaceholderNil(b bool)`

 SetPlaceholderNil sets the value for Placeholder to be an explicit nil

### UnsetPlaceholder
`func (o *NumberFormFieldParams) UnsetPlaceholder()`

UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
### GetRequired

`func (o *NumberFormFieldParams) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *NumberFormFieldParams) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *NumberFormFieldParams) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *NumberFormFieldParams) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *NumberFormFieldParams) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *NumberFormFieldParams) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


