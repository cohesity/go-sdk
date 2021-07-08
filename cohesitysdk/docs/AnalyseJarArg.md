# AnalyseJarArg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JarName** | Pointer to **NullableString** | Name of the JAR to be analysed. | [optional] 
**JarPath** | Pointer to **NullableString** | Path of the jar file. | [optional] 
**JarRelativePath** | Pointer to **NullableString** |  | [optional] 
**SaveEntities** | Pointer to **NullableBool** | If this flag is true, then also save mapper and reducers in scribe. | [optional] 

## Methods

### NewAnalyseJarArg

`func NewAnalyseJarArg() *AnalyseJarArg`

NewAnalyseJarArg instantiates a new AnalyseJarArg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyseJarArgWithDefaults

`func NewAnalyseJarArgWithDefaults() *AnalyseJarArg`

NewAnalyseJarArgWithDefaults instantiates a new AnalyseJarArg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJarName

`func (o *AnalyseJarArg) GetJarName() string`

GetJarName returns the JarName field if non-nil, zero value otherwise.

### GetJarNameOk

`func (o *AnalyseJarArg) GetJarNameOk() (*string, bool)`

GetJarNameOk returns a tuple with the JarName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJarName

`func (o *AnalyseJarArg) SetJarName(v string)`

SetJarName sets JarName field to given value.

### HasJarName

`func (o *AnalyseJarArg) HasJarName() bool`

HasJarName returns a boolean if a field has been set.

### SetJarNameNil

`func (o *AnalyseJarArg) SetJarNameNil(b bool)`

 SetJarNameNil sets the value for JarName to be an explicit nil

### UnsetJarName
`func (o *AnalyseJarArg) UnsetJarName()`

UnsetJarName ensures that no value is present for JarName, not even an explicit nil
### GetJarPath

`func (o *AnalyseJarArg) GetJarPath() string`

GetJarPath returns the JarPath field if non-nil, zero value otherwise.

### GetJarPathOk

`func (o *AnalyseJarArg) GetJarPathOk() (*string, bool)`

GetJarPathOk returns a tuple with the JarPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJarPath

`func (o *AnalyseJarArg) SetJarPath(v string)`

SetJarPath sets JarPath field to given value.

### HasJarPath

`func (o *AnalyseJarArg) HasJarPath() bool`

HasJarPath returns a boolean if a field has been set.

### SetJarPathNil

`func (o *AnalyseJarArg) SetJarPathNil(b bool)`

 SetJarPathNil sets the value for JarPath to be an explicit nil

### UnsetJarPath
`func (o *AnalyseJarArg) UnsetJarPath()`

UnsetJarPath ensures that no value is present for JarPath, not even an explicit nil
### GetJarRelativePath

`func (o *AnalyseJarArg) GetJarRelativePath() string`

GetJarRelativePath returns the JarRelativePath field if non-nil, zero value otherwise.

### GetJarRelativePathOk

`func (o *AnalyseJarArg) GetJarRelativePathOk() (*string, bool)`

GetJarRelativePathOk returns a tuple with the JarRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJarRelativePath

`func (o *AnalyseJarArg) SetJarRelativePath(v string)`

SetJarRelativePath sets JarRelativePath field to given value.

### HasJarRelativePath

`func (o *AnalyseJarArg) HasJarRelativePath() bool`

HasJarRelativePath returns a boolean if a field has been set.

### SetJarRelativePathNil

`func (o *AnalyseJarArg) SetJarRelativePathNil(b bool)`

 SetJarRelativePathNil sets the value for JarRelativePath to be an explicit nil

### UnsetJarRelativePath
`func (o *AnalyseJarArg) UnsetJarRelativePath()`

UnsetJarRelativePath ensures that no value is present for JarRelativePath, not even an explicit nil
### GetSaveEntities

`func (o *AnalyseJarArg) GetSaveEntities() bool`

GetSaveEntities returns the SaveEntities field if non-nil, zero value otherwise.

### GetSaveEntitiesOk

`func (o *AnalyseJarArg) GetSaveEntitiesOk() (*bool, bool)`

GetSaveEntitiesOk returns a tuple with the SaveEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveEntities

`func (o *AnalyseJarArg) SetSaveEntities(v bool)`

SetSaveEntities sets SaveEntities field to given value.

### HasSaveEntities

`func (o *AnalyseJarArg) HasSaveEntities() bool`

HasSaveEntities returns a boolean if a field has been set.

### SetSaveEntitiesNil

`func (o *AnalyseJarArg) SetSaveEntitiesNil(b bool)`

 SetSaveEntitiesNil sets the value for SaveEntities to be an explicit nil

### UnsetSaveEntities
`func (o *AnalyseJarArg) UnsetSaveEntities()`

UnsetSaveEntities ensures that no value is present for SaveEntities, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


