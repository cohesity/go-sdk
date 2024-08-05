# Template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compress** | Pointer to **NullableBool** | Specifies whether to enable compression in storage domain. | [optional] 
**Dedup** | Pointer to **NullableBool** | Specifies whether to enable dedup in storage domain. | [optional] 
**DefaultTemplateName** | Pointer to **NullableString** | Used for uniquely indentifying a default template. | [optional] [readonly] 
**Id** | Pointer to **NullableInt64** | Specifies an id of the view template. | [optional] [readonly] 
**IsDefault** | Pointer to **NullableBool** | Specifies if the tempate is custom or static. | [optional] [readonly] 
**Name** | Pointer to **NullableString** | Specifies the name of the view template. | [optional] 
**ViewParams** | Pointer to [**TemplateViewParams**](TemplateViewParams.md) |  | [optional] 

## Methods

### NewTemplate

`func NewTemplate() *Template`

NewTemplate instantiates a new Template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateWithDefaults

`func NewTemplateWithDefaults() *Template`

NewTemplateWithDefaults instantiates a new Template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompress

`func (o *Template) GetCompress() bool`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *Template) GetCompressOk() (*bool, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *Template) SetCompress(v bool)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *Template) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### SetCompressNil

`func (o *Template) SetCompressNil(b bool)`

 SetCompressNil sets the value for Compress to be an explicit nil

### UnsetCompress
`func (o *Template) UnsetCompress()`

UnsetCompress ensures that no value is present for Compress, not even an explicit nil
### GetDedup

`func (o *Template) GetDedup() bool`

GetDedup returns the Dedup field if non-nil, zero value otherwise.

### GetDedupOk

`func (o *Template) GetDedupOk() (*bool, bool)`

GetDedupOk returns a tuple with the Dedup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedup

`func (o *Template) SetDedup(v bool)`

SetDedup sets Dedup field to given value.

### HasDedup

`func (o *Template) HasDedup() bool`

HasDedup returns a boolean if a field has been set.

### SetDedupNil

`func (o *Template) SetDedupNil(b bool)`

 SetDedupNil sets the value for Dedup to be an explicit nil

### UnsetDedup
`func (o *Template) UnsetDedup()`

UnsetDedup ensures that no value is present for Dedup, not even an explicit nil
### GetDefaultTemplateName

`func (o *Template) GetDefaultTemplateName() string`

GetDefaultTemplateName returns the DefaultTemplateName field if non-nil, zero value otherwise.

### GetDefaultTemplateNameOk

`func (o *Template) GetDefaultTemplateNameOk() (*string, bool)`

GetDefaultTemplateNameOk returns a tuple with the DefaultTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTemplateName

`func (o *Template) SetDefaultTemplateName(v string)`

SetDefaultTemplateName sets DefaultTemplateName field to given value.

### HasDefaultTemplateName

`func (o *Template) HasDefaultTemplateName() bool`

HasDefaultTemplateName returns a boolean if a field has been set.

### SetDefaultTemplateNameNil

`func (o *Template) SetDefaultTemplateNameNil(b bool)`

 SetDefaultTemplateNameNil sets the value for DefaultTemplateName to be an explicit nil

### UnsetDefaultTemplateName
`func (o *Template) UnsetDefaultTemplateName()`

UnsetDefaultTemplateName ensures that no value is present for DefaultTemplateName, not even an explicit nil
### GetId

`func (o *Template) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Template) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Template) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Template) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Template) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Template) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsDefault

`func (o *Template) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Template) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Template) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Template) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *Template) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *Template) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetName

`func (o *Template) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Template) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Template) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Template) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Template) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Template) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetViewParams

`func (o *Template) GetViewParams() TemplateViewParams`

GetViewParams returns the ViewParams field if non-nil, zero value otherwise.

### GetViewParamsOk

`func (o *Template) GetViewParamsOk() (*TemplateViewParams, bool)`

GetViewParamsOk returns a tuple with the ViewParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewParams

`func (o *Template) SetViewParams(v TemplateViewParams)`

SetViewParams sets ViewParams field to given value.

### HasViewParams

`func (o *Template) HasViewParams() bool`

HasViewParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


