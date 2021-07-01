# HdfsRecoverJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HdfsExcludePattern** | Pointer to **[]string** | Any path/Glob pattern from HDFS that is to excluded. | [optional] 
**HdfsRecoverPattern** | Pointer to **[]string** | Any path/Glob pattern from HDFS that is to recovered. | [optional] 
**TargetDirectory** | Pointer to **NullableString** | A target directory where all the recovered entities are created | [optional] 

## Methods

### NewHdfsRecoverJobParams

`func NewHdfsRecoverJobParams() *HdfsRecoverJobParams`

NewHdfsRecoverJobParams instantiates a new HdfsRecoverJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHdfsRecoverJobParamsWithDefaults

`func NewHdfsRecoverJobParamsWithDefaults() *HdfsRecoverJobParams`

NewHdfsRecoverJobParamsWithDefaults instantiates a new HdfsRecoverJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHdfsExcludePattern

`func (o *HdfsRecoverJobParams) GetHdfsExcludePattern() []string`

GetHdfsExcludePattern returns the HdfsExcludePattern field if non-nil, zero value otherwise.

### GetHdfsExcludePatternOk

`func (o *HdfsRecoverJobParams) GetHdfsExcludePatternOk() (*[]string, bool)`

GetHdfsExcludePatternOk returns a tuple with the HdfsExcludePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsExcludePattern

`func (o *HdfsRecoverJobParams) SetHdfsExcludePattern(v []string)`

SetHdfsExcludePattern sets HdfsExcludePattern field to given value.

### HasHdfsExcludePattern

`func (o *HdfsRecoverJobParams) HasHdfsExcludePattern() bool`

HasHdfsExcludePattern returns a boolean if a field has been set.

### SetHdfsExcludePatternNil

`func (o *HdfsRecoverJobParams) SetHdfsExcludePatternNil(b bool)`

 SetHdfsExcludePatternNil sets the value for HdfsExcludePattern to be an explicit nil

### UnsetHdfsExcludePattern
`func (o *HdfsRecoverJobParams) UnsetHdfsExcludePattern()`

UnsetHdfsExcludePattern ensures that no value is present for HdfsExcludePattern, not even an explicit nil
### GetHdfsRecoverPattern

`func (o *HdfsRecoverJobParams) GetHdfsRecoverPattern() []string`

GetHdfsRecoverPattern returns the HdfsRecoverPattern field if non-nil, zero value otherwise.

### GetHdfsRecoverPatternOk

`func (o *HdfsRecoverJobParams) GetHdfsRecoverPatternOk() (*[]string, bool)`

GetHdfsRecoverPatternOk returns a tuple with the HdfsRecoverPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsRecoverPattern

`func (o *HdfsRecoverJobParams) SetHdfsRecoverPattern(v []string)`

SetHdfsRecoverPattern sets HdfsRecoverPattern field to given value.

### HasHdfsRecoverPattern

`func (o *HdfsRecoverJobParams) HasHdfsRecoverPattern() bool`

HasHdfsRecoverPattern returns a boolean if a field has been set.

### SetHdfsRecoverPatternNil

`func (o *HdfsRecoverJobParams) SetHdfsRecoverPatternNil(b bool)`

 SetHdfsRecoverPatternNil sets the value for HdfsRecoverPattern to be an explicit nil

### UnsetHdfsRecoverPattern
`func (o *HdfsRecoverJobParams) UnsetHdfsRecoverPattern()`

UnsetHdfsRecoverPattern ensures that no value is present for HdfsRecoverPattern, not even an explicit nil
### GetTargetDirectory

`func (o *HdfsRecoverJobParams) GetTargetDirectory() string`

GetTargetDirectory returns the TargetDirectory field if non-nil, zero value otherwise.

### GetTargetDirectoryOk

`func (o *HdfsRecoverJobParams) GetTargetDirectoryOk() (*string, bool)`

GetTargetDirectoryOk returns a tuple with the TargetDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDirectory

`func (o *HdfsRecoverJobParams) SetTargetDirectory(v string)`

SetTargetDirectory sets TargetDirectory field to given value.

### HasTargetDirectory

`func (o *HdfsRecoverJobParams) HasTargetDirectory() bool`

HasTargetDirectory returns a boolean if a field has been set.

### SetTargetDirectoryNil

`func (o *HdfsRecoverJobParams) SetTargetDirectoryNil(b bool)`

 SetTargetDirectoryNil sets the value for TargetDirectory to be an explicit nil

### UnsetTargetDirectory
`func (o *HdfsRecoverJobParams) UnsetTargetDirectory()`

UnsetTargetDirectory ensures that no value is present for TargetDirectory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


