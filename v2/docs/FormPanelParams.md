# FormPanelParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**[]FormFieldParams**](FormFieldParams.md) | Array of fields | [optional] 
**Id** | Pointer to **NullableString** | Id to identify the panel. This is also be used for assigning component html ids which can be leveraged for writing automation against the panel. | [optional] 
**Optional** | Pointer to **NullableBool** | Specifies whether the panel is optional and kept behind a toggle slider/collapsed state | [optional] 
**Title** | Pointer to **NullableString** | Title for the panel to be shown on UI screen | [optional] 

## Methods

### NewFormPanelParams

`func NewFormPanelParams() *FormPanelParams`

NewFormPanelParams instantiates a new FormPanelParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormPanelParamsWithDefaults

`func NewFormPanelParamsWithDefaults() *FormPanelParams`

NewFormPanelParamsWithDefaults instantiates a new FormPanelParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *FormPanelParams) GetFields() []FormFieldParams`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *FormPanelParams) GetFieldsOk() (*[]FormFieldParams, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *FormPanelParams) SetFields(v []FormFieldParams)`

SetFields sets Fields field to given value.

### HasFields

`func (o *FormPanelParams) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *FormPanelParams) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormPanelParams) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormPanelParams) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormPanelParams) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *FormPanelParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *FormPanelParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetOptional

`func (o *FormPanelParams) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *FormPanelParams) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *FormPanelParams) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *FormPanelParams) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### SetOptionalNil

`func (o *FormPanelParams) SetOptionalNil(b bool)`

 SetOptionalNil sets the value for Optional to be an explicit nil

### UnsetOptional
`func (o *FormPanelParams) UnsetOptional()`

UnsetOptional ensures that no value is present for Optional, not even an explicit nil
### GetTitle

`func (o *FormPanelParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FormPanelParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FormPanelParams) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FormPanelParams) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *FormPanelParams) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *FormPanelParams) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


