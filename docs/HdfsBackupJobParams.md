# HdfsBackupJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HdfsExcludePattern** | Pointer to **[]string** | Any path/Glob pattern from HDFS that is to excluded. | [optional] 
**HdfsProtectPattern** | Pointer to **[]string** | Any path/Glob pattern from HDFS that is to protected. | [optional] 

## Methods

### NewHdfsBackupJobParams

`func NewHdfsBackupJobParams() *HdfsBackupJobParams`

NewHdfsBackupJobParams instantiates a new HdfsBackupJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHdfsBackupJobParamsWithDefaults

`func NewHdfsBackupJobParamsWithDefaults() *HdfsBackupJobParams`

NewHdfsBackupJobParamsWithDefaults instantiates a new HdfsBackupJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHdfsExcludePattern

`func (o *HdfsBackupJobParams) GetHdfsExcludePattern() []string`

GetHdfsExcludePattern returns the HdfsExcludePattern field if non-nil, zero value otherwise.

### GetHdfsExcludePatternOk

`func (o *HdfsBackupJobParams) GetHdfsExcludePatternOk() (*[]string, bool)`

GetHdfsExcludePatternOk returns a tuple with the HdfsExcludePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsExcludePattern

`func (o *HdfsBackupJobParams) SetHdfsExcludePattern(v []string)`

SetHdfsExcludePattern sets HdfsExcludePattern field to given value.

### HasHdfsExcludePattern

`func (o *HdfsBackupJobParams) HasHdfsExcludePattern() bool`

HasHdfsExcludePattern returns a boolean if a field has been set.

### SetHdfsExcludePatternNil

`func (o *HdfsBackupJobParams) SetHdfsExcludePatternNil(b bool)`

 SetHdfsExcludePatternNil sets the value for HdfsExcludePattern to be an explicit nil

### UnsetHdfsExcludePattern
`func (o *HdfsBackupJobParams) UnsetHdfsExcludePattern()`

UnsetHdfsExcludePattern ensures that no value is present for HdfsExcludePattern, not even an explicit nil
### GetHdfsProtectPattern

`func (o *HdfsBackupJobParams) GetHdfsProtectPattern() []string`

GetHdfsProtectPattern returns the HdfsProtectPattern field if non-nil, zero value otherwise.

### GetHdfsProtectPatternOk

`func (o *HdfsBackupJobParams) GetHdfsProtectPatternOk() (*[]string, bool)`

GetHdfsProtectPatternOk returns a tuple with the HdfsProtectPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsProtectPattern

`func (o *HdfsBackupJobParams) SetHdfsProtectPattern(v []string)`

SetHdfsProtectPattern sets HdfsProtectPattern field to given value.

### HasHdfsProtectPattern

`func (o *HdfsBackupJobParams) HasHdfsProtectPattern() bool`

HasHdfsProtectPattern returns a boolean if a field has been set.

### SetHdfsProtectPatternNil

`func (o *HdfsBackupJobParams) SetHdfsProtectPatternNil(b bool)`

 SetHdfsProtectPatternNil sets the value for HdfsProtectPattern to be an explicit nil

### UnsetHdfsProtectPattern
`func (o *HdfsBackupJobParams) UnsetHdfsProtectPattern()`

UnsetHdfsProtectPattern ensures that no value is present for HdfsProtectPattern, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


