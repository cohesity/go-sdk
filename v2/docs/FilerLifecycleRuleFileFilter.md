# FilerLifecycleRuleFileFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileExtensions** | Pointer to **[]string** | Specifies the file&#39;s selection based on their extension. Eg: .pdf, .txt, etc. Note: Provide extensions here with the initial &#39;.&#39; character, example .pdf and not pdf. Extensions are case-insensitive, i.e. .pdf extension in filter will delete all files have .pdf, .PDF, .pDF, etc. | [optional] 
**FileSize** | Pointer to [**FilerLifecycleRuleFilterFileSize**](FilerLifecycleRuleFilterFileSize.md) |  | [optional] 

## Methods

### NewFilerLifecycleRuleFileFilter

`func NewFilerLifecycleRuleFileFilter() *FilerLifecycleRuleFileFilter`

NewFilerLifecycleRuleFileFilter instantiates a new FilerLifecycleRuleFileFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilerLifecycleRuleFileFilterWithDefaults

`func NewFilerLifecycleRuleFileFilterWithDefaults() *FilerLifecycleRuleFileFilter`

NewFilerLifecycleRuleFileFilterWithDefaults instantiates a new FilerLifecycleRuleFileFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileExtensions

`func (o *FilerLifecycleRuleFileFilter) GetFileExtensions() []string`

GetFileExtensions returns the FileExtensions field if non-nil, zero value otherwise.

### GetFileExtensionsOk

`func (o *FilerLifecycleRuleFileFilter) GetFileExtensionsOk() (*[]string, bool)`

GetFileExtensionsOk returns a tuple with the FileExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtensions

`func (o *FilerLifecycleRuleFileFilter) SetFileExtensions(v []string)`

SetFileExtensions sets FileExtensions field to given value.

### HasFileExtensions

`func (o *FilerLifecycleRuleFileFilter) HasFileExtensions() bool`

HasFileExtensions returns a boolean if a field has been set.

### SetFileExtensionsNil

`func (o *FilerLifecycleRuleFileFilter) SetFileExtensionsNil(b bool)`

 SetFileExtensionsNil sets the value for FileExtensions to be an explicit nil

### UnsetFileExtensions
`func (o *FilerLifecycleRuleFileFilter) UnsetFileExtensions()`

UnsetFileExtensions ensures that no value is present for FileExtensions, not even an explicit nil
### GetFileSize

`func (o *FilerLifecycleRuleFileFilter) GetFileSize() FilerLifecycleRuleFilterFileSize`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *FilerLifecycleRuleFileFilter) GetFileSizeOk() (*FilerLifecycleRuleFilterFileSize, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *FilerLifecycleRuleFileFilter) SetFileSize(v FilerLifecycleRuleFilterFileSize)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *FilerLifecycleRuleFileFilter) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


