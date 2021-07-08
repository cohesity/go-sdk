# FilesToDirectoryMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePattern** | Pointer to **NullableString** | Source file name. The file name can be a regex matching source files. | [optional] 
**TargetDirectory** | Pointer to **NullableString** | Target directtory for the source file pattern. | [optional] 

## Methods

### NewFilesToDirectoryMapping

`func NewFilesToDirectoryMapping() *FilesToDirectoryMapping`

NewFilesToDirectoryMapping instantiates a new FilesToDirectoryMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesToDirectoryMappingWithDefaults

`func NewFilesToDirectoryMappingWithDefaults() *FilesToDirectoryMapping`

NewFilesToDirectoryMappingWithDefaults instantiates a new FilesToDirectoryMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePattern

`func (o *FilesToDirectoryMapping) GetFilePattern() string`

GetFilePattern returns the FilePattern field if non-nil, zero value otherwise.

### GetFilePatternOk

`func (o *FilesToDirectoryMapping) GetFilePatternOk() (*string, bool)`

GetFilePatternOk returns a tuple with the FilePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePattern

`func (o *FilesToDirectoryMapping) SetFilePattern(v string)`

SetFilePattern sets FilePattern field to given value.

### HasFilePattern

`func (o *FilesToDirectoryMapping) HasFilePattern() bool`

HasFilePattern returns a boolean if a field has been set.

### SetFilePatternNil

`func (o *FilesToDirectoryMapping) SetFilePatternNil(b bool)`

 SetFilePatternNil sets the value for FilePattern to be an explicit nil

### UnsetFilePattern
`func (o *FilesToDirectoryMapping) UnsetFilePattern()`

UnsetFilePattern ensures that no value is present for FilePattern, not even an explicit nil
### GetTargetDirectory

`func (o *FilesToDirectoryMapping) GetTargetDirectory() string`

GetTargetDirectory returns the TargetDirectory field if non-nil, zero value otherwise.

### GetTargetDirectoryOk

`func (o *FilesToDirectoryMapping) GetTargetDirectoryOk() (*string, bool)`

GetTargetDirectoryOk returns a tuple with the TargetDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDirectory

`func (o *FilesToDirectoryMapping) SetTargetDirectory(v string)`

SetTargetDirectory sets TargetDirectory field to given value.

### HasTargetDirectory

`func (o *FilesToDirectoryMapping) HasTargetDirectory() bool`

HasTargetDirectory returns a boolean if a field has been set.

### SetTargetDirectoryNil

`func (o *FilesToDirectoryMapping) SetTargetDirectoryNil(b bool)`

 SetTargetDirectoryNil sets the value for TargetDirectory to be an explicit nil

### UnsetTargetDirectory
`func (o *FilesToDirectoryMapping) UnsetTargetDirectory()`

UnsetTargetDirectory ensures that no value is present for TargetDirectory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


