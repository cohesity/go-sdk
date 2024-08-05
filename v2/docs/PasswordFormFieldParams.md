# PasswordFormFieldParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Description for the field to be shown on UI screen. | [optional] 
**Placeholder** | Pointer to **NullableString** | Placeholder value for the form field. | [optional] 
**Required** | Pointer to **NullableBool** | Specifies whether the field is mandatory. | [optional] 

## Methods

### NewPasswordFormFieldParams

`func NewPasswordFormFieldParams() *PasswordFormFieldParams`

NewPasswordFormFieldParams instantiates a new PasswordFormFieldParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordFormFieldParamsWithDefaults

`func NewPasswordFormFieldParamsWithDefaults() *PasswordFormFieldParams`

NewPasswordFormFieldParamsWithDefaults instantiates a new PasswordFormFieldParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PasswordFormFieldParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PasswordFormFieldParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PasswordFormFieldParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PasswordFormFieldParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PasswordFormFieldParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PasswordFormFieldParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPlaceholder

`func (o *PasswordFormFieldParams) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *PasswordFormFieldParams) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *PasswordFormFieldParams) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *PasswordFormFieldParams) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### SetPlaceholderNil

`func (o *PasswordFormFieldParams) SetPlaceholderNil(b bool)`

 SetPlaceholderNil sets the value for Placeholder to be an explicit nil

### UnsetPlaceholder
`func (o *PasswordFormFieldParams) UnsetPlaceholder()`

UnsetPlaceholder ensures that no value is present for Placeholder, not even an explicit nil
### GetRequired

`func (o *PasswordFormFieldParams) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *PasswordFormFieldParams) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *PasswordFormFieldParams) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *PasswordFormFieldParams) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *PasswordFormFieldParams) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *PasswordFormFieldParams) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


