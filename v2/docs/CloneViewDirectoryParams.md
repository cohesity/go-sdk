# CloneViewDirectoryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceDirectoryPath** | **NullableString** | Specifies the path of source directory. This is the full path including the Storage Domain name and View name. | 
**TargetDirectoryName** | **NullableString** | Specifies the name of the target directory. | 
**TargetParentDirectoryPath** | **NullableString** | Specifies the path of parent directory of the target directory. This is the full path including the Storage Domain name and View Name. | 

## Methods

### NewCloneViewDirectoryParams

`func NewCloneViewDirectoryParams(sourceDirectoryPath NullableString, targetDirectoryName NullableString, targetParentDirectoryPath NullableString, ) *CloneViewDirectoryParams`

NewCloneViewDirectoryParams instantiates a new CloneViewDirectoryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneViewDirectoryParamsWithDefaults

`func NewCloneViewDirectoryParamsWithDefaults() *CloneViewDirectoryParams`

NewCloneViewDirectoryParamsWithDefaults instantiates a new CloneViewDirectoryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceDirectoryPath

`func (o *CloneViewDirectoryParams) GetSourceDirectoryPath() string`

GetSourceDirectoryPath returns the SourceDirectoryPath field if non-nil, zero value otherwise.

### GetSourceDirectoryPathOk

`func (o *CloneViewDirectoryParams) GetSourceDirectoryPathOk() (*string, bool)`

GetSourceDirectoryPathOk returns a tuple with the SourceDirectoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDirectoryPath

`func (o *CloneViewDirectoryParams) SetSourceDirectoryPath(v string)`

SetSourceDirectoryPath sets SourceDirectoryPath field to given value.


### SetSourceDirectoryPathNil

`func (o *CloneViewDirectoryParams) SetSourceDirectoryPathNil(b bool)`

 SetSourceDirectoryPathNil sets the value for SourceDirectoryPath to be an explicit nil

### UnsetSourceDirectoryPath
`func (o *CloneViewDirectoryParams) UnsetSourceDirectoryPath()`

UnsetSourceDirectoryPath ensures that no value is present for SourceDirectoryPath, not even an explicit nil
### GetTargetDirectoryName

`func (o *CloneViewDirectoryParams) GetTargetDirectoryName() string`

GetTargetDirectoryName returns the TargetDirectoryName field if non-nil, zero value otherwise.

### GetTargetDirectoryNameOk

`func (o *CloneViewDirectoryParams) GetTargetDirectoryNameOk() (*string, bool)`

GetTargetDirectoryNameOk returns a tuple with the TargetDirectoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDirectoryName

`func (o *CloneViewDirectoryParams) SetTargetDirectoryName(v string)`

SetTargetDirectoryName sets TargetDirectoryName field to given value.


### SetTargetDirectoryNameNil

`func (o *CloneViewDirectoryParams) SetTargetDirectoryNameNil(b bool)`

 SetTargetDirectoryNameNil sets the value for TargetDirectoryName to be an explicit nil

### UnsetTargetDirectoryName
`func (o *CloneViewDirectoryParams) UnsetTargetDirectoryName()`

UnsetTargetDirectoryName ensures that no value is present for TargetDirectoryName, not even an explicit nil
### GetTargetParentDirectoryPath

`func (o *CloneViewDirectoryParams) GetTargetParentDirectoryPath() string`

GetTargetParentDirectoryPath returns the TargetParentDirectoryPath field if non-nil, zero value otherwise.

### GetTargetParentDirectoryPathOk

`func (o *CloneViewDirectoryParams) GetTargetParentDirectoryPathOk() (*string, bool)`

GetTargetParentDirectoryPathOk returns a tuple with the TargetParentDirectoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetParentDirectoryPath

`func (o *CloneViewDirectoryParams) SetTargetParentDirectoryPath(v string)`

SetTargetParentDirectoryPath sets TargetParentDirectoryPath field to given value.


### SetTargetParentDirectoryPathNil

`func (o *CloneViewDirectoryParams) SetTargetParentDirectoryPathNil(b bool)`

 SetTargetParentDirectoryPathNil sets the value for TargetParentDirectoryPath to be an explicit nil

### UnsetTargetParentDirectoryPath
`func (o *CloneViewDirectoryParams) UnsetTargetParentDirectoryPath()`

UnsetTargetParentDirectoryPath ensures that no value is present for TargetParentDirectoryPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


