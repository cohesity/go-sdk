# FilenamePatternToDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Directory** | Pointer to **NullableString** | Specifies the directory where to keep the files matching the pattern. | [optional] 
**FilenamePattern** | Pointer to **NullableString** | Specifies a pattern to be matched with filenames. This can be a regex expression. | [optional] 

## Methods

### NewFilenamePatternToDirectory

`func NewFilenamePatternToDirectory() *FilenamePatternToDirectory`

NewFilenamePatternToDirectory instantiates a new FilenamePatternToDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilenamePatternToDirectoryWithDefaults

`func NewFilenamePatternToDirectoryWithDefaults() *FilenamePatternToDirectory`

NewFilenamePatternToDirectoryWithDefaults instantiates a new FilenamePatternToDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectory

`func (o *FilenamePatternToDirectory) GetDirectory() string`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *FilenamePatternToDirectory) GetDirectoryOk() (*string, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *FilenamePatternToDirectory) SetDirectory(v string)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *FilenamePatternToDirectory) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### SetDirectoryNil

`func (o *FilenamePatternToDirectory) SetDirectoryNil(b bool)`

 SetDirectoryNil sets the value for Directory to be an explicit nil

### UnsetDirectory
`func (o *FilenamePatternToDirectory) UnsetDirectory()`

UnsetDirectory ensures that no value is present for Directory, not even an explicit nil
### GetFilenamePattern

`func (o *FilenamePatternToDirectory) GetFilenamePattern() string`

GetFilenamePattern returns the FilenamePattern field if non-nil, zero value otherwise.

### GetFilenamePatternOk

`func (o *FilenamePatternToDirectory) GetFilenamePatternOk() (*string, bool)`

GetFilenamePatternOk returns a tuple with the FilenamePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilenamePattern

`func (o *FilenamePatternToDirectory) SetFilenamePattern(v string)`

SetFilenamePattern sets FilenamePattern field to given value.

### HasFilenamePattern

`func (o *FilenamePatternToDirectory) HasFilenamePattern() bool`

HasFilenamePattern returns a boolean if a field has been set.

### SetFilenamePatternNil

`func (o *FilenamePatternToDirectory) SetFilenamePatternNil(b bool)`

 SetFilenamePatternNil sets the value for FilenamePattern to be an explicit nil

### UnsetFilenamePattern
`func (o *FilenamePatternToDirectory) UnsetFilenamePattern()`

UnsetFilenamePattern ensures that no value is present for FilenamePattern, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


