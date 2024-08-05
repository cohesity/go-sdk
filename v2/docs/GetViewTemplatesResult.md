# GetViewTemplatesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Templates** | Pointer to [**[]Template**](Template.md) | Array of view template. Specifies the list of view templates returned in this response. | [optional] 

## Methods

### NewGetViewTemplatesResult

`func NewGetViewTemplatesResult() *GetViewTemplatesResult`

NewGetViewTemplatesResult instantiates a new GetViewTemplatesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetViewTemplatesResultWithDefaults

`func NewGetViewTemplatesResultWithDefaults() *GetViewTemplatesResult`

NewGetViewTemplatesResultWithDefaults instantiates a new GetViewTemplatesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplates

`func (o *GetViewTemplatesResult) GetTemplates() []Template`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *GetViewTemplatesResult) GetTemplatesOk() (*[]Template, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *GetViewTemplatesResult) SetTemplates(v []Template)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *GetViewTemplatesResult) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### SetTemplatesNil

`func (o *GetViewTemplatesResult) SetTemplatesNil(b bool)`

 SetTemplatesNil sets the value for Templates to be an explicit nil

### UnsetTemplates
`func (o *GetViewTemplatesResult) UnsetTemplates()`

UnsetTemplates ensures that no value is present for Templates, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


