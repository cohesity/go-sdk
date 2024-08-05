# ViewIntent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultTemplateName** | Pointer to **NullableString** | Used for uniquely indentifying a default template | [optional] 
**TemplateId** | Pointer to **NullableInt64** | Specifies the template id from which the View is created. | [optional] 
**TemplateName** | Pointer to **NullableString** | Specifies the template name from which the View is created. | [optional] 

## Methods

### NewViewIntent

`func NewViewIntent() *ViewIntent`

NewViewIntent instantiates a new ViewIntent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewIntentWithDefaults

`func NewViewIntentWithDefaults() *ViewIntent`

NewViewIntentWithDefaults instantiates a new ViewIntent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultTemplateName

`func (o *ViewIntent) GetDefaultTemplateName() string`

GetDefaultTemplateName returns the DefaultTemplateName field if non-nil, zero value otherwise.

### GetDefaultTemplateNameOk

`func (o *ViewIntent) GetDefaultTemplateNameOk() (*string, bool)`

GetDefaultTemplateNameOk returns a tuple with the DefaultTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTemplateName

`func (o *ViewIntent) SetDefaultTemplateName(v string)`

SetDefaultTemplateName sets DefaultTemplateName field to given value.

### HasDefaultTemplateName

`func (o *ViewIntent) HasDefaultTemplateName() bool`

HasDefaultTemplateName returns a boolean if a field has been set.

### SetDefaultTemplateNameNil

`func (o *ViewIntent) SetDefaultTemplateNameNil(b bool)`

 SetDefaultTemplateNameNil sets the value for DefaultTemplateName to be an explicit nil

### UnsetDefaultTemplateName
`func (o *ViewIntent) UnsetDefaultTemplateName()`

UnsetDefaultTemplateName ensures that no value is present for DefaultTemplateName, not even an explicit nil
### GetTemplateId

`func (o *ViewIntent) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ViewIntent) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ViewIntent) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ViewIntent) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *ViewIntent) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *ViewIntent) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetTemplateName

`func (o *ViewIntent) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *ViewIntent) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *ViewIntent) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *ViewIntent) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### SetTemplateNameNil

`func (o *ViewIntent) SetTemplateNameNil(b bool)`

 SetTemplateNameNil sets the value for TemplateName to be an explicit nil

### UnsetTemplateName
`func (o *ViewIntent) UnsetTemplateName()`

UnsetTemplateName ensures that no value is present for TemplateName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


