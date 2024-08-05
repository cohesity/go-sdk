# FileFolderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Specifies the environment of the object. | [optional] 
**HdfsParams** | Pointer to [**[]HdfsFileFolderParams**](HdfsFileFolderParams.md) | Specifies the parameters for Hdfs. | [optional] 
**PaginationInfo** | Pointer to [**PaginationInfo**](PaginationInfo.md) |  | [optional] 

## Methods

### NewFileFolderInfo

`func NewFileFolderInfo() *FileFolderInfo`

NewFileFolderInfo instantiates a new FileFolderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileFolderInfoWithDefaults

`func NewFileFolderInfoWithDefaults() *FileFolderInfo`

NewFileFolderInfoWithDefaults instantiates a new FileFolderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *FileFolderInfo) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FileFolderInfo) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FileFolderInfo) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FileFolderInfo) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *FileFolderInfo) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *FileFolderInfo) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetHdfsParams

`func (o *FileFolderInfo) GetHdfsParams() []HdfsFileFolderParams`

GetHdfsParams returns the HdfsParams field if non-nil, zero value otherwise.

### GetHdfsParamsOk

`func (o *FileFolderInfo) GetHdfsParamsOk() (*[]HdfsFileFolderParams, bool)`

GetHdfsParamsOk returns a tuple with the HdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsParams

`func (o *FileFolderInfo) SetHdfsParams(v []HdfsFileFolderParams)`

SetHdfsParams sets HdfsParams field to given value.

### HasHdfsParams

`func (o *FileFolderInfo) HasHdfsParams() bool`

HasHdfsParams returns a boolean if a field has been set.

### GetPaginationInfo

`func (o *FileFolderInfo) GetPaginationInfo() PaginationInfo`

GetPaginationInfo returns the PaginationInfo field if non-nil, zero value otherwise.

### GetPaginationInfoOk

`func (o *FileFolderInfo) GetPaginationInfoOk() (*PaginationInfo, bool)`

GetPaginationInfoOk returns a tuple with the PaginationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationInfo

`func (o *FileFolderInfo) SetPaginationInfo(v PaginationInfo)`

SetPaginationInfo sets PaginationInfo field to given value.

### HasPaginationInfo

`func (o *FileFolderInfo) HasPaginationInfo() bool`

HasPaginationInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


