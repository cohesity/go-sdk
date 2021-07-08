# CassandraRecoverJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CassandraAdditionalInfo** | Pointer to [**CassandraAdditionalParams**](CassandraAdditionalParams.md) |  | [optional] 
**CassandraSourceVersion** | Pointer to **NullableString** | Cassandra source version | [optional] 
**SelectedDataCenterVec** | Pointer to **[]string** | The data centers selected for recovery. | [optional] 
**StagingDirectoryVec** | Pointer to **[]string** | Cassandra staging directory | [optional] 
**Suffix** | Pointer to **NullableString** | A suffix that is to be applied to all recovered entities TODO (faizan.khan) : Remove this. | [optional] 

## Methods

### NewCassandraRecoverJobParams

`func NewCassandraRecoverJobParams() *CassandraRecoverJobParams`

NewCassandraRecoverJobParams instantiates a new CassandraRecoverJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraRecoverJobParamsWithDefaults

`func NewCassandraRecoverJobParamsWithDefaults() *CassandraRecoverJobParams`

NewCassandraRecoverJobParamsWithDefaults instantiates a new CassandraRecoverJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCassandraAdditionalInfo

`func (o *CassandraRecoverJobParams) GetCassandraAdditionalInfo() CassandraAdditionalParams`

GetCassandraAdditionalInfo returns the CassandraAdditionalInfo field if non-nil, zero value otherwise.

### GetCassandraAdditionalInfoOk

`func (o *CassandraRecoverJobParams) GetCassandraAdditionalInfoOk() (*CassandraAdditionalParams, bool)`

GetCassandraAdditionalInfoOk returns a tuple with the CassandraAdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraAdditionalInfo

`func (o *CassandraRecoverJobParams) SetCassandraAdditionalInfo(v CassandraAdditionalParams)`

SetCassandraAdditionalInfo sets CassandraAdditionalInfo field to given value.

### HasCassandraAdditionalInfo

`func (o *CassandraRecoverJobParams) HasCassandraAdditionalInfo() bool`

HasCassandraAdditionalInfo returns a boolean if a field has been set.

### GetCassandraSourceVersion

`func (o *CassandraRecoverJobParams) GetCassandraSourceVersion() string`

GetCassandraSourceVersion returns the CassandraSourceVersion field if non-nil, zero value otherwise.

### GetCassandraSourceVersionOk

`func (o *CassandraRecoverJobParams) GetCassandraSourceVersionOk() (*string, bool)`

GetCassandraSourceVersionOk returns a tuple with the CassandraSourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraSourceVersion

`func (o *CassandraRecoverJobParams) SetCassandraSourceVersion(v string)`

SetCassandraSourceVersion sets CassandraSourceVersion field to given value.

### HasCassandraSourceVersion

`func (o *CassandraRecoverJobParams) HasCassandraSourceVersion() bool`

HasCassandraSourceVersion returns a boolean if a field has been set.

### SetCassandraSourceVersionNil

`func (o *CassandraRecoverJobParams) SetCassandraSourceVersionNil(b bool)`

 SetCassandraSourceVersionNil sets the value for CassandraSourceVersion to be an explicit nil

### UnsetCassandraSourceVersion
`func (o *CassandraRecoverJobParams) UnsetCassandraSourceVersion()`

UnsetCassandraSourceVersion ensures that no value is present for CassandraSourceVersion, not even an explicit nil
### GetSelectedDataCenterVec

`func (o *CassandraRecoverJobParams) GetSelectedDataCenterVec() []string`

GetSelectedDataCenterVec returns the SelectedDataCenterVec field if non-nil, zero value otherwise.

### GetSelectedDataCenterVecOk

`func (o *CassandraRecoverJobParams) GetSelectedDataCenterVecOk() (*[]string, bool)`

GetSelectedDataCenterVecOk returns a tuple with the SelectedDataCenterVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedDataCenterVec

`func (o *CassandraRecoverJobParams) SetSelectedDataCenterVec(v []string)`

SetSelectedDataCenterVec sets SelectedDataCenterVec field to given value.

### HasSelectedDataCenterVec

`func (o *CassandraRecoverJobParams) HasSelectedDataCenterVec() bool`

HasSelectedDataCenterVec returns a boolean if a field has been set.

### SetSelectedDataCenterVecNil

`func (o *CassandraRecoverJobParams) SetSelectedDataCenterVecNil(b bool)`

 SetSelectedDataCenterVecNil sets the value for SelectedDataCenterVec to be an explicit nil

### UnsetSelectedDataCenterVec
`func (o *CassandraRecoverJobParams) UnsetSelectedDataCenterVec()`

UnsetSelectedDataCenterVec ensures that no value is present for SelectedDataCenterVec, not even an explicit nil
### GetStagingDirectoryVec

`func (o *CassandraRecoverJobParams) GetStagingDirectoryVec() []string`

GetStagingDirectoryVec returns the StagingDirectoryVec field if non-nil, zero value otherwise.

### GetStagingDirectoryVecOk

`func (o *CassandraRecoverJobParams) GetStagingDirectoryVecOk() (*[]string, bool)`

GetStagingDirectoryVecOk returns a tuple with the StagingDirectoryVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingDirectoryVec

`func (o *CassandraRecoverJobParams) SetStagingDirectoryVec(v []string)`

SetStagingDirectoryVec sets StagingDirectoryVec field to given value.

### HasStagingDirectoryVec

`func (o *CassandraRecoverJobParams) HasStagingDirectoryVec() bool`

HasStagingDirectoryVec returns a boolean if a field has been set.

### SetStagingDirectoryVecNil

`func (o *CassandraRecoverJobParams) SetStagingDirectoryVecNil(b bool)`

 SetStagingDirectoryVecNil sets the value for StagingDirectoryVec to be an explicit nil

### UnsetStagingDirectoryVec
`func (o *CassandraRecoverJobParams) UnsetStagingDirectoryVec()`

UnsetStagingDirectoryVec ensures that no value is present for StagingDirectoryVec, not even an explicit nil
### GetSuffix

`func (o *CassandraRecoverJobParams) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *CassandraRecoverJobParams) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *CassandraRecoverJobParams) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *CassandraRecoverJobParams) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *CassandraRecoverJobParams) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *CassandraRecoverJobParams) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


